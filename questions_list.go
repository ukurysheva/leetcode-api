package leetcode_api

import (
	"encoding/json"
	"io"
	"sync"
)

func (l *LeetcodeAPI) GetQuestionsList(params GetQuestionsListParams) ([]Question, error) {
	payload := makeQuestionListGraphqlPayload(params)
	response, err := callAPI(payload)
	if err != nil {
		return nil, err
	}

	responseBody, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	var responseObj QuestionsListResponse
	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		return nil, err
	}

	questions := make([]Question, 0, len(responseObj.Data.QuestionsList.Questions))

	mx := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(len(responseObj.Data.QuestionsList.Questions))

	for _, q := range responseObj.Data.QuestionsList.Questions {
		go func(q QuestionFromApi) {
			defer wg.Done()
			mx.Lock()
			defer mx.Unlock()

			questionModel := l.processQuestion(q)

			questions = append(questions, *questionModel)

		}(q)
	}

	wg.Wait()

	return questions, err
}

type QuestionsListResponse struct {
	Data struct {
		QuestionsList QuestionsList `json:"problemsetQuestionList"`
	} `json:"data"`
}

type QuestionsList struct {
	Total     int               `json:"total"`
	Questions []QuestionFromApi `json:"questions"`
}

type QuestionFromApi struct {
	AcRate             float64     `json:"acRate"`
	Difficulty         string      `json:"difficulty"`
	FreqBar            interface{} `json:"freqBar"`
	FrontendQuestionID string      `json:"frontendQuestionId"`
	IsFavor            bool        `json:"isFavor"`
	PaidOnly           bool        `json:"paidOnly"`
	Status             interface{} `json:"status"`
	Title              string      `json:"title"`
	TitleSlug          string      `json:"titleSlug"`
	TopicTags          []TopicTag  `json:"topicTags"`
	HasSolution        bool        `json:"hasSolution"`
	HasVideoSolution   bool        `json:"hasVideoSolution"`
}
