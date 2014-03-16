package handlebars

%%{
  machine handlebars;

  open  = '{' :> '{';
  close = '}' :> '}';

  unescaped_open  = open :> '{';
  unescaped_close = close :> '}';

  # Note the current pointer position.
  action mark {
    x = fpc;
  }

  # Create a text node from MARK to FPC
  action make_text {
    if fpc > x {
      text := data[x:fpc]
      node := stack.Peek()
      node.Append(NewTextNode(text))

      // ???
      x = fpc
    }
  }



  action start_expr {
    m = fpc
  }

  action end_expr {
    expr = data[m:fpc]
  }


  action set_escaped   { m_esc = true }
  action set_unescaped { m_esc = false }

  action make_mustache {
    log("M")
    node := stack.Peek()
    node.Append(NewMustacheNode(expr, m_esc))
  }

  action make_block_open {
    log("#")
    child := NewBlockNode(expr)
    stack.Push(child)
  }

  # TODO: Assert that the expr of the block we're closing is the same as the
  #       block we're popping off the stack. Currently we're not checking.
  action make_block_close {
    log("/")
    child := stack.Pop()
    parent := stack.Peek()
    parent.Append(child)
  }



  action error {
    log("!")
    panic(fmt.Sprintf("Error at: %d", fpc))
  }


  expr = (
    space*
    lower+ >start_expr %end_expr
    space*
  );

  unescaped_var = (
    unescaped_open
    expr >set_unescaped
    unescaped_close >make_mustache %mark
  ) >make_text;

  var = (
    open
    expr >set_escaped
    close >make_mustache %mark
  ) >make_text; # create text element from mark to start of var

  block_open = (
    open
    '#'
    expr
    close >make_block_open %mark
  ) >make_text;

  block_close = (
    open
    '/'
    expr
    close >make_block_close %mark
  ) >make_text;

  text = (any+ -- open);

  statement = (
    unescaped_var
    | var
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

func Compile(source string) *BlockNode {
  fmt.Printf("\n\n %#v\n", source)
  root := NewBlockNode("")
  stack := NewStack()
  stack.Push(root)

  x := 0 // mark
  m := 0 // start of expr
  expr := "" // last expression

  // current mustache properties
  // initialized in init_mustache
  var m_esc bool

  // Ragel vars
  cs   := 0           // Current state
  p    := 0           // Data pointer
  pe   := len(source) // Data end pointer
  eof  := pe          // End of file pointer
  data := source      // array containting the data to process

  log := func(label string) {
    fmt.Printf(label + strings.Repeat(" ", p) + "^\n")
  }

  // -- BEGIN RAGEL GENERATED STUFF --------------------------------------------
  %% write data;
  %% write init;
  %% write exec;
  // ---------------------------------------------------------------------------

  _ = m_esc
  return root
}
