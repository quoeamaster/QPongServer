package datastore

func NewSpecEntity() *Spec {
	return &Spec{}
}

func (o *Spec) AddBackgroundImagePath(path string) *Spec {
	o.BackgroundImagePath = path
	return o
}

func (o *Spec) AddTitle(t *TextBlock) *Spec {
	o.Title = t
	return o
}

func (o *Spec) AddSubTitle(t *TextBlock) *Spec {
	o.SubTitle = t
	return o
}

func (o *Spec) AddDescription(t *TextBlock) *Spec {
	o.Description = t
	return o
}
