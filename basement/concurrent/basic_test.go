package concurrent

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func Test_Goroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Millisecond * 500)
}

func Test_Goroutine2(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(in int) {
			fmt.Println(in)
		}(i)
	}
	time.Sleep(time.Millisecond * 500)
}

func Test_Goroutine3(t *testing.T) {

	num := 10
	ch := make(chan int, num)

	for i := 0; i < 10; i++ {
		go func() {
			cur := <-ch
			fmt.Println(cur)
		}()
	}

	for i := 0; i <= num; i++ {
		ch <- i
	}

	<-ch
}

func Test_Goroutine4(t *testing.T) {

	ch := make(chan int)

	end := make(chan struct{})

	go func() {
		ch <- 0
	}()

	for i := 0; i < 10; i++ {
		go func() {
			cur := <-ch
			fmt.Println(cur)
			cur += 1

			if cur <= 9 {
				ch <- cur
			} else {
				end <- struct{}{}
			}
		}()
	}

	<-end
}

func Test_Use_baise(t *testing.T) {
	var count int32 = 0

	trigger := func(i int32, fn func()) {
		for {
			//if atomic.CompareAndSwapInt32(&count, i, i + 1) {
			//	fn()
			//	break
			//}
			if n := atomic.LoadInt32(&count); n == i {
				fn()
				atomic.AddInt32(&count, 1)
				break
			}
			// 自旋消耗CPU
			time.Sleep(time.Nanosecond)
		}
	}

	for i := int32(0); i < 10; i++ {
		go func(in int32) {
			trigger(in, func() {
				println(in)
			})
		}(i)
	}

	trigger(10, func() {})
}

func Test_Use_Sync_WaitGroup(t *testing.T) {

}
