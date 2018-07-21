package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Repo is...
type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	var name string
	fmt.Printf("name: ")
	fmt.Scanf("%s", &name)
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", name)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("GET Request is error")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Read Body is error")
	}
	var repos []Repo
	json.Unmarshal(data, &repos)
	for _, repo := range repos {
		fmt.Println(repo.ID, repo.Name, repo.URL)
	}
}
