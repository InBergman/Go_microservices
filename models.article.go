package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "First content"},
	article{ID: 2, Title: "Article 2", Content: "Second content"},
}

func getAllArticles() []article {
	return articleList
}
