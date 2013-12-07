package todo

import "fmt"

type Task struct {
	id    int64
	Title string
	Done  bool
}

func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("empty title")
	}
	return &Task{0, title, false}, nil
}

type TaskManager struct {
	tasks  []*Task
	lastID int64
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (m *TaskManager) Save(task *Task) error {
	if task.id == 0 {
		m.lastID++
		task.id = m.lastID
		m.tasks = append(m.tasks, cloneTask(task))
		return nil
	}

	for i, t := range m.tasks {
		if t.id == task.id {
			m.tasks[i] = cloneTask(task)
			return nil
		}
	}
	return fmt.Errorf("unknown task")
}

func cloneTask(t *Task) *Task {
	c := *t
	return &c
}

func (m *TaskManager) All() []*Task {
	return m.tasks
}
