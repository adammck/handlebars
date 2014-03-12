package handlebars

import (
  "testing"
  "github.com/stretchr/testify/assert"
)


func TestTextNode(t *testing.T) {
  n := NewTextNode("blah")
  assert.Equal(t, `"blah"`, n.String())
}


func TestMustacheNode(t *testing.T) {
  n := NewMustacheNode("var")
  assert.Equal(t, `{{var}}`, n.String())
}


func TestBlockNode(t *testing.T) {
  n := NewBlockNode("if aaa")
  n.Append(NewTextNode("bbb"))
  n.Append(NewTextNode("ccc"))

  assert.Equal(t, `{{#if aaa ["bbb", "ccc"]}}`, n.String())
}
