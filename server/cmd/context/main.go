package main

import (
	"context"
	"fmt"
	"time"
)

type paramKey struct {
}

func main() {
	c := context.WithValue(context.Background(), paramKey{}, "abc")
	c, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()
	mainTask(c)
}

func mainTask(c context.Context) {
	fmt.Printf("main task started with param %q\n", c.Value(paramKey{}))
	smallTask(context.Background(), "task1")
	smallTask(c, "task2")
}

func smallTask(c context.Context, name string) {
	fmt.Printf("%s started with param %q\n", name, c.Value(paramKey{}))
	select {
	case <-time.After(6 * time.Second):
		fmt.Printf("%s done\n", name)
	case <-c.Done():
		fmt.Printf("%s cancelled\n")
	}
}
