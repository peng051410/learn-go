package lock_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestMutexCopyErrorCase(t *testing.T) {
	var wg sync.WaitGroup
	i := 0
	var mu sync.Mutex

	wg.Add(1)
	go func(mu1 sync.Mutex) {
		mu1.Lock()
		i = 10
		time.Sleep(10 * time.Second)
		fmt.Printf("goroutine1: i = %d\n", i)
		mu1.Unlock()
		wg.Done()
	}(mu)

	time.Sleep(1 * time.Second)

	mu.Lock()
	i = 1
	fmt.Printf("main goroutine: i = %d\n", i)
	mu.Unlock()

	wg.Wait()
}

func TestMutexCopyRightCase(t *testing.T) {
	var wg sync.WaitGroup
	i := 0
	var mu sync.Mutex

	wg.Add(1)
	go func(mu1 *sync.Mutex) {
		mu1.Lock()
		i = 10
		time.Sleep(10 * time.Second)
		fmt.Printf("goroutine1: i = %d\n", i)
		mu1.Unlock()
		wg.Done()
	}(&mu)

	time.Sleep(1 * time.Second)

	mu.Lock()
	i = 1
	fmt.Printf("main goroutine: i = %d\n", i)
	mu.Unlock()

	wg.Wait()
}

type signal struct{}

var ready bool

func worker(i int) {
	fmt.Printf("worker %d: start working\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: finished\n", i)
}

func spawnGroup(f func(i int), num int, mu *sync.Mutex) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				mu.Lock()
				if !ready {
					mu.Unlock()
					time.Sleep(100 * time.Millisecond)
					continue
				}
				mu.Unlock()
				fmt.Printf("worker %d: ready\n", i)
				f(i)
				wg.Done()
				return
			}
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

func TestStateCheckWithoutCond(t *testing.T) {
	fmt.Println("start a group of workers")
	mu := &sync.Mutex{} //need use pointer
	c := spawnGroup(worker, 5, mu)

	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers is ready to work")

	mu.Lock()
	ready = true
	mu.Unlock()

	<-c
	fmt.Println("the group of workers has finished")
}

func spawnGroupWithCond(f func(i int), num int, groupSignal *sync.Cond) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			groupSignal.L.Lock()
			for !ready {
				groupSignal.Wait()
			}
			groupSignal.L.Unlock()
			fmt.Printf("worker %d: ready\n", i)
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

func TestStateCheckWithCond(t *testing.T) {
	fmt.Println("start a group of workers")
	groupSignal := sync.NewCond(&sync.Mutex{})
	c := spawnGroupWithCond(worker, 5, groupSignal)

	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers is start to work")

	groupSignal.L.Lock()
	ready = true
	groupSignal.Broadcast()
	groupSignal.L.Unlock()

	<-c
	fmt.Println("the group of workers has finished")
}

func TestAtomic(t *testing.T) {

	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 5000; i++ {
		wg.Add(1)

		go func() {
			atomic.AddUint64(&ops, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("ops: %d\n", ops)
}

func TestDeadLock(t *testing.T) {
	var mu1 sync.Mutex
	var mu2 sync.Mutex

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		mu1.Lock()
		time.Sleep(1 * time.Second)
		mu2.Lock()
		mu2.Unlock()
		mu1.Unlock()
		wg.Done()
	}()

	mu2.Lock()
	time.Sleep(1 * time.Second)
	mu1.Lock()
	mu1.Unlock()
	mu2.Unlock()

	wg.Wait()
}
