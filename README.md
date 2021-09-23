# About
Goで bitFlyer Lightningのpublic APIから、ビットコイン(BTC)、イーサリアム(ETH)、リップル(XRP)、ステラルーメン(MXL)、モナコイン(MONA)のTickerの値を取得して、HTMLのテンプレートで表示するソースコードです。

bitFlyer Lightning Document<br>
https://lightning.bitflyer.com/docs?lang=ja

Ticker json file(BTC_JPY)<br>
https://lightning.bitflyer.com/v1/ticker

youtube解説
https://www.youtube.com/watch?v=bTugqJgSCc8&t=88s


# Quick Run
Go言語をインストールしてください。<br>
$ sudo apt install golang-go

実行<br>
$ go run main.go

※main.goの44行目のtemplate file(html)の絶対パスは適時置き換えて使用ください。<br>
※独自ドメインを取得して外部公開する際は、80番ポートで起動します。<br>
```
log.Fatal(http.ListenAndServe(":80",nil))
```
