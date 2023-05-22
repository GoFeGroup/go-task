package go_task

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type queue struct {
	list.List
}

// Add 按照执行顺序添加
func (q *queue) Add(task *Task) *list.Element {
	next := task.NextTick()
	for item := q.Back(); item != nil; item = item.Prev() {
		t := item.Value.(*Task)
		if next >= t.NextTick() {
			return q.InsertAfter(task, item)
		}
	}
	return q.PushFront(task)
}

type Manager struct {
	items      map[string]*Task
	readyQueue queue // 准备队列
	lock       sync.RWMutex
}

func (m *Manager) runTask(task *Task) {
	task.Run()

	m.lock.Lock()
	if _, ok := m.items[task.key]; ok { // 如果没有在items里删除，则继续进行任务
		task.e = m.readyQueue.Add(task)
	}
	m.lock.Unlock()
}

func (m *Manager) scheduleOnce() {
	for m.readyQueue.Len() > 0 {
		item := m.readyQueue.Front()
		task := item.Value.(*Task)
		if !task.Can() {
			break
		}
		m.readyQueue.Remove(item)
		go m.runTask(task)
	}
}

func (m *Manager) schedule() {
	for {
		time.Sleep(time.Second) // 每秒执行一次
		m.lock.Lock()
		m.scheduleOnce()
		m.lock.Unlock()
	}
}

func (m *Manager) Add(task *Task) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if _, ok := m.items[task.key]; ok {
		return fmt.Errorf("already existed: %v", task.key)
	}
	m.items[task.key] = task
	go m.runTask(task)
	return nil
}

func (m *Manager) Delete(key string) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	if task, ok := m.items[key]; !ok {
		return fmt.Errorf("not existed key: %v", key)
	} else {
		task.ctx.setCanceled()
		m.readyQueue.Remove(task.e)
		delete(m.items, key)
	}
	return nil
}

func NewManager() *Manager {
	mgr := &Manager{
		items: make(map[string]*Task, 0),
	}
	go mgr.schedule()
	return mgr
}
