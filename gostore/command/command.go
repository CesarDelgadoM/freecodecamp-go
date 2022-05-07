package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CesarDelgadoM/gostore/store"
	"github.com/CesarDelgadoM/gostore/structures"
)

func StartCLI() {
	reader := bufio.NewReader(os.Stdin)
	items := &structures.Stack{}
	fmt.Printf("[Transactional Key-Value Store - GO]\n")
	for {
		fmt.Printf("> ")
		text, _ := reader.ReadString('\n')
		operation := strings.Fields(text)
		switch operation[0] {
		case "BEGIN":
			items.Push()
		case "ROLLBACK":
			items.RollBack()
		case "COMMIT":
			items.Commit()
			items.Pop() //pasamos a la siguiente transaccion de la pila
		case "END":
			items.Pop()
		case "SET":
			Set(operation[1], operation[2], items)
		case "GET":
			Get(operation[1], items)
		case "DELETE":
			Delete(operation[1], items)
		case "COUNT":
			items.Count()
		case "GLOBAL":
			store.PrintGlobalStore()
		case "STOP":
			os.Exit(0)
		default:
			fmt.Printf("ERROR: Unrecognised operation %s\n", operation[0])
		}
	}
}

func Get(key string, s *structures.Stack) {
	activeTransaction := s.Peek()
	if activeTransaction != nil {
		if val, ok := activeTransaction.Store[key]; ok {
			fmt.Printf("%s\n", val)
		}
	} else {
		if val, ok := store.GlobalStore[key]; ok {
			fmt.Printf("%s\n", val)
		} else {
			fmt.Printf("%s not set\n", key)
		}
	}
}

func Set(key string, value string, s *structures.Stack) {
	activeTransaction := s.Peek()
	if activeTransaction != nil {
		activeTransaction.Store[key] = value
	} else {
		store.GlobalStore[key] = value
	}
}

func Delete(key string, s *structures.Stack) {
	activeTransaction := s.Peek()
	if activeTransaction != nil {
		delete(activeTransaction.Store, key)
	} else {
		delete(store.GlobalStore, key)
	}
}
