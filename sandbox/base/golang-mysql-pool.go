//数据库连接池测试
package main

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "log"
  "net/http"
)

var db *sql.DB

func init() {
  db, _ = sql.Open("mysql", "readonly:readonly@tcp(10.137.255.71:3306)/test?charset=utf8")
  db.SetMaxOpenConns(2000)
  db.SetMaxIdleConns(1000)
  db.Ping()
}

func main() {
  startHttpServer()
}

func startHttpServer() {
  http.HandleFunc("/pool", pool)
  err := http.ListenAndServe(":9090", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

func pool(w http.ResponseWriter, r *http.Request) {
  rows, err := db.Query("SELECT * FROM runtimes")
  defer rows.Close()
  checkErr(err)

  columns, _ := rows.Columns()
  scanArgs := make([]interface{}, len(columns))
  values := make([]interface{}, len(columns))
  for j := range values {
    scanArgs[j] = &values[j]
  }

  record := make(map[string]string)
  for rows.Next() {
    //将行数据保存到record字典
    err = rows.Scan(scanArgs...)
    for i, col := range values {
      if col != nil {
        record[columns[i]] = string(col.([]byte))
      }
    }
  }

  //fmt.Println(record)
  fmt.Println(record)
  // fmt.Fprintln(w, "finish")
  fmt.Fprintln(w, "finish", record)
}

func checkErr(err error) {
  if err != nil {
    fmt.Println(err)
    panic(err)
  }
}
