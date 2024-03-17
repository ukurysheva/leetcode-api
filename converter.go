package leetcode_api

func (l *LeetcodeAPI) processQuestion(q QuestionFromApi) *Question {
	return &Question{
		AcceptanceRate:   q.AcRate,
		Difficulty:       Difficulty(q.Difficulty),
		FreqBar:          q.FreqBar,
		QuestionID:       q.FrontendQuestionID,
		PaidOnly:         q.PaidOnly,
		Status:           q.Status,
		Title:            q.Title,
		TitleSlug:        q.TitleSlug,
		TopicTags:        q.TopicTags,
		HasSolution:      q.HasSolution,
		HasVideoSolution: q.HasVideoSolution,
	}
}
