package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/axgle/mahonia"
)

var workResultLock sync.WaitGroup

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func downloadImg(requestUrl string, name string, dirPath string) {
	image, err := http.Get(requestUrl)
	check(err)
	imageByte, err := ioutil.ReadAll(image.Body)
	fmt.Println(name)
	defer image.Body.Close()
	filePath := filepath.Join(dirPath, name+".jpg")
	err = ioutil.WriteFile(filePath, imageByte, 0644)
	check(err)
	fmt.Println(requestUrl + "\t下载成功")
}

func spider(i int, dirPath string) {
	defer workResultLock.Done()
	url := fmt.Sprintf("http://www.xiaohuar.com/daxue/%d.html", i)
	response, err2 := http.Get(url)
	check(err2)
	content, err3 := ioutil.ReadAll(response.Body)
	check(err3)
	defer response.Body.Close()
	html := string(content)
	html = mahonia.NewDecoder("utf-8").ConvertString(html)
	match := regexp.MustCompile(`<img .*alt="(.*?)".*src="(.*?)" style="width: 550px; .*?>`)
	matchedStr := match.FindAllString(html, -1)
	for _, matchStr := range matchedStr {
		var imgUrl string
		name := match.FindStringSubmatch(matchStr)[1]
		src := match.FindStringSubmatch(matchStr)[2]
		if strings.HasPrefix(src, "http") != true {
			var buffer bytes.Buffer
			buffer.WriteString("http://www.xiaohuar.com")
			buffer.WriteString(src)
			imgUrl = buffer.String()
		} else {
			imgUrl = src
		}
		downloadImg(imgUrl, name, dirPath)
	}
}

func main() {
	start := time.Now()
	dir := filepath.Dir(os.Args[0])
	dirPath := filepath.Join(dir, "images")
	err1 := os.MkdirAll(dirPath, os.ModePerm)
	check(err1)
	//spider(6,dir_path)
	for i := 170; i < 190; i++ {
		workResultLock.Add(1)
		go spider(i, dirPath)
	}
	workResultLock.Wait()
	fmt.Println(time.Now().Sub(start))
}
