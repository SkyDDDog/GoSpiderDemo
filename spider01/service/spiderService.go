package service

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"spider01/entity"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Crawl() []entity.Fzu {
	url := "https://www.fzu.edu.cn/index/fdyw.htm"
	//url := "https://www.fzu.edu.cn/index/fdyw/69.htm"
	var client http.Client
	result := make([]entity.Fzu, 1)

	for true {
		isStart := false
		isStop := false
		req, _ := http.NewRequest("GET", url, nil)
		resp, _ := client.Do(req)
		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		doc.Find(".n_list ul li").
			Each(func(i int, s *goquery.Selection) {
				link, _ := s.Find("a").Attr("href")
				// 详情页
				req, _ := http.NewRequest("GET", link, nil)
				resp, _ := client.Do(req)
				newDoc, _ := goquery.NewDocumentFromReader(resp.Body)
				nav := newDoc.Find(".n_news_detail .ar_article_box .nav01")
				title, _ := nav.Find("h3").Html()
				var publish_time time.Time
				var author string
				var read_num int
				nav.Find("h6").Find("span").
					Each(func(i int, s *goquery.Selection) {
						switch i {
						case 0:
							date := strings.ReplaceAll(s.Text(), "发布日期:  ", "")
							timeLayoutStr := "2006-01-02"
							publish_time, _ = time.ParseInLocation(timeLayoutStr, date, time.Local)
							if publish_time.After(time.Date(2022, 11, 1, 0, 0, 0, 0, time.Local)) {
								isStart = true
							}
							if publish_time.Before(time.Date(2022, 9, 1, 0, 0, 0, 0, time.Local)) {
								isStop = true
							}
						case 1:
							author = strings.ReplaceAll(s.Text(), "作者： ", "")
						case 2:
							raw := s.Text()
							raw = strings.ReplaceAll(raw, "阅读： _showDynClicks(\"wbnews\", ", "")
							raw = strings.ReplaceAll(raw, ")", "")
							r := strings.Split(raw, ", ")
							owner := r[0]
							clickid := r[1]
							clickUrl := "https://news.fzu.edu.cn/system/resource/code/news/click/dynclicks.jsp?clickid=" + clickid + "&owner=" + owner
							req, _ := http.NewRequest("GET", clickUrl, nil)
							resp, _ := client.Do(req)
							clickDoc, _ := goquery.NewDocumentFromReader(resp.Body)
							read_num, _ = strconv.Atoi(clickDoc.Text())
						}
					})
				//html, _ := newDoc.Find("#vsb_content_2 DIV").Html()
				content := newDoc.Find("#vsb_content_2 DIV").Text()
				//fmt.Println(content)
				//fmt.Println(nav.Html())
				obj := entity.Fzu{
					Link:         link,
					Author:       author,
					Publish_time: publish_time,
					Title:        title,
					Content:      content,
					Read_num:     read_num,
				}
				if isStart {
					result = append(result, obj)
				}

				if isStop {
					return
				}
			})
		if isStop {
			break
		}
		//fmt.Println(isStop)
		prefix := strings.ReplaceAll(url, "fdyw.htm", "")
		reg := regexp.MustCompile(`\d*.htm`)
		prefix = reg.ReplaceAllString(prefix, "")
		next, _ := doc.Find(".pb_sys_common .p_next a").Attr("href")
		url = prefix + next
		//fmt.Println(url)
	}

	return result
}

func GetFzuPages(urlChan chan<- string) {
	var client http.Client
	url := "https://www.fzu.edu.cn/index/fdyw.htm"
	isStart := false
	isStop := false
	for true {
		req, _ := http.NewRequest("GET", url, nil)
		resp, _ := client.Do(req)
		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		doc.Find(".n_list ul li").
			Each(func(i int, s *goquery.Selection) {
				link, _ := s.Find("a").Attr("href")
				//详情页
				req, _ := http.NewRequest("GET", link, nil)
				resp, _ := client.Do(req)
				newDoc, _ := goquery.NewDocumentFromReader(resp.Body)
				nav := newDoc.Find(".n_news_detail .ar_article_box .nav01")
				var publish_time time.Time
				nav.Find("h6").Find("span").
					Each(func(i int, s *goquery.Selection) {
						switch i {
						case 0:
							date := strings.ReplaceAll(s.Text(), "发布日期:  ", "")
							timeLayoutStr := "2006-01-02"
							publish_time, _ = time.ParseInLocation(timeLayoutStr, date, time.Local)
							if !isStart && publish_time.Before(time.Date(2022, 11, 1, 0, 0, 0, 0, time.Local)) {
								isStart = true
							}
							if publish_time.Before(time.Date(2022, 9, 1, 0, 0, 0, 0, time.Local)) {
								isStop = true
								break
							} else if isStart {
								urlChan <- link
							}
						}
					})
				if isStop {
					return
				}
			})
		if isStop {
			urlChan <- "Closed"
			close(urlChan)
			break
		}
		prefix := strings.ReplaceAll(url, "fdyw.htm", "")
		reg := regexp.MustCompile(`\d*.htm`)
		prefix = reg.ReplaceAllString(prefix, "")
		next, _ := doc.Find(".pb_sys_common .p_next a").Attr("href")
		url = prefix + next
	}
}

func FzuSpider(wg *sync.WaitGroup, urlChan <-chan string, result chan<- entity.Fzu) {
	var client http.Client
	for url := range urlChan {
		wg.Add(1)
		if url == "Closed" {
			close(result)
			return
		}
		// 详情页
		req, _ := http.NewRequest("GET", url, nil)
		resp, _ := client.Do(req)
		newDoc, _ := goquery.NewDocumentFromReader(resp.Body)
		temp := newDoc.Find(".n_news_detail .ar_article_box .nav01")
		title, _ := temp.Find("h3").Html()
		var publish_time time.Time
		var author string
		var read_num int
		temp.Find("h6").Find("span").
			Each(func(i int, s *goquery.Selection) {
				switch i {
				case 0:
					date := strings.ReplaceAll(s.Text(), "发布日期:  ", "")
					timeLayoutStr := "2006-01-02"
					publish_time, _ = time.ParseInLocation(timeLayoutStr, date, time.Local)
				case 1:
					author = strings.ReplaceAll(s.Text(), "作者： ", "")
				case 2:
					raw := s.Text()
					raw = strings.ReplaceAll(raw, "阅读： _showDynClicks(\"wbnews\", ", "")
					raw = strings.ReplaceAll(raw, ")", "")
					r := strings.Split(raw, ", ")
					owner := r[0]
					clickid := r[1]
					clickUrl := "https://news.fzu.edu.cn/system/resource/code/news/click/dynclicks.jsp?clickid=" + clickid + "&owner=" + owner
					req, _ := http.NewRequest("GET", clickUrl, nil)
					resp, _ := client.Do(req)
					clickDoc, _ := goquery.NewDocumentFromReader(resp.Body)
					read_num, _ = strconv.Atoi(clickDoc.Text())
				}
			})
		content := newDoc.Find(".ar_article DIV").Text()
		obj := entity.Fzu{
			Link:         url,
			Author:       author,
			Publish_time: publish_time,
			Title:        title,
			Content:      content,
			Read_num:     read_num,
		}
		result <- obj
	}
}
