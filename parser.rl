package handlebars

%%{
  machine handlebars;

  open  = '{' :> '{';
  close = '}' :> '}';
  unescapedOpen  = open :> '{';
  unescapedClose = close :> '}';

  # ----------------------------------------------------------------------------

  # Note the current pointer position.
  action mark {
    x = fpc;
  }

  # Create a text node from MARK to FPC
  action makeText {
    if fpc > x {
      text := data[x:fpc]
      node := stack.Peek()
      node.Append(NewTextNode(text))

      // ???
      x = fpc
    }
  }


  action startExpr {
    m = fpc
  }

  action endExpr {
    expr = data[m:fpc]
  }

  action setEscaped   { m_esc = true }
  action setUnescaped { m_esc = false }

  action makeMustache {
    log("M")
    node := stack.Peek()
    node.Append(NewMustacheNode(expr, m_esc))
  }

  action makeBlockOpen {
    log("#")
    child := NewBlockNode(expr)
    stack.Push(child)
  }

  # TODO: Assert that the expr of the block we're closing is the same as the
  #       block we're popping off the stack. Currently we're not checking.
  action makeCloseBlock {
    log("/")
    child := stack.Pop()
    parent := stack.Peek()
    parent.Append(child)
  }

  action error {
    log("!")
    panic(fmt.Sprintf("Error at: %d", fpc))
  }

  # ----------------------------------------------------------------------------

  expr = (
    space*
    [a-z\.]+ >startExpr %endExpr
    space*
  );

  unescapedVar = (
    unescapedOpen
    expr >setUnescaped
    unescapedClose >makeMustache %mark
  ) >makeText;

  var = (
    open
    expr >setEscaped
    close >makeMustache %mark
  ) >makeText; # create text element from mark to start of var

  openBlock = (
    open
    '#'
    expr
    close >makeBlockOpen %mark
  ) >makeText;

  closeBlock = (
    open
    '/'
    expr
    close >makeCloseBlock %mark
  ) >makeText;

  text = (any+ -- open);

  statement = (
    unescapedVar
    | var
    | openBlock
    | closeBlock
    | text
  );

  main := statement* %eof(makeText) $err(error);
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
