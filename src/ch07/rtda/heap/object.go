package heap

type Object struct {
	class  *Class
	fields Slots
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (o *Object) Fields() Slots {
	return o.fields
}

func (o *Object) IsInstanceOf(dst *Class) bool {
	return dst.isAssignableFrom(o.class)
}

func (o *Object) Class() *Class {
	return o.class
}
