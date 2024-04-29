package main

import (  "net/http"

// 	"database/sql"
//   "fmt"
//   "log"
//   "os"

//   "github.com/gorilla/mux"
//   _ "github.com/go-sql-driver/mysql"
//   "github.com/joho/godotenv"
//    "encoding/json"
)

func main(){
	http.ListenAndServe(":9090", nil)
}