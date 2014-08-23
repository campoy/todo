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

package task

import "testing"

func newTaskOrFatal(t *testing.T, title string) *Task {
	task, err := NewTask(title)
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
		t.Errorf("expected 'empty title' error, got nil")
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
		t.Errorf("expected %v, got %v", task, all[0])
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
	if *all[0] != *learnTDD && *all[1] != *learnTDD {
		t.Errorf("missing task: %v", learnTDD)
	}
}

func TestSaveModifyAndRetrieve(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")
	m := NewTaskManager()
	m.Save(task)

	task.Done = true
	if m.All()[0].Done {
		t.Errorf("saved task wasn't done")
	}
}

func TestSaveTwiceAndRetrieve(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")
	m := NewTaskManager()
	m.Save(task)
	m.Save(task)

	all := m.All()
	if len(all) != 1 {
		t.Errorf("expected 1 task, got %v", len(all))
	}
	if *all[0] != *task {
		t.Errorf("expected task %v, got %v", task, all[0])
	}
}

func TestSaveAndFind(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")
	m := NewTaskManager()
	m.Save(task)

	nt, ok := m.Find(task.ID)
	if !ok {
		t.Errorf("didn't find task")
	}
	if *task != *nt {
		t.Errorf("expected %v, got %v", task, nt)
	}
}
