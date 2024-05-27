package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go print(5, "halo")
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)

	var s1, s2, s3 string
	fmt.Scanln(&s1, &s2, &s3)

	fmt.Println(s1) 
	fmt.Println(s2) 
	fmt.Println(s3) 
}
