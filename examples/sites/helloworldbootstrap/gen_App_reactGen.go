// Code generated by reactGen. DO NOT EDIT.

package main

import "github.com/lijianying10/react"

type AppElem struct {
	react.Element
}

func buildApp(cd react.ComponentDef) react.Component {
	return AppDef{ComponentDef: cd}
}

func buildAppElem(children ...react.Element) *AppElem {
	return &AppElem{
		Element: react.CreateElement(buildApp, nil, children...),
	}
}

func (a AppDef) RendersElement() react.Element {
	return a.Render()
}
