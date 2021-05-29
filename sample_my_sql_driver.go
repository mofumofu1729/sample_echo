package main

import (
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

type Country struct {
    ID          int     `db:"id"`
    Name        string  `db:"name"`
    Ranking     int     `db:"ranking"`
    GroupName   string  `db:"group_name"`
}

type CountryList []Country

func main() {
    var countryList CountryList

    db, err := sqlx.Open("mysql", "yusuke:password@/worldcup2014")
    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Queryx("SELECT * FROM countries")
    if err != nil {
        log.Fatal(err)
    }

    var country Country
    for rows.Next() {
        err := rows.StructScan(&country)
        if err != nil {
            log.Fatal(err)
        }
        countryList = append(countryList, country)
    }

    for _, country = range countryList {
        fmt.Println(country)
    }
}
