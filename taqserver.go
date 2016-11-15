package main

import (
	"flag"
	"html/template"
	"io"
	"net/http"
	"path/filepath"

	"github.com/kriox26/taqcompiler"
)

var port = flag.String("port", ":8080", "port where the server runs")
var templ = template.Must(template.ParseFiles(filepath.Join("templates", "homepage.html")))

func main() {
	flag.Parse()

	http.HandleFunc("/", homepage)

	http.ListenAndServe(*port, nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	var compiler *taqcompiler.Compiler
	if r.Method == "POST" {
		if inProgram := r.PostFormValue("inProgram"); inProgram != "" {
			compiler = taqcompiler.NewCompilerFromString(inProgram)
			compiler.Compile()
		} else {
			inProgramFile, _, err := r.FormFile("inProgramFile")
			if err != nil {
				io.WriteString(w, err.Error())
			} else {
				compiler = taqcompiler.NewCompilerFromFile(inProgramFile)
				compiler.Compile()
			}
		}
	}

	templ.Execute(w, compiler)
}
