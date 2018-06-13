package global

import (
	"context"
	"sync"
)

//TaskManager  任务管理
type TaskManager struct {
	TaskName   string
	Queue      chan interface{}
	l          sync.RWMutex
	maxsize    int
	OnExecTask func(interface{})
	Cancel     context.CancelFunc
	StopFlag   int //1开启 0关闭
	GSize      int //开启Goroutine 数量
	CloseOnce  sync.Once
}

//NewTaskManager 新建一个TaskManager
//队列名称 name
//缓存长度 bufsize
//工作数量 gSize
//工作函数 onExecTask
func NewTaskManager(name string, bufSize int, gSize int, onExecTask func(interface{})) *TaskManager {
	return &TaskManager{
		maxsize:    bufSize,
		TaskName:   name,
		Queue:      make(chan interface{}, bufSize),
		OnExecTask: onExecTask,
		StopFlag:   0,
		GSize:      gSize,
		//CloseOnce:       make(sync.Once),
	}
}

//Start 启动一个TaskTaskManager
func (task *TaskManager) Start(ctx context.Context) {
	task.StopFlag = 1
	for i := 1; i <= task.GSize; i++ {
		task.l.Lock()
		Info(F(BUS, "TaskManager", "启动 TaskManager -%v %v"), task.TaskName, i)
		go task.ExecTask(ctx)
		task.l.Unlock()
	}
}

//Count  包含队列长度
func (task *TaskManager) Count() int {
	task.l.RLock()
	defer task.l.RUnlock()
	return len(task.Queue)
}

//Enqueue 入队
func (task *TaskManager) Enqueue(p interface{}) {
	if task.StopFlag == 0 {
		return
	}
	task.Queue <- p
}

//Dequeue 出队
// func (task *TaskManager) Dequeue() interface{} {
// 	return <-task.Queue
// }

//ExecTask 执行任务
func (task *TaskManager) ExecTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			task.Stop()
			return
		case job := <-task.Queue: //出队
			if task.OnExecTask != nil {
				task.OnExecTask(job) //执行job
			}
		}
	}
}

//Stop  停止任务管理
func (task *TaskManager) Stop() {
	task.l.Lock()
	defer task.l.Unlock()
	task.CloseOnce.Do(func() {
		close(task.Queue)
		task.StopFlag = 0
	})
	//task.Cancel()

}
