// Code generated by reactGen. DO NOT EDIT.

package react

// LabelProps defines the properties for the <label> element
type LabelProps struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTMLDef
	For                     string
	ID                      string
	Key                     string

	OnChange
	OnClick

	Role  string
	Style *CSS
}

func (l *LabelProps) assign(v *_LabelProps) {

	v.ClassName = l.ClassName

	v.DangerouslySetInnerHTML = l.DangerouslySetInnerHTML

	v.For = l.For

	if l.ID != "" {
		v.ID = l.ID
	}

	if l.Key != "" {
		v.Key = l.Key
	}

	if l.OnChange != nil {
		v.o.Set("onChange", l.OnChange.OnChange)
	}

	if l.OnClick != nil {
		v.o.Set("onClick", l.OnClick.OnClick)
	}

	v.Role = l.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = l.Style.hack()

}
