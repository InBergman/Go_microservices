package main

import "testing"

func testAllArticles(test *testing.T) {
	allList := getAllArticles()
	if len(allList) != len(articleList) {
		test.Fail()
	}

	for index, value := range allList {
		if v.ID != articleList[i].ID ||
		v.Title != articleList[i].Title ||
		v.Content != articleList[i].Content 
		{
			test.Fail()
			break
		}
	}
}
