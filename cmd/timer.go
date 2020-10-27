package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()

	fmt.Println(fmt.Sprintf("%s: %s", time.Now().UTC(), id))
	for _ = range time.Tick(time.Second * 5) {
		fmt.Println(fmt.Sprintf("%s: %s", time.Now().UTC(), id))
	}
}
