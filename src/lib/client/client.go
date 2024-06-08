package client

import (
	"emad10101/authGo/config"
	"fmt"
	"log"
	"net/http"
	"os"
)

var portvar int = 9094

var c *config.YAML = config.ReadConfigFile()

func Start() {
	client := http.NewServeMux()
	client.HandleFunc("/login", loginHandler)
	client.HandleFunc("/user-signup", ServeUserSignUpPage)
	log.Printf("Client is running at: %s", c.Client.Url)
	http.ListenAndServe(fmt.Sprintf(":%d", portvar), client)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling Login request")
	writeHtml(w, r, "./lib/client/static/html/login.html")
}

// Responds with HTML file from the path `filename`
func writeHtml(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
