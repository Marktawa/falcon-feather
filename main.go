package main

import "fmt"
import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello</h1>")
	})

	http.ListenAndServe(":80", nil)
}

/* Step 1
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}
*/