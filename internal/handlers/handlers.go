package handlers

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	_, err := template.ParseFiles("../index.html")
	if err != nil {
		http.Error(w, "cannot parse template", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "cannot get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "cannot read file", http.StatusBadRequest)
		return
	}

	st, err := service.Conv(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ext := filepath.Ext(header.Filename)
	name := time.Now().UTC().String() + ext
	name = strings.ReplaceAll(name, ":", "-")

	outFile, err := os.Create("../" + name)
	if err != nil {
		http.Error(w, "cannot create file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = outFile.WriteString(st)
	if err != nil {
		http.Error(w, "cannot write file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte(st))
}
