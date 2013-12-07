package todo

import "testing"

func TestNewTask(t *testing.T) {
	title := "learn Go"
	task := NewTask(title)
	if task.Title != title {
		t.Errorf("expected title %q, got %q", title, task.Title)
	}
}
