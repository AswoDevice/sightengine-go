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
	models string
	url string
}

func New(user string, secret string, models ...Model) *Client {
	return &Client{
		apiUser: user,
		apiSecret: secret,
		endpoint: "https://api.sightengine.com/",
		http: &http.Client{},
		models: join(models, ","),
		url: "1.0/check.json",
	}
}

func (client *Client) CheckUrl(imageUrl string) (*Response, error) {
	var response *Response
	req, err := http.NewRequest("GET", client.getUrl(), nil)
	if err == nil {
		query := client.newQuery(req)
		query.Add("url", imageUrl)

		req.URL.RawQuery = query.Encode()

		response, err = client.request(req)
	}

	return response, err
}

func (client *Client) CheckFile(imagePath string) (*Response, error) {
	file, err := os.Open(imagePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	return client.checkReader(file, filepath.Base(file.Name()))
}

func (client *Client) CheckBytes(imageSrc []byte, filename string) (*Response, error) {
	return client.checkReader(bytes.NewReader(imageSrc), filename)
}

func (client *Client) checkReader(reader io.Reader, filename string) (*Response, error) {
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
		query := client.newQuery(req)

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

func (client *Client) newQuery(req *http.Request) url.Values {
	query := req.URL.Query()
	query.Add("api_user", client.apiUser)
	query.Add("api_secret", client.apiSecret)
	query.Add("models", client.models)
	return query
}

func (client *Client) getUrl() string {
	return client.endpoint + client.url
}