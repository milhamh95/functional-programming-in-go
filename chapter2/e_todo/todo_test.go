package e_todo

import "testing"

func TestTodo_Write(t *testing.T) {
	todo := Todo{
		Db: &Db{
			AuthorizationF: func() bool {
				return true
			},
		},
	}

	todo.Write("hello")
	if todo.Text != "hello" {
		t.Errorf("expected 'hello' but got %v\n", todo.Text)
	}

	todo.Append(" world")
	if todo.Text != "hello world" {
		t.Errorf("expected 'hello world' but got %v\n", todo.Text)
	}
}
