// Пакет handlers реализует обработчики мультиплексора.

package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Ошибка при загрузке шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Ошибка при отображении страницы: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error retrieving file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read the data from the file
	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fileContent := string(fileData)
	convertedContent, err := service.DetectAndConvert(fileContent)
	if err != nil {
		http.Error(w, fmt.Sprintf("Input: %s. Error converting content: %v ", fileContent, err.Error()), http.StatusOK)
		return
	}

	// Generate a filename based on current UTC time
	fileExt := filepath.Ext(handler.Filename)
	timestamp := time.Now().UTC().String()
	newFilename := timestamp + fileExt

	// Create a local file
	outputFile, err := os.Create(newFilename)
	if err != nil {
		http.Error(w, "Error creating output file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Write the converted content to the file
	_, err = outputFile.WriteString(convertedContent)
	if err != nil {
		http.Error(w, "Error writing to output file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Input: %s\nConversion result:\n%s\n\nSaved to file: %s", fileContent, convertedContent, newFilename)
}
