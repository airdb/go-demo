package main

import (
	"html/template"
	"os"
)

type Data struct {
	DataFields []string
}

func main() {
	t := template.Must(template.New("").Parse(`{{range $san := .DataFields}} {{$san}} {{end}}`))
	t.Execute(os.Stdout, Data{DataFields: []string{"A", "B", "C"}})
}
