package main

import (
    "html/template"
    "net/http"
    "strconv"
    "sync"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

var (
    items = make(map[int]string)
    nextID = 1
    mu     sync.Mutex
)

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/create", createHandler)
    http.HandleFunc("/update", updateHandler)
    http.HandleFunc("/delete", deleteHandler)
    http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("HX-Request") == "" {
        tpl.ExecuteTemplate(w, "index.html", items)
    } else {
        tpl.ExecuteTemplate(w, "item_list.html", items)
    }
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        mu.Lock()
        defer mu.Unlock()
        items[nextID] = r.FormValue("item")
        nextID++
        tpl.ExecuteTemplate(w, "item_list.html", items)
    } else {
        tpl.ExecuteTemplate(w, "form.html", map[string]interface{}{
            "Action": "/create",
            "Item":   "",
            "ID":     0,
        })
    }
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.FormValue("id"))
    if r.Method == http.MethodPost {
        mu.Lock()
        defer mu.Unlock()
        items[id] = r.FormValue("item")
        tpl.ExecuteTemplate(w, "item_list.html", items)
    } else {
        mu.Lock()
        item := items[id]
        mu.Unlock()
        tpl.ExecuteTemplate(w, "form.html", map[string]interface{}{
            "Action": "/update",
            "Item":   item,
            "ID":     id,
        })
    }
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.URL.Query().Get("id"))
    mu.Lock()
    delete(items, id)
    mu.Unlock()
    tpl.ExecuteTemplate(w, "item_list.html", items)
}
