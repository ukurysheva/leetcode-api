package leetcode_api

import (
	"encoding/json"
	"fmt"
	"io"
)

type GetQuestionsListParams struct {
	Offset       int
	Limit        int
	Filters      Filters
	CategorySlug CategorySlug
}

type Filters struct {
	OrderBy    OrderBy
	OrderType  OrderType
	Difficulty string
}

type (
	CategorySlug string
	OrderBy      string
	OrderType    string
	Difficulty   string
)

const (
	OrderByID                   OrderBy = "FRONTEND_ID"
	OrderByAcceptanceRate       OrderBy = "AC_RATE"
	OrderByAcceptanceDifficulty OrderBy = "DIFFICULTY"

	OrderTypeAscending  OrderType = ""
	OrderTypeDescending OrderType = "DESCENDING"

	CategorySlugAlgorithms  CategorySlug = "algorithms"
	CategorySlugAll         CategorySlug = "all-code-essentials"
	CategorySlugDatabase    CategorySlug = "database"
	CategorySlugShell       CategorySlug = "shell"
	CategorySlugConcurrency CategorySlug = "concurrency"
	CategorySlugJavaScript  CategorySlug = "javascript"
	CategorySlugPandas      CategorySlug = "pandas"

	DifficultyEasy   Difficulty = "EASY"
	DifficultyMedium Difficulty = "MEDIUM"
	DifficultyHard   Difficulty = "HARD"
)

func GetQuestionsList(params GetQuestionsListParams) (*QuestionsListData, error) {
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

	return &responseObj.Data, err
}

type QuestionsListResponse struct {
	Data QuestionsListData `json:"data"`
}

type QuestionsListData struct {
	QuestionList QuestionList `json:"problemsetQuestionList"`
}

func makeQuestionListGraphqlPayload(params GetQuestionsListParams) string {
	variables := fmt.Sprintf(graphqlGetQuestionsListVariables,
		params.CategorySlug,
		params.Offset,
		params.Limit,
		params.Filters.OrderBy,
		params.Filters.OrderType,
		params.Filters.Difficulty,
	)

	return makeGraphqlQuery(graphqlGetQuestionsListQuery, variables)
}

const graphqlGetQuestionsListQuery = `query problemsetQuestionList(
  $categorySlug: String
  $limit: Int
  $skip: Int
  $filters: QuestionListFilterInput
) {
  problemsetQuestionList: questionList(
    categorySlug: $categorySlug
    limit: $limit
    skip: $skip
    filters: $filters
  ) {
    total: totalNum
    questions: data {
      acRate
      difficulty
      freqBar
      frontendQuestionId: questionFrontendId
      isFavor
      paidOnly: isPaidOnly
      status
      title
      titleSlug
      topicTags {
        name
        id
        slug
      }
      hasSolution
      hasVideoSolution
    }
  }
}`

const graphqlGetQuestionsListVariables = `{
	"categorySlug": "%s",
	"skip": %v,
	"limit": %v,
	"filters": {
		"orderBy": "%s",
		"sortOrder": "%s",
		"difficulty": "%s"
	}
}`
