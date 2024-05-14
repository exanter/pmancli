package main

import (
    "fmt"
    "github.com/davecgh/go-spew/spew"

    cfgparse "pmancli/internal/CfgParse"
)


func main() {
    var bmark []cfgparse.Bookmark
    var en []cfgparse.Environment

    bmark = cfgparse.ParseBookmarks("./bookmark-defs.json")
    en = cfgparse.ParseEnvironments("./env-defs.json")

    fmt.Printf("Dumping bookmarks...\n")
    spew.Dump(bmark)
    fmt.Printf("\nDumping Environments...\n")
    spew.Dump(en)

    var newbmark cfgparse.Bookmark

    newbmark.Id = "2"
    newbmark.Name = "Second Bookmark"
    newbmark.Request = "https://www.konectauto.com/"
    newbmark.Method = "GET"
    newbmark.Body.Mode = ""

    cfgparse.RunBookmark(bmark[0])

    cfgparse.AddBookmark(newbmark, bmark, "./bookmark-defs.json")
}
