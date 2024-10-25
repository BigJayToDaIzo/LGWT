package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Oy! %s\n", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "interwebz!\n")
}

func main() {
	//stdout ezmode
	Greet(os.Stdout, "Boi!")

	// buffer ezmode 2
	b := bytes.Buffer{}
	Greet(&b, "buuuf gang ryze uuhp!")
	fmt.Printf("Buffer say: %s\n", b.String())

	// interwebz ezmode 3
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
