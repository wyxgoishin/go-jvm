package heap

const (
	ACC_PUBLIC       = 0x0001 // class field method
	ACC_PRIVATE      = 0x0002 //       field method
	ACC_PROTECTED    = 0x0004 //       field method
	ACC_STATIC       = 0x0008 //       field method
	ACC_FINAL        = 0x0010 // class field method
	ACC_SUPER        = 0x0020 // class
	ACC_SYNCHRONIZED = 0x0020 //             method
	ACC_VOLATILE     = 0x0040 //       field
	ACC_BRIDGE       = 0x0040 //             method
	ACC_TRANSIENT    = 0x0080 //       field
	ACC_VARARGS      = 0x0080 //             method
	ACC_NATIVE       = 0x0100 //             method
	ACC_INTERFACE    = 0x0200 // class
	ACC_ABSTRACT     = 0x0400 // class       method
	ACC_STRICT       = 0x0800 //             method
	ACC_SYNTHETIC    = 0x1000 // class field method
	ACC_ANNOTATION   = 0x2000 // class
	ACC_ENUM         = 0x4000 // class field
)

func IsPublic(accFlags uint16) bool {
	return accFlags&ACC_PUBLIC > 0
}

func IsPrivate(accFlags uint16) bool {
	return accFlags&ACC_PRIVATE > 0
}

func IsProtected(accFlags uint16) bool {
	return accFlags&ACC_PROTECTED > 0
}

func IsStatic(accFlags uint16) bool {
	return accFlags&ACC_STATIC > 0
}

func IsFinal(accFlags uint16) bool {
	return accFlags&ACC_FINAL > 0
}

func IsSuper(accFlags uint16) bool {
	return accFlags&ACC_SUPER > 0
}

func IsSynchronized(accFlags uint16) bool {
	return accFlags&ACC_SYNCHRONIZED > 0
}

func IsVolatile(accFlags uint16) bool {
	return accFlags&ACC_VOLATILE > 0
}

func IsBridge(accFlags uint16) bool {
	return accFlags&ACC_BRIDGE > 0
}

func IsTransient(accFlags uint16) bool {
	return accFlags&ACC_TRANSIENT > 0
}

func IsVarArgs(accFlags uint16) bool {
	return accFlags&ACC_VARARGS > 0
}

func IsNative(accFlags uint16) bool {
	return accFlags&ACC_NATIVE > 0
}

func IsInterface(accFlags uint16) bool {
	return accFlags&ACC_INTERFACE > 0
}

func IsAbstract(accFlags uint16) bool {
	return accFlags&ACC_ABSTRACT > 0
}

func IsStrict(accFlags uint16) bool {
	return accFlags&ACC_STRICT > 0
}

func IsSynthetic(accFlags uint16) bool {
	return accFlags&ACC_SYNTHETIC > 0
}

func IsAnnotation(accFlags uint16) bool {
	return accFlags&ACC_ANNOTATION > 0
}

func IsEnum(accFlags uint16) bool {
	return accFlags&ACC_ENUM > 0
}
