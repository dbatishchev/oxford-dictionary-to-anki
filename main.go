package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"oxford/api"
	"time"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

func main() {
	readData()
}

func readData() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()
		fmt.Println(word)

		yResp, err := getYandexResponse(word)
		if err != nil {
			fmt.Println(err)
			continue
		}

		dictResult, err := getOxfordResponse(word)
		if err != nil {
			fmt.Println(err)
			continue
		}

		t, err := template.ParseFiles("./tmpl/word.html")
		if err != nil {
			fmt.Println(err)
			continue
		}

		templateContext := struct {
			YandexResults     []string
			DictionaryResults *api.DictionaryResult
		}{
			yResp.Text,
			dictResult,
		}

		buf := bytes.NewBufferString("")
		err = t.Execute(buf, templateContext)
		if err != nil {
			fmt.Println(err)
			continue
		}

		m := minify.New()
		m.AddFunc("text/html", html.Minify)
		if err := m.Minify("text/html", os.Stdout, buf); err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func getOxfordResponse(word string) (*api.DictionaryResult, error) {
	searchResp, err := api.GetSearchResponse(word)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Second * 1)

	wordID := searchResp.Results[0].ID
	dictResp, err := api.GetDictionaryResponse(wordID)
	if err != nil {
		return nil, err
	}

	dictResult := dictResp.Results[0]

	time.Sleep(time.Second * 1)

	return &dictResult, nil
}

func getYandexResponse(word string) (*api.YandexResponse, error) {
	yResp, err := api.GetYandexResponse(word)
	if err != nil {
		return nil, err
	}

	return yResp, nil
}
