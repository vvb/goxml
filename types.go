package goxml


type Attr struct {
	Name string
	Value string
}

type Elem struct {
	Tag string
	Attr []Attr
	Child []*Elem
	Parent *Elem
}

func (e *Elem) createAttr(name, value string) {
	e.Attr = append(e.Attr, Attr{Name:name, Value:value})
}

func (e *Elem) createElem(tag string) *Elem {
	elem := newElement(tag, e)
	e.childAdd(elem)
	return elem
}

func (e *Elem) childAdd(child *Elem) {
	e.Child = append(e.Child, child)
}

func newElement(tag string, parent *Elem) *Elem {
	return &Elem{
		Tag:    tag,
		Attr:   make([]Attr, 0),
		Child:  make([]*Elem, 0),
		Parent: parent,
	}
}