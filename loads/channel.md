## channel

单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

虽然可以使用共享内存进行数据交换，但是共享内存在不同的`goroutine`中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

Go语言的并发模型是`CSP（Communicating Sequential Processes）`，提倡**通过通信共享内存**而不是**通过共享内存而实现通信**。

如果说`goroutine`是Go程序并发的执行体，`channel`就是它们之间的连接。`channel`是可以让一个`goroutine`发送特定值到另一个`goroutine`的通信机制。

Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

## 定义一个channel变量

```go
var 变量名 chan 元素类型
```

```go
	var ch1 chan int
	var ch2 chan string
```

### 创建channel

通道是引用类型，通道类型的空值是`nil`

```go
var ch chan int
fmt.Pringtln(ch) // nil
```

声明的通道后需要使用`make`函数初始化之后才能使用。



## 无缓冲通道和缓冲通道

### 无缓冲通道

无缓冲的通道又称为阻塞的通道。我们来看一下下面的代码：

```go
package main

import "fmt"

func recv(ch chan bool){
	ret:=<-ch
	fmt.Println(ret)
}


func main() {
	var ch = make(chan bool) // 无缓冲通道
	go recv(ch) // 起一个goroutine，等待接收值 阻塞
	ch <- true
	fmt.Println("main函数结束")
}
```

无缓冲通道上的发送操作会阻塞，知道另一个`goroutine`在该通道上执行接收操作，这时值才能发送成功，两个`goroutine`将继续执行。相反，如果接收操作先执行，接收方的`goroutine`将阻塞，直到另一个`goroutine`在该通道上发送一个值。

使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为`同步通道`。

### 缓冲通道

解决上面问题的方法还有一种就是使用有缓冲区的通道。我们可以在使用make函数初始化通道的时候为其指定通道的容量，例如：

```go
package main

import "fmt"

func recv(ch chan bool){
	ret:=<-ch
	fmt.Println(ret)
}


func main() {
	var ch = make(chan bool,1) // 有缓冲通道
	ch <- true
	// len :数据量 cap：容量
	fmt.Println(len(ch),cap(ch))

	go recv(ch)
	ch <- false
	fmt.Println(len(ch),cap(ch))
	fmt.Println("main函数结束")
}
```

只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。

我们可以使用内置的`len`函数获取通道内元素的数量，使用`cap`函数获取通道的容量，虽然我们很少会这么做。



### 取值时通道判断是否关闭

使用value,ok:=<-ch1的方式，ok=false，通道关闭

```go
package main

import "fmt"

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	var ch1 = make(chan int, 100)
	go send(ch1)
	// 利用for循环,取通道ch1中接收值
	for {
		ret,ok:=<-ch1 // 使用 value，ok :<-ch1 取值的方式，查看当前通道是否关闭 ok-false

		if !ok{
			break
		}
		fmt.Println(ret)
	}


}
```

使用for range的方式

```go
package main

import "fmt"

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	var ch1 = make(chan int, 100)
	go send(ch1)

	for ret:= range ch1{
		fmt.Println(ret)
	}

}
```



练习题

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生产者消费者模型
// 使用goroutine 和 channel实现一个简单的生产者消费模型
// 生产者：产生随机数
// 消费者：计算每个随机数的每个数字的数字的和   123 = 6
// 一个生产者 20个消费者

type Item struct {
	id  int64
	num int64
}

type Result struct {
	item *Item
	sum int64
}

func producer(ch chan *Item){
	var id int64
	for {
		id ++
		num:=rand.Int63()
		tmp := &Item{
			id:  id,
			num: num,
		}
		ch <-tmp
		fmt.Println(tmp)
	}

}

func calc(num int64)int64{
	var sum int64
	for num>0{
		sum = sum+num%10
		num = num/10
	}
	return sum
}

func startWorker(n int,itemChan chan *Item,resultChan chan *Result){
	for i:=0;i<n;i++{
		go consumer(itemChan,resultChan)
	}
}

func consumer(itemChan chan *Item,resultChan chan *Result){
	for item := range itemChan{
		sum := calc(item.num)
		ret := &Result{
			item: item,
			sum:  sum,
		}
		resultChan <- ret
		time.Sleep(time.Microsecond*5)
	}
}

func printResult(resultChan chan *Result){
	for ret := range resultChan{
		fmt.Printf("id:%v num:%v,sum:%v\n",ret.item.id,ret.item.num,ret.sum)
	}
}

func main() {
	var itemChan = make(chan *Item,100)
	var resultChan = make(chan *Result,100)

	go producer(itemChan)

	startWorker(20,itemChan,resultChan)
	printResult(resultChan)
}

```

