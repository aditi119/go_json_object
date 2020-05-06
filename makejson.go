package main

import ( 
          "fmt"
          "encoding/json"
)


func main(){

var n,add string

fmt.Print("\nEnter person name:")
fmt.Scan(&n)
fmt.Print("\nEnter address:")
fmt.Scan(&add)
mymap := map[string]string { "name":n , "address" :add }

barr , err := json.Marshal(mymap)
if err != nil {
		fmt.Println("error:", err)
	}

  fmt.Println(string(barr))

}