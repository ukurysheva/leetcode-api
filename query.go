package leetcode_api

import "fmt"

func makeGraphqlQuery(query string, variables string) string {
	return fmt.Sprintf(graphqlQuery, query, variables)
}

// const
const graphqlQuery = "{\"query\": \"%s\",\n\"variables\": %s}"
