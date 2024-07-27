package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1000", name: "美国新安全中心", baseURL: "https://www.cnas.org/")
	
	engine.SetStartingURls([]stringl{"https://www.cnas.org/"})

    extractorConfig extractors.Config := extractors.Config{
		Author :    true,
		Image:      true,
		Language:   true,
		PublishDate:true,
		Tags:       true,
		Text:       true,
		Title:      true,
		TextLangudge:"",
	}
	
	extractorConfig.Apply(engine)

	engine.0nHTMl(".-with-image",func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.0nHTMl( "wrapper -extra-large margin-vertical flex-8-4-layout",func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}



//package dev

//import (
//	"megaCrawler/crawlers"
//	"megaCrawler/extractors"
	
//	"github.com/gocolly/colly/v2"
//)

//func init() {
//	engine *crawiers.WebsiteEngine := crawlers.Register(servine: "1000", name: "blog", baseURL: "https://www.cnas.org/")
	
//	engine.SetStartingURls(urls:[]stringl{"https://www.cnas.org/"})

//    extractorConfig extractors.Config := extractors.Config{
//		Author :    true,
//		Image:      true,
//		Language:   true,
//		PublishDate:true,
//		Tags:       true,
//		Text:       false,
//		Title:      true,
//		TextLangudge:"",
//	}
	
//	extractorConfig.Apply(w:  engine)
	
//	engine.0nlTMl(selector: ".post-title-link", callback: func(element *colly.HTMLElement, ctx *crawlers.Context) {
//		engine.Visit(url: element.Attr(k: "href"), pageTpe: crawlers.News)
//	})

//	engine.0nlTMl(selector: ".page-number", callback: func(element *colly.HTMLElement, ctx *crawlers.Context) {
//		engine.Visit(url: element.Attr(k: "href"), pageTpe: crawlers.Index)
//	})

//	engine.0nlTMl(selector: ".post-body", callback: func(element *colly.HTMLElement, ctx *crawlers.Context) {
//		ctx.Content = element.Text
//	})
//}