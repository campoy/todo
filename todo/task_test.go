package todo

import "testing"

func newTaskOrFatal(t *testing.T, title string) *Task {
	task, err := NewTask("learn Go")
	if err != nil {
		t.Fatalf("new task: %v", err)
	}
	return task
}

func TestNewTask(t *testing.T) {
	title := "learn Go"
	task := newTaskOrFatal(t, title)
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
	task := newTaskOrFatal(t, "learn Go")

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

func TestSaveAndRetrieveTwoTasks(t *testing.T) {
	learnGo := newTaskOrFatal(t, "learn Go")
	learnTDD := newTaskOrFatal(t, "learn TDD")

	m := NewTaskManager()
	m.Save(learnGo)
	m.Save(learnTDD)

	all := m.All()
	if len(all) != 2 {
		t.Errorf("expected 2 tasks, got %v", len(all))
	}
	if *all[0] != *learnGo && *all[1] != *learnGo {
		t.Errorf("missing task: %v", learnGo)
	}
	if *all[0] != *learnTDD && *all[2] != *learnTDD {
		t.Errorf("missing task: %v", learnTDD)
	}
}

func TestSaveModifyAndRetrieve(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")
	m := NewTaskManager()
	m.Save(task)

	task.Done = true

	all := m.All()
	if all[0].Done {
		t.Errorf("saved task wasn't done")
	}
}
