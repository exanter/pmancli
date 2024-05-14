package CfgParse

import (
    "fmt"
    "log"
    "os"
    "io"
    "encoding/json"
)

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
    return result
}


func AddEnvironment(nenv Environment, envlist []Environment, path string) {
    bfile, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer bfile.Close()

    envlist = append(envlist, nenv)

    benv, _ := json.Marshal(envlist)
    wres, _ := bfile.Write([]byte(benv))
    fmt.Printf("wres (Environment): %v\n", wres)
}
