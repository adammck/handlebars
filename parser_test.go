package handlebars

import (
	"testing"
	"reflect"
)



func example(t *testing.T, tmpl string, expected *BlockNode) {
	actual := Compile(tmpl)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %#v, expected %#v", actual, expected)
	}
}


func TestParseText(t *testing.T) {
	tmpl     := "blah"
	expected := &BlockNode{"", []Node{NewTextNode("blah")}}
	example(t, tmpl, expected)
}

func TestParseMustache(t *testing.T) {
	tmpl     := "{{hello}}"
	expected := &BlockNode{"", []Node{&MustacheNode{"hello"}}}
	example(t, tmpl, expected)
}

func TestParseMustacheWhitespace(t *testing.T) {
	tmpl     := "{{ hello  }}"
	expected := &BlockNode{"", []Node{&MustacheNode{"hello"}}}
	example(t, tmpl, expected)
}

func TestSimpleParser(t *testing.T) {
	tmpl := "abc{{alpha}}{{beta}}ghi"
	expected := &BlockNode{"", []Node{
		NewTextNode("abc"),
		&MustacheNode{"alpha"},
		&MustacheNode{"beta"},
		NewTextNode("ghi"),
	}}

	example(t, tmpl, expected)
}

func TestSimpleBlockParser(t *testing.T) {
	tmpl := "{{#list}}aaa{{/list}}"
	expected := &BlockNode{"", []Node{
		&BlockNode{"list", []Node{
			NewTextNode("aaa"),
		}},
	}}

	example(t, tmpl, expected)
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

	example(t, tmpl, expected)
}
