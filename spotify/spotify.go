package spotify

import (
	"bytes"
	"errors"
	"encoding/base64"
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
	"log"

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

	req.Header.SetMethodBytes([]byte(method))
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.SetBody(reqBody)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if reqError := fasthttp.Do(req, resp); reqError != nil {
		fmt.Printf("API request failed: %s\n", err)
		return apiResponse{url, 0, []byte{}}, reqError
	}

	body := resp.Body()

	status := resp.StatusCode()
	if status != fasthttp.StatusOK {
		return apiResponse{url, status, []byte{}}, errors.New(fmt.Sprintf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode()))
    }

    contentType := resp.Header.Peek("Content-Type")
    if bytes.Index(contentType, []byte("application/json")) != 0 {
        return apiResponse{url, 0, []byte{}}, errors.New(fmt.Sprintf("Expected content type application/json but got %s\n", contentType))
	}


	result = apiResponse{url, status, body}
	return result, nil
}

type SpotifyClient struct {
	ClientId string
	ApiSecret string
	AccessToken string
}

func (c *SpotifyClient) Authenticate() error {
	decodedHeader := c.ClientId + ":" + c.ApiSecret
	encodedHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(decodedHeader))
	body := []byte("grant_type=client_credentials")
	resp, err := apiRequest("https://accounts.spotify.com/api/token", "POST", body, encodedHeader)
	if err != nil {
		return err
	}

	json := string(resp.response)
	var p fastjson.Parser
	v, err := p.Parse(json)
	if err != nil { log.Fatal(err) }
	data := string(v.GetStringBytes("access_token"))
	fmt.Println(data)
	c.AccessToken = data
	return nil
}

func (c SpotifyClient) getAlbum() Album  {
	myAlbum := Album{title: "Madvillainy", artist: "Madvillain", runtime: "46:00", tracks: []Track{}, release: "2004"}
	return myAlbum
}