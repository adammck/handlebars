package handlebars

import (
	"bytes"
	"fmt"
	"strings"
	"strconv"
)

type Node interface {
	String() string
	Execute(interface{}) string
}

// TextNodes are just wrapped strings. The stuff in between the {{expressions}}
// and {{#blocks}}.
type TextNode struct {
	str string
}

func NewTextNode(str string) *TextNode {
	return &TextNode{str}
}

func (n *TextNode) String() string {
	return fmt.Sprintf("%#v", n.str)
}

func (n *TextNode) Execute(context interface{}) string {
	return n.str
}

type MustacheNode struct {
	expr    string
	escaped bool
}

func NewMustacheNode(expr string, escaped bool) *MustacheNode {
	return &MustacheNode{expr, escaped}
}

func (n *MustacheNode) String() string {
	return "{{" + n.expr + "}}"
}

func toString(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v

	case int:
		return strconv.Itoa(v)

	// Handlebars.js returns the string "true" for true, and an empty string for
	// false and undefined.
	case bool:
		if v {
			return "true"
		} else {
			return ""
		}

	// Handlebars.js doesn't actually explode when given a hash or array to
	// render. It just casts them to string. I'm not sure what to do about that,
	// since the string representation of them will be very different in Go. So
	// this is just a reminder to fix it later.
	default:
		panic(fmt.Sprintf("invalid value type: %T", v))
	}
}

func lookup(key string, context interface{}) string {

	// If the key is just a dot, render the current context.
	if key == "." {
		return toString(context)
	}

	// Assert that the context is a map[string]. Nothing else makes sense.
	ctx, ok := context.(map[string]interface{})
	if !ok {
		panic(fmt.Sprintf("Invalid map type %T", context))
	}

	// If the key is a single path segment (i.e. it doesn't contain any dots),
	// just fetch the corresponding value and render that. If the key doesn't
	// exist, just render an empty string.
	if !strings.Contains(key, ".") {
		v, ok := ctx[key]
		if ok {
			return toString(v)
		} else {
			return ""
		}
	}

	// We know that the key has a dot, so fetch the first path segment and it's
	// corresponding subcontext, and recurse with those.
	parts := strings.SplitN(key, ".", 2)
	k := parts[0]
	return lookup(parts[1], ctx[k])
}

func (n *MustacheNode) Execute(context interface{}) string {
	return lookup(n.expr, context)
}

type BlockNode struct {
	expr  string
	nodes []Node
}

func NewBlockNode(expr string) *BlockNode {
	return &BlockNode{expr, make([]Node, 0)}
}

func (n *BlockNode) String() string {

	// TODO: Surely there is a more succinct way of doing this.
	children := make([]string, len(n.nodes))
	for i, node := range n.nodes {
		children[i] = node.String()
	}

	return "{{#" + n.expr + " [" + strings.Join(children, ", ") + "]}}"
}

func (n *BlockNode) Append(node Node) {
	n.nodes = append(n.nodes, node)
}

func (n *BlockNode) Execute(context interface{}) string {
	var buffer bytes.Buffer

	for i := range n.nodes {
		buffer.WriteString(n.nodes[i].Execute(context))
	}

	return buffer.String()
}
