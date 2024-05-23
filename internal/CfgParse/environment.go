package CfgParse

import (
    "fmt"
    "log"
    "os"
    "io"
    //"encoding/json"
    "gopkg.in/yaml.v3"
)

func ParseEnvironments(envdir string) []Environment {
    var result []Environment

    entries, err := os.ReadDir(envdir)
    if err != nil {
        log.Fatal(err)
    }

    for _, e := range entries {
        var envResult Environment

        path := envdir + "/" + e.Name()
        efile, err := os.Open(path)
        if err != nil {
            log.Fatal(err)
        }
        defer efile.Close()

        ebyteValue, _ := io.ReadAll(efile)
        if en := yaml.Unmarshal(ebyteValue, &envResult); en != nil {
            fmt.Println("Other error:", en)
        }

        result = append(result, envResult)
    }

    return result
}


func AddEnvironment(nenv Environment, envlist []Environment, path string) {
    bfile, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer bfile.Close()

    envlist = append(envlist, nenv)

    benv, _ := yaml.Marshal(envlist)
    wres, _ := bfile.Write([]byte(benv))
    fmt.Printf("wres (Environment): %v\n", wres)
}
