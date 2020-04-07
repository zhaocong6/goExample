package main

import "fmt"

func main() {
	queue := make(chan string, 5)

	queue <- "one"
	queue <- "two"
	close(queue)

	for job := range queue{
		fmt.Println(job)
	}
}