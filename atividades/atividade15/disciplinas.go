package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type Disciplina struct {
	Titulo    string `json:"titulo"`
	Professor string `json:"professor"`
	Semestre  string `json:"semestre"`
	Alunos    []struct {
		Nome     string    `json:"nome"`
		Notas    []float64 `json:"notas"`
		Media    float64   `json:"media"`
		Situacao string    `json:"situacao"`
	} `json:"alunos"`
}

func main() {
	var disciplina Disciplina

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	file, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	if err != nil {
		log.Fatal("JSON read from file failed: %s", err)
	}

	if err = json.Unmarshal([]byte(file), &disciplina); err != nil {
		log.Fatal("JSON unmarshaling failed: %s", err)
	}

	t := template.Must(template.ParseFiles("disciplinas.tmpl"))

	if err := t.Execute(os.Stdout, disciplina); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%+v\n", disciplina)
}
