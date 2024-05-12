package CfgParse

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

func ParseEnvironments(path string) []Environment {
    var result []Environment

    efile, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer efile.Close()

    ebyteValue, _ := io.ReadAll(efile)
    if en := json.Unmarshal(ebyteValue, &result); en != nil {
        if ute, ok := en.(*json.UnmarshalTypeError); ok {
            fmt.Printf("unmarshalTypeError %v - %v - %v\n", ute.Value, ute.Type, ute.Offset)
        } else {
            fmt.Println("Other error:", en)
        }
    }

    fmt.Printf("\nEnvironment read in:\n")
    spew.Dump(result)
    return result
}


func ParseBookmarks(path string) []Bookmark {
    var result []Bookmark

    bfile, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer bfile.Close()

    byteValue, _ := io.ReadAll(bfile)
    if e := json.Unmarshal(byteValue, &result); e != nil {
        if ute, ok := e.(*json.UnmarshalTypeError); ok {
            fmt.Printf("unmarshalTypeError %v - %v - %v\n", ute.Value, ute.Type, ute.Offset)
        } else {
            fmt.Println("Other error:", e)
        }
    }

    fmt.Printf("Bookmark read in:\n")
    spew.Dump(result)
    return result
}
