package api

import (
	"io"
	"log"
	"net/http"
)

type API_Server struct {
	addr string
}

func New_Api_Server(addr string) *API_Server {
	return &API_Server{
		addr: addr,
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Get_Github_README(repoName string) []byte {

	var full_link string = "https://raw.githubusercontent.com/Chris-Coleongco/" + repoName + "/main/README.md" // case in the file name DOES MATTER

	resp, resp_err := http.Get(full_link)

	if resp_err != nil {
		log.Println("response failed")
	}

	defer resp.Body.Close()

	body, err_body := io.ReadAll(resp.Body)

	if err_body != nil {
		log.Println("couldn't get body :c")
	}

	println(string(body))

	return body

}

func (s *API_Server) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("/repo/{repoName}", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		repoName := r.PathValue("repoName")
		w.Write([]byte(repoName)) // use reponame to curl github page and get the README.md of the repo.
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		page_data := Get_Github_README(repoName)

		w.Write([]byte(page_data))

	})

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Println("server started!")

	return server.ListenAndServe()
}
