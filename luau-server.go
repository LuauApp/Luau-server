package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jkern888/Luau-server/model"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func loginHandler(response http.ResponseWriter, request *http.Request) {
	db, _ := sqlx.Connect("mysql", "luau:test@127.0.0.1/luau")

	request.ParseForm()
	name := request.Form.Get("name")
	pass := request.Form.Get("pass")

	if name {
		row, err := db.Queryx("select * from users where name = '?' and password = '?'", name, pass)

		if err == 0 {
			return
		}
		user := User{}
		row.StructScan(user)

		myFoo := model.User{name, 3}
		b, _ := json.MarshalIndent(myFoo, "", "    ")
		response.Write(b)
	}
}
