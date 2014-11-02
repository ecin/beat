package beat

import (
  "testing"
)

func TestNewEchoNestClient(t *testing.T) {
  apiKey := "key"
  client := NewEchoNestClient(apiKey)

  if client.ApiKey != apiKey {
    t.Error("EchoNestClient constructor didn't assign API key")
  }
}

func TestGetTrack(t *testing.T) {
  apiKey := "GHGRDKHEKSC8WEQZW"
  client := NewEchoNestClient(apiKey)

  ids := []string{
    "spotify:track:3GhntU8mCuMuW5NXHvaTOx",
    "spotify:track:1ko1hqVxyzvRlAsbklLIbV",
    "spotify:track:6iejJ6Siz6lHcgcdsGNAaY",
  }

  tracks := client.GetTrack(ids)

  if len(tracks) != 3 {
    t.Error("Expected 2 tracks to be returned.")
  }

  trackTests := []struct {
    title string
    artist string
    tempo float64
  }{
    {"Air", "Rogue", 148.056000},
    {"Prism", "Case & Point", 127.964},
    {"Duality", "Fractal", 159.964},
  }

  for index, test := range trackTests {
    track := tracks[index]
    title := track.Title
    artist := track.Artist
    tempo := track.AudioSummary.Tempo

    if title != test.title {
      t.Errorf("track.Title => %q, want %q", title, test.title)
    }

    if track.Artist != test.artist {
      t.Errorf("track.Artist => %q, want %q", artist, test.artist)
    }

    if track.AudioSummary.Tempo != test.tempo {
      t.Errorf("track.Tempo => %q, want %q", tempo, test.tempo)
    }
  }
}
