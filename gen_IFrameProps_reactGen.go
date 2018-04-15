// Code generated by reactGen. DO NOT EDIT.

package react

// IFrameProps are the props for a <iframe> component
type IFrameProps struct {
	AriaSet
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	ID  string
	Key string

	OnChange
	OnClick

	Ref
	Role   string
	Src    string
	SrcDoc string
	Style  *CSS
}

func (i *IFrameProps) assign(_v *_IFrameProps) {

	if i.AriaSet != nil {
		for dk, dv := range i.AriaSet {
			_v.o.Set("aria-"+dk, dv)
		}
	}

	_v.ClassName = i.ClassName

	_v.DangerouslySetInnerHTML = i.DangerouslySetInnerHTML

	if i.DataSet != nil {
		for dk, dv := range i.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if i.ID != "" {
		_v.ID = i.ID
	}

	if i.Key != "" {
		_v.Key = i.Key
	}

	if i.OnChange != nil {
		_v.o.Set("onChange", i.OnChange.OnChange)
	}

	if i.OnClick != nil {
		_v.o.Set("onClick", i.OnClick.OnClick)
	}

	if i.Ref != nil {
		_v.o.Set("ref", i.Ref.Ref)
	}

	_v.Role = i.Role

	_v.Src = i.Src

	_v.SrcDoc = i.SrcDoc

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = i.Style.hack()

}
