# Map & Mapb Implementation Plan

## 1. Update `config` struct (repl.go)

Add `next` and `previous` fields to track pagination:

```go
type config struct {
    commands map[string]cliCommand
    next     *string
    previous *string
}
```

Initialize `next` in `main.go` so the first `map` call has a URL:

```go
baseURL := "https://pokeapi.co/api/v2/location-area/"
cfg := &config{
    next: &baseURL,
    commands: map[string]cliCommand{ ... },
}
```

## 2. Define response structs

The PokeAPI `location-area` endpoint (no ID) returns:

```json
{"count": 1054, "next": "...?offset=20&limit=20", "previous": null, "results": [{"name": "canalave-city-area", "url": "..."}]}
```

Create a struct to unmarshal this (can live in `map_command.go` or a new file):

```go
type locationAreaResponse struct {
    Count    int     `json:"count"`
    Next     *string `json:"next"`
    Previous *string `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}
```

## 3. Implement `mapCommand` (map_command.go)

```go
func mapCommand(c *config) error {
    if c.next == nil {
        fmt.Println("you're on the last page")
        return nil
    }

    resp, err := http.Get(*c.next)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    var data locationAreaResponse
    err = json.Unmarshal(body, &data)
    if err != nil {
        return err
    }

    // Update pagination state
    c.next = data.Next
    c.previous = data.Previous

    // Print location names
    for _, loc := range data.Results {
        fmt.Println(loc.Name)
    }
    return nil
}
```

Imports needed: `"fmt"`, `"net/http"`, `"io"`, `"encoding/json"`

## 4. Implement `mapbCommand` (mapb_command.go)

Same logic but uses `c.previous`:

```go
func mapbCommand(c *config) error {
    if c.previous == nil {
        fmt.Println("you're on the first page")
        return nil
    }

    // Same HTTP + unmarshal logic as mapCommand, but use *c.previous as the URL
    // Update c.next and c.previous from the response
    // Print location names
}
```

Consider extracting the HTTP+unmarshal logic into a shared helper to avoid duplication.

## 5. Register commands in `main.go`

Add both commands to the map:

```go
"map": {
    name:        "map",
    description: "Display the next 20 location areas",
    callback:    mapCommand,
},
"mapb": {
    name:        "mapb",
    description: "Display the previous 20 location areas",
    callback:    mapbCommand,
},
```

## 6. Test

Run the CLI and verify:
- `map` prints 20 locations starting with `canalave-city-area`
- A second `map` prints the next 20
- `mapb` after that goes back to the first page
- `mapb` on the first page prints "you're on the first page"
- Run `go test ./...` from the repo root
