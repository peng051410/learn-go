package goroutine_benchmark_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"sync"
	"testing"
	"time"
)

func spawn(f func() error) <-chan error {
	c := make(chan error)
	go func() {
		c <- f()
	}()

	return c
}

func TestGraceExit(t *testing.T) {
	c := spawn(func() error {
		time.Sleep(1 * time.Second)
		return errors.New("timeout")
	})
	log.Println(<-c)
}

func TestChannelDefine(t *testing.T) {
	var ch chan int
	t.Log(ch)
	assert.Equal(t, nil, ch)
}

func TestChannelInit(t *testing.T) {
	ch := make(chan int)
	t.Log(ch)
	assert.True(t, ch != nil)
}

func TestChannelInitWithBuffer(t *testing.T) {
	ch := make(chan int, 10)
	t.Log(ch)
	assert.True(t, ch != nil)
}

// fatal error: all goroutines are asleep - deadlock!
func TestChannelSyncDeadlock(t *testing.T) {
	ch := make(chan int)
	ch <- 10
	n := <-ch
	t.Log(n)
}

func TestChannelSync(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 10 //send 10 to channel ch
	}()
	n := <-ch //receive data from channel ch
	t.Log(n)
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 10
	t.Log(<-ch)
}

// fatal error: all goroutines are asleep - deadlock!
func TestChannelFullDeadlock(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 10
	ch <- 11
}

func TestSendOnlyChannel(t *testing.T) {
	ch := make(chan<- int, 1)
	ch <- 10
	//t.Log(<-ch) //invalid operation: <-ch (receive from send-only type chan<- int)
}

func TestReceiveOnlyChannel(t *testing.T) {
	ch := make(<-chan int, 1)
	//ch <- 10 //invalid operation: ch <- 10 (send to receive-only type <-chan int)
	t.Log(ch)
}

func TestProduceConsume(t *testing.T) {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()
	wg.Wait()
}

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func consume(ch <-chan int) {
	for i := range ch {
		log.Println(i)
	}
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int, 5)
	close(ch)
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			assert.Equal(t, "send on closed channel", err.(error).Error())
		}
	}()
	ch <- 10 //panic: send on closed channel
}

func TestChannelSelect(t *testing.T) {
	ch := make(chan int, 5)
	ch2 := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			ch2 <- i * 2
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()

	for {
		select {
		case n := <-ch:
			log.Println(n)
			if n > 5 {
				return
			}
		case n, ok := <-ch2:
			if !ok {
				log.Println("ch2 closed")
				return
			}
			log.Println(n)
			if n > 5 {
				return
			}
		default:
			log.Println("default")
			time.Sleep(1 * time.Second)
		}
	}
}

type signal struct{}

func worker() {
	println("worker is working...")
	time.Sleep(1 * time.Second)
}

func spawn1(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		c <- signal{}
	}()

	return c
}

func TestChannelAsSignal(t *testing.T) {
	println("start a worker")
	c := spawn1(worker)
	<-c
	println("worker is done")
}

func workerWithNum(i int) {
	fmt.Printf("worker %d is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d is done\n", i)
}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal) <-chan signal {

	c := make(chan signal)
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-groupSignal
			fmt.Printf("worker %d start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

func TestChannelAsMultiSignal(t *testing.T) {
	fmt.Printf("start %d workers\n", 10)
	groupSignal := make(chan signal)
	c := spawnGroup(workerWithNum, 10, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	close(groupSignal)
	<-c
	fmt.Println("the group of workers is done")
}

type counter struct {
	sync.Mutex
	i int
}

var cter counter

func Increase() int {
	cter.Lock()
	defer cter.Unlock()
	cter.i++
	return cter.i
}

func TestCounterWithShareMemory(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := Increase()
			fmt.Printf("goroutine %d: counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

type counterWithChannel struct {
	c chan int
	i int
}

func NewCounter() *counterWithChannel {
	cter := &counterWithChannel{
		c: make(chan int),
	}

	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	return cter
}

func (c *counterWithChannel) Increase() int {
	return <-c.c
}

func TestCounterWithChannel(t *testing.T) {
	cter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cter.Increase()
			fmt.Printf("goroutine %d: counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func TestBufferedChannelAsCountingSemaphore(t *testing.T) {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- i + 1
		}
		close(jobs)
	}()

	var wg sync.WaitGroup
	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			active <- struct{}{}
			fmt.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			<-active
			wg.Done()
		}(j)
	}
	wg.Wait()
}

func producer(c chan<- int) {

	var i int = 1
	for {
		time.Sleep(2 * time.Second)
		ok := trySend(c, i)
		if ok {
			fmt.Printf("[producer] send %d to channel\n", i)
			i++
			continue
		}
		fmt.Printf("[producer] try send [%d], but channel is full\n", i)
	}
}

func tryRecv(c <-chan int) (int, bool) {
	select {
	case v := <-c:
		return v, true
	default: // channel is empty
		return 0, false
	}
}

func trySend(c chan<- int, v int) bool {
	select {
	case c <- v:
		return true
	default: // channel is full
		return false
	}
}

func consumer(c <-chan int) {
	for {
		i, ok := tryRecv(c)
		if !ok {
			fmt.Printf("[consumer] try recv, but channel is empty\n")
			time.Sleep(1 * time.Second)
			continue
		}

		fmt.Printf("[consumer] recv %d from channel\n", i)
		if i >= 3 {
			fmt.Printf("[consumer] recv %d from channel, and exit\n", i)
			return
		}
	}
}

// len(channel) only can use on one sender and multi receiver or multi sender and one receiver
func TestChangeChannelState(t *testing.T) {

	var wg sync.WaitGroup
	c := make(chan int, 3)
	wg.Add(2)
	go func() {
		producer(c)
		wg.Done()
	}()

	go func() {
		consumer(c)
		wg.Done()
	}()

	wg.Wait()
}

func TestNilChannelRead(t *testing.T) {
	var c chan int
	<-c // fatal error: all goroutines are asleep - deadlock!
}

func TestNilChannelWrite(t *testing.T) {
	var c chan int
	c <- 1 // fatal error: all goroutines are asleep - deadlock!
}

func TestNilChannelApply(t *testing.T) {

	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 5
		close(ch1)
	}()

	go func() {
		time.Sleep(7 * time.Second)
		ch2 <- 7
		close(ch2)
	}()

	for {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				fmt.Println("ch1 is closed")
			} else {
				fmt.Printf("ch1 recv %d\n", v)
			}
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				fmt.Println("ch2 is closed")
			} else {
				fmt.Printf("ch2 recv %d\n", v)
			}
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}
	fmt.Println("program is end")
}

func MyAdd(tt int, ch chan<- int) {
	switch tt {
	case 1:
		ch <- 8
	case 0:
		ch <- 3
	default:
		ch <- -1
	}
}

func TestMultiAdd(t *testing.T) {
	ch := make(chan int)
	for i := 0; i < 2; i++ {
		go MyAdd(i, ch)
	}

	t.Log("start calculate sum")
	sum := 0
	for i := 0; i < 2; i++ {
		sum += <-ch
	}
	t.Log(sum)
}

func TestMultiAddWithWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	var sum = 0
	for i := 1; i < 3; i++ {
		go func(i int) {
			defer wg.Done()
			sum += i
		}(i)
	}
	wg.Wait()
	t.Log(sum)
}
