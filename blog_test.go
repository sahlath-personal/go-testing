package main

import "testing"

func TestAddArticle(t *testing.T) {
	blog := Init()
	blog.addArticle(Article{"test title", "test body"})
	if blog.Articles[0].Title != "test title" {
		t.Fatalf("Article was not added")
	}
}

func TestFetchArticlesInBlog(t *testing.T) {
	blog := Init()
	blog.addArticle(Article{"test tile", "test body"})
	articles := blog.fetchArticlesInBlog()
	if len(articles) == 0 {
		t.Errorf("fetch not working")
	}
}
