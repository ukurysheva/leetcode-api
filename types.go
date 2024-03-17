package leetcode_api

type Filters struct {
	OrderBy    OrderBy
	OrderType  OrderType
	Difficulty Difficulty
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
