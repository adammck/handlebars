package handlebars

import (
  "testing"
  "github.com/stretchr/testify/assert"
)


func TestTextNode(t *testing.T) {
  n := NewTextNode("blah")
  assert.Equal(t, `"blah"`, n.String())
}


func TestMustacheNodeString(t *testing.T) {
  n := NewMustacheNode("var")
  assert.Equal(t, `{{var}}`, n.String())
}


func TestMustacheNodeExecute(t *testing.T) {
  ctx := map[string]string{"a": "aaa", "b": "bbb"}
  assert.Equal(t, `aaa`, NewMustacheNode("a").Execute(ctx))
  assert.Equal(t, `bbb`, NewMustacheNode("b").Execute(ctx))
  assert.Equal(t, ``,    NewMustacheNode("c").Execute(ctx))
}


func TestBlockNode(t *testing.T) {
  n := NewBlockNode("if aaa")
  n.Append(NewTextNode("bbb"))
  n.Append(NewTextNode("ccc"))

  assert.Equal(t, `{{#if aaa ["bbb", "ccc"]}}`, n.String())
}
