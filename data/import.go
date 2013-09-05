package data

import (
    "encoding/json"
    "io/ioutil"
)

func ImportComponent(path string, comp interface{}) error {
    bytes, err := ioutil.ReadFile(path)
    if err == nil {
        json.Unmarshal(bytes, comp)
    }
    return err
}
