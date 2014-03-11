package handlebars




type Node interface {
  String() string
}




type TextNode string

func (n TextNode) String() string {
  return string(n)
}




type MustacheNode struct {
  expr string
}

func (n MustacheNode) String() string {
  return n.expr
}




type BlockNode struct {
  expr string
  nodes []Node
}

func NewBlockNode(expr string) *BlockNode {
  return &BlockNode{expr, make([]Node, 0)}
}

func (n BlockNode) String() string {
  return n.expr
}

func (n *BlockNode) Append(node Node) {
  n.nodes = append(n.nodes, node)
}
