package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

var url = "http://localhost:8080"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	ID      int    `json:"id"`
}

type TodoList struct {
	List []Todo `json:"events"`
}

func TestAddTodoSuccess(t *testing.T) {
	api := url + "/api/v1/event"

	// Request JSON
	deadline := "2019-06-11T14:00:00+09:00"
	title := "レポート提出"
	memo := ""

	jsonStr := `{"deadline":"` + deadline + `","title":"` + title + `","memo":"` + memo + `"}`
	req, _ := http.NewRequest(
		"POST",
		api,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	rsp, err := client.Do(req)
	if err != nil {
		t.Fatal("Request failed")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != 200 {
		t.Fatal("Request failed with status code " + strconv.Itoa(rsp.StatusCode))
	}

	body, _ := ioutil.ReadAll(rsp.Body)
	var rspMsg Response
	if err := json.Unmarshal(body, &rspMsg); err != nil {
		t.Fatal("Response type is not Todo")
	} else if rspMsg.Status != "success" || rspMsg.Message != "registered" {
		t.Fatal("Response type or values is not correct")
	}
}

func TestAddTodoFail(t *testing.T) {
	api := url + "/api/v1/event"

	// Request JSON
	deadline := "2019/06/11/14:00:00+09:00"
	title := "レポート提出"
	memo := ""

	jsonStr := `{"deadline":"` + deadline + `","title":"` + title + `","memo":"` + memo + `"}`
	req, _ := http.NewRequest(
		"POST",
		api,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	rsp, err := client.Do(req)
	if err != nil {
		t.Fatal("Request failed")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != 400 {
		t.Fatal("Request failed with status code " + strconv.Itoa(rsp.StatusCode))
	}
}

func TestGetListSuccess(t *testing.T) {
	api := url + "/api/v1/event"
	req, _ := http.NewRequest("GET", api, nil)

	client := new(http.Client)
	rsp, err := client.Do(req)
	if err != nil {
		t.Fatal("Request failed")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != 200 {
		t.Fatal("Request failed with status code " + strconv.Itoa(rsp.StatusCode))
	}

	body, _ := ioutil.ReadAll(rsp.Body)
	var list TodoList
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatal("Response type is not Todo")
	}
}

func TestGetTodoSuccess(t *testing.T) {
	id := "1"
	api := url + "/api/v1/event/" + id
	req, _ := http.NewRequest("GET", api, nil)

	client := new(http.Client)
	rsp, err := client.Do(req)
	if err != nil {
		t.Fatal("Request failed")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != 200 {
		t.Fatal("Request failed with status code " + strconv.Itoa(rsp.StatusCode))
	}

	body, _ := ioutil.ReadAll(rsp.Body)
	var todo Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		t.Fatal("Response type is not Todo")
	}
}

func TestGetTodoFail(t *testing.T) {
	id := "-1"
	api := url + "/api/v1/event/" + id
	req, _ := http.NewRequest("GET", api, nil)

	client := new(http.Client)
	rsp, err := client.Do(req)
	if err != nil {
		t.Fatal("Request failed")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != 404 {
		t.Fatal("Request failed with status code " + strconv.Itoa(rsp.StatusCode))
	}
}
