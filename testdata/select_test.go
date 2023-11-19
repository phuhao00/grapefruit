package testdata

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "channel 1"
	}()

	go func() {
		time.Sleep(6 * time.Second)
		ch2 <- "channel 2"
	}()

	// select 会等待 ch1 或 ch2 任意一个可读取数据时执行
	select {
	case msg1 := <-ch1:
		fmt.Println("Received from channel 1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from channel 2:", msg2)
	case <-time.After(4 * time.Second):
		fmt.Println("Timeout: No channel was ready!")
	}

}

func TestTicker(t *testing.T) {
	tk := time.NewTicker(time.Second * 2).C
	for {
		select {
		case <-tk:
			//todo
			fmt.Println(time.Now().Unix())
		}
	}

}
