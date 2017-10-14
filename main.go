package main

import (
	"bufio"
	"fmt"
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

		sense := dictResp.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions[0]
		fmt.Println(sense)

		time.Sleep(time.Second * 1)
	}
}
