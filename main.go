package main

import (
	"bytes"
	"fmt"
	"io"
)

type Stuck struct {
	stuck []interface{}
}

func (s *Stuck) Push(T interface{}) {
	s.stuck = append(s.stuck, T)
}
func (s *Stuck) Pop() interface{} {
	index := len(s.stuck) - 1
	popped := s.stuck[index]
	s.stuck = s.stuck[:index]
	return popped
}

func main() {
	var stack Stuck
	message := "ReverseMeNOW"
	for _, v := range message {
		stack.Push(v)
	}
	var buf bytes.Buffer
	for _, _ = range message {
		value, ok := stack.Pop().(rune)
		if ok {
			io.WriteString(&buf, string(value))
		}
	}
	fmt.Println(buf.String())
}
