package main

import (
	"fmt"
	"time"
)

func main() {
	words := []string{"gopher", "cat", "dog", "test"}
	for _, word := range words {
		fmt.Println(word)
		select {
		case ans := <-input():
			if ans == word {
				fmt.Println("正解")
			} else {
				fmt.Println("不正解")
			}
		case <-time.After(5 * time.Second):
			fmt.Println("time up")
		}
	}
}

func input() <-chan string {
	ch := make(chan string)
	go func() {
		var ans string
		fmt.Print(">")
		fmt.Scanln(&ans)
		ch <- ans
	}()
	return ch
}
