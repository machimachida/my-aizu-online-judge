package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ゲーム内の言葉(シーザー暗号)を総当たりで調べるプログラム。たまに辞書サイトに登録されていない言葉があるので、その場合は目視で全件確認する。
func main() {
	f := flag.String("word", "", "processing word")
	nodic := flag.Bool("nodic", false, "don't access dictionary")
	flag.Parse()
	res := make([][]rune, 26)
	var existAns bool

	for i := rune(0); i < 26; i++ {
		word := []rune(*f)
		res[i] = make([]rune, len(word))
		for j := 0; j < len(word); j++ {
			r := word[j]
			r -= i
			if r < 97 {
				r += 25
			}
			res[i][j] = r
		}
		if !*nodic {
			resstr := string(res[i])
			if ok, _ := isExistInDictionary(resstr); ok {
				fmt.Println(i, resstr)
				existAns = true
			}
		}
	}
	if !existAns {
		for i, word := range res {
			fmt.Println(i, string(word))
		}
	}
}

var URL string = "https://api.dictionaryapi.dev/api/v2/entries/en/"

func isExistInDictionary(word string) (bool, error) {
	req, err := http.NewRequest(http.MethodGet, URL+word, nil)
	if err != nil {
		return false, err
	}
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	byte, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var res []map[string]any
	err = json.Unmarshal(byte, &res)
	if err != nil || res == nil {
		return false, err
	}
	return true, nil
}
