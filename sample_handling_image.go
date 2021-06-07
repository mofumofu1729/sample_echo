package main

import (
    "fmt"
    "os"
    "image/png"
    "flag"
)

func main() {
    flag.Parse()
    path := flag.Args()[0]

    fmt.Printf("open: %s\n", path)
    f, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    img, err := png.Decode(f)
    if err != nil {
        panic(err)
    }

    bound := img.Bounds()
    fmt.Printf("x : %d px, Y : %d px\n", bound.Dx(), bound.Dy())
}
