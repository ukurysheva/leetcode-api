package leetcode_api

import (
	"encoding/json"
	"github.com/microcosm-cc/bluemonday"
	"io"
)

func (l *LeetcodeAPI) GetQuestionContentByTitleSlug(titleSlug string) (*QuestionContent, error) {
	payload := makeQuestionContentGraphqlPayload(titleSlug)
	response, err := callAPI(payload)
	if err != nil {
		return nil, err
	}
	responseBody, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	var responseObj QuestionContentResponseBody
	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		return nil, err
	}

	// Do this once for each unique policy, and use the policy for the life of the program
	// Policy creation/editing is not safe to use in multiple goroutines
	p := bluemonday.StripTagsPolicy()

	// The policy can then be used to sanitize lots of input and it is safe to use the policy in multiple goroutines
	html := p.Sanitize(
		responseObj.Data.Question.Content,
	)

	content := &QuestionContent{Raw: html}

	return content, err
}

type QuestionContentResponseBody struct {
	Data struct {
		Question struct {
			Content      string        `json:"content"`
			MysqlSchemas []interface{} `json:"mysqlSchemas"`
			DataSchemas  []interface{} `json:"dataSchemas"`
		} `json:"question"`
	} `json:"data"`
}
