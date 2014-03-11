package handlebars

%%{
  machine handlebars;

  open  = '{' :> '{';
  close = '}' :> '}';


  # Note the current pointer position.
  action mark {
    x = fpc;
  }

  # Create a text node from MARK to FPC
  action make_text {
    if fpc > x {
      text := data[x:fpc]
      log("T", x, fpc);

      node := stack.Peek()
      node.Append(TextNode(text))

      // ???
      x = fpc
    }
  }

  action start_mustache {
    m = fpc
  }

  action make_mustache {
    text := data[m:fpc]
    log("M", m, fpc);
    node := stack.Peek()
    node.Append(MustacheNode{text})
  }


  action start_block_open {
    m = fpc
  }

  action make_block_open {
    text := data[m:fpc]
    log("#", m, fpc);
    child := BlockNode{text, make([]Node, 0)}
    stack.Push(&child)
  }


  action start_block_close {
    m = fpc
  }

  action make_block_close {
    log("/", m, fpc);
    child := stack.Pop()
    parent := stack.Peek()
    parent.Append(*child)
  }



  action error {
    panic("Error at" + string(fpc))
  }

  var = (
    open
    space*                                # zero or more spaces
    lower+ >start_mustache %make_mustache #
    space*                                # more optional spaces
    close %mark                           # mark after close, since that's where the next text block starts
  ) >make_text;                           # create text element from mark to start of var

  block_open = (
    open
    '#'
    space*
    lower+ >start_block_open %make_block_open
    space*
    close %mark
  ) >make_text;

  block_close = (
    open
    '/'
    space*
    lower+ >start_block_close %make_block_close
    space*
    close %mark
  ) >make_text;

  text = (any+ -- open);

  statement = (
    var
    | block_open
    | block_close
    | text
  );

  main := statement* %eof(make_text) $err(error);
}%%


import (
  "fmt"
  "strings"
)




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




func log(label string, start int, end int) {
  fmt.Printf(label + strings.Repeat(" ", start + 1) + strings.Repeat("─", (end - start)) + "\n")
}

func Compile(source string) *BlockNode {
  fmt.Printf("\n\nC%#v\n", source)
  root := NewBlockNode("")

  stack := NewStack()
  stack.Push(root)

  x := 0 // mark
  m := 0 // start of identifier

  // Ragel vars
  cs   := 0           // Current state
  p    := 0           // Data pointer
  pe   := len(source) // Data end pointer
  eof  := pe          // End of file pointer
  data := source      // array containting the data to process

  // -- BEGIN RAGEL GENERATED STUFF --------------------------------------------
  %% write data;
  %% write init;
  %% write exec;
  // ---------------------------------------------------------------------------

  return root
}
