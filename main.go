package main

import (
	"net/http"
	"os"
	"fmt"
	"html/template"
	//"github.com/path/to/timestamp"
	_"github.com/golang/protobuf/ptypes/timestamp"
)

type Coin struct {
	Symbol string
	Id string
	Name string
	Created  string
  }

  var templates=template.Must(template.ParseFiles("searchCoin.html","addCoin.html","Index.html"))

func searchCoinHandler(w http.ResponseWriter, r *http.Request) {
       err := templates.ExecuteTemplate(w,"searchCoin.html",map[string] string{"Title":"Search Coin Page"})
       if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addCoinHandler(w http.ResponseWriter, r *http.Request) {
       err := templates.ExecuteTemplate(w,"addCoin.html",map[string] string{"Title":"Add Coin Page"})
       if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getSavedCoinHandler(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
    selDB, err := db.Query("SELECT * FROM  dashboard_coin")
    if err != nil {
        panic(err.Error())
    }
    coin := Coin{}
	res := []Coin{}
	//var res[] Coin
    for selDB.Next() {
		var name, symbol,id string
		var created string
        err = selDB.Scan(&symbol, &id, &name, &created)
        if err != nil {
            panic(err.Error())
		}
		coin.Symbol = symbol
        coin.Id = id
        coin.Name = name
        coin.Created = created
		res = append(res, coin)
	}
	fmt.Println(res)
    templates.ExecuteTemplate(w, "Index.html",res)
	defer db.Close();
}

  func buildPage(w http.ResponseWriter, r *http.Request){
	 tmpl := template.Must(template.ParseFiles("main.html","navigation.html"))
     tmpl.Execute(w,map[string] string{"Title":"Main Page","nav":"Navigation Page","action":"action"})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", buildPage)
	mux.HandleFunc("/addCoin/",addCoinHandler)
	mux.HandleFunc("/searchCoin/",searchCoinHandler)
	mux.HandleFunc("/getSavedCoin/",getSavedCoinHandler)
	http.ListenAndServe(":"+port, mux)
}