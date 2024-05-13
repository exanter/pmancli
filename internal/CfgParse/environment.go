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
