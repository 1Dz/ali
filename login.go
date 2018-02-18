package main

import (

	"net/http/cookiejar"
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
	"fmt"
)

func main(){
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		panic(err)
	}
	client := http.Client{Jar: jar}
	resp, err := client.PostForm("https://passport.aliexpress.com/", url.Values{"loginId":{"daiz@bk.ru"},
	"password":{"system"}})
	if err != nil{
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}
