package channelbase

import (
	"fmt"
	"testing"
	"time"
)

func Test_Chan_Select(t *testing.T) {

	// 准备好几个通道。
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道，并向它发送元素值。
	//index := rand.Intn(3)
	//fmt.Printf("The index: %d\n", index)
	//intChannels[index] <- index
	//
	//intChannels[1] <- 1
	//intChannels[0] <- 0

	go func() {
		time.AfterFunc(time.Duration(1)*time.Second, func() {
			intChannels[1] <- 1
		})
	}()

	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	// 重上到下一次开始执行，
	// 阻塞则求值失败
	// 没有合适的，会走default, 没有default，则select会被阻塞, 如果有合适的分支，则会被唤醒
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
		//default:
		//	fmt.Println("No candidate case is selected!")
	}

}

func Test_Chan_Select2(t *testing.T) {

}
