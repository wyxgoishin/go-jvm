package utils

import "os"

const (
	VERSION = "0.0.12"
)

const (
	JAVA_HOME         = "JAVA_HOME"
	Jre               = "jre"
	Lib               = "lib"
	Ext               = "ext"
	JarExtension      = ".jar"
	PathListSeparator = string(os.PathListSeparator)
)

// error
const (
	ErrorJreNotFound   = "jre folder not found"
	ErrorClassNotFound = "given class not found in classpath"
)

var (
	OpCode2String = map[byte]string{
		0x00: "nop",
		0x01: "aconst_null",
		0x02: "iconst_m1",
		0x03: "iconst_0",
		0x04: "iconst_1",
		0x05: "iconst_2",
		0x06: "iconst_3",
		0x07: "iconst_4",
		0x08: "iconst_5",
		0x09: "lconst_0",
		0x0a: "lconst_1",
		0x0b: "fconst_0",
		0x0c: "fconst_1",
		0x0d: "fconst_2",
		0x0e: "dconst_0",
		0x0f: "dconst_1",
		0x10: "bi_push",
		0x11: "si_push",
		0x12: "ldc",
		0x13: "ldc_w",
		0x14: "ldc2_w",
		0x15: "iload",
		0x16: "lload",
		0x17: "fload",
		0x18: "dload",
		0x19: "aload",
		0x1a: "iload_0",
		0x1b: "iload_1",
		0x1c: "iload_2",
		0x1d: "iload_3",
		0x1e: "lload_0",
		0x1f: "lload_1",
		0x20: "lload_2",
		0x21: "lload_3",
		0x22: "fload_0",
		0x23: "fload_1",
		0x24: "fload_2",
		0x25: "fload_3",
		0x26: "dload_0",
		0x27: "dload_1",
		0x28: "dload_2",
		0x29: "dload_3",
		0x2a: "aload_0",
		0x2b: "aload_1",
		0x2c: "aload_2",
		0x2d: "aload_3",
		0x2e: "iaload",
		0x2f: "laload",
		0x30: "faload",
		0x31: "daload",
		0x32: "aaload",
		0x33: "baload",
		0x34: "caload",
		0x35: "saload",
		0x36: "istore",
		0x37: "lstore",
		0x38: "fstore",
		0x39: "dstore",
		0x3a: "astore",
		0x3b: "istore_0",
		0x3c: "istore_1",
		0x3d: "istore_2",
		0x3e: "istore_3",
		0x3f: "lstore_0",
		0x40: "lstore_1",
		0x41: "lstore_2",
		0x42: "lstore_3",
		0x43: "fstore_0",
		0x44: "fstore_1",
		0x45: "fstore_2",
		0x46: "fstore_3",
		0x47: "dstore_0",
		0x48: "dstore_1",
		0x49: "dstore_2",
		0x4a: "dstore_3",
		0x4b: "astore_0",
		0x4c: "astore_1",
		0x4d: "astore_2",
		0x4e: "astore_3",
		0x4f: "iastore",
		0x50: "lastore",
		0x51: "fastore",
		0x52: "dastore",
		0x53: "aastore",
		0x54: "bastore",
		0x55: "castore",
		0x56: "sastore",
		0x57: "pop",
		0x58: "pop2",
		0x59: "dup",
		0x5a: "dup_x1",
		0x5b: "dup_x2",
		0x5c: "dup2",
		0x5d: "dup2_x1",
		0x5e: "dup2_x2",
		0x5f: "swap",
		0x60: "iadd",
		0x61: "ladd",
		0x62: "fadd",
		0x63: "dadd",
		0x64: "isub",
		0x65: "lsub",
		0x66: "fsub",
		0x67: "dsub",
		0x68: "imul",
		0x69: "lmul",
		0x6a: "fmul",
		0x6b: "dmul",
		0x6c: "idiv",
		0x6d: "ldiv",
		0x6e: "fdiv",
		0x6f: "ddiv",
		0x70: "irem",
		0x71: "lrem",
		0x72: "frem",
		0x73: "drem",
		0x74: "ineg",
		0x75: "lneg",
		0x76: "fneg",
		0x77: "dneg",
		0x78: "ishl",
		0x79: "lshl",
		0x7a: "ishr",
		0x7b: "lshr",
		0x7c: "iushr",
		0x7d: "lushr",
		0x7e: "iand",
		0x7f: "land",
		0x80: "ior",
		0x81: "lor",
		0x82: "ixor",
		0x83: "lxor",
		0x84: "iinc",
		0x85: "i2l",
		0x86: "i2f",
		0x87: "i2d",
		0x88: "l2i",
		0x89: "l2f",
		0x8a: "l2d",
		0x8b: "f2i",
		0x8c: "f2l",
		0x8d: "f2d",
		0x8e: "d2i",
		0x8f: "d2l",
		0x90: "d2f",
		0x91: "i2b",
		0x92: "i2c",
		0x93: "i2s",
		0x94: "lcmp",
		0x95: "fcmpl",
		0x96: "fcmpg",
		0x97: "dcmpl",
		0x98: "dcmpg",
		0x99: "ifeq",
		0x9a: "ifne",
		0x9b: "iflt",
		0x9c: "ifge",
		0x9d: "ifgt",
		0x9e: "ifle",
		0x9f: "ificmp_eq",
		0xa0: "ificmp_ne",
		0xa1: "ificmp_lt",
		0xa2: "ificmp_ge",
		0xa3: "ificmp_gt",
		0xa4: "ificmp_le",
		0xa5: "ifacmp_eq",
		0xa6: "ifacmp_ne",
		0xa7: "jsr",
		0xa8: "ret",
		0xa9: "tableswitch",
		0xaa: "lookupswitch",
		0xac: "ireturn",
		0xad: "lreturn",
		0xae: "freturn",
		0xaf: "dreturn",
		0xb0: "areturn",
		0xb1: "_return",
		0xb2: "get_static",
		0xb3: "put_static",
		0xb4: "get_field",
		0xb5: "put_field",
		0xb6: "invoke_virtual",
		0xb7: "invoke_specila",
		0xb8: "invoke_static",
		0xb9: "invoke_interface",
		0xba: "invoke_dynamic",
		0xbb: "new",
		0xbc: "new_array",
		0xbd: "anew_array",
		0xbe: "arraylength",
		0xbf: "athrow",
		0xc0: "checkcast",
		0xc1: "instanceof",
		0xc2: "monitor_enter",
		0xc3: "monitor_exit",
		0xc4: "wide",
		0xc5: "multi_anew_array",
		0xc6: "ifnull",
		0xc7: "ifnonnull",
		0xc8: "goto_w",
		0xc9: "jsr_w",
		0xca: "breakpoint",
		0xfe: "impdep1",
		0xff: "impdep2",
	}
	String2OpCode = map[string]byte{
		"nop":              0x00,
		"aconst_null":      0x01,
		"iconst_m1":        0x02,
		"iconst_0":         0x03,
		"iconst_1":         0x04,
		"iconst_2":         0x05,
		"iconst_3":         0x06,
		"iconst_4":         0x07,
		"iconst_5":         0x08,
		"lconst_0":         0x09,
		"lconst_1":         0x0a,
		"fconst_0":         0x0b,
		"fconst_1":         0x0c,
		"fconst_2":         0x0d,
		"dconst_0":         0x0e,
		"dconst_1":         0x0f,
		"bi_push":          0x10,
		"si_push":          0x11,
		"ldc":              0x12,
		"ldc_w":            0x13,
		"ldc2_w":           0x14,
		"iload":            0x15,
		"lload":            0x16,
		"fload":            0x17,
		"dload":            0x18,
		"aload":            0x19,
		"iload_0":          0x1a,
		"iload_1":          0x1b,
		"iload_2":          0x1c,
		"iload_3":          0x1d,
		"lload_0":          0x1e,
		"lload_1":          0x1f,
		"lload_2":          0x20,
		"lload_3":          0x21,
		"fload_0":          0x22,
		"fload_1":          0x23,
		"fload_2":          0x24,
		"fload_3":          0x25,
		"dload_0":          0x26,
		"dload_1":          0x27,
		"dload_2":          0x28,
		"dload_3":          0x29,
		"aload_0":          0x2a,
		"aload_1":          0x2b,
		"aload_2":          0x2c,
		"aload_3":          0x2d,
		"iaload":           0x2e,
		"laload":           0x2f,
		"faload":           0x30,
		"daload":           0x31,
		"aaload":           0x32,
		"baload":           0x33,
		"caload":           0x34,
		"saload":           0x35,
		"istore":           0x36,
		"lstore":           0x37,
		"fstore":           0x38,
		"dstore":           0x39,
		"astore":           0x3a,
		"istore_0":         0x3b,
		"istore_1":         0x3c,
		"istore_2":         0x3d,
		"istore_3":         0x3e,
		"lstore_0":         0x3f,
		"lstore_1":         0x40,
		"lstore_2":         0x41,
		"lstore_3":         0x42,
		"fstore_0":         0x43,
		"fstore_1":         0x44,
		"fstore_2":         0x45,
		"fstore_3":         0x46,
		"dstore_0":         0x47,
		"dstore_1":         0x48,
		"dstore_2":         0x49,
		"dstore_3":         0x4a,
		"astore_0":         0x4b,
		"astore_1":         0x4c,
		"astore_2":         0x4d,
		"astore_3":         0x4e,
		"iastore":          0x4f,
		"lastore":          0x50,
		"fastore":          0x51,
		"dastore":          0x52,
		"aastore":          0x53,
		"bastore":          0x54,
		"castore":          0x55,
		"sastore":          0x56,
		"pop":              0x57,
		"pop2":             0x58,
		"dup":              0x59,
		"dup_x1":           0x5a,
		"dup_x2":           0x5b,
		"dup2":             0x5c,
		"dup2_x1":          0x5d,
		"dup2_x2":          0x5e,
		"swap":             0x5f,
		"iadd":             0x60,
		"ladd":             0x61,
		"fadd":             0x62,
		"dadd":             0x63,
		"isub":             0x64,
		"lsub":             0x65,
		"fsub":             0x66,
		"dsub":             0x67,
		"imul":             0x68,
		"lmul":             0x69,
		"fmul":             0x6a,
		"dmul":             0x6b,
		"idiv":             0x6c,
		"ldiv":             0x6d,
		"fdiv":             0x6e,
		"ddiv":             0x6f,
		"irem":             0x70,
		"lrem":             0x71,
		"frem":             0x72,
		"drem":             0x73,
		"ineg":             0x74,
		"lneg":             0x75,
		"fneg":             0x76,
		"dneg":             0x77,
		"ishl":             0x78,
		"lshl":             0x79,
		"ishr":             0x7a,
		"lshr":             0x7b,
		"iushr":            0x7c,
		"lushr":            0x7d,
		"iand":             0x7e,
		"land":             0x7f,
		"ior":              0x80,
		"lor":              0x81,
		"ixor":             0x82,
		"lxor":             0x83,
		"iinc":             0x84,
		"i2l":              0x85,
		"i2f":              0x86,
		"i2d":              0x87,
		"l2i":              0x88,
		"l2f":              0x89,
		"l2d":              0x8a,
		"f2i":              0x8b,
		"f2l":              0x8c,
		"f2d":              0x8d,
		"d2i":              0x8e,
		"d2l":              0x8f,
		"d2f":              0x90,
		"i2b":              0x91,
		"i2c":              0x92,
		"i2s":              0x93,
		"lcmp":             0x94,
		"fcmpl":            0x95,
		"fcmpg":            0x96,
		"dcmpl":            0x97,
		"dcmpg":            0x98,
		"ifeq":             0x99,
		"ifne":             0x9a,
		"iflt":             0x9b,
		"ifge":             0x9c,
		"ifgt":             0x9d,
		"ifle":             0x9e,
		"ificmp_eq":        0x9f,
		"ificmp_ne":        0xa0,
		"ificmp_lt":        0xa1,
		"ificmp_ge":        0xa2,
		"ificmp_gt":        0xa3,
		"ificmp_le":        0xa4,
		"ifacmp_eq":        0xa5,
		"ifacmp_ne":        0xa6,
		"jsr":              0xa7,
		"ret":              0xa8,
		"tableswitch":      0xa9,
		"lookupswitch":     0xaa,
		"ireturn":          0xac,
		"lreturn":          0xad,
		"freturn":          0xae,
		"dreturn":          0xaf,
		"areturn":          0xb0,
		"_return":          0xb1,
		"get_static":       0xb2,
		"put_static":       0xb3,
		"get_field":        0xb4,
		"put_field":        0xb5,
		"invoke_virtual":   0xb6,
		"invoke_special":   0xb7,
		"invoke_static":    0xb8,
		"invoke_interface": 0xb9,
		"invoke_dynamic":   0xba,
		"new":              0xbb,
		"new_array":        0xbc,
		"anew_array":       0xbd,
		"arraylength":      0xbe,
		"athrow":           0xbf,
		"checkcast":        0xc0,
		"instanceof":       0xc1,
		"monitor_enter":    0xc2,
		"monitor_exit":     0xc3,
		"wide":             0xc4,
		"multi_anew_array": 0xc5,
		"ifnull":           0xc6,
		"ifnonnull":        0xc7,
		"goto_w":           0xc8,
		"jsr_w":            0xc9,
		"breakpoint":       0xca,
		"impdep1":          0xfe,
		"impdep2":          0xff,
	}
)
