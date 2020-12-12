package main

import "net/http"

func main() {

	http.HandleFunc("/query/v1", func(res http.ResponseWriter, req *http.Request) {
		//var userinfo = `{"name":"dalong","age":333}`
		var userinfo2 = `[{"name":"dalong","age":333},{"name":"dalong2","age":33}]`
		res.Header().Add("Content-Type", "application/json")
		res.Write([]byte(userinfo2))
	})
	http.ListenAndServe(":8090", nil)
}
