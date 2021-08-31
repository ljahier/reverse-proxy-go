// listen on 80 port
// check origin domain
// keep connection alive
// hashmap or for loop to get the config in the config file

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Name       string `json:"name"`
	ServerPort string `json:"server_port"`
}

func main() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	var config []Config
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &config)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < len(config); i++ {
			if config[i].Name == r.Host {
				// keep connection alive and make http call to config[i].Port and return the response
				resp, err := http.Get(config[i].ServerPort)
				if err == nil {
					defer resp.Body.Close()
					body, err := io.ReadAll(resp.Body)
					if err == nil {
						fmt.Fprint(w, "", string(body))
						return
					}
				}
			}
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	http.ListenAndServe(":80", nil)
}
