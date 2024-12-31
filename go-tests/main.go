package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

var serverUrl string = "http://localhost:8080"

func testGet() {

	resp, err := http.Get(serverUrl + "/users/31872957-10b0-407b-959e-1497bbdb5152")
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("testGet passed")
	} else {
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("testGet not passed - " + buf.String())
	}

}

func testGetList() {

	resp, err := http.Get(serverUrl + "/users")
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("testGet passed")
	} else {
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("testGetList not passed - " + buf.String())
	}

}

func testPost() {

	data := []byte(`{"id":"31872957-10b0-407b-959e-1497bbdb5152", "name": "wertj"}`)
	r := bytes.NewReader(data)
	resp, err := http.Post(serverUrl+"/users", "application/json", r)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("testPost passed")
	} else {
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("testPost not passed - " + buf.String())
	}

}

func testUpdate() {

	data := []byte(`{"id":"31872957-10b0-407b-959e-1497bbdb5152", "name": "weghjkjyt456rtj"}`)
	req, err := http.NewRequest(http.MethodPut, serverUrl+"/users", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("testUpdate passed")
	} else {
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("testUpdate not passed - " + buf.String())
	}

}

func testDelete() {

	mockId := "31872957-10b0-407b-959e-1497bbdb5152"

	data := []byte(`{"id":"31872957-10b0-407b-959e-1497bbdb5152", "name": "weghjkjyt456rtj"}`)
	req, err := http.NewRequest(http.MethodDelete, serverUrl+"/users/"+mockId, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("testDelete passed")
	} else {
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("testDelete not passed - " + buf.String())
	}

}

func main() {

	testPost()
	testGet()
	testUpdate()
	testGetList()
	testDelete()

}
