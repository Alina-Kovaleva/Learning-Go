package main

import (
	"fmt"
	"log"
	"time"
)

func sum(nums []int, pipe chan<- int) {
	sum := 0
	log.Printf("Goroutine starts: %s\n", time.Now().String())
	for _, num := range nums {
		sum += num
		// fmt.Println(num, sum)
	}
	pipe <- sum
	log.Printf("Goroutine ends: %s\n", time.Now().String())
}

func main() {
	s := []int{0, 4, 7, -6, -3, 8, 2, 6, 9, 56, 89}
	c := make(chan int, 2)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	p1, p2 := <-c, <-c
	allSum := p1 + p2

	fmt.Println("First part sum = ", p1)
	fmt.Println("Second part sum = ", p2)
	fmt.Println("Sum = ", allSum)

}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sum(nums []int, pipe chan<- int) {
// 	sum := 0
// 	for _, num := range nums {
// 		sum += num
// 		time.Sleep(time.Second)
// 		// fmt.Println(nums[0])
// 		fmt.Println(num, sum)
// 	}
// 	pipe <- sum
// }

// func main() {
// 	s := []int{0, 4, 7, -6, -3, 8, 2, 6, 9, 56, 89}
// 	// c1 := make(chan int)
// 	// c2 := make(chan int)
// 	c := make(chan int, 2)

// 	// go sum(s[:len(s)/2], c1)
// 	// go sum(s[len(s)/2:], c2)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)

// 	// p1, p2 := <-c1, <-c2
// 	p1, p2 := <-c, <-c
// 	allSum := p1 + p2

// 	fmt.Println("First part sum = ", p1)
// 	fmt.Println("Second part sum = ", p2)
// 	fmt.Println("Sum = ", allSum)

// }
