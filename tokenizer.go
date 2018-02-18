package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

func tokenize(){
	doc, err := goquery.NewDocument("https://tmall.aliexpress.com/item/Asus-X540YA-AMD-15-6-E1-6010-500-2-NoODD-Win10-90NB0CN1-M09280/1000005110967.html?spm=a2g02.9897535.products.5.3bc7fe3caLj1aD&traffic_analysisId=")
	if err != nil{
		panic(err)
	}
	f := doc.Find("#j-sku-discount-price")
	f.Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	/*span := doc.Find("span")
	span.Each(func(i int, s *goquery.Selection) {
		if s.HasClass("p-price"){
			fmt.Println(s.Text())
		}
	})*/
}

