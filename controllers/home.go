package controllers

import (
	"fmt"
	"net/http"
	"text/template"
	"encoding/json"
	"bytes"
	"mime/multipart"
	"sync"
	// "strings"
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

type PostResponse struct {
	StatusCode 	int
	Filename 	 	string
}

func sendFile(fh *multipart.FileHeader, c chan PostResponse, wg *sync.WaitGroup) {

	defer wg.Done()
	f, err := fh.Open()
  if err != nil {
		fmt.Println(err)
	}
  
	var buf bytes.Buffer
	fileSize, err := buf.ReadFrom(f)
	if err != nil {
		fmt.Println(err)
	}

	url := "http://localhost:8080/api/upload"
	
	fi := FileInfo {
		Filename: fh.Filename,
		Size: fileSize,
		ContentType: fh.Header["Content-Type"][0], 
		Contents: buf.String(),
	}

	file, err := json.Marshal(fi)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(file))
  req.Header.Set("X-Custom-Header", "myvalue")
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
	
	response := PostResponse{
		StatusCode: resp.StatusCode,
		Filename: fh.Filename,
	}

	c <- response
}



func PostFileHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		c := make(chan PostResponse, 10)
		var wg sync.WaitGroup
		var succResponses []string
		var failResponses []string

		r.ParseMultipartForm(32 << 20) // 32MB is the default used by FormFile
		fhs := r.MultipartForm.File["uploadFile"]

		for _, fh := range fhs {
			wg.Add(1)
			go sendFile(fh, c, &wg)
		}

		wg.Wait()
		for i := 0; i < len(fhs); i++ {
			response := <- c
			if (response.StatusCode == 200) {
				succResponses = append(succResponses, response.Filename)
			} else {
				failResponses = append(failResponses, response.Filename)
			}
		}
		close(c)
		
		if (len(failResponses) == 0) {
			fmt.Fprintf(w, "Opplasting vellykket!")
			return
		}	else {
			fmt.Fprintf(w, "Følgene filer ble ikke lastet opp: ")
			return
		}
 	}
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
