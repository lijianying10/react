/*

Package jsx allows you to render blocks of HTML as github.com/lijianying10/react elements.
It is a temporary runtime solution for what will become a compile-time
transpilation, much like JSX's relationship with Javascript.

For more information see https://github.com/myitcv/react/wiki

*/
package jsx

import (
	"fmt"
	"strings"

	"github.com/lijianying10/react"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var htmlCache = make(map[string][]react.Element)

// HTML is a runtime JSX-like parsereact. It parses the supplied HTML string into
// github.com/lijianying10/react element values. It exists as a stop-gap runtime solution to
// full JSX-like support within the GopherJS compilereact. It should only be used
// where the argument is a compile-time constant string (TODO enforce this
// within reactVet). HTML will panic in case s cannot be parsed as a valid HTML
// fragment
//
func HTML(s string) []react.Element {
	s = strings.TrimSpace(s)

	if v, ok := htmlCache[s]; ok {
		return v
	}

	toParse := getParse(s)

	var res []react.Element

	for _, v := range toParse {
		res = append(res, parse(v))
	}

	htmlCache[s] = res

	return res
}

// HTMLReactElementTemplate Parse Component map key as HTML node name. The key must be lower-case.
func HTMLReactElementTemplate(s string, comp map[string]react.Element) []react.Element {
	toParse := getParse(s)

	var res []react.Element

	for _, v := range toParse {
		if val, ok := comp[v.Data]; ok {
			res = append(res, val)
		} else {
			res = append(res, parse(v))
		}
	}

	return res
}

func getParse(s string) []*html.Node {
	// a dummy div for parsing the fragment
	div := &html.Node{
		Type:     html.ElementNode,
		Data:     "div",
		DataAtom: atom.Div,
	}

	elems, err := html.ParseFragment(strings.NewReader(s), div)
	if err != nil {
		panic(fmt.Errorf("failed to parse HTML %q: %v", s, err))
	}

	var toParse []*html.Node
	var toWalk []*html.Node
	for _, v := range elems {
		if v.Type == html.TextNode && strings.TrimSpace(v.Data) == "" {
			continue
		}
		toParse = append(toParse, v)
		toWalk = append(toWalk, v)
	}

	var v *html.Node

	for len(toWalk) > 0 {
		v, toWalk = toWalk[0], toWalk[1:]

		c := v.FirstChild

		for c != nil {
			if c.Type == html.TextNode && strings.TrimSpace(c.Data) == "" {
				v.RemoveChild(c)
			}

			toWalk = append(toWalk, c)
			c = c.NextSibling
		}
	}
	return toParse
}

// HTMLElem is a convenience wrapper around HTML where only a single root
// element is expected. HTMLElem will panic if more than one HTML element
// results
//
func HTMLElem(s string) react.Element {
	res := HTML(s)

	if v := len(res); v != 1 {
		panic(fmt.Errorf("expected single element result from %q; got %v", s, v))
	}

	return res[0]
}

// HTMLElemReactElementTemplate Parse Component map key as HTML node name. The key must be lower-case.
func HTMLElemReactElementTemplate(s string, comp map[string]react.Element) react.Element {
	res := HTMLReactElementTemplate(s, comp)

	if v := len(res); v != 1 {
		panic(fmt.Errorf("expected single element result from %q; got %v", s, v))
	}

	return res[0]
}
