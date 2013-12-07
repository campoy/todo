package todo

import "testing"

func TestNewTask(t *testing.T) {
	title := "learn Go"
	task, err := NewTask(title)
	if err != nil {
		t.Errorf("new task: %v", err)
	}
	if task.Title != title {
		t.Errorf("expected title %q, got %q", title, task.Title)
	}
	if task.Done {
		t.Errorf("new task is done")
	}
}

func TestNewTaskEmptyTitle(t *testing.T) {
	_, err := NewTask("")
	if err == nil {
		t.Errorf("task with empty title created")
	}
}

func TestSaveTaskAndRetrieve(t *testing.T) {
	task, err := NewTask("learn Go")
	if err != nil {
		t.Errorf("new task: %v", err)
	}

	m := NewTaskManager()
	m.Save(task)

	all := m.All()
	if len(all) != 1 {
		t.Errorf("expected 1 task, got %v", len(all))
	}
	if *all[0] != *task {
		t.Errorf("expected %v, got %v", task, all)
	}
}
