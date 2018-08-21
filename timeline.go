package grf

import (
	"time"
)

type TimeLine struct {
	tsk_num  uint16
	duration uint32
	repeat   bool
	t_slot   []*Task
}

func (tl *TimeLine) CreateRepeatTimeline(duration uint32) *TimeLine {
	tl.duration = duration
	tl.repeat = true
	return tl
}

func (tl *TimeLine) AddTask(tsk *Task) *TimeLine {
	tl.t_slot = append(tl.t_slot, tsk)
	tl.tsk_num = tl.tsk_num + 1
	return tl
}

func (tl *TimeLine) runAllTask() {
	for i := 0; i < int(tl.tsk_num); i++ {
		tl.t_slot[i].fn()
	}
}

func (tl *TimeLine) Start() {
	if tl.repeat {
		ticker := time.NewTicker(time.Duration(tl.duration) * time.Millisecond)
		quit := make(chan struct{})
		for {
			select {
			case <-ticker.C:
				tl.runAllTask()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	} else {
		tl.runAllTask()
	}
}
