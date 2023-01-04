package main

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Blog struct {
	Articles []Article
}

func Init() *Blog {
	return &Blog{}
}

func (b *Blog) addArticle(article Article) {
	b.Articles = append(b.Articles, article)
}

func (b *Blog) fetchArticlesInBlog() []Article {
	return b.Articles
}
