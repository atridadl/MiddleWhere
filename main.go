package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func frontEndHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>MiddleWhere?!?!?!?!?!</h1>")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Health Check"))
}

func webHookHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == os.Getenv("SECURE_TOKEN") || os.Getenv("SECURE_TOKEN") == "" {
		// Create a Bearer string by appending string access token
		// Create a new request using http
		req, err := http.NewRequest(os.Getenv("ENDPOINT_REQUEST_METHOD"), os.Getenv("ENDPOINT_URL"), nil)
		// add authorization header to the req
		authToken := os.Getenv("ENDPOINT_AUTH_TOKEN")
		if authToken != "" {
			var bearer = "Bearer " + authToken
			req.Header.Add("Authorization", bearer)
		}

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Internal Server Error!"))
			log.Println("Error while reading the response bytes:", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Success!"))
		log.Println("Success: ", body)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Unauthorized!"))
	}
}

func main() {
	loadEnv()

	http.HandleFunc("/", frontEndHandler)

	http.HandleFunc("/api/healthcheck", healthCheckHandler)

	http.HandleFunc("/api/triggerWebHook", webHookHandler)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
