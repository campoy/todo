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

// +build appengine

// A Google App Engine application providing a web UI for task management.
// DISCLAIMER: This is not intended to be used, since instances can be
// restarted at any time and all tasks would be lost.
package todo

import "github.com/campoy/todo/server"

func init() {
	server.RegisterHandlers()
}
