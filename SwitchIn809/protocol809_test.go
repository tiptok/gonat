package SwitchIn809

import (
	"fmt"
	"testing"
	"time"
)

func TestParseMsg(t *testing.T) {
	//server809 := &SwitchIn809.Tcp809Server{}
	// c :=conn.NewConn()

	// protocol :=protocol809{}
	// protocol.ParseMsg()

	LastReadTime, _ := time.Parse("2006-01-02 15:04:05", "2000-01-01 00:00:00")
	fmt.Println(LastReadTime)
}
