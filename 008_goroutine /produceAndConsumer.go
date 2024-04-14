package main

import(
	"fmt"
	"sync"
)


func producer(ch chan int, wg *sync.WaitGroup)  {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- i
	}	
	close(ch)
}


func counsumer(ch chan int, wg *sync.WaitGroup)  {
	defer wg.Done()
	// for v := range ch {
	// 	fmt.Println("recived: ", v)
	// }
	// close(ch)
	//当通道已经被关闭时，range 语句会自动结束循环，并且会将通道中剩余的值都读取完毕。
	//因此，消费者 goroutine 会在通道被关闭后继续尝试从已关闭的通道中读取数据，这时会触发 panic

    for {
        num, ok := <-ch // 从 Channel 中接收数据
        if !ok {
            // 通道已经关闭，退出循环
            fmt.Println("Channel closed. Exiting consumer.")
            return
        }
        fmt.Println("Received:", num)
    }
}

func main()  {
	ch := make(chan int)
	var wg  sync.WaitGroup

	wg.Add(1)
	go producer(ch, &wg)

	wg.Add(1)
	go counsumer(ch, &wg)

	wg.Wait()

}