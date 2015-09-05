package heron

import (
    "fmt"
    "bytes"
    "net/http"
    "io/ioutil"
    "path/filepath"
    "html/template"
)

var tplFuncMaps = template.FuncMap {
    "include": func(filename string, data map[string] interface{}) template.HTML {
        var buf bytes.Buffer
        templatePath := filepath.Join(Config.TemplatePath, filename)
        tpl, err := template.ParseFiles(templatePath)
        if err != nil {
            panic(err)
        }
        err = tpl.Execute(&buf, data)
        if err != nil {
            panic(err)
        }
        return template.HTML(buf.Bytes())
    },
}

func parseTemplate(file, baseFile string, data map[string]interface{}) []byte {
    var buf bytes.Buffer
    tpl := template.New(file).Funcs(tplFuncMaps)
    baseTemplatePath := filepath.Join(Config.TemplatePath, baseFile)
    baseBytes, err := ioutil.ReadFile(baseTemplatePath)
    if err != nil {
        panic(err)
    }
    tpl, err = tpl.Parse(string(baseBytes))
    if err != nil {
        panic(err)
    }
    templatePath := filepath.Join(Config.TemplatePath, file)
    tpl, err = tpl.ParseFiles(templatePath)
    if err != nil {
        panic(err)
    }
    err = tpl.Execute(&buf, data)
    if err != nil {
        panic(err)
    }
    return buf.Bytes()
}

func renderTemplate(w http.ResponseWriter, file, baseFile string, data map[string]interface{}) {
    page := parseTemplate(file, baseFile, data)
    fmt.Fprintf(w, string(page))
}