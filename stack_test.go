package handlebars

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	if stack.Len() != 0 {
		t.Errorf("expected new stack to be empty")
	}

	one := stack.Pop()
	if one != nil {
		t.Errorf("expected Pop on empty stack to return nil")
	}

	stack.Push(NewBlockNode("aaa"))
	if stack.Len() != 1 {
		t.Errorf("expected stack have one element")
	}

	stack.Push(NewBlockNode("bbb"))
	if stack.Len() != 2 {
		t.Errorf("expected stack have two elements")
	}

	two := stack.Pop()
	if two.expr != "bbb" {
		t.Errorf("expected Pop to return last-pushed element")
	}
}
