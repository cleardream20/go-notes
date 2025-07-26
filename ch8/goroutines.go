package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	// fibTest()
	// sumChannel()
	// pipelineEx()
	// bufferedCh()
	// waitEx()
	// errHandling()
	// rocketLaunch()
	multiGorExit()
}

func fibTest() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibnacci(%d) = %d\n", n, fibN)
}

// 整一个小动画，省得算的无聊
// range `-\|/`？但是它会转哎😋🥳
func spinner(delay time.Duration) {
	for {
		for _,  r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func sumChannel() {
	s := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int) // 创建channel
	go sum(s[:len(s)/2], c) // gorou1
	go sum(s[len(s)/2:], c) // gorou2

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func pipelineEx() {
    naturals := make(chan int)
    squares := make(chan int)

    // Counter
    go func() {
        for x := 0; x < 10; x++ {
            naturals <- x
        }
        close(naturals)
    }()

    // Squarer
    go func() {
        // chan int 也可以for range?!
        for x := range naturals {
            squares <- x * x
        }
        close(squares)
    }()

    // Printer
    for x := range squares {
        fmt.Println(x)
    }
}

func bufferedCh() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // 由发送方关闭channel
	// <- 1 2 3 <-
	for x := range ch {
		fmt.Println(x) // 1 2 3
	}
}

func waitEx() {
    values := []int{1, 2, 3, 4, 5}
    
    // 创建等待组
    var wg sync.WaitGroup
    
    for _, v := range values {
        wg.Add(1) // 为每个goroutine增加计数
        
        go func(val int) {
            defer wg.Done() // 完成后减少计数
            
            // 处理val
            fmt.Println(val * val)
        }(v) // 注意这里传递v的副本
    }
    
    wg.Wait() // 等待所有goroutine完成
}

func process(val int) (int, error) {
    if val == 0 {
        return 0, errors.New("invalid value")
    }
    return val * val, nil
}

func errHandling() {
    values := []int{1, 2, 3, 0, 5}
    
    var wg sync.WaitGroup
    errChan := make(chan error, len(values))
    resultChan := make(chan int, len(values))
    
    for _, v := range values {
        wg.Add(1)
        
        go func(val int) {
            defer wg.Done()
            
            res, err := process(val)
            if err != nil {
                errChan <- err
                return
            }
            resultChan <- res
        }(v)
    }
    
    // 等待所有goroutine完成
    go func() {
        wg.Wait()
        close(resultChan)
        close(errChan)
    }()
    
    // 处理结果
    for res := range resultChan {
        fmt.Println("Result:", res)
    }
    
    // 处理错误
    for err := range errChan {
        fmt.Println("Error:", err)
    }
}

func rocketLaunch() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.  Press return to abort.")
    select {
    case <-time.After(5 * time.Second):
        // Do nothing.
    case <-abort:
        fmt.Println("Launch aborted!")
        return
    }

	fmt.Println("起飞！！")
}

func multiGorExit() {
    var wg sync.WaitGroup       // 1. 创建WaitGroup用于等待goroutine结束
    stopCh := make(chan struct{}) // 2. 创建退出信号通道
    
    // 3. 启动3个worker goroutine
    for i := 0; i < 3; i++ {
        wg.Add(1)  // 增加WaitGroup计数器
        go func(id int) {
            defer wg.Done()  // goroutine结束时减少计数器
            
            // 4. 阻塞等待退出信号
            <-stopCh
            fmt.Printf("Worker %d exiting\n", id)
        }(i)  // 注意这里传递i的副本
    }
    
    // 5. 主goroutine等待2秒
    time.Sleep(2 * time.Second)
    
    // 6. 关闭通道通知所有worker退出
    close(stopCh)
    
    // 7. 等待所有worker完成
    wg.Wait()
}