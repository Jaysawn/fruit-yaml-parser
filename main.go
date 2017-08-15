package main

import (
    "fmt"
    "log"
    "io/ioutil"

    "github.com/ghodss/yaml"
)

//yaml tags are superflous since names are the same
type Box struct {
    Id int `yaml:"id"`
    Apples int `yaml:"Apples"`
    Oranges int `json:"Oranges"`
}

type Crate struct {
    Id int
    Box Box
}

func main() {
    crate := Crate{}

    yamlFile, err := ioutil.ReadFile("fruit.yaml")
    if err != nil {
        log.Printf("yamlFile open error: $%v.", err)
    }

    err = yaml.Unmarshal(yamlFile, &crate)
    if err != nil {
        log.Fatalf("yamlFile marshal error: %v", err)
    }

    fmt.Println("Crate Id: ", crate.Id)
    fmt.Println("Box Id:", crate.Box.Id)
    fmt.Println("Box Apple Count", crate.Box.Apples)
    fmt.Println("Box Orange Count:", crate.Box.Oranges)
}
