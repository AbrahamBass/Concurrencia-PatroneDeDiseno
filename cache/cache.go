package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpresiveFibonacci(n int) int {
	fmt.Printf("Calculate Expesive Fibonacci %v\n", n)
	time.Sleep(5 * time.Second)
	return n
}

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Lock       sync.Mutex
}

func (s *Service) Work(job int) {
	s.Lock.Lock()
	exist := s.InProgress[job]
	if exist {
		s.Lock.Unlock()
		response := make(chan int)
		defer close(response)

		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()
		fmt.Printf("Response Job: %v\n", job)
		resp := <-response
		fmt.Printf("Response Done: %v", resp)
	}
	s.Lock.Unlock()

	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Calculate Fibonacci %v\n", job)
	result := ExpresiveFibonacci(job)

	s.Lock.Lock()
	pendingsWorkers, exists := s.IsPending[job]
	s.Lock.Unlock()

	if exists {
		for _, worker := range pendingsWorkers {
			worker <- result
		}

		fmt.Printf("Result Send %v\n", job)
	}

	s.Lock.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func main() {
	service := NewService()
	jobs := []int{3, 2, 3, 4, 5, 2, 8}

	var wg sync.WaitGroup
	wg.Add(len(jobs))
	for _, n := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(n)
	}

	wg.Wait()
}
