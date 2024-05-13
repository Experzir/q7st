package services

import (
	context "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type beefServer struct {
}

func NewBeefServer() BeefServer {
	return &beefServer{}
}

func (*beefServer) mustEmbedUnimplementedBeefServer() {}

func (beef *beefServer) CountBeef(context.Context, *BeefRequest) (*BeefResponse, error) {
	count, err := Count()
	if err != nil {
		return nil, err
	}

	data := map[string]map[string]int{"beef": count}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	res := BeefResponse{
		Beef: jsonBytes,
	}
	return &res, nil
}

func Count() (map[string]int, error) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	text := string(body)

	words := strings.Fields(text)

	wordCount := make(map[string]int)

	for _, word := range words {
		word = strings.ToLower(word)
		word = strings.Trim(word, ",.?!")
		wordCount[word]++
	}

	// dataByte, err := json.Marshal(wordCount)
	// if err != nil {
	// 	return nil, err
	// }

	// return dataByte, err
	return wordCount, nil
}
