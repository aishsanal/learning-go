package main

import (
	"fmt"
	"net/http"

	"github.com/aishsanal/learning-go/hello-world/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		n, err := fmt.Fprintf(w, "Hello, world!!")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println("Number of bytes printed:", n)
// 	})

// 	_ = http.ListenAndServe(":8080", nil)
// }

/*func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func about(w http.ResponseWriter, r *http.Request) {
	sum := add(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

func divide(w http.ResponseWriter, r *http.Request) {
	res, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, res))
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	res := x / y
	return res, nil
}

func add(x, y int) int {
	return x + y
}
*/
