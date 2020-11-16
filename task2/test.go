package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type WebSite struct {
	url    string
	header map[string]string
}

func (keyword WebSite) GetHtml() string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", keyword.url, nil)

	for key, val := range keyword.header {
		req.Header.Add(key, val)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return string(body)
}

func parse() {
	header := map[string]string{
		"Host":                      "flura.cn",
		"Connection":                "keep-alive",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:82.0) Gecko/20100101 Firefox/82.0",
		"Referer":                   "http://flura.cn/",
	}

	fmt.Println("正在抓取")
	for i := 1; i <= 5; i++ {
		url := "http://flura.cn"
		url0 := "http://flura.cn"

		if i != 1 {
			url = "http://flura.cn" + "/page/" + strconv.Itoa(i) + "/"
		}

		fmt.Println(url)

		WebSit := WebSite{url, header}
		WebSit.header["Referer"] = url

		html := WebSit.GetHtml()
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			log.Fatalln(err)
		}

		rb := "C:/Users/xuanke/Desktop/web/"
		ra := rb + strconv.Itoa(i) + ".html"
		rf, err := os.Create(ra)
		if err != nil {
			panic(err)
		}
		defer rf.Close()
		rf.WriteString(html)

		var title string
		var tim, tip, summary [20]string

		//时间

		dom.Find("article div[class=post-meta] span[class=post-time]").Each(func(i int, selection *goquery.Selection) {
			str := strings.Replace(selection.Text(), " ", "", -1)
			str = strings.Replace(str, "\n", "", -1)
			tim[i] = str
		})

		//标签

		dom.Find("article div[class=post-meta] a").Each(func(i int, selection *goquery.Selection) {
			str := strings.Replace(selection.Text(), " ", "", -1)
			str = strings.Replace(str, "\n", "", -1)
			tip[i] = str
		})

		//概要

		dom.Find("article div[class=post-body] ").Each(func(i int, selection *goquery.Selection) {
			str := strings.Replace(selection.Text(), " ", "", -1)
			str = strings.Replace(str, "\n", "", -1)
			summary[i] = str
		})

		dom.Find("article header h1 a").Each(func(i int, selection *goquery.Selection) {
			link, _ := selection.Attr("href")
			title = selection.Text()

			b := "C:/Users/xuanke/Desktop/web/"
			a := b + title + ".txt"
			f, err := os.Create(a)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			//写入

			f.WriteString("\xEF\xBB\xBF")
			f.WriteString("题目：" + title + "\n")
			f.WriteString("时间：" + tim[i] + "\n")
			f.WriteString("标志：" + tip[i] + "\n")
			f.WriteString("概要：" + summary[i] + "\n")

			var Article WebSite
			Article = WebSit
			Article.header["Referer"] = url0 + link

			fmt.Println(WebSit.header["Referer"])
			Article.url = Article.header["Referer"]

			//获取具体内容

			htmll := Article.GetHtml()
			domm, err := goquery.NewDocumentFromReader(strings.NewReader(htmll))
			if err != nil {
				log.Fatalln(err)
			}

			domm.Find("body main div[class=post-body]").Each(func(ii int, selection2 *goquery.Selection) {
				f.WriteString(selection2.Text())
			})

		})
	}

}

func main() {
	t1 := time.Now()
	parse()
	usedt := time.Since(t1)
	fmt.Println("爬虫结束，总耗时: ", usedt)
}
