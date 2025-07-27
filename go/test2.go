package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	a := 1
	b := &a
	df(b)
	fmt.Println(a)
	// 考察点 ：指针的使用、值传递与引用传递的区别。

	// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	c := []int{1, 2, 3}
	d := &c
	df1(d)
	fmt.Println(c)
	// 考察点 ：指针运算、切片操作。

	//编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	//考察点 ： go 关键字的使用、协程的并发执行。
	var wg sync.WaitGroup
	wg.Add(2) // 等待两个goroutine完成

	// 打印奇数的goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Printf("奇数: %d\n", i)
		}
	}()

	// 打印偶数的goroutine
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Printf("偶数: %d\n", i)
		}
	}()

	wg.Wait() // 等待所有goroutine完成
	fmt.Println("打印完成")

	//编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	// 创建一个无缓冲的整数通道
	ch := make(chan int)

	// 启动生产者协程
	go producer(ch)

	// 启动消费者协程
	go consumer(ch)

	// 等待用户输入以防止主程序退出
	fmt.Println("按回车键退出...")

	// 创建一个缓冲大小为10的整数通道
	ch1 := make(chan int, 10)

	wg.Add(2) // 等待生产者和消费者两个协程完成

	// 启动生产者协程
	go func() {
		defer wg.Done()
		defer close(ch1) // 发送完成后关闭通道

		for i := 1; i <= 100; i++ {
			ch1 <- i // 发送数据到通道
			fmt.Printf("生产者发送: %d\n", i)
		}
		fmt.Println("生产者完成")
	}()

	// 启动消费者协程
	go func() {
		defer wg.Done()

		for num := range ch1 { // 循环接收直到通道关闭
			fmt.Printf("消费者接收: %d\n", num)
		}
		fmt.Println("消费者完成")
	}()

	wg.Wait() // 等待所有协程完成
	fmt.Println("程序结束")

	var counter int64 // 使用int64类型的计数器

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程递增计数器1000次
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1) // 原子递增操作
			}
		}()
	}

	wg.Wait() // 等待所有协程完成
	fmt.Printf("最终计数器值: %d\n", counter)
}
func df(b *int) {
	*b = 10
}
func df1(b *[]int) {
	e := *b
	for i := range e {
		e[i] *= 2
	}
}

// 生产者函数：生成1到10的整数并发送到通道
func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("生产者发送: %d\n", i)
		ch <- i // 发送数据到通道
	}
	close(ch) // 发送完成后关闭通道
}

// 消费者函数：从通道接收整数并打印
func consumer(ch <-chan int) {
	for num := range ch { // 循环接收直到通道关闭
		fmt.Printf("消费者接收: %d\n", num)
	}
}
