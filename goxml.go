package goxml

import (
	"io"
	"encoding/xml"
)

func ParseXml(reader io.Reader) *Elem {
	dec := xml.NewDecoder(reader)
	top := newElement("top", nil)
	parent := top
	for {
		t, err := dec.RawToken()
		if t == nil {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			continue
		}

		switch {
		case err == io.EOF:
			break
		case err != nil:
			continue
		}

		switch t := t.(type) {
		case xml.StartElement:
			elem := parent.createElem(t.Name.Local)
			for _, a := range t.Attr {
				elem.createAttr(a.Name.Local, a.Value)
			}
			parent = elem
		case xml.EndElement:
			parent = parent.Parent
		case xml.CharData:
		case xml.Comment:
		case xml.Directive:
		case xml.ProcInst:
		}
	}
	return top
}


