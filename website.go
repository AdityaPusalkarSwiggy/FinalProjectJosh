package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type StatusChecker interface {
	CheckWebsiteStatus(ctx context.Context, url string)
}

type HTTPChecker struct {
	Timeout int
}

func UpdateWebsiteStatus() {
	h := HTTPChecker{Timeout: 10}
	for {
		fmt.Println("\nStatus Check at:", time.Now().Format("2006-01-02 3:04:05 PM"))
		for key := range websiteStatus {
			go h.CheckWebsiteStatus(context.Background(), key)
		}
		time.Sleep(15 * time.Second)
	}
}

func (h HTTPChecker) CheckWebsiteStatus(ctx context.Context, url string) {
	client := http.Client{
		Timeout: time.Duration(h.Timeout) * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		websiteStatus[url] = "DOWN"
	} else {
		fmt.Println(url, "Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
		if resp.StatusCode <= 299 {
			websiteStatus[url] = "UP"
		} else {
			websiteStatus[url] = "DOWN"
		}
		defer resp.Body.Close()
	}
}
