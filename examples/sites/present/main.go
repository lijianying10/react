// Template generated by reactGen

package main

import (
	"github.com/lijianying10/react"

	"honnef.co/go/js/dom"
)

//go:generate reactGen

var document = dom.GetWindow().Document()

func main() {
	domTarget := document.GetElementByID("app")

	react.Render(App(), domTarget)
}
