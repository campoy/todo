package todo

import "fmt"

type Task struct {
	Title string
	Done  bool
}

func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("empty title")
	}
	return &Task{title, false}, nil
}

type TaskManager struct {
	tasks []*Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (m *TaskManager) Save(task *Task) {
	copy := *task
	m.tasks = append(m.tasks, &copy)
}

func (m *TaskManager) All() []*Task {
	return m.tasks
}
