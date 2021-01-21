package channelbase

import (
	"fmt"
	"testing"
	"time"
)

func Test_channel(t *testing.T) {

	// FIFO
	var c = make(chan string, 2)

	go func() {
		msg := "Hello Go Channel"
		time.Sleep(time.Duration(2) * time.Second)
		c <- msg
		c <- msg + "1"
	}()

	start := time.Now()
	m := <-c

	t.Logf("get msg from chan: %s, cost: %d s", m, time.Now().Second()-start.Second())
}

func Test_FIFO_Chan(t *testing.T) {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3

	// 第四个值将会被阻塞
	// fatal error: all goroutines are asleep - deadlock!
	//c <- 4

	first := <-c

	t.Logf("first element: %d", first)
}

func Test_Chane_Block(t *testing.T) {
	// 示例1。
	ch1 := make(chan int, 1)
	ch1 <- 1
	//ch1 <- 2 // 通道已满，因此这里会造成阻塞。

	// 示例2。
	ch2 := make(chan int, 1)
	//elem, ok := <-ch2 // 通道已空，因此这里会造成阻塞。
	//_, _ = elem, ok
	ch2 <- 1

	// 示例3。
	var ch3 chan int
	//ch3 <- 1 // 通道的值为nil，因此这里会造成永久的阻塞！
	//<-ch3 // 通道的值为nil，因此这里会造成永久的阻塞！
	_ = ch3
}

func Test_PC(t *testing.T) {
	ch1 := make(chan int, 2)
	// 发送方。
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender: sending element %v...\n", i)
			ch1 <- i
		}
		fmt.Println("Sender: close the channel...")
		close(ch1)
	}()

	// 接收方。
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n", elem)
	}

	fmt.Println("End.")

}

func Test_OutPutData(t *testing.T) {

	odd := make(chan int)
	even := make(chan int)

	end := make(chan bool)

	// Producer
	go func() {
		for i := 0; i < 10; i++ {
			if i&1 > 0 {
				odd <- i
			} else {
				even <- i
			}
		}
		close(odd)
		close(even)

		end <- true
	}()

	output := func(key string, c chan int) {
		for {
			ele, ok := <-c
			if !ok {
				fmt.Println("Receiver: closed channel")
				break
			}
			fmt.Printf("%s: %d\n", key, ele)
		}
	}

	// println
	go output("odd", odd)
	go output("even", even)

	<-end
	fmt.Println("End")
}

func Test_OutPutData2(t *testing.T) {

	odd := make(chan int)
	even := make(chan int)

	end := make(chan bool)

	output := func(key string, rec chan int, send chan int) {
		for {
			ele, _ := <-rec
			fmt.Printf("%s: %d\n", key, ele)
			if ele == 10 {
				end <- true
				break
			}

			ele += 1
			send <- ele

			if ele == 10 {
				break
			}
		}
	}

	// println
	go output("odd", odd, even)
	go output("even", even, odd)

	odd <- 0

	<-end
	fmt.Println("End")
}
