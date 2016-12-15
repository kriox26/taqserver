package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/kriox26/taqcompiler"
)

func init() {
	http.HandleFunc("/", homepage)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	var templ = template.Must(template.ParseFiles(filepath.Join("templates", "homepage.html")))
	compiler := taqcompiler.NewCompilerFromString("")
	if r.Method == "POST" {
		if inProgram := r.PostFormValue("inProgram"); inProgram != "" {
			fmt.Println("QUE ESTA PASANDO")
			compiler = taqcompiler.NewCompilerFromString(inProgram)
			compiler.Compile()
		} else {
			inProgramFile, _, err := r.FormFile("inProgramFile")
			defer inProgramFile.Close()
			if err != nil {
				io.WriteString(w, err.Error())
			} else {
				compiler = taqcompiler.NewCompilerFromFile(inProgramFile)
				compiler.Compile()
			}
		}
	}

	if err := templ.Execute(w, compiler); err != nil {
		log.Println("error: ", err)
		return
	}
}
