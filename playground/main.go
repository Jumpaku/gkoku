package main

import (
	"fmt"
	"time"
)

func main() {
	{
		t := time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)
		fmt.Println(t)
	}
	{
		t := time.Date(1, 1, 0, 1, 1, 1, 1, time.UTC)
		fmt.Println(t)
	}
	{
		t := time.Date(1, 1, -1, 1, 1, 1, 1, time.UTC)
		fmt.Println(t)
	}
	{
		t := time.Date(1, 1, 1, 0, 1, 1, 1, time.UTC)
		fmt.Println(t)
	}
	{
		t := time.Date(1, 1, 1, 24, 1, 1, 1, time.UTC)
		fmt.Println(t)
	}
}
