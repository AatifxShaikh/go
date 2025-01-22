package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4"}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		dbCall(i)
	}
	fmt.Printf("\n Total execution time:%v", time.Since(t0))
}
func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The roast from the database is:", dbData[i])
}
