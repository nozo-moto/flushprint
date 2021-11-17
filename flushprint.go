package flushprint

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var (
	writer = bufio.NewWriter(os.Stdout)
	mux    sync.Mutex
)

func Print(a ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	fmt.Fprintln(writer, fmt.Sprint(a...))
	err := writer.Flush()
	if err != nil {
		panic(err)
	}
}
