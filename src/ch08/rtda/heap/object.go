package heap

type Object struct {
	class *Class
	data  interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (o *Object) Fields() Slots {
	return o.data.(Slots)
}

func (o *Object) IsInstanceOf(dst *Class) bool {
	return dst.isAssignableFrom(o.class)
}

func (o *Object) Class() *Class {
	return o.class
}

func (o *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := o.Class().getField(name, descriptor, false)
	slots := o.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (o *Object) GetRefVar(name, descriptor string) *Object {
	filed := o.class.getField(name, descriptor, false)
	slots := o.data.(Slots)
	return slots.GetRef(filed.slotId)
}
