package CfgParse

import (
    "fmt"
    "log"
    "os"
    "io"
    "encoding/json"
    "net/http"
    "net/http/httputil"
    // "context"
    // "time"
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


func RunBookmark(bmark Bookmark) {
    client := &http.Client{}
    // ctx := context.Background()
    // ctx, _ = context.WithCancel(ctx)

    // Probably use channels here 
    var req, _ = http.NewRequest(bmark.Method, bmark.Request, nil)
    req.Header.Add("User-Agent", "pmancli/0.0.1")
    for _, header := range bmark.Headers {
        req.Header.Add(header.Name, header.Val)
    }

    dump, _ := httputil.DumpRequestOut(req, false)
    fmt.Println(string(dump))

    res, _ := client.Do(req)
    body, _ := io.ReadAll(res.Body)
    res.Body.Close()
    fmt.Printf("%s\n", body)
}


func AddBookmark(bmark Bookmark, bmarklist []Bookmark, path string) {
    bfile, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer bfile.Close()

    bmarklist = append(bmarklist, bmark)

    bjson, _ := json.Marshal(bmarklist)
    wres, _ := bfile.Write([]byte(bjson))
    fmt.Printf("wres: %v\n", wres)
}
