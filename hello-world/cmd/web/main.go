package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aishsanal/learning-go/hello-world/pkg/config"
	"github.com/aishsanal/learning-go/hello-world/pkg/handlers"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var (
	app     config.AppConfig
	session *scs.SessionManager
)

func main() {
	app.IsProd = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.IsProd
	session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = session

	ts, err := handlers.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot start application")
	}

	app.TemplateCache = ts
	app.UseCache = false
	handlers.SetConfig(app)

	repository := *handlers.CreateRepository(app)
	handlers.SetRepository(&repository)

	// http.HandleFunc("/", handlers.Home)
	// http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	// http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
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
