package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var v [6]int
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<6; i++{
		v[i] = rand.Intn(60)
	}
	fmt.Println(v)
}