package pkg

import (
	"bytes"
	"github.com/bmsandoval/congenial-octo-bassoon/internal/modules"
	"io"
	"log"
	"os"
	"sync"
)

// from medium article https://medium.com/@hau12a1/golang-capturing-log-println-and-fmt-println-output-770209c791b4
// used to call my file parser and capture the os output for the sake of unit testing
func CaptureOutput(f func(string, ...modules.ModuleInterface), input string, modules ...modules.ModuleInterface) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f(input, modules...)
	writer.Close()
	return <-out
}
