package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
)

func ValidateFileHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, fmt.Sprintf("Requested Method %s is not allowed", r.Method),
			http.StatusMethodNotAllowed)
		return
	}

	multiPartReader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, fmt.Sprintln(err), http.StatusBadRequest)
		return
	}

	for multiPartPart, err := multiPartReader.NextPart(); err == nil; multiPartPart, err = multiPartReader.NextPart() {
		if strings.ToUpper(multiPartPart.FormName()) != "UPLOADFILE" {
			continue
		}

	}
}

type FileInfo struct {
	Filename    string
	Size        int64
	ContentType string
	Contents    string
}

type PostResponse struct {
	StatusCode int
	Filename   string
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

	// url := "http://localhost:8080/api/upload"
	url := "https://luft-184208.appspot.com/api/upload"

	var filename string

	// Remove Windows-style path
	pathAndFilename := strings.Split(fh.Filename, "\\")
	if len(pathAndFilename) > 0 {
		filename = pathAndFilename[len(pathAndFilename)-1]
	}

	// Remove unix path
	pathAndFilename = strings.Split(filename, "/")
	if len(pathAndFilename) > 0 {
		filename = pathAndFilename[len(pathAndFilename)-1]
	}

	invalidFilename := false
	var extStr string

	// Change file extenion to .csv (if not already csv)
	filenameWithoutExt := filename
	extIdx := strings.LastIndexByte(filename, '.')
	if extIdx > 0 {
		extStr = filename[extIdx:]
		filenameWithoutExt = filename[:extIdx]
	}
	if strings.ToUpper(extStr) != ".CSV" {
		extStr = ".csv"
		invalidFilename = true
	}

	// Remove ALL non-letter or digit characters from filename and replace with _
	filenameChars := []byte(filenameWithoutExt)
	for i := 0; i < len(filenameChars); i++ {
		ch := filenameChars[i]
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '.' {
			continue
		}
		filenameChars[i] = '_'
		invalidFilename = true
	}
	if invalidFilename {
		filename = string(filenameChars) + extStr
	}

	fi := FileInfo{

		Filename:    filename,
		Size:        fileSize,
		ContentType: fh.Header["Content-Type"][0],
		Contents:    buf.String(),
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
		Filename:   fh.Filename,
	}

	c <- response
}

func PostFileHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		c := make(chan PostResponse, 10)
		var wg sync.WaitGroup
		var failResponses []string

		r.ParseMultipartForm(32 << 20) // 32MB is the default used by FormFile
		fhs := r.MultipartForm.File["uploadFile"]

		for _, fh := range fhs {
			wg.Add(1)
			go sendFile(fh, c, &wg)
		}

		wg.Wait()
		for i := 0; i < len(fhs); i++ {
			response := <-c
			if response.StatusCode != 200 {
				failResponses = append(failResponses, response.Filename)
			}
		}
		close(c)

		if len(failResponses) > 0 {
			http.Error(w, strings.Join(failResponses, ", "), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Opplasting vellykket!")
	}
}
