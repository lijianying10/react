// Code generated by reactGen. DO NOT EDIT.

package react

// IProps are the props for a <i> component
type IProps struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	ID                      string
	Key                     string

	OnChange
	OnClick

	Role  string
	Src   string
	Style *CSS
}

func (i *IProps) assign(v *_IProps) {

	v.ClassName = i.ClassName

	v.DangerouslySetInnerHTML = i.DangerouslySetInnerHTML

	if i.ID != "" {
		v.ID = i.ID
	}

	if i.Key != "" {
		v.Key = i.Key
	}

	if i.OnChange != nil {
		v.o.Set("onChange", i.OnChange.OnChange)
	}

	if i.OnClick != nil {
		v.o.Set("onClick", i.OnClick.OnClick)
	}

	v.Role = i.Role

	v.Src = i.Src

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = i.Style.hack()

}
