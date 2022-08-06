package lang

import (
	"go-jvm/src/ch11/instruction/base"
	"go-jvm/src/ch11/native"
	"go-jvm/src/ch11/rtda"
	"go-jvm/src/ch11/rtda/heap"
	"runtime"
	"time"
)

func init() {
	native.Register("java/lang/System", "arraycopy",
		"(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
	native.Register("java/lang/System", "initProperties",
		"(Ljava/util/Properties;)Ljava/util/Properties;", initProperties)
	native.Register("java/lang/System", "setIn0",
		"(Ljava/io/InputStream;)V", setIn0)
	native.Register("java/lang/System", "setOut0",
		"(Ljava/io/PrintStream;)V", setOut0)
	native.Register("java/lang/System", "setErr0",
		"(Ljava/io/PrintStream;)V", setErr0)
	native.Register("java/lang/System", "currentTimeMillis",
		"()J", currentTimeMillis)
}

// public static native void arraycopy(Object src, int srcPos, Object dst, int dstPos, int length)
func arraycopy(frame *rtda.Frame) {
	localVars := frame.LocalVars
	src := localVars.GetRef(0)
	srcPos := localVars.GetInt(1)
	dst := localVars.GetRef(2)
	dstPos := localVars.GetInt(3)
	length := localVars.GetInt(4)

	if src == nil || dst == nil {
		panic("java.lang.NullPointerException")
	}
	if !checkArrayCopy(src, dst) {
		panic("java.lang.ArrayStoreException")
	}
	if srcPos < 0 || dstPos < 0 || length < 0 || srcPos+length > src.ArrayLength() || dstPos+length > dst.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}
	heap.ArrayCopy(src, srcPos, dst, dstPos, length)
}

func checkArrayCopy(src, dst *heap.Object) bool {
	srcClass := src.Class()
	dstClass := dst.Class()
	if !srcClass.IsArray() || !dstClass.IsArray() {
		return false
	}
	if srcClass.ComponentClass().IsPrimitive() || dstClass.ComponentClass().IsPrimitive() {
		return srcClass == dstClass
	}
	return srcClass == dstClass
}

//ToDo:

// private static native Properties initProperties(Properties props);
// (Ljava/util/Properties;)Ljava/util/Properties;
func initProperties(frame *rtda.Frame) {
	vars := frame.LocalVars
	props := vars.GetRef(0)

	stack := frame.OperandStack
	stack.PushRef(props)

	// public synchronized Object setProperty(String key, String value)
	setPropMethod := props.Class().GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	thread := frame.Thread()
	for key, val := range _sysProps() {
		jvKey := heap.JvString(frame.Method().Class().ClassLoader(), key)
		jvVal := heap.JvString(frame.Method().Class().ClassLoader(), val)
		ops := rtda.NewOperandStack(3)
		ops.PushRef(props)
		ops.PushRef(jvKey)
		ops.PushRef(jvVal)
		shimFrame := rtda.NewShimFrame(thread, ops)
		thread.PushFrame(shimFrame)

		base.InvokeMethod(shimFrame, setPropMethod)
	}
}

func _sysProps() map[string]string {
	return map[string]string{
		"java.version":         "1.8.0",
		"java.vendor":          "jvm.go",
		"java.vendor.url":      "https://github.com/zxh0/jvm.go",
		"java.home":            "todo",
		"java.class.version":   "52.0",
		"java.class.path":      "todo",
		"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
		"os.name":              runtime.GOOS,   // todo
		"os.arch":              runtime.GOARCH, // todo
		"os.version":           "",             // todo
		"file.separator":       "/",            // todo os.PathSeparator
		"path.separator":       ":",            // todo os.PathListSeparator
		"line.separator":       "\n",           // todo
		"user.name":            "",             // todo
		"user.home":            "",             // todo
		"user.dir":             ".",            // todo
		"user.country":         "CN",           // todo
		"file.encoding":        "UTF-8",
		"sun.stdout.encoding":  "UTF-8",
		"sun.stderr.encoding":  "UTF-8",
	}
}

// private static native void setIn0(InputStream in);
// (Ljava/io/InputStream;)V
func setIn0(frame *rtda.Frame) {
	vars := frame.LocalVars
	in := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetStaticRefVar("in", "Ljava/io/InputStream;", in)
}

// private static native void setOut0(PrintStream out);
// (Ljava/io/PrintStream;)V
func setOut0(frame *rtda.Frame) {
	vars := frame.LocalVars
	out := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetStaticRefVar("out", "Ljava/io/PrintStream;", out)
}

// private static native void setErr0(PrintStream err);
// (Ljava/io/PrintStream;)V
func setErr0(frame *rtda.Frame) {
	vars := frame.LocalVars
	err := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetStaticRefVar("err", "Ljava/io/PrintStream;", err)
}

// public static native long currentTimeMillis();
// ()J
func currentTimeMillis(frame *rtda.Frame) {
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	stack := frame.OperandStack
	stack.PushLong(millis)
}
