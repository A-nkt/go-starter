// このプログラムでは、API Serverを立てる
//　このAPI Serverでは、endpointに対してクエリパラメータvalueの値を２倍して返す。
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	var Iid int
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if len(id) != 0 {
			Iid, _ = strconv.Atoi(id)
			fmt.Println(Iid * 2)
		}
	})
	http.ListenAndServe(":3000", nil)
}
