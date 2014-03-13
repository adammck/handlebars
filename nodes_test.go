package handlebars

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextNode(t *testing.T) {
	n := NewTextNode("blah")
	assert.Equal(t, `"blah"`, n.String())
}

func TestTextNodeExecute(t *testing.T) {
	ctx := map[string]string{"a": "AAA"}
	assert.Equal(t, `a`, NewTextNode("a").Execute(ctx))
}

func TestMustacheNodeString(t *testing.T) {
	n := NewMustacheNode("var")
	assert.Equal(t, `{{var}}`, n.String())
}

func TestMustacheNodeExecute(t *testing.T) {
	ctx := map[string]string{"a": "aaa", "b": "bbb"}
	assert.Equal(t, `aaa`, NewMustacheNode("a").Execute(ctx))
	assert.Equal(t, `bbb`, NewMustacheNode("b").Execute(ctx))
	assert.Equal(t, ``, NewMustacheNode("c").Execute(ctx))
}

func TestBlockNode(t *testing.T) {
	n := NewBlockNode("if aaa")
	n.Append(NewTextNode("bbb"))
	n.Append(NewTextNode("ccc"))

	assert.Equal(t, `{{#if aaa ["bbb", "ccc"]}}`, n.String())
}

func TestBlockNodeExecute(t *testing.T) {
	n := NewBlockNode("")
	n.Append(NewTextNode("aaa "))
	n.Append(NewMustacheNode("x"))
	n.Append(NewTextNode(" bbb "))
	n.Append(NewMustacheNode("y"))
	n.Append(NewTextNode(" ccc"))

	ctx := map[string]string{"x": "XXX", "y": "YYY"}
	assert.Equal(t, `aaa XXX bbb YYY ccc`, n.Execute(ctx))
}
