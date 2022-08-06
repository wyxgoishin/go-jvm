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

func (self *MethodDescriptor) addParameterType(t string) {
	pLen := len(self.parameterTypes)
	if pLen == cap(self.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, self.parameterTypes)
		self.parameterTypes = s
	}

	self.parameterTypes = append(self.parameterTypes, t)
}
