package SwitchIn809

import (
	"fmt"
	"testing"

	"github.com/tiptok/gonat/model"
)

func TestJ1203(t *testing.T) {
	pos := &model.UP_EXG_MSG_HISTORY_LOCATION{}
	fmt.Println(&pos)
	TransData(*pos)
}

func TransData(location model.UP_EXG_MSG_HISTORY_LOCATION) {
	fmt.Println(location)
}
