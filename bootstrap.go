package golight

import (
	"time"
)

func Bootstrap() {
	On(Red)
	time.Sleep(300 * time.Millisecond)
	Off(Red)
	On(Yellow)
	time.Sleep(300 * time.Millisecond)
	Off(Yellow)
	On(Green)
	time.Sleep(300 * time.Millisecond)
	Off(Green)
}
