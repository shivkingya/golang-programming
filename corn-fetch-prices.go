package main

import (
	"io/ioutil"
	"fmt"
	"time"
	"net/http"
	"net/url"
	"encoding/json"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "admin"
    dbName := "sampledb"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

func main() {
		fmt.Println("Job Started")
	fetchPriceAndSaveIntoDB()
}

//loop over DB available coins and fetch prices from api https://api.coingecko.com/api/v3/simple/price
//save into table 'prices'
func fetchPriceAndSaveIntoDB(){
	//db connection get
	db := dbConn()
	//loop over available db coins
    selDB, err := db.Query("SELECT id,symbol FROM  dashboard_coin")
    if err != nil {
        panic(err.Error())
	}
	for selDB.Next() {
		var id,symbol string
        err = selDB.Scan(&id,&symbol)
        if err != nil {
            panic(err.Error())
		}
		baseUrl := "https://api.coingecko.com/api/v3/simple/price"
		url, _ := url.Parse(baseUrl)
		queryString := url.Query()
		queryString.Set("ids", id)
		queryString.Set("vs_currencies", "usd")
        // add query to url
        url.RawQuery = queryString.Encode()
	 //fetch prces from api
	response, err := http.Get(url.String())
	if err != nil {
	   panic(err.Error())
   }
   responseData, err := ioutil.ReadAll(response.Body)
   if err != nil {
	   panic(err.Error())
   }

   var p2 interface{}
   json.Unmarshal([]byte(responseData), &p2)
   m := p2.(map[string]interface{})
  // fmt.Println(m)

  var price float64
   for _, element := range m {
	m2 := element.(map[string]interface{})
		for _, name := range m2 {
			fmt.Println("\t", name)
			fmt.Printf("var1 = %T\n", name)
			price = name.(float64)   //type assertion
		}	
   }

   t := time.Now()
    created_dt := t.Format("2006-01-02 15:04:05")
    fmt.Println(created_dt)
	insertCoin, err := db.Prepare("INSERT INTO prices(symbol,id,price,created_dt) VALUES(?,?,?,?)")
    if err != nil {
        panic(err.Error())
    }
    insertCoin.Exec(symbol,id,price,created_dt)
 
 fmt.Println("Insertion done successfully")
 fmt.Println("Updating current_price into Dashboard_coin")
 updateCoin, err := db.Prepare("UPDATE dashboard_coin SET created=?, price_current=? WHERE id=?")
 if err != nil {
	 panic(err.Error())
 }
 updateCoin.Exec(created_dt, price, id)
 fmt.Println(" Done Updating current_price into Dashboard_coin ")
  }
}