/* Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

package framework

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestScheduler(t *testing.T) {

	Convey("Applications with state", t, func() {
		Mesos = new(FakeMesos)
		scheduler := NewScheduler(NewSchedulerConfig())

		application := &Application{
			Type:          "run-once",
			ID:            "foo",
			Cpu:           0.5,
			Mem:           512,
			LaunchCommand: "sleep 10",
		}
		contextChan := make(chan *ApplicationRunStatus)

		Convey("should work properly", func() {
			So(scheduler.applicationsWithState(StateIdle), ShouldHaveLength, 0)

			scheduler.stageApplication(application, contextChan)
			So(scheduler.applicationsWithState(StateIdle), ShouldHaveLength, 1)
			scheduler.applications[application.ID].State = StateRunning
			So(scheduler.applicationsWithState(StateIdle), ShouldHaveLength, 0)
			So(scheduler.applicationsWithState(StateRunning), ShouldHaveLength, 1)
		})

	})

	Convey("Application ID from Task ID", t, func() {
		Convey("should extract proper application ID", func() {
			So(applicationIDFromTaskID("foobar.ip-123-123-123-123.f81d4fae-7dec-11d0-a765-00a0c91e6bf6"), ShouldEqual, "foobar")
			So(func() { applicationIDFromTaskID("foobar-ip-123-123-123-123-f81d4fae-7dec-11d0-a765-00a0c91e6bf6") }, ShouldPanic)
		})
	})

	Convey("Hostname from Task ID", t, func() {
		Convey("should extract proper hostname", func() {
			So(hostnameFromTaskID("foobar.ip-123-123-123-123.f81d4fae-7dec-11d0-a765-00a0c91e6bf6"), ShouldEqual, "ip-123-123-123-123")
			So(func() { hostnameFromTaskID("foobar-ip-123-123-123-123-f81d4fae-7dec-11d0-a765-00a0c91e6bf6") }, ShouldPanic)
		})
	})

}
