package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
	port = "3003"
)

func StartApplication() {

	mapUrls()

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	fmt.Println(fmt.Sprintf("listening on port: %s", port) )

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
