package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var config Config = configInit()

func main() {

	//build base elasticsearch URL
	url := fmt.Sprintf("%s://%s:%d", config.protocol, config.host, config.port)

	//use context struct
	ctx := ElasticCtx{url}
	_ = ctx

	fmt.Println(config.protocol, config.host, config.port, config.index, config.action)

	createTestIndex(ctx, config.index)
	prepareElasticActionTestRun(ctx, config.index)
	// executeElasticAction(ctx, config.index, config.action)
}

type Config struct {
	protocol string
	host     string
	port     int
	index    string
	url      string
	action   string
}

type ElasticCtx struct {
	url string
}

func configInit() Config {
	//config setup
	protocol := flag.String("protocol", "http", "HTTP vs. HTTPS")
	host := flag.String("host", "localhost", "hostname or IP for elasticsearch HTTP")
	port := flag.Int("port", 9200, "HTTP Port for ElasticSearch")
	index := flag.String("index", "foo", "Elasticsearch Index to perform action on")
	action := flag.String("action", "open", "Action to perform e.g. open or close index")
	flag.Parse()

	// deference yr pointers to get the actual option values
	// fmt.Println(*protocol, *host, *port)

	//initial Config struct
	config := Config{*protocol, *host, *port, *index, *index, *action}
	// fmt.Println(config.host)

	return config
}

func executeElasticAction(es ElasticCtx, index string, action string) string {

	url := fmt.Sprintf("%s/%s/_%s", es.url, config.index, config.action)
	fmt.Println("URL:>", url)

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

func prepareElasticActionTestRun(es ElasticCtx, index string) {
	prepareElasticAction(es, index, "close")
	prepareElasticAction(es, index, "open")
	prepareElasticAction(es, index, "close")
	prepareElasticAction(es, index, "open")
	fmt.Printf("done with test run on %s\n", index)
	return
}

func prepareElasticAction(es ElasticCtx, index string, action string) string {

	fmt.Printf("action will be: '%s' on index: '%s' \n", action, index)
	switch action {
	case "open":
		fmt.Printf("opened %s", index)
		action := "open"
		executeElasticAction(es, index, action)
	case "close":
		fmt.Printf("closed %s", index)
		action := "close"
		executeElasticAction(es, index, action)
	case "state":
		fmt.Printf("state of %s is", index)
	}

	return action
}

func createTestIndex(es ElasticCtx, index string) {

	url := fmt.Sprintf("%s/%s", es.url, index)
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
