package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func scan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("scanning")
	defer r.Body.Close()

	cmd := exec.Command("tesseract", "stdin", "stdout")
	cmd.Stdin = r.Body
	cmd.Stdout = w
	err := cmd.Run()
	fmt.Println(err)

}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/scan", scan)
	http.ListenAndServe(":8080", nil)
}
