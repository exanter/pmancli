package main

import (
	"fmt"
    "os"
    "log"
    // "database/sql"
	// "github.com/davecgh/go-spew/spew"
    //sqlite "github.com/mattn/go-sqlite3"
	cfgparse "pmancli/internal/CfgParse"
)

type pmanContext struct {
    Bookmarks []cfgparse.Bookmark
    Environments []cfgparse.Environment
}

var configDirBase string = ".pmancli"
var bookmarksDir string = "bookmarks"
var environmentsDir string = "environments"
var configDir string = ""

/* ---
 * Intialize the environtment:
 *  ~/.pmancli
 *      bookmarks
 *          <collectionname>.json
 *      environments
 *          <collectionname>.json
 *      results
 *          queryresults.db
 */
func InitializeEnvironment() {
    configDir = os.Getenv("HOME") + "/" + configDirBase

    _, err := os.Stat(configDir)
    if os.IsNotExist(err) {
        err = os.Mkdir(configDir, 0750)
        if err != nil && !os.IsExist(err) {
            log.Fatal(err)
        }
    }

    bookmarksDir = configDir + "/bookmarks"
    var envdirs = [3]string{bookmarksDir, configDir+"/environments", configDir+"/results"}
    for _, edir := range envdirs {
        _, err := os.Stat(edir)
        if os.IsNotExist(err) {
            err = os.Mkdir(edir, 0750)
            if err != nil && !os.IsExist(err) {
                log.Fatal(err)
            }
        }
    }

    var resultsDb = configDir + "/results/results.db"
    _, err = os.Stat(resultsDb)
    if os.IsNotExist(err) {
        f, ferr := os.Create(resultsDb)
        if ferr != nil {
            log.Fatal(ferr.Error())
        }
        f.Close()
    }

    //sqlfh, _ := sql.Open("sqlite3", resultsDb)
    //defer sqlfh.Close()

    /*
        db, sqlerr := sqlite.Open(resultsDb, ":memory")
        if sqlerr != nil {
            log.Fatal("Failed to create sqlite3 results db: ", sqlerr)
        }    
    */
}

func main() {
    context := pmanContext{}

    InitializeEnvironment()

    context.Environments = cfgparse.ParseEnvironments(configDir+"/"+environmentsDir)
    context.Bookmarks = cfgparse.ParseBookmarks(bookmarksDir)

    fmt.Printf("Environments:\n+%v\n", context.Environments)

    /*
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
    */
}
