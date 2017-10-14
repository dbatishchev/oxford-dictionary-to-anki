package api

import (
	"encoding/json"
	"net/http"
)

const YandexBase = "https://translate.yandex.net/api/v1.5/tr.json/translate"
const YandexAPIKey = "trnsl.1.1.20171014T132434Z.7fdcc037a26b0a8a.6b0fa46fc14f40d5e97fa08b780f52f637c5aa85"

type YandexResponse struct {
	Code string   `json:"code"`
	Lang string   `json:"lang"`
	Text []string `json:"text"`
}

// CreateYandexTranslateRequest new req
func CreateYandexTranslateRequest(word string) (*http.Request, error) {
	req, err := http.NewRequest("GET", Base, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("text", word)
	q.Add("lang", "en-ru")
	q.Add("key", YandexAPIKey)
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func GetYandexResponse(word string) (*YandexResponse, error) {
	yReq, err := CreateYandexTranslateRequest(word)
	if err != nil {
		return nil, err
	}

	yData, err := DoRequest(yReq)
	if err != nil {
		return nil, err
	}

	yResp := &YandexResponse{}
	err = json.Unmarshal(yData, yResp)
	if err != nil {
		return nil, err
	}

	return yResp, nil
}
