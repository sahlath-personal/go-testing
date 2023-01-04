package main

import "fmt"

func main() {
	fmt.Println("Good to Go from Sahla!!!")
	blog := Init()
	fmt.Println(blog)
	article := Article{"my first article", "content of my first article"}
	blog.addArticle(article)
	fmt.Println(blog)
}
