package datastore


func NewDesignEntity() *Design {
	design := Design{}

	return &design
}

func (o *Design) AddSpec(spec *Spec) *Design {
	o.Spec = spec
	return o
}