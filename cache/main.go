package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func Fibonacci(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	return Fibonacci(n-1) + Fibonacci(n-2)
// }

// type Memory struct {
// 	f     Function
// 	cache map[int]FunctionResult
// 	m sync.Mutex
// }

// type Function func(key int) (interface{}, error)

// type FunctionResult struct {
// 	value interface{}
// 	err   error
// }

// func NewCache(f Function) *Memory {
// 	return &Memory{
// 		f:     f,
// 		cache: make(map[int]FunctionResult),
// 	}
// }

// func (m *Memory) Get(key int) (interface{}, error) {
// 	m.m.Lock()
// 	result, exists := m.cache[key]
// 	m.m.Unlock()

// 	if !exists {
// 		m.m.Lock()
// 		result.value, result.err = m.f(key)
// 		m.cache[key] = result
// 		m.m.Unlock()
// 	}
// 	return result.value, result.err
// }

// func GetFibonacci(n int) (interface{}, error) {
// 	return Fibonacci(n), nil
// }

// func main() {
// 	cache := NewCache(GetFibonacci)
// 	fibo := []int{42, 40, 41, 42, 38}
// 	var wg sync.WaitGroup
// 	for _, n := range fibo {
// 		wg.Add(1)
// 		go func(index int) {
// 			defer wg.Done()
// 			start := time.Now()
// 			value, err := cache.Get(index)
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Printf("%v %v %v\n", index, time.Since(start), value)
// 		}(n)

// 	}

// 	wg.Wait()
// }
