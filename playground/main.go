package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2022, 3, 31, 0, 0, 0, 0, time.UTC)
	fmt.Println(t)
	{
		t := t.AddDate(0, -1, 0)
		fmt.Println(t)
		t = t.AddDate(0, 0, -t.Day())
		fmt.Println(t)
	}
	{
		t := t.AddDate(0, 1, 0)
		fmt.Println(t)
		t = t.AddDate(0, 0, -t.Day())
		fmt.Println(t)
	}
}
