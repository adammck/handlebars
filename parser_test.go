package handlebars

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseText(t *testing.T) {
	tmpl := "blah"
	expected := &BlockNode{"", []Node{NewTextNode("blah")}}
	assert.Equal(t, Compile(tmpl), expected)
}

func TestParseMustache(t *testing.T) {
	tmpl := "{{hello}}"
	expected := &BlockNode{"", []Node{NewMustacheNode("hello", true)}}
	assert.Equal(t, Compile(tmpl), expected)
}

func TestParseMustacheWhitespace(t *testing.T) {
	tmpl := "{{ hello  }}"
	expected := &BlockNode{"", []Node{NewMustacheNode("hello", true)}}
	assert.Equal(t, Compile(tmpl), expected)
}

func TestParseMustacheUnescaped(t *testing.T) {
	n := NewMustacheNode("omg", false)

	tmpl := "{{{ omg }}}"
	expected := &BlockNode{"", []Node{n}}
	assert.Equal(t, Compile(tmpl), expected)
}

func TestSimpleParser(t *testing.T) {
	tmpl := "abc{{alpha}}{{beta}}ghi"
	expected := &BlockNode{"", []Node{
		NewTextNode("abc"),
		NewMustacheNode("alpha", true),
		NewMustacheNode("beta", true),
		NewTextNode("ghi"),
	}}

	assert.Equal(t, Compile(tmpl), expected)
}

func TestSimpleBlockParser(t *testing.T) {
	tmpl := "{{#list}}aaa{{/list}}"
	expected := &BlockNode{"", []Node{
		&BlockNode{"list", []Node{
			NewTextNode("aaa"),
		}},
	}}

	assert.Equal(t, Compile(tmpl), expected)
}

func TestNestedBlockParser(t *testing.T) {
	tmpl := "aaa{{#alpha}}bbb{{#beta}}ccc{{/beta}}ddd{{/alpha}}eee"
	expected := &BlockNode{"", []Node{
		NewTextNode("aaa"),
		&BlockNode{"alpha", []Node{
			NewTextNode("bbb"),
			&BlockNode{"beta", []Node{
				NewTextNode("ccc"),
			}},
			NewTextNode("ddd"),
		}},
		NewTextNode("eee"),
	}}

	assert.Equal(t, Compile(tmpl), expected)
}
