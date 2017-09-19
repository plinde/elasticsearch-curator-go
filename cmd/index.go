// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	indexName string
	Verbose   bool
)

type RESTParams struct {
	method   string
	protocol string
	host     string
	port     int
	index    string
	action   string
}

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Index related operations",
	ValidArgs: []string{
		"close",
		"open",
	},

	Long: `Perform index index related operations in Elasticsearch`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("elasticsearch-curator-go v0.1")
	},
}

func makeRESTCall(restParams RESTParams) error {
	url := fmt.Sprintf("%s://%s:%d/%s%s", restParams.protocol, restParams.host, restParams.port, restParams.index, restParams.action)

	if Verbose {
		fmt.Println("Params:> ", restParams)
		fmt.Println("URL:>", url)
	}

	req, err := http.NewRequest(restParams.method, url, bytes.NewBuffer(nil))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if Verbose {
		fmt.Println("response Status:", resp.Status)
	}
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	if Verbose {
		fmt.Println("response Body:", string(body))
	}

	return err
}

var createIndexCmd = &cobra.Command{
	Use:   "create",
	Short: "creates empty target index",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		method := "PUT"
		protocol := "http"
		host := "localhost"
		port := 9200
		index := args[0]
		action := ""
		flag.Parse()

		restParams := RESTParams{method, protocol, host, port, index, action}
		err := makeRESTCall(restParams)
		return err

	},
}

var deleteIndexCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete target index",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		method := "DELETE"
		protocol := "http"
		host := "localhost"
		port := 9200
		index := args[0]
		action := ""
		flag.Parse()

		restParams := RESTParams{method, protocol, host, port, index, action}
		err := makeRESTCall(restParams)
		return err

	},
}

var openIndexCmd = &cobra.Command{
	Use:   "open",
	Short: "opens target index",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		method := "POST"
		protocol := "http"
		host := "localhost"
		port := 9200
		index := args[0]
		action := "/_open"
		flag.Parse()

		restParams := RESTParams{method, protocol, host, port, index, action}
		err := makeRESTCall(restParams)
		return err

	},
}

var closeIndexCmd = &cobra.Command{
	Use:   "close",
	Short: "closes target index",
	RunE: func(cmd *cobra.Command, args []string) error {

		method := "POST"
		protocol := "http"
		host := "localhost"
		port := 9200
		index := args[0]
		action := "/_close"
		flag.Parse()

		restParams := RESTParams{method, protocol, host, port, index, action}
		err := makeRESTCall(restParams)
		return err

	},
}

func init() {
	RootCmd.AddCommand(indexCmd)
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	RootCmd.AddCommand(versionCmd)
	indexCmd.AddCommand(openIndexCmd)
	indexCmd.AddCommand(closeIndexCmd)
	indexCmd.AddCommand(createIndexCmd)
	indexCmd.AddCommand(deleteIndexCmd)
	openIndexCmd.Flags().String("open", "", "open index")
	closeIndexCmd.Flags().String("close", "", "close index")
	createIndexCmd.Flags().String("create", "", "create index")
	deleteIndexCmd.Flags().String("delete", "", "delete index")

}
