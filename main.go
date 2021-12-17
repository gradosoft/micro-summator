package main

import (
    "net/http"
    "encoding/json" 
    "strconv"
    "strings"
    "log" 
)

func Handler(wr http.ResponseWriter, req *http.Request) {
    
    type Data struct{
        Status string
        Result int
    }

    res := Data{"OK", 0}

    req.ParseForm()
    for _, v := range req.Form {
        if val, err := strconv.Atoi(strings.Join(v," ")); err != nil{ 
            res.Status = "Error" //Int only
            res.Result = 0
            break
        } else {
            res.Result += val
        }
    }

    wr.Header().Set("Content-Type", "application/json")
    wr.WriteHeader(http.StatusOK)
    json.NewEncoder(wr).Encode(res)
}

func main() {
    http.HandleFunc("/", Handler)
    log.Fatal(http.ListenAndServe("localhost:8020", nil)) 
}