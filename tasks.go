package grf

type Task struct {
	id             uint32
	fn             func()
	last_exec_time float32
	avg_exec_time  float32
}

var id uint32

func CreateTask(fn func()) *Task {
	tsk := Task{}
	tsk.id = tsk.id + 1
	tsk.fn = fn
	return &tsk
}
