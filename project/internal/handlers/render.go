package handlers

import (
    "html/template"
    "net/http"
    "path/filepath"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    layouts := filepath.Join("web", "templates", "layout", "*.html")
    content := filepath.Join("web", "templates", tmpl)

    // Парсим все файлы макета и конкретный шаблон
    files, err := filepath.Glob(layouts)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    files = append(files, content)

    ts, err := template.ParseFiles(files...)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = ts.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
