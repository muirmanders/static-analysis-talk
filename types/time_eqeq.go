package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	sfTZ, _ := time.LoadLocation("America/Los_Angeles")
	nnTZ, _ := time.LoadLocation("Europe/Moscow")

	nowSF := now.In(sfTZ)
	nowNN := now.In(nnTZ)

	fmt.Println("==", nowSF == nowNN)
	fmt.Println("Equal", nowSF.Equal(nowNN))
}
