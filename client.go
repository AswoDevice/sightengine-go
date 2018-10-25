package sightengine_go

import (
	"strings"
	"net/http"
	"os"
	"bytes"
	"mime/multipart"
	"path/filepath"
	"io"
	"net/url"
	"encoding/json"
)

type Check struct {
	apiUser string
	apiSecret string
	endpoint string
	http *http.Client
	models string
	url string
}

func newCheck(user string, secret string, models []Model) *Check {
	return &Check{
		apiUser: user,
		apiSecret: secret,
		endpoint: "https://api.sightengine.com/",
		http: &http.Client{},
		models: join(models, ","),
		url: "1.0/check.json",
	}
}

func (check *Check) SetUrl(imageUrl string) (*Response, error) {
	var response *Response
	req, err := http.NewRequest("GET", check.getUrl(), nil)
	if err == nil {
		query := check.newQuery(req)
		query.Add("url", imageUrl)

		req.URL.RawQuery = query.Encode()

		response, err = check.request(req)
	}

	return response, err
}

func (check *Check) SetFile(imagePath string) (*Response, error) {
	file, err := os.Open(imagePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("media", filepath.Base(file.Name()))

	if err != nil {
		return nil, err
	}

	io.Copy(part, file)
	writer.Close()

	var response *Response
	req, err := http.NewRequest("POST", check.getUrl(), body)
	if err == nil {
		query := check.newQuery(req)

		req.URL.RawQuery = query.Encode()
		req.Header.Add("Content-Type", writer.FormDataContentType())

		response, err = check.request(req)
	}

	return response, err
}

func (check *Check) request(req *http.Request) (*Response, error) {
	var response Response
	var resp *http.Response
	resp, err := check.http.Do(req)

	if err == nil {
		if resp.StatusCode == http.StatusOK {
			err = json.NewDecoder(resp.Body).Decode(&response)
		}
	}

	return &response, err
}

func (check *Check) newQuery(req *http.Request) url.Values {
	query := req.URL.Query()
	query.Add("api_user", check.apiUser)
	query.Add("api_secret", check.apiSecret)
	query.Add("models", check.models)
	return query
}

func (check *Check) getUrl() string {
	return check.endpoint + check.url
}