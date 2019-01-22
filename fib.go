package main

import "fmt"

func main() {
	// Making two buffered channels
	jobs := make(chan int,100)
	results := make(chan int,100)

	go worker(jobs,results)
	go worker(jobs,results)
	go worker(jobs,results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	//close channel after loop
	// Avoids deadlock error
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
 	
 } 

// worker has two channels 
// jobs channel contains ints that will undergo fib func
 // results channel contains results of this process
func worker(jobs <-chan int,results chan<- int){
// Over the range of numbers carry out fib func
	for n := range jobs{
		results <- fib(n)
	}
}

func fib(n int) int{
	if n <=1 {
		return n
	}
	return fib(n-1)+fib(n-2)
	
}