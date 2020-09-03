package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
	"os"
	"encoding/json"
)

type Coin struct {
	Id string
	Symbol string
	Name string
  }

  var coin []Coin

func main() {
    response, err := http.Get("https://api.coingecko.com/api/v3/coins/list")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
	}
	
	json.Unmarshal([]byte(responseData), &coin)
   var i int
	for i := 0; i < len(coin); i++ {
		fmt.Println("[",coin[i].Id," , ",coin[i].Symbol," , ",coin[i].Name," ]")
		if coin[i].Symbol == "tiox" {
             insert(coin[i].Id, coin[i].Symbol, coin[i].Name)
             break
		}
      }

      if i<len(coin){
        fmt.Println("Insertion done successfully")
     }else{
       fmt.Println("Coin not found")
     }    
}


