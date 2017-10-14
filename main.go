package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"oxford/api"
	"time"
)

func main() {
	readData()
}

func readData() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()
		fmt.Println(word)

		searchResp, err := api.GetSearchResponse(word)
		if err != nil {
			fmt.Println(err)
			continue
		}

		time.Sleep(time.Second * 1)

		wordID := searchResp.Results[0].ID
		dictResp, err := api.GetDictionaryResponse(wordID)
		if err != nil {
			fmt.Println(err)
			continue
		}

		dictResult := dictResp.Results[0]

		t, err := template.ParseFiles("./tmpl/word.html")
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = t.Execute(os.Stdout, dictResult)
		if err != nil {
			fmt.Println(err)
			continue
		}

		time.Sleep(time.Second * 1)
	}
}
