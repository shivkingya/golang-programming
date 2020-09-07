package main

import (
	"net/http"
	_"os"
	"fmt"
	"html/template"
	"github.com/gorilla/mux"
	_"github.com/golang/protobuf/ptypes/timestamp"
)

type Coin struct {
	Symbol string
	Id string
//Name string
    Price string
	Created  string
	Price_current string
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
    selDB, err := db.Query("SELECT * FROM  prices")
    if err != nil {
        panic(err.Error())
    }
    coin := Coin{}
	res := []Coin{}
	//var res[] Coin
    for selDB.Next() {
		var price, symbol,id string
		var created,price_current string
        err = selDB.Scan(&symbol, &id, &price, &created, &price_current)
        if err != nil {
            panic(err.Error())
		}
		coin.Symbol = symbol
        coin.Id = id
        coin.Price = price
		coin.Created = created
		coin.Price_current = price_current
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
	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = "3000"
	//}
	router := mux.NewRouter()
	router.HandleFunc("/", buildPage)
	router.HandleFunc("/addCoin/",addCoinHandler)
	router.HandleFunc("/searchCoin/",searchCoinHandler)
	router.HandleFunc("/getSavedCoin/",getSavedCoinHandler)
	http.ListenAndServe(":8081", router)
}