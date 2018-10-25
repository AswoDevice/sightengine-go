package sightengine_go

import (
	"net/http"
	"os"
	"bytes"
	"mime/multipart"
	"path/filepath"
	"io"
	"net/url"
	"encoding/json"
)

type Client struct {
	apiUser string
	apiSecret string
	endpoint string
	http *http.Client
	url string
}

func New(user string, secret string) *Client {
	return &Client{
		apiUser: user,
		apiSecret: secret,
		endpoint: "https://api.sightengine.com/",
		http: &http.Client{},
		url: "1.0/check.json",
	}
}

func (client *Client) CheckUrl(imageUrl string, models ...Model) (*Response, error) {
	var response *Response
	req, err := http.NewRequest("GET", client.getUrl(), nil)
	if err == nil {
		query := client.newQuery(req, models)
		query.Add("url", imageUrl)

		req.URL.RawQuery = query.Encode()

		response, err = client.request(req)
	}

	return response, err
}

func (client *Client) CheckFile(imagePath string, models ...Model) (*Response, error) {
	file, err := os.Open(imagePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	return client.checkReader(file, filepath.Base(file.Name()), models)
}

func (client *Client) CheckBytes(imageSrc []byte, filename string, models ...Model) (*Response, error) {
	return client.checkReader(bytes.NewReader(imageSrc), filename, models)
}

func (client *Client) checkReader(reader io.Reader, filename string, models []Model) (*Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("media", filename)

	if err != nil {
		return nil, err
	}

	io.Copy(part, reader)
	writer.Close()

	var response *Response
	req, err := http.NewRequest("POST", client.getUrl(), body)
	if err == nil {
		query := client.newQuery(req, models)

		req.URL.RawQuery = query.Encode()
		req.Header.Add("Content-Type", writer.FormDataContentType())

		response, err = client.request(req)
	}

	return response, err
}

func (client *Client) request(req *http.Request) (*Response, error) {
	var response Response
	var resp *http.Response
	resp, err := client.http.Do(req)

	if err == nil {
		if resp.StatusCode == http.StatusOK {
			err = json.NewDecoder(resp.Body).Decode(&response)
		}
	}

	return &response, err
}

func (client *Client) newQuery(req *http.Request, models []Model) url.Values {
	query := req.URL.Query()
	query.Add("api_user", client.apiUser)
	query.Add("api_secret", client.apiSecret)
	query.Add("models", join(models, ","))
	return query
}

func (client *Client) getUrl() string {
	return client.endpoint + client.url
}