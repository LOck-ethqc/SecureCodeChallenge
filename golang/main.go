package main

import (
    "fmt"
    "net/http"
    "path/filepath"
    "strings"
)

func main() {
    http.HandleFunc("/download", downloadHandler)
    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}

func sanitizePath(filename string) string {
    // If the filename still contains "../", continue sanitizing
    if strings.Contains(filename, "../") {
        filename = strings.Replace(filename, "../", "", -1)
        return sanitizePath(filename) // Recursively call sanitizePath
    }
    return filename
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
    // Gets the query parameter value
    filename := r.URL.Query().Get("filename")
    // Checks if the query parameter value was empty(No filename requested)
    if filename == "" {
        http.Error(w, "Missing filename parameter", http.StatusBadRequest)
        return
    }
    fmt.Println("Query request:", filename) 

    // Preventing Path Traversal
    // -1 means that it ensures that it replaces all occurances of "../"
    sanitizedFilename := sanitizePath(filename)
    cleanedFilename := filepath.Clean(sanitizedFilename)
    baseDir := "public/files/"
    fullPath := filepath.Join(baseDir, cleanedFilename)
    // /public/files/<The requested file name>
    fmt.Println("Query Request (POST-sanitization):", fullPath)

    w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", filepath.Base(fullPath)))
    w.Header().Set("Content-Type", "text/plain")
    http.ServeFile(w, r, fullPath)
}

