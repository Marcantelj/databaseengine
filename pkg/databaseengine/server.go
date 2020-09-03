package databaseengine

import(
	"fmt"
	"net/http"
)

type Server struct{

}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HELLO WORLD")
}

func (s Server) Run() error {

	mux := http.NewServeMux()

	mux.HandleFunc("/", serveIndex)

	srv := http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	return srv.ListenAndServe()
}
