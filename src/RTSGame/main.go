package RTSGame

import (
	"fmt"
	"time"
)

func Hello() {
	fmt.Println("Hi from rts game")

	count := 1.0

	for {
		count += 0.5
		fmt.Printf("Hi from rts game %f\n", count)
		time.Sleep(100_000_000)
	}
}
