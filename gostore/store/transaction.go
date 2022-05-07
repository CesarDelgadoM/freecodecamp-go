package store

import "fmt"

type Transaction struct {
	Store map[string]string
	Next  *Transaction
}

var GlobalStore = make(map[string]string)

func PrintGlobalStore() {
	for key, value := range GlobalStore {
		fmt.Printf("key: %s value: %s\n", key, value)
	}

}
