package task

import (
	"github.com/google/uuid"
)

type status int

type duration int64

const (
	backlog status = iota
	queued
	shpd
)

type Task struct {
	ID   uuid.UUID
	Desc string
	Tit  string
	// 	slug           string
	// tags           []string
	// status         status
	// recurring      bool
	// createdAt      time.Time
	// startOn        time.Time
	// finishBy       time.Time
	// timeEstimation duration
	// missions       []string
}

func (t Task) Title() string       { return t.Tit }
func (t Task) Description() string { return t.Desc }

// func NewTask() *Task {
// 	return &Task{
// 		id: uuid.New(),
//     createdAt: time.Now(),
//     status: backlog,
// 	}
// }

//	func (t *Task) Tags() []string {
//		return t.tags
//	}
//
//	func (t *Task) Recurring() bool {
//		return t.recurring
//	}
//
//	func (t *Task) StartOn() time.Time {
//		return t.startOn
//	}
//
//	func (t *Task) FinishBy() time.Time {
//		return t.finishBy
//	}
//
//	func (t *Task) TimeEstimation() duration {
//		return t.timeEstimation
//	}
//
//	func (t *Task) Missions() []string {
//		return t.missions
//	}
func (t Task) FilterValue() string {
	return t.Tit
}
