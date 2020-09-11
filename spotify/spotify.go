package spotify

import (
	"errors"
	"encoding/base64"
	"fmt"
	"github.com/valyala/fasthttp"
)

type Track struct {
	title string
	artist string
	runtime string
}

type Album struct {
	title string
	artist string
	runtime string
	tracks []Track
	release string
}

type apiResponse struct {
	url string
	status int
	response []byte
}

func apiRequest(url string, method string, reqBody []byte, authHeader string) (result apiResponse, err error)  {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	req.SetBody(reqBody)
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if reqError := fasthttp.Do(req, resp); reqError != nil {
		fmt.Printf("API request failed: %s\n", err)
		return (nil, reqError)
	}

	status := resp.StatusCode()
	if status != fasthttp.StatusOK {
        return (nil, errors.New(fmt.Sprintf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())))
    }

    contentType := resp.Header.Peek("Content-Type")
    if bytes.Index(contentType, []byte("application/json")) != 0 {
        return (nil, errors.New(fmt.Sprintf("Expected content type application/json but got %s\n", contentType)))
	}

	body := resp.Body()
	result := apiResponse{url, status, body}
	return (result, nil)
}

type SpotifyClient struct {
	apiKey string
	apiSecret string
}

func (c SpotifyClient) Authenticate() error {
	decodedHeader := c.apiKey + ":" + c.apiSecret
	encodedHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(decodedHeader))
	body := []byte("grant_type=client_credentials")
	resp, err := apiRequest("https://accounts.spotify.com/api/token", "POST", body, encodedHeader)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return
}

func (c SpotifyClient) getAlbum() Album  {
	myAlbum := Album{title: "Madvillainy", artist: "Madvillain", runtime: "46:00", tracks: [], release: "2004"}
	return myAlbum
}