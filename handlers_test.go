package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

// To run these test cases properly, the server must be running
func TestPostWebsite(t *testing.T) {
	data := Websites{
		Websites: []string{"http://www.swiggy.in", "http://www.google.co.in", "http://www.thisisreallyfake.co.in"},
	}
	payloadBuf := new(bytes.Buffer)
	if err := json.NewEncoder(payloadBuf).Encode(&data); err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest("POST", "http://127.0.0.1:3000/website", payloadBuf)
	if err != nil {
		t.Error(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Errorf("Expected Status Code 200, Got %v", resp.StatusCode)
	}
}

func TestGetWebsite(t *testing.T) {
	var websiteResponse = make(map[string]string)
	url := "http://127.0.0.1:3000/website"
	resp, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(body))
	if err := json.Unmarshal([]byte(string(body)), &websiteResponse); err != nil {
		t.Error(err)
	}
}

func TestGetWebsiteQuery(t *testing.T) {
	var websiteResponse = make(map[string]string)
	name := "http://www.swiggy.in"
	url := "http://127.0.0.1:3000/website?name=" + name
	resp, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(body))
	if err := json.Unmarshal([]byte(string(body)), &websiteResponse); err != nil {
		t.Error(err)
	}
	if len(websiteResponse) != 1 {
		t.Error("Incorrect Number of Websites Returned")
	}
}

// curl -X "POST" -d '{"websites":["http://www.google.co.in","http://www.swiggy.in", "http://www.thisisreallyfake.co.in", "http://www.apple.co.in/fake", "http://www.amazon.com/everythingisfree"]}' 'http://127.0.0.1:3000/website'
// curl -X "GET" 'http://127.0.0.1:3000/website'
// curl -X "GET" 'http://127.0.0.1:3000/website?name=http://www.swiggy.in'
