package testutils

import "github.com/lijianying10/react"

//go:generate reactGen

// WrapperDef is the definition of the Wrapper component
type WrapperDef struct {
	react.ComponentDef
}

// Wrapper creates instances of the Wrapper component
func Wrapper(children ...react.Element) *WrapperElem {
	return buildWrapperElem(children...)
}

// Render renders the Wrapper component
func (h WrapperDef) Render() react.Element {
	return react.Div(nil, h.Children()...)
}
