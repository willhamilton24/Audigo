package spotify

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/url"
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

type SpotifyAlbum struct {
	AlbumType string `json:"album_type"`
	Artists   []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	Copyrights       []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"copyrights"`
	ExternalIds struct {
		Upc string `json:"upc"`
	} `json:"external_ids"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Genres []interface{} `json:"genres"`
	Href   string        `json:"href"`
	ID     string        `json:"id"`
	Images []struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"images"`
	Label                string `json:"label"`
	Name                 string `json:"name"`
	Popularity           int64  `json:"popularity"`
	ReleaseDate          string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	TotalTracks          int64  `json:"total_tracks"`
	Tracks               struct {
		Href  string `json:"href"`
		Items []struct {
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			DiscNumber       int64    `json:"disc_number"`
			DurationMs       int64    `json:"duration_ms"`
			Explicit         bool     `json:"explicit"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href        string `json:"href"`
			ID          string `json:"id"`
			IsLocal     bool   `json:"is_local"`
			Name        string `json:"name"`
			PreviewURL  string `json:"preview_url"`
			TrackNumber int64  `json:"track_number"`
			Type        string `json:"type"`
			URI         string `json:"uri"`
		} `json:"items"`
		Limit    int64       `json:"limit"`
		Next     interface{} `json:"next"`
		Offset   int64       `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int64       `json:"total"`
	} `json:"tracks"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type SpotifyAlbums struct {
	Albums []SpotifyAlbum
}


type apiResponse struct {
	url string
	status int
	response []byte
}

func apiRequest(url string, method string, reqBody []byte, authHeader string) (result apiResponse, e error)  {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)

	req.Header.SetMethodBytes([]byte(method))
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if method == "POST" {
		req.SetBody(reqBody)
	}


	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if reqError := fasthttp.Do(req, resp); reqError != nil {
		fmt.Printf("API request failed: %s\n", reqError)
		return apiResponse{url, 0, []byte{}}, reqError
	}

	body := resp.Body()
	fmt.Println(resp)

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

func makeQueryString(params map[string]string) string {
	queryString := "?"
	for key, value := range params {
		queryString += key +"="+value+"&"
	}
	return queryString[0:len(queryString)-2]
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
	if err != nil {
		log.Fatal(err)
		return err
	}
	data := string(v.GetStringBytes("access_token"))
	fmt.Println(data)
	c.AccessToken = "Bearer " + data
	return nil
}

//
// TRACKS
//

