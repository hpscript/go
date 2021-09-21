package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"html/template"
	"log"
	"fmt"
	"html"
)

func apiHandler(w http.ResponseWriter, r *http.Request){

	var code string // 変数として宣言
	switch html.EscapeString(r.URL.Path) {
	case "/btc":
		code = "BTC_JPY" // ビットコイン
	case "/eth":
		code = "ETH_JPY" // イーサリアム
	case "/xrp":
		code = "XRP_JPY" // リップル
	case "/xlm":
		code = "XML_JPY" // ステラルーメン
	case "/mona":
		code = "MONA_JPY" // モナコイン
	default:
		code = "BTC_JPY"
	}

	uri := "https://api.bitflyer.com/v1/getticker?product_code=" + code
	req, _ := http.NewRequest("GET", uri, nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	var ticker Ticker
	json.Unmarshal([]byte(byteArray),&ticker)


	t := template.Must(template.ParseFiles("/home/vagrant/go/src/github.com/me/api/src/bitflyer.html.tpl"))
	if err := t.ExecuteTemplate(w, "bitflyer.html.tpl", fmt.Sprintf("%.0f\n", ticker.Ltp)); err != nil {
		log.Fatal(err)
	}
}


func main() {
	http.HandleFunc("/btc", apiHandler)
	http.HandleFunc("/eth", apiHandler)
	http.HandleFunc("/xrp", apiHandler)
	http.HandleFunc("/xlm", apiHandler)
	http.HandleFunc("/mona", apiHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}


type Ticker struct {
	Code string `json:"product_code"`
	Ltp float64 `json:"ltp"`
}