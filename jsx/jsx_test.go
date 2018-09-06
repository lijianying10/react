// +build js

package jsx_test

import (
	"testing"

	"honnef.co/go/js/dom"

	"github.com/lijianying10/react"
	"github.com/lijianying10/react/jsx"
	"github.com/lijianying10/react/testutils"
)

func TestIssue_10(t *testing.T) {
	const class = "test"

	div := react.Div(nil, jsx.HTML(`
		<nav><ul><li class="`+class+`">Testing</li>
		</ul></nav>
	`)...)

	x := testutils.Wrapper(div)
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLLIElement); !ok {
		t.Fatal("Failed to find <li> element")
	}
}
