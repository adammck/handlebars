package handlebars

import (
  "fmt"
  "strings"
)




type Node interface {
  String() string
}




// TextNodes are just wrapped strings. The stuff in between the {{expressions}}
// and {{#blocks}}.
type TextNode struct {
  str string
}

func NewTextNode(str string) *TextNode {
  return &TextNode{str}
}

func (n TextNode) String() string {
  return fmt.Sprintf("%#v", n.str)
}




type MustacheNode struct {
  expr string
}

func NewMustacheNode(expr string) *MustacheNode {
  return &MustacheNode{expr}
}

func (n MustacheNode) String() string {
  return "{{" + n.expr + "}}"
}




type BlockNode struct {
  expr string
  nodes []Node
}

func NewBlockNode(expr string) *BlockNode {
  return &BlockNode{expr, make([]Node, 0)}
}

func (n BlockNode) String() string {

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
