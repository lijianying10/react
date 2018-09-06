package main

import (
	"math/rand"
	"time"

	"github.com/lijianying10/react"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type AppDef struct {
	react.ComponentDef
}

func App() *AppElem {
	return buildAppElem()
}

func (a AppDef) Render() react.Element {
	return Latency()
}
