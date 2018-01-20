package lib

import (
	"time"
	"bytes"
	"net/http"
	"log"
)

func GetJson(url string) ([]byte, error) {
	log.Println("Url is ", url)
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	return buf.Bytes(), err
}