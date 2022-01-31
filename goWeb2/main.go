package main

import (
		"fmt"
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

		// "user-edit"へのリクエストを関数で処理する
    http.HandleFunc("/user-edit", req_handler.HandlerUserEdit)

    // "user-update"へのリクエストを関数で処理する
    http.HandleFunc("/user-update", req_handler.HandlerUserUpdate)

		// css・js・イメージファイル等の静的ファイル格納パス
    http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset/"))))

    // サーバーを起動
    //http.ListenAndServe(":8080", nil)
    err := http.ListenAndServeTLS(":10443", "crt/localhost.crt", "crt/localhost.key", nil)
    if err != nil {
        fmt.Printf("ERROR : %s", err)
    }
}
