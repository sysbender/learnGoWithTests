package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	//fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)

}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "Elodie")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler)))
}
