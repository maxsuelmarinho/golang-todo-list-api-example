package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	address := flag.String("server", "http://localhost:8080", "HTTP Gateway url, e.g. http://localhost:8080")
	flag.Parse()

	t := time.Now().In(time.UTC)
	pfx := t.Format(time.RFC3339Nano)

	var body string

	// call create
	createResponse, err := http.Post(*address+"/v1/todo", "application/json", strings.NewReader(fmt.Sprintf(`
		{
			"api": "v1",
			"toDo": {
				"title": "title (%s)",
				"description": "description (%s)",
				"reminder": "%s"
			}
		}
	`, pfx, pfx, pfx)))
	if err != nil {
		log.Fatalf("failed to call Create method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(createResponse.Body)
	createResponse.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Create response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Create response: Code=%d, Body=%s\n\n", createResponse.StatusCode, body)

	// Call Read
	var created struct {
		API string `json:"api"`
		ID  string `json:"id"`
	}
	err = json.Unmarshal(bodyBytes, &created)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON response of Create method: %v", err)
		fmt.Println("error: ", err)
	}

	readResponse, err := http.Get(fmt.Sprintf("%s%s/%s", *address, "/v1/todo", created.ID))
	if err != nil {
		log.Fatalf("failed to call Read method: %v", err)
	}
	bodyBytes, err = ioutil.ReadAll(readResponse.Body)
	readResponse.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Read response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Read response: Code=%d, Body=%s\n\n", readResponse.StatusCode, body)

	// Call Update
	updateRequest, err := http.NewRequest("PUT", fmt.Sprintf("%s%s/%s", *address, "/v1/todo", created.ID),
		strings.NewReader(fmt.Sprintf(`
		{
			"api": "v1",
			"toDo": {
				"title": "title (%s) updated",
				"description: "description (%s) updated",
				"reminder": "%s"
			}
		}
	`, pfx, pfx, pfx)))
	updateRequest.Header.Set("Content-Type", "application/json")
	updateResponse, err := http.DefaultClient.Do(updateRequest)
	if err != nil {
		log.Fatalf("failed to call Update method: %v", err)
	}
	bodyBytes, err = ioutil.ReadAll(updateResponse.Body)
	updateResponse.Body.Close()
	updateRequest.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Update response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Update response: Code=%d, Body=%s\n\n", updateResponse.StatusCode, body)

	// Call ReadAll
	readAllResponse, err := http.Get(fmt.Sprintf("%s%s", *address, "/v1/todo"))
	if err != nil {
		log.Fatalf("failed to call ReadAll method: %v", err)
	}
	bodyBytes, err = ioutil.ReadAll(readAllResponse.Body)
	readAllResponse.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read ReadAll response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("ReadAll response: Code=%d, Body=%s\n\n", readAllResponse.StatusCode, body)

	// Call Delete
	deleteRequest, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s/%s", *address, "/v1/todo", created.ID), nil)
	deleteRequest.Header.Set("Content-Type", "application/json")
	deleteResponse, err := http.DefaultClient.Do(deleteRequest)
	if err != nil {
		log.Fatalf("failed to call Delete method: %v", err)
	}
	bodyBytes, err = ioutil.ReadAll(updateResponse.Body)
	deleteResponse.Body.Close()
	deleteRequest.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Update response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Update response: Code=%d, Body=%s\n\n", deleteResponse.StatusCode, body)
}
