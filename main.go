package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
    payload := struct {
        FontColor string
        BgColor   string
    }{
        FontColor: GetEnvOrDefault("APP_FONT_COLOR", "white"),
        BgColor:   GetEnvOrDefault("APP_BG_COLOR", "blue"),
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if err := tpl.ExecuteTemplate(w, "index.gohtml", payload); err != nil {
            log.Printf("index template kiszolgálása sikertelen. %v\n", err)
        }
    })

    http.Handle("/favico.ico", http.NotFoundHandler())

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Szerver indítása sikertelen. %v\n", err)
    }
}

func GetEnvOrDefault(env, def string) string {
    if envvar := os.Getenv(env); len(envvar) == 0 {
        return def
    } else {
        return envvar
    }
}
