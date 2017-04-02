// Code generated by reactGen. DO NOT EDIT.

package react

// PProps are the props for a <div> component
type PProps struct {
	ID                      string
	Key                     string
	ClassName               string
	Role                    string
	OnChange                func(e *SyntheticEvent)
	OnClick                 func(e *SyntheticMouseEvent)
	DangerouslySetInnerHTML *DangerousInnerHTMLDef
}

func (p *PProps) assign(v *_PProps) {

	if p.ID != "" {
		v.ID = p.ID
	}

	if p.Key != "" {
		v.Key = p.Key
	}

	v.ClassName = p.ClassName

	v.Role = p.Role

	v.OnChange = p.OnChange

	v.OnClick = p.OnClick

	v.DangerouslySetInnerHTML = p.DangerouslySetInnerHTML

}
