package goroutin_benchmark_test

import (
	"sync"
	"testing"
)

var cs = 0

var mu sync.Mutex
var c = make(chan struct{}, 1)

func criticalSectionSyncByMutex() {
	mu.Lock()
	cs++
	mu.Unlock()
}

func criticalSectionSyncByChannel() {
	c <- struct{}{}
	cs++
	<-c
}

func BenchmarkCriticalSectionSyncByMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		criticalSectionSyncByMutex()
	}
}

func BenchmarkCriticalSectionSyncByChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		criticalSectionSyncByChannel()
	}
}

func BenchmarkCriticalSectionSyncByMutexInParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			criticalSectionSyncByMutex()
		}
	})
}

func BenchmarkCriticalSectionSyncByChannelInParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			criticalSectionSyncByChannel()
		}
	})
}
