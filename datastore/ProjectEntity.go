package datastore

func (o *Project) AddDesign(design *Design) *Project {
	o.Design = design
	return o
}

func (o *Project) SetTemplates(t []*Template) *Project {
	o.Templates = t
	return o
}

func (o *Project) AddTemplate(t *Template) *Project {
	if len(o.Templates)==0 {
		o.Templates = make([]*Template, 0)
	}
	// force the slice to increment the capacity / length / size
	o.Templates[len(o.Templates)] = t

	return o
}