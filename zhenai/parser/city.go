package parser

import (
	"go-crawl/engine"
	"go-crawl/model"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const lastPageItemSelection string = ".paging-item:last-child>a"
const itemType string = "zhenai"

func ParseCityAndProfile(doc *goquery.Document) engine.ParseResult {
	result := engine.ParseResult{}

	doc.Find(".content>table>tbody").Each(func(i int, s *goquery.Selection) {
		var item engine.Item
		var profile model.Profile

		if url, ok := s.Find("tr>th>a").Attr(htmlHrefAttr); ok {
			strs := strings.Split(url, "/")
			item.Id = strs[len(strs)-1]
			item.Url = url
		}
		item.Type = itemType

		nickName := s.Find("tr>th>a").Text()
		profile.Name = nickName
		s.Find("tr>td").Each(func(i int, s *goquery.Selection) {
			strs := strings.Split(s.Text(), "：")
			strs[0] = strings.Replace(strs[0], " ", "", -1)

			switch strs[0] {
			case "性别":
				profile.Gender = strs[1]
			case "居住地":
				profile.House = strs[1]
			case "年龄":
				profile.Age, _ = strconv.Atoi(strs[1])
			case "婚况":
				profile.Marriage = strs[1]
			default:
				if ok, _ := regexp.Match("月.*薪", []byte(strs[0])); ok {
					profile.Income = strs[1]
				} else if ok, _ := regexp.Match("学.*历", []byte(strs[0])); ok {
					profile.Education = strs[1]
				} else if ok, _ := regexp.Match("身.*高", []byte(strs[0])); ok {
					profile.Height, _ = strconv.Atoi(strs[1])
				}
			}
		})
		item.Payload = profile
		result.Items = append(result.Items, item)
	})

	doc.Find(lastPageItemSelection).Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr(htmlHrefAttr)
		if !ok || s.Text() != "下一页" {
			return
		}
		result.Requests = append(result.Requests, engine.Request{Url: url, ParserFunc: ParseCityAndProfile})
	})
	return result
}
