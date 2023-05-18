package go_task

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"testing"
	"time"
)

func TestManger_001(t *testing.T) {
	m := NewManager()

	fmt.Println("before add ")
	for i := 0; i < 10; i++ {
		x := i
		err := m.Add(NewTask(fmt.Sprintf("abcf-%v", i), func() {
			fmt.Printf("%d func %d\n", time.Now().Unix(), x)
		}, 3+i))
		Throw(err)
	}
	fmt.Println("after add ")

	go func() {
		time.Sleep(time.Second * 15)
		Throw(m.Delete("abcf-0"))
	}()

	// dump
	for item := m.readyQueue.Front(); item != nil; item = item.Next() {
		task := item.Value.(*Task)
		fmt.Printf("%d, %v --> %d\n", time.Now().Unix(), task.key, task.NextTick())
	}
	fmt.Println("after dump ")

	// waiting to be interrupted
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func Throw(e any) {
	if e != nil {
		fmt.Printf("===> %s\n", e)
		fmt.Printf("===> %s\n", debug.Stack())
		os.Exit(-1)
	}
}
