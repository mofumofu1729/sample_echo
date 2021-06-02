package main

import (
    "log"
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

type Country struct {
    ID          int     `db:"id"`
    Name        string  `db:"name"`
    Ranking     int     `db:"ranking"`
    GroupName   string  `db:"group_name"`
}

func main() {
    db, err := sql.Open("mysql", "yusuke:password@/worldcup2014")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM countries;")
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    var c Country
    for rows.Next() {
        err := rows.Scan(&c.ID, &c.Name, &c.Ranking, &c.GroupName)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Name: %s, Ranking: %d, GroupId: %s\n", c.ID, c.Name, c.Ranking, c.GroupName)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
}
