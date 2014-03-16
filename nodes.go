package handlebars

import (
	"bytes"
	"fmt"
	"strings"
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

func (n *MustacheNode) Execute(context interface{}) string {
	val, ok := context.(map[string]string)[n.expr]
	if ok {
		return string(val)
	} else {
		return ""
	}
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
