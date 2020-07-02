package main

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	var code int
	go func() {
		code = m.Run()
	}()

	time.Sleep(3 * time.Second)
	if code != 0 {
		log.Fatal(fmt.Errorf("error occurred, exit code: %d", code))
	}
}
