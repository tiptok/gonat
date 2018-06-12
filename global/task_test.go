package global

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestStartTask(t *testing.T) {
	TaskRec := NewTaskManager("BizDBTdf.Rec.Task", 1000, 1, onReceiveEntity)
	ctx, _ := context.WithCancel(context.Background())
	TaskRec.Start(ctx)
	for i := 0; i <= 100; i++ {
		TaskRec.Enqueue(i)
	}
	time.Sleep(time.Second * 10)
}

func onReceiveEntity(val interface{}) {
	log.Printf("Rec %v", val)
}
