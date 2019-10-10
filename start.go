package main

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var baseUrl = "https://news.yahoo.co.jp/ranking/access/news"

type Article struct {
	Title string
	Url   string
	Date  time.Time
}

type ArticleList []Article

func main() {
	//fmt.Println("Go!!")
	doc, err := goquery.NewDocument(baseUrl)
	if err != nil {
		fmt.Printf("Failed")
	}
	title := doc.Find("title").Text()
	// println : 改行付き出力
	fmt.Println(title)
	fmt.Println(doc)
	// articleList, err := getList(baseUrl)
	// if err != nil {
	// 	panic(err)
	// }
}
