package gr

import (
	"github.com/gopherjs/gopherjs/js"
)

// A Modifier modifies an element, adding attributes, style or child elements etc.
type Modifier interface {
	Modify(element *Element)
}

type Element struct {
	tag        string
	properties map[string]interface{}
	style      map[string]interface{}

	// An element is either a text node or a container with children.
	text     string
	children []Component

	// This is the actual ReactJS element.
	// ReactElement, ReactText or a ReactFragment
	element *js.Object
}

func NewElement(tag string) *Element {
	return &Element{tag: tag}
}

func (e *Element) Node() *js.Object {
	if e.element == nil {
		e.element = e.createElement()
	}

	return e.element
}

func (e *Element) Modify(in *Element) {
	in.children = append(in.children, e)
}

type Modifiers []Modifier

func (mods Modifiers) Modify(e *Element) {
	for _, m := range mods {
		if m != nil {
			m.Modify(e)
		}
	}
}

// TODO(bep)
var counter = 1

func (e *Element) createElement() *js.Object {
	// TODO(bep) figure a better way to add keys.

	if e.properties == nil {
		e.properties = make(map[string]interface{})
	}

	if _, ok := e.properties["key"]; !ok {
		e.properties["key"] = counter
		counter++
	}

	if len(e.style) != 0 {
		// TODO(bep) merge with existing?
		e.properties["style"] = e.style
	}

	if e.text != "" {
		return react.Call("createElement", e.tag, e.properties, e.text)
	}

	if len(e.children) > 0 {
		children := make([]interface{}, len(e.children))

		for _, c := range e.children {
			children = append(children, c.Node())
		}

		return createElement(e.tag, e.properties, children)
	}

	return createElement(e.tag, e.properties)
}

func createElement(tag string, props map[string]interface{}, children ...interface{}) *js.Object {
	if len(children) == 0 {
		return react.Call("createElement", tag, props)
	}
	return react.Call("createElement", tag, props, children)
}
