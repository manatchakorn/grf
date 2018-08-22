package grf

import (
	"time"
)

type timeline struct {
	tsk_num  uint16
	duration uint32
	repeat   bool
	t_slot   []*Task
}

type bgtimeline struct {
	tsk_num  uint16
	duration uint32
	repeat   bool
	t_slot   []*Task
}

var Timeline *timeline
var BgTimeline *bgtimeline

func CreateTimeline() *timeline {
	if Timeline == nil {
		Timeline = &timeline{}
	}
	return Timeline
}

func (tl *timeline) CreateRepeatTimeline(duration uint32) *timeline {
	tl.duration = duration
	tl.repeat = true
	return tl
}

func (tl *timeline) AddTask(tsk *Task) *timeline {
	tl.t_slot = append(tl.t_slot, tsk)
	tl.tsk_num = tl.tsk_num + 1
	return tl
}

func (tl *timeline) runAllTask() {
	for i := 0; i < int(tl.tsk_num); i++ {
		tl.t_slot[i].fn()
	}
}

func (tl *timeline) Start() {
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

func CreateBackgroundTimeline() *bgtimeline {
	if BgTimeline == nil {
		BgTimeline = &bgtimeline{}
	}
	return BgTimeline
}

func (tl *bgtimeline) CreateRepeatTimeline(duration uint32) *bgtimeline {
	tl.duration = duration
	tl.repeat = true
	return tl
}

func (tl *bgtimeline) AddTask(tsk *Task) *bgtimeline {
	tl.t_slot = append(tl.t_slot, tsk)
	tl.tsk_num = tl.tsk_num + 1
	return tl
}

func (tl *bgtimeline) runAllTask() {
	for i := 0; i < int(tl.tsk_num); i++ {
		go tl.t_slot[i].fn()
	}
}

func (tl *bgtimeline) Start() {
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
