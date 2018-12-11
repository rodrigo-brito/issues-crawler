package util

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rodrigo-brito/issues-crawler/config"
	"gopkg.in/resty.v1"
)

const baseURL = "https://api.github.com/graphql"

type GraphQLData struct {
	Query      string                 `json:"query"`
	Parameters map[string]interface{} `json:"variables"`
}

func (g *GraphQLData) Marshal() []byte {
	result, err := json.Marshal(g)
	if err != nil {
		log.Print(err)
	}
	return result
}

func GraphQL(data GraphQLData) (response *resty.Response, err error) {
	req := resty.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", config.GetKey())).
		SetBody(data.Marshal())

	return req.Post(baseURL)
}
