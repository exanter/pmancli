package CfgParse

import (
    "fmt"
    "log"
    "os"
    "io"
    "encoding/json"
)

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
    return result
}