type SpotifyTracks struct {
	Href  string `json:"href"`
	Items []struct {
	Artists []struct {
	ExternalUrls struct {
	Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href string `json:"href"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	URI  string `json:"uri"`
	} `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	DiscNumber       int64    `json:"disc_number"`
	DurationMs       int64    `json:"duration_ms"`
	Explicit         bool     `json:"explicit"`
	ExternalUrls     struct {
	Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href        string `json:"href"`
	ID          string `json:"id"`
	IsLocal     bool   `json:"is_local"`
	Name        string `json:"name"`
	PreviewURL  string `json:"preview_url"`
	TrackNumber int64  `json:"track_number"`
	Type        string `json:"type"`
	URI         string `json:"uri"`
	} `json:"items"`
	Limit    int64       `json:"limit"`
	Next     interface{} `json:"next"`
	Offset   int64       `json:"offset"`
	Previous interface{} `json:"previous"`
	Total    int64       `json:"total"`
}


//
// ALBUMS
//

func (c SpotifyClient) GetAlbum(id string) (album SpotifyAlbum, e error)   {
	resp, err := apiRequest("https://api.spotify.com/v1/albums/" + id, "GET", []byte{}, c.AccessToken)
	if err != nil { return album, err }
	err2 := json.Unmarshal(resp.response, &album)
	if err2 != nil { return album, err2 }
	return album, nil
}


func (c SpotifyClient) GetAlbumTracks(id string, options map[string]string) (tracks SpotifyTracks, e error) {
	queryString := ""
	if options != nil {
		queryString = makeQueryString(options)
	}

	resp, err := apiRequest("https://api.spotify.com/v1/albums/" + id + "/tracks" + queryString, "GET", []byte{}, c.AccessToken)
	if err != nil { return tracks, err }
	fmt.Println(string(resp.response))
	err2 := json.Unmarshal(resp.response, &tracks)
	if err2 != nil { return tracks, err2 }
	return tracks, nil
}

func (c SpotifyClient) GetAlbums(ids []string) (albums SpotifyAlbums, e error) {
	fmt.Println("GETTING ALBUMS")
	if len(ids) < 20 {
		queryString := "?ids="
		for i := 0; i < len(ids) - 1; i++ {
			queryString += ids[i] + ","
		}
		queryString += ids[len(ids) - 1]
		fmt.Println("QS" + queryString)

		resp, err := apiRequest("https://api.spotify.com/v1/albums/" + queryString, "GET", []byte{}, c.AccessToken)
		if err != nil { return albums, err }
		//fmt.Println(string(resp.response))
		err2 := json.Unmarshal(resp.response, &albums)
		if err2 != nil { return albums, err2 }
		return albums, nil
	} else {
		return albums, errors.New("Maximum of 20 IDs Per Call")
	}
}

//
// ARISTS
//

type ArtistAlbums struct {
	Href  string `json:"href"`
	Items []struct {
		AlbumGroup string `json:"album_group"`
		AlbumType  string `json:"album_type"`
		Artists    []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		AvailableMarkets []string `json:"available_markets"`
		ExternalUrls     struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href   string `json:"href"`
		ID     string `json:"id"`
		Images []struct {
			Height int64  `json:"height"`
			URL    string `json:"url"`
			Width  int64  `json:"width"`
		} `json:"images"`
		Name                 string `json:"name"`
		ReleaseDate          string `json:"release_date"`
		ReleaseDatePrecision string `json:"release_date_precision"`
		TotalTracks          int64  `json:"total_tracks"`
		Type                 string `json:"type"`
		URI                  string `json:"uri"`
	} `json:"items"`
	Limit    int64       `json:"limit"`
	Next     string      `json:"next"`
	Offset   int64       `json:"offset"`
	Previous interface{} `json:"previous"`
	Total    int64       `json:"total"`
}

func (c SpotifyClient) GetArtistAlbums(id string) (albums ArtistAlbums, e error) {
	resp, err := apiRequest("https://api.spotify.com/v1/artists/1vCWHaC5f2uS3yhpwWbIA6/albums", "GET", []byte{}, c.AccessToken)
	if err != nil { return albums,err}
	fmt.Println(string(resp.response))
	err2 := json.Unmarshal(resp.response, &albums)
	if err2 != nil { return albums, err2 }
	return albums, nil
}

//
// SEARCH
//

type SpotifySearchResults struct {
	Artists struct {
		Href  string `json:"href"`
		Items []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Followers struct {
				Href  interface{} `json:"href"`
				Total int64       `json:"total"`
			} `json:"followers"`
			Genres []string `json:"genres"`
			Href   string   `json:"href"`
			ID     string   `json:"id"`
			Images []struct {
				Height int64  `json:"height"`
				URL    string `json:"url"`
				Width  int64  `json:"width"`
			} `json:"images"`
			Name       string `json:"name"`
			Popularity int64  `json:"popularity"`
			Type       string `json:"type"`
			URI        string `json:"uri"`
		} `json:"items"`
		Limit    int64       `json:"limit"`
		Next     interface{} `json:"next"`
		Offset   int64       `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int64       `json:"total"`
	} `json:"artists"`
	Albums struct {
		Href  string `json:"href"`
		Items []SpotifyAlbum `json: "items"`
		Limit    int64       `json:"limit"`
		Next     interface{} `json:"next"`
		Offset   int64       `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int64       `json:"total"`
	}
}

func (c SpotifyClient) Search(search string, catagory string) (results string, e error)  {
	keywords := url.QueryEscape(search)
	resp, err := apiRequest("https://api.spotify.com/v1/search?q=" + keywords + "&type=" + catagory, "GET", []byte{}, c.AccessToken)
	if err != nil { return results,err}
	fmt.Println(string(resp.response))
	//err2 := json.Unmarshal(resp.response, &results)
	return results, e
}