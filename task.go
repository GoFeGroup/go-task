package go_task

import (
	"container/list"
	"time"
)

// Task 任务
type Task struct {
	key        string
	e          *list.Element
	fn         func(*Context) // 任务的回调函数
	interval   int            // 间隔
	createTime time.Time      // 任务创建时间
	nextTime   time.Time      // 任务启动时间
	ctx        *Context
}

// Can 任务是否达到了启动的条件
func (t *Task) Can() bool {
	return time.Now().After(t.nextTime)
}

func (t *Task) Run() {
	t.fn(t.ctx)
	t.nextTime = time.Now().Add(time.Second * time.Duration(t.interval))
}

func (t *Task) NextTick() int64 {
	return int64(t.nextTime.Sub(time.Now()) / time.Second)
}

func NewTask(key string, fn func(*Context), interval int) *Task {
	t := &Task{
		key:        key,
		fn:         fn,
		interval:   interval,
		createTime: time.Now(),
		nextTime:   time.Now().Add(time.Second * time.Duration(interval)),
		ctx:        newContext(),
	}
	return t
}
