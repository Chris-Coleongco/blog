package api

import (
	"encoding/json"
	"fmt"
	//"fmt"
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

func parse_github_api_json(body []byte) {
	var data []interface{}

	// Unmarshal JSON into the slice
	err := json.Unmarshal(body, &data)

	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	fmt.Println(len(data))
	for key, item := range data {
		// Check if the item is a map
		itemMap, err := item.(map[string]interface{})
		if err {
			// Extract 'html_url' from the map
			htmlURL, err := itemMap["html_url"].(string)
			if err {
				fmt.Println("GitHub Page URL:", htmlURL)
			} else {
				fmt.Println("html_url not found or not a string")
			}
		} else {
			fmt.Printf("Item at key %d is not a map\n", key)
		}
	}

}

func get_all_github_repos() []byte {

	var full_link string = "https://api.github.com/users/Chris-Coleongco/repos"

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

	// ge the svn_link in the body

	parse_github_api_json(body)

	return body

}

func (s *API_Server) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		w.Write(get_all_github_repos())
	})
	router.HandleFunc("/repo/{repoName}", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		repoName := r.PathValue("repoName")
		w.Write([]byte(repoName)) // use reponame to curl github page and get the README.md of the repo.
		w.Header().Set("Access-Control-Allow-Origin", "*")
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
