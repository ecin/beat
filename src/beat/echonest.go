package beat

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
)

const (
  TRACK_PROFILE_ENDPOINT = "http://developer.echonest.com/api/v4/track/profile"
  SONG_PROFILE_ENDPOINT = "http://developer.echonest.com/api/v4/song/profile"
)

type TrackRequest struct {
  ApiKey string `json:"api_key"`
  Id string     `json:"id"`

  Buckets string `json:"bucket,omitempty"`
  format string `json:"format"`
}

type Track struct {
  Status string `json:"status"`
  Artist string `json:"artist_name"`
  AudioSummary struct {
    Tempo float64 `json:"tempo"`
  } `json:"audio_summary"`
  Title string `json:"title"`
}

type TrackResponse struct {
  Response struct {
    Tracks []Track `json:"songs"`
  } `json:"response"`
}

type EchoNestClient struct {
  ApiKey string
  // perhaps http client
}

func NewEchoNestClient(apiKey string) EchoNestClient {
  client := EchoNestClient{
    apiKey,
  }

  return client
}

func (client EchoNestClient) GetTrack(trackIds []string) []Track {
  params := url.Values{}
  params.Add("api_key", client.ApiKey)
  params.Add("bucket", "audio_summary")
  params.Add("format", "json")

  for _, trackId := range trackIds {
    params.Add("track_id", trackId)
  }

  var endpoint = fmt.Sprintf("%s?%s", TRACK_PROFILE_ENDPOINT, params.Encode())
  endpoint = fmt.Sprintf("%s?%s", SONG_PROFILE_ENDPOINT, params.Encode())

  println(params.Encode())

  // TODO: handle error codes
  response, _ := http.Get(endpoint)
  body, _ := ioutil.ReadAll(response.Body)
  fmt.Printf("%s\n", body)

  var trackResponse TrackResponse
  json.Unmarshal(body, &trackResponse)

  fmt.Printf("%v\n", trackResponse)

  return trackResponse.Response.Tracks
}
