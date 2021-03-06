package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func randwait(n int) {
	time.Sleep(time.Duration(rand.Intn(n) * 100) * time.Millisecond)
}

func main() {

	q := make([]int,0)
	done := make(chan int, 0)
	m := sync.Mutex{}
	cv := sync.NewCond(&m)

	items := 0

	fmt.Println("Hello")

	// Producer
	go func(id int) {
		randwait(5)
		m.Lock()
		q = append(q, id)
		items++
		m.Unlock()
		cv.Signal()
		print("Producer:", id, "\n")
		done <- 1
	}(1)

	// Producer
	go func(id int) {
		randwait(5)
		m.Lock()
		q = append(q, id)
		items++
		m.Unlock()
		cv.Signal()
		print("Producer:", id, "\n")
		done <- 1
	}(2)

	// Consumer
	go func(id int) {
		randwait(5)
		var tmp int
		m.Lock()
		for items == 0 {
			cv.Wait()
		}
		items--
		tmp = q[0]
		q = q[1:]
		m.Unlock()
		print("Consumer:", id, " ", tmp, "\n")
		done <- 1
	}(1)

	// Consumer
	go func(id int) {
		randwait(5)
		var tmp int
		m.Lock()
		for items == 0 {
			cv.Wait()
		}
		items--
		tmp = q[0]
		q = q[1:]
		m.Unlock()
		print("Consumer:", id, " ", tmp, "\n")
		done <- 1
	}(2)


	for i:=0; i<4; i++ {
		<- done
	}
}
