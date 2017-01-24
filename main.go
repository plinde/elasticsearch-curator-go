package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	indexName := "foo"
	action := "close"
	_ = action

	createTestIndex("foo")
	prepareElasticActionTestRun(indexName)

	// executeElasticAction(indexName, action)
	// executeElasticAction(indexName, action)

}

func executeElasticAction(indexName string, action string) string {

	url := "http://localhost:9200/" + indexName + "/_" + action + ""
	fmt.Println("URL:\n", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return string(body)
}

func prepareElasticActionTestRun(indexName string) {
	prepareElasticAction(indexName, "close")
	prepareElasticAction(indexName, "open")
	prepareElasticAction(indexName, "close")
	prepareElasticAction(indexName, "open")
	fmt.Printf("done with test run on %s\n", indexName)
	return
}

func prepareElasticAction(indexName string, action string) string {

	fmt.Printf("action will be: '%s' on index: '%s' \n", action, indexName)
	switch action {
	case "open":
		fmt.Printf("opened %s", indexName)
		action := "open"
		executeElasticAction(indexName, action)
	case "close":
		fmt.Printf("closed %s", indexName)
		action := "close"
		executeElasticAction(indexName, action)
	case "state":
		fmt.Printf("state of %s is", indexName)
	}

	return action
}

func createTestIndex(indexName string) {

	url := "http://localhost:9200/" + indexName + "/bar/1"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return

}
