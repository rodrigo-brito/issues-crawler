package crawler

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/rodrigo-brito/issues-crawler/util"
)

const (
	query = `
query Issues($owner: String!, $project: String!) {
  repository(owner: $owner, name: $project) {
    databaseId
    issues(last: 100) {
      nodes {
        title
        databaseId
        createdAt
        closedAt
        bodyText
        reactionGroups {
          content
          users(first: 100) {
            pageInfo {
              hasNextPage
              nextPage
            }
            totalCount
          }
        }
      }
    }
  }
}
	`
)

type Crawler struct {
	Context context.Context
}

type Time time.Time

func (t *Time) UnmarshalJSON(value []byte) error {
	parse, err := time.Parse(time.RFC3339, string(strings.Trim(string(value), `"`)))
	if err != nil {
		return err
	}
	*t = Time(parse)
	return nil
}

type Issue struct {
	ID             int    `json:"databaseId"`
	Title          string `json:"title"`
	CreatedAt      *Time  `json:"createdAt"`
	ClosedAt       *Time  `json:"closedAt"`
	BodyText       string `json:"bodyText"`
	ReactionGroups []struct {
		Content string `json:"content"`
		Users   struct {
			PageInfo struct {
				HasNextPage bool `json:"hasNextPage"`
			} `json:"pageInfo"`
			TotalCount int `json:"totalCount"`
		} `json:"users"`
	} `json:"reactionGroups"`
}

type APIResponse struct {
	Data struct {
		Repository struct {
			Issues struct {
				Nodes []*Issue `json:"nodes"`
			} `json:"issues"`
		} `json:"repository"`
	} `json:"data"`
}

func (c *Crawler) Fetch(ownerName, projectName string) (issues []*Issue, err error) {
	response, err := util.GraphQL(util.GraphQLData{
		Query: query,
		Parameters: map[string]interface{}{
			"owner":   ownerName,
			"project": projectName,
		},
	})

	if response.IsSuccess() {
		data := new(APIResponse)
		err := json.Unmarshal(response.Body(), data)
		if err != nil {
			return nil, err
		}

		return data.Data.Repository.Issues.Nodes, nil
	}

	return nil, nil
}

func NewCrawler(ctx context.Context) *Crawler {
	return &Crawler{
		Context: ctx,
	}
}
