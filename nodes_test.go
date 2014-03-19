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
	ctx := map[string]interface{}{"a": "AAA"}
	assert.Equal(t, `a`, NewTextNode("a").Execute(ctx))
}

func TestMustacheNodeString(t *testing.T) {
	n := NewMustacheNode("var", true)
	assert.Equal(t, `{{var}}`, n.String())
}

func TestMustacheNodeExecute(t *testing.T) {
	ctx := map[string]interface{}{"a": "aaa", "b": "bbb"}
	assert.Equal(t, `aaa`, NewMustacheNode("a", true).Execute(ctx))
	assert.Equal(t, `bbb`, NewMustacheNode("b", true).Execute(ctx))
	assert.Equal(t, ``, NewMustacheNode("c", true).Execute(ctx))
}

func TestMustacheNodeExecuteDot(t *testing.T) {
	node := NewMustacheNode(".", true)
	assert.Equal(t, `dot`, node.Execute("dot"))
	assert.Equal(t, `1`, node.Execute(1))
	assert.Equal(t, `true`, node.Execute(true))
	assert.Equal(t, ``, node.Execute(false))
}

func TestMustacheNodeExecuteNestedPath(t *testing.T) {
	ctx := map[string]interface{}{"c": map[string]interface{}{"d": "EEE"}}
	assert.Equal(t, `EEE`, NewMustacheNode("c.d", true).Execute(ctx))
	assert.Equal(t, ``, NewMustacheNode("c.x", true).Execute(ctx))
	//assert.Equal(t, ``, NewMustacheNode("c.d.e", true).Execute(ctx)) // invalid map type
	//assert.Equal(t, ``, NewMustacheNode("c", true).Execute(ctx)) // invalid value type
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
	n.Append(NewMustacheNode("x", true))
	n.Append(NewTextNode(" bbb "))
	n.Append(NewMustacheNode("y", true))
	n.Append(NewTextNode(" ccc"))

	ctx := map[string]interface{}{"x": "XXX", "y": "YYY"}
	assert.Equal(t, `aaa XXX bbb YYY ccc`, n.Execute(ctx))
}
