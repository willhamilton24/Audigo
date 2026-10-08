# Audigo
![Concept](https://img.shields.io/badge/Concept-F-red) ![Effort to Cashflow](https://img.shields.io/badge/Effort_to_Cashflow-90%2F100-red)
A Go Wrapper for Music APIs - Last.fm, Spotify, Apple Music, Soundcloud, Tidal, Discogs

Under construction, check back later!

## Docs

### Spotify
**spotify.SpotifyClient{ClientId string, ApiSecret string, AccessToken string)**
- ClientId: Spotify API Client Id
- ApiSecret: Spotify API Secret
- AccessToken: Spotify API Access Token if you have a valid one, otherwise pass empty string

*SpotifyClient.Authenticate() error*
- Uses SpotifyClient's ClientId and ApiSecret to generate a new API Access Token.
- IMPORTANT:This Must Be Run Before Any Other SpotifyClient method if you did not pass a valid AccessToken to the constructor

#### Albums

*SpotifyClient.GetAlbum(id string) (album SpotifyAlbum, e error)*
- id: Spotify Album Id

*SpotifyClient.GetAlbumTracks(id string, options map[string]string) (tracks SpotifyTracks, e error)*
- id: Spotify Album Id
- options: Query string options

*SpotifyClient.GetAlbums(ids []string) (album SpotifyAlbums, e error)*
- ids: Slice of Spotify Album Ids

#### Artists

*SpotifyClient.GetArtistAlbums(id string) (albums ArtistAlbums, e error)*
- id: Spotify Artist Id


#### Search

*SpotifyClient.Search(term string, catagory string) (results SpotifySearchResults, e error)* 
- term: The Search Term
- catagory: "artist" or "album"


## 💰 Path to Revenue
Open-source API client libraries are free by convention — official and mature community SDKs already exist for Spotify/Last.fm — so direct revenue is unrealistic; any path runs through a paid service built on top (e.g. a hosted music-metadata aggregation API).

### Release TODOs
- [ ] Finish the advertised coverage (Apple Music, SoundCloud, Tidal, Discogs are listed but unimplemented)
- [ ] Modernize to current Go practices (modules with go.mod, context support, error wrapping) and publish on pkg.go.dev
- [ ] Add tests and CI; remove the "Under construction" banner
- [ ] If monetizing: pivot to a hosted unified music-metadata API (one key, many providers) with usage-based pricing
- [ ] Alternatively accept donations/sponsorship (GitHub Sponsors) as an OSS library
- [ ] Promote in Go and music-dev communities to build adoption first
