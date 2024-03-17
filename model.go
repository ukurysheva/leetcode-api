package leetcode_api

type Question struct {
	AcceptanceRate   float64     `json:"acRate"`
	Difficulty       Difficulty  `json:"difficulty"`
	FreqBar          interface{} `json:"freqBar"`
	QuestionID       string      `json:"frontendQuestionId"`
	IsFavor          bool        `json:"isFavor"`
	PaidOnly         bool        `json:"paidOnly"`
	Status           interface{} `json:"status"`
	Title            string      `json:"title"`
	TitleSlug        string      `json:"titleSlug"`
	TopicTags        []TopicTag  `json:"topicTags"`
	HasSolution      bool        `json:"hasSolution"`
	HasVideoSolution bool        `json:"hasVideoSolution"`
}

type TopicTag struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Slug string `json:"slug"`
}

type QuestionsListData struct {
	QuestionList QuestionList `json:"problemsetQuestionList"`
}

type QuestionList struct {
	Total     int        `json:"total"`
	Questions []Question `json:"questions"`
}

type QuestionContent struct {
	Raw      string `json:"content"`
	Text     string
	Examples []string
}
