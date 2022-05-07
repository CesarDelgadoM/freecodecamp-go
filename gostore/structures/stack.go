package structures

import (
	"fmt"

	"github.com/CesarDelgadoM/gostore/store"
)

type Stack struct {
	head *store.Transaction
	size int
}

func (s *Stack) Push() {
	temp := store.Transaction{
		Store: make(map[string]string),
	}
	temp.Next = s.head
	s.head = &temp
	s.size++
}

func (s *Stack) Pop() {
	if s.head != nil {
		s.head = s.head.Next
		s.size--
	} else {
		fmt.Printf("ERROR: No active transactions\n")
	}
}

func (s *Stack) Peek() *store.Transaction {
	return s.head
}

func (s *Stack) Commit() {
	activeTransaction := s.Peek()
	if activeTransaction != nil {
		for key, value := range activeTransaction.Store {
			store.GlobalStore[key] = value
		}
	} else {
		fmt.Printf("INFO: Nothing to commit\n")
	}
}

func (s *Stack) RollBack() {
	if s.head != nil {
		for key := range s.head.Store {
			delete(s.head.Store, key)
		}
	} else {
		fmt.Printf("ERROR: No active transaction\n")
	}
}

func (s *Stack) Count() int {
	return s.size
}
