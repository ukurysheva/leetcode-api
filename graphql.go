package leetcode_api

import "fmt"

const (
	graphqlGetQuestionsListQuery = `query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\n  problemsetQuestionList: questionList(\n    categorySlug: $categorySlug\n    limit: $limit\n    skip: $skip\n    filters: $filters\n  ) {\n    total: totalNum\n    questions: data {\n      acRate\n      difficulty\n      freqBar\n      frontendQuestionId: questionFrontendId\n      isFavor\n      paidOnly: isPaidOnly\n      status\n      title\n      titleSlug\n      topicTags {\n        name\n        id\n        slug\n      }\n      hasSolution\n      hasVideoSolution\n    }\n  }\n}\n    `

	graphqlGetQuestionsListVariables = `{
		"categorySlug": "%s",
		"skip": %v,
		"limit": %v,
		"filters": {
			"orderBy": "%s",
			"sortOrder": "%s",
			"difficulty": "%s"
		}
	}`

	graphqlGetQuestionContentQuery     = `query questionContent($titleSlug: String!) { question(titleSlug: $titleSlug) { content mysqlSchemas dataSchemas }}`
	graphqlGetQuestionContentVariables = `{
        "titleSlug": "%s"
    }`
)

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

func makeQuestionContentGraphqlPayload(titleSlug string) string {
	variables := fmt.Sprintf(graphqlGetQuestionContentVariables, titleSlug)

	return makeGraphqlQuery(graphqlGetQuestionContentQuery, variables)
}
