package main

import (
    "fmt"
    "log"
    "os"
    "io"
    "encoding/json"

    "github.com/davecgh/go-spew/spew"
)

type Header struct {
    Name string `json:"name"`
    Val string `json:"val"`
}

type URLEncoded struct {
    Var string `json:"var"`
    Val string `json:"val"`
}

type BookmarkBody struct {
    Mode string `json:"mode"`
    Urlencoded []URLEncoded `json:"urlencoded"`
    Jsonencoded string `json:"jsonencoded"`
}

type Bookmark struct {
    Id string `json:"id"` 
    Name string `json:"name"`
    Request string `json:"request"`
    Method string `json:"method"`
    Headers []Header `json:"headers"` 
    Body BookmarkBody `json:"body"`
}

type KeyVal struct {
    Key string
    Val string
}

type Environment struct {
    Name string `json:"Name"`
    Vars map[string]string  `json:"vars"`
}

func main() {
    var bmark Bookmark
    var en Environment

    bfile, err := os.Open("./bookmark-defs.json")
    if err != nil {
        log.Fatal(err)
    }
    defer bfile.Close()

    byteValue, _ := io.ReadAll(bfile)
    if e := json.Unmarshal(byteValue, &bmark); e != nil {
        if ute, ok := e.(*json.UnmarshalTypeError); ok {
            fmt.Printf("unmarshalTypeError %v - %v - %v\n", ute.Value, ute.Type, ute.Offset)
        } else {
            fmt.Println("Other error:", e)
        }
    }

    fmt.Printf("Bookmark read in:\n")
    spew.Dump(bmark)

    efile, err := os.Open("./env-defs.json")
    if err != nil {
        log.Fatal(err)
    }
    defer efile.Close()

    ebyteValue, _ := io.ReadAll(efile)
    if er := json.Unmarshal(ebyteValue, &en); er != nil {
        if ute, ok := er.(*json.UnmarshalTypeError); ok {
            fmt.Printf("unmarshalTypeError %v - %v - %v\n", ute.Value, ute.Type, ute.Offset)
        } else {
            fmt.Println("Other error:", er)
        }
    }

    fmt.Printf("\nEnvironment read in:\n")
    spew.Dump(en)
}

