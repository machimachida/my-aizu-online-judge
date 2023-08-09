package main

import (
	"fmt"
	"os"
	"practice-go/design-pattern/abstract-factory/factory"
	"practice-go/design-pattern/abstract-factory/listfactory"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ./design-pattern/abstract-factory/main.go FactoryName")
		fmt.Println("Usage: go run ./design-pattern/abstract-factory/main.go ListFactory")
		fmt.Println("Usage: go run ./design-pattern/abstract-factory/main.go TableFactory")
	}

	var fac factory.Factory
	switch os.Args[1] {
	case "ListFactory":
		fac = listfactory.ListFactory{}
	case "TableFactory":
	default:
		fmt.Println("factory " + os.Args[1] + " is not found")
		return
	}

	asahi := fac.CreateLink("朝日新聞", "http://www.asahi.com/")
	yomiuri := fac.CreateLink("読売新聞", "http://www.yomiuri.co.jp/")
	usYahoo := fac.CreateLink("Yahoo!", "http://www.yahoo.com/")
	jpYahoo := fac.CreateLink("Yahoo!Japan", "http://www.yahoo.co.jp/")
	excite := fac.CreateLink("Excite", "http://www.excite.com/")
	google := fac.CreateLink("Google", "http://www.google.com/")

	trayNews := fac.CreateTray("新聞")
	trayNews.Add(asahi)
	trayNews.Add(yomiuri)

	trayYahoo := fac.CreateTray("Yahoo!")
	trayYahoo.Add(usYahoo)
	trayYahoo.Add(jpYahoo)

	traySearch := fac.CreateTray("サーチエンジン")
	traySearch.Add(trayYahoo)
	traySearch.Add(excite)
	traySearch.Add(google)

	page := fac.CreatePage("Linkage", "結城 浩")
	page.Add(trayNews)
	page.Add(trayYahoo)
	page.Add(traySearch)
	page.Output()
}
