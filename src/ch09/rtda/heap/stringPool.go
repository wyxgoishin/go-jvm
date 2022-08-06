package heap

import "unicode/utf16"

var internedStrings = map[string]*Object{}

// string is stored in utf16 in Java String Object
func JvString(classloader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	chars := stringToUtf16(goStr)
	jvChars := &Object{class: classloader.LoadClass("[C"), data: chars}
	jvStr := classloader.LoadClass("java/lang/String").NewObject()
	jvStr.SetInstanceRefVar("value", "[C", jvChars)
	internedStrings[goStr] = jvStr
	return jvStr
}

func stringToUtf16(str string) []uint16 {
	runes := []rune(str)
	return utf16.Encode(runes)
}

func utf16ToString(chars []uint16) string {
	runes := utf16.Decode(chars)
	return string(runes)
}

func GoString(jvStr *Object) string {
	charArr := jvStr.GetInstanceRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

func InternedString(jvStr *Object) *Object {
	goStr := GoString(jvStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	internedStrings[goStr] = jvStr
	return jvStr
}
