package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SearchResult struct {
	Word         string `json:"word"`
	ID           string `json:"id"`
	InflectionID string `json:"inflection_id"`
	MatchString  string `json:"matchString"`
	Region       string `json:"region"`
	MatchType    string `json:"matchType"`
}
type SearchResponse struct {
	Results []SearchResult `json:"results"`
}

type Example struct {
	Text string `json:"text"`
}
type Sense struct {
	ID          string    `json:"id"`
	Definitions []string  `json:"definitions"`
	Examples    []Example `json:"examples"`
	Subsenses   []Sense   `json:"subsenses"`
}
type Entry struct {
	Etymologies     []string `json:"etymologies"`
	Senses          []Sense  `json:"senses"`
	HomographNumber string   `json:"homographNumber"`
}
type LexicalEntry struct {
	Entries         []Entry `json:"entries"`
	Language        string  `json:"language"`
	LexicalCategory string  `json:"lexicalCategory"`
	Text            string  `json:"text"`
}
type DictionaryResult struct {
	ID             string         `json:"id"`
	Languange      string         `json:"language"`
	Type           string         `json:"type"`
	Word           string         `json:"word"`
	LexicalEntries []LexicalEntry `json:"lexicalEntries"`
}
type DictionaryResponse struct {
	Results []DictionaryResult `json:"results"`
}

const Base = "https://od-api.oxforddictionaries.com/api/v1"
const SearchEndpoint = "/search/en"
const DictionaryEndpoint = "/entries/en/"
const AppID = "074bca24"
const AppKey = "8b712bc4bb6ba238e877adbbc7d4cd1a"

// CreateSearchRequest new req
func CreateSearchRequest(niddle string) (*http.Request, error) {
	req, err := http.NewRequest("GET", Base+SearchEndpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("app_id", AppID)
	req.Header.Set("app_key", AppKey)

	q := req.URL.Query()
	q.Add("q", niddle)
	req.URL.RawQuery = q.Encode()

	return req, nil
}

// CreateDictionaryRequest new req
func CreateDictionaryRequest(wordID string) (*http.Request, error) {
	req, err := http.NewRequest("GET", Base+DictionaryEndpoint+wordID, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("app_id", AppID)
	req.Header.Set("app_key", AppKey)

	return req, nil
}

// DoRequest do request
func DoRequest(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, fmt.Errorf("%s returned %d\n", resp.StatusCode)
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func GetSearchResponse(word string) (*SearchResponse, error) {
	searchReq, err := CreateSearchRequest(word)
	if err != nil {
		return nil, err
	}

	searchData, err := DoRequest(searchReq)
	if err != nil {
		return nil, err
	}

	searchResp := &SearchResponse{}
	err = json.Unmarshal(searchData, searchResp)
	if err != nil {
		return nil, err
	}
	if len(searchResp.Results) == 0 {
		return nil, errors.New("No search results returned from API")
	}

	return searchResp, nil
}

func GetDictionaryResponse(wordID string) (*DictionaryResponse, error) {
	dictReq, err := CreateDictionaryRequest(wordID)
	if err != nil {
		return nil, err
	}

	dictData, err := DoRequest(dictReq)
	if err != nil {
		return nil, err
	}

	dictResp := &DictionaryResponse{}
	err = json.Unmarshal(dictData, dictResp)
	if err != nil {
		return nil, err
	}
	if len(dictResp.Results) == 0 {
		return nil, errors.New("No dictionary results returned from API")
	}

	return dictResp, nil
}
