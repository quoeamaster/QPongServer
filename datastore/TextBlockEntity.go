package datastore

func NewTextBlockEntity() *TextBlock {
	return &TextBlock{}
}

func (o *TextBlock) AddDimension(d *Dimension) *TextBlock {
	o.Dimen = d
	return o
}

func (o *TextBlock) AddText(s string) *TextBlock {
	o.Text = s
	return o
}
