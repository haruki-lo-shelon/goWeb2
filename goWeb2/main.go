package main

import (
    "net/http"
    "goWeb2/req_handler" // 実装したHTTPリクエストハンドラパッケージの読み込み
)

func main() {

		// "user-list"へのリクエストを関数で処理する
    http.HandleFunc("/user-list", req_handler.HandlerUserList)

		// "user-form"へのリクエストを関数で処理する
    http.HandleFunc("/user-form", req_handler.HandlerUserForm)

    // "user-confirm"へのリクエストを関数で処理する
    http.HandleFunc("/user-confirm", req_handler.HandlerUserConfirm)

    // "user-registered"へのリクエストを関数で処理する
    http.HandleFunc("/user-registered", req_handler.HandlerUserRegistered)

		// css・js・イメージファイル等の静的ファイル格納パス
    http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset/"))))

    // サーバーを起動
    http.ListenAndServe(":8080", nil)
}
