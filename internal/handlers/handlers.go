package handlers

import (
	"errors"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func indexPath() (string, error) {

	p := "index.html"
	if _, err := os.Stat(p); err == nil {
		return p, nil
	}

	p = filepath.Join("..", "index.html")
	if _, err := os.Stat(p); err == nil {
		return p, nil
	}

	return "", errors.New("index.html not found")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := indexPath()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	p, err := indexPath()
	if err != nil {
		http.Error(w, "index not found", http.StatusInternalServerError)
		return
	}

	if _, err := template.ParseFiles(p); err != nil {
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

	in := strings.TrimSpace(string(data))

	out, err := service.Conv(in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out = strings.TrimSpace(out)

	ext := filepath.Ext(header.Filename)
	name := time.Now().UTC().String() + ext
	name = strings.ReplaceAll(name, ":", "-")

	outFile, err := os.Create("../" + name)
	if err != nil {
		http.Error(w, "cannot create file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	if _, err := outFile.WriteString(out); err != nil {
		http.Error(w, "cannot write file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte(in + "\n" + out))
}
