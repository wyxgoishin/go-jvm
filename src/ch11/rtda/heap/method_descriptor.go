package heap

//ToDo: reference https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.3.3
// Format: (FieldTypes)ReturnType
// FieldType/ReturnType:
//   BasicType: B,C,S,I,J,F,D
//   ReferenceType: L<xxx>;
//   ArrayType: [<xxx>;
//   VoidType: V
type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func (methodDescriptor *MethodDescriptor) addParameterType(t string) {
	pLen := len(methodDescriptor.parameterTypes)
	if pLen == cap(methodDescriptor.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, methodDescriptor.parameterTypes)
		methodDescriptor.parameterTypes = s
	}

	methodDescriptor.parameterTypes = append(methodDescriptor.parameterTypes, t)
}
