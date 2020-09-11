package main

import (
    "fmt"
    "github.com/AudiGo/spotify"
    "github.com/joho/godotenv"
    "log"
    "os"
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main() {
    id, _ := os.LookupEnv("CLIENT_ID") // Your Spotify ClientId
    secret, _ := os.LookupEnv("CLIENT_SECRET") // Your Spotify Secret
    client := spotify.SpotifyClient{id, secret, ""}
    fmt.Println(client.ClientId)
    if err := client.Authenticate(); err != nil {
        fmt.Println(err)
    }
}