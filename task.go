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
	createTime int64          // 任务创建时间
	nextTime   int64          // 任务启动时间
	ctx        *Context
}

// Can 任务是否达到了启动的条件
func (t *Task) Can() bool {
	return time.Now().Unix() > t.nextTime
}

func (t *Task) Run() {
	t.fn(t.ctx)
	t.nextTime = time.Now().Unix() + int64(t.interval)
}

func (t *Task) NextTick() int64 {
	return t.nextTime - time.Now().Unix()
}

func NewTask(key string, fn func(*Context), interval int) *Task {
	return NewTaskWithContext(key, fn, interval, NewContext())
}

func NewTaskWithContext(key string, fn func(*Context), interval int, ctx *Context) *Task {
	t := &Task{
		key:        key,
		fn:         fn,
		interval:   interval,
		createTime: time.Now().Unix(),
		nextTime:   time.Now().Unix() + int64(interval),
		ctx:        ctx,
	}
	return t
}
