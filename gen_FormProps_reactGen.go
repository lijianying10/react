// Code generated by reactGen. DO NOT EDIT.

package react

// FormProps defines the properties for the <form> element
type FormProps struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTMLDef
	ID                      string
	Key                     string

	OnChange
	OnClick

	Role  string
	Style *CSS
}

func (f *FormProps) assign(v *_FormProps) {

	v.ClassName = f.ClassName

	v.DangerouslySetInnerHTML = f.DangerouslySetInnerHTML

	if f.ID != "" {
		v.ID = f.ID
	}

	if f.Key != "" {
		v.Key = f.Key
	}

	if f.OnChange != nil {
		v.o.Set("onChange", f.OnChange.OnChange)
	}

	if f.OnClick != nil {
		v.o.Set("onClick", f.OnClick.OnClick)
	}

	v.Role = f.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = f.Style.hack()

}
