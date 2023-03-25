package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	countOfCats, _ := strconv.Atoi(os.Getenv("COUNT_OF_CATS"))
	for i := 0; i < countOfCats; i++ {
		fmt.Println("🐈")
		time.Sleep(time.Second)
	}
}
