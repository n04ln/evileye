package wayback

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func NewTarekomi(url string) (string, error) {
	// WARN : 日本語のURLはバグるので英語のURLのみ
	resp, err := http.Get(
		fmt.Sprintf("https://web.archive.org/save/%s", url))
	if err != nil {
		log.Fatal(err)
		return "", errors.New("connot ")
	}

	var loc []string
	var ok bool
	if loc, ok = resp.Header["Content-Location"]; !ok || len(loc) == 0 {
		log.Println(resp.Header)
		log.Fatal(
			fmt.Sprintf("response is ", resp.Status))

		errstr := fmt.Sprintf("status : %v, connot regist request url", resp.Status)

		return "", errors.New(errstr)
	}

	newurl := fmt.Sprintf("http://web.archive.org%v\n", loc[0])

	return newurl, nil
}
