// Copyright 2014 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

// The package provides a TaskManager implemented using TDD techniques.
// The tests were developed before the code was written.
package task

import "fmt"

type Task struct {
	ID    int64  // Unique identifier
	Title string // Description
	Done  bool   // Is this task done?
}

// NewTask creates a new task given a title, that can't be empty.
func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("empty title")
	}
	return &Task{0, title, false}, nil
}

// TaskManager manages a list of tasks in memory.
type TaskManager struct {
	tasks  []*Task
	lastID int64
}

// NewTaskManager returns an empty TaskManager.
func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

// Save saves the given Task in the TaskManager.
func (m *TaskManager) Save(task *Task) error {
	if task.ID == 0 {
		m.lastID++
		task.ID = m.lastID
		m.tasks = append(m.tasks, cloneTask(task))
		return nil
	}

	for i, t := range m.tasks {
		if t.ID == task.ID {
			m.tasks[i] = cloneTask(task)
			return nil
		}
	}
	return fmt.Errorf("unknown task")
}

// cloneTask creates and returns a deep copy of the given Task.
func cloneTask(t *Task) *Task {
	c := *t
	return &c
}

// All returns the list of all the Tasks in the TaskManager.
func (m *TaskManager) All() []*Task {
	return m.tasks
}

// Find returns the Task with the given id in the TaskManager and a boolean
// indicating if the id was found.
func (m *TaskManager) Find(ID int64) (*Task, bool) {
	for _, t := range m.tasks {
		if t.ID == ID {
			return t, true
		}
	}
	return nil, false
}
