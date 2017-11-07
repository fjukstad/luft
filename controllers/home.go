package controllers

import (
	"fmt"
	"net/http"
	"text/template"
	"encoding/json"
	"bytes"
)


var indexTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html",
	"views/index.html", "views/footer.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := indexTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

var liveTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html",
	"views/live.html", "views/footer.html"))

func LiveHandler(w http.ResponseWriter, r *http.Request) {
	err := liveTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

var historyTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html",
	"views/history.html", "views/footer.html"))

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	err := historyTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

var uploadTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html",
	"views/upload.html", "views/footer.html"))

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := uploadTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

type FileInfo struct {
	Filename 		string
	Size 				int64
	ContentType string
	Contents 		string
}


func PostFileHandler(w http.ResponseWriter, r *http.Request) {
	// _, header, err := r.FormFile("uploadFile")
	file, header, err := r.FormFile("uploadFile")
	if err != nil {
		fmt.Println(err)
	}

	var buf bytes.Buffer
	fileSize, err := buf.ReadFrom(file)
	if err != nil {
		fmt.Println(err)
	}

	url := "http://localhost:8080/api/upload"
	
	fi := FileInfo {
		Filename: header.Filename,
		Size: fileSize,
		ContentType: header.Header["Content-Type"][0], 
		Contents: buf.String(),
	}

	f, err := json.Marshal(fi)
	if err != nil {
		fmt.Println(err)
	}


	req, err := http.NewRequest("POST", url, bytes.NewBuffer(f))
  req.Header.Set("X-Custom-Header", "myvalue")
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()

	http.Redirect(w,r, "/alert", http.StatusFound)
}

var alertTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html",
	"views/alert.html", "views/footer.html"))

func AlertHandler(w http.ResponseWriter, r *http.Request) {
	err := alertTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
