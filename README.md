# About
Use the Sightengine Moderation API to instantly moderate images. See http://sightengine.com for more information.

Before starting, please make sure you have created an account on https://sightengine.com

# Installation

You can use the client using the `go` command.

    $ go get github.com/AswoDevice/sightengine-go

# Initialize the client

You will need your API USER and API SECRET to initialize the client. You can find both of them on your Sightengine account.
```go
sg "github.com/AswoDevice/sightengine-go"

client := sg.New({api_user}, {api_secret}, {models})
```

## Moderate an image through a public URL:

```go
// Detect nudity in an image
client := sg.New("api_user", "api_secret", sg.Nudity)
resp, err := client.CheckUrl("http://img09.deviantart.net/2bd0/i/2009/276/c/9/magic_forrest_wallpaper_by_goergen.jpg")

// Detect nudity, weapons, alcohol, drugs, likely fruadulant users, celebrities and faces in an image, along with image properties and type
client := sg.New("api_user", "api_secret", sg.Nudity, sg.Type, sg.Properties, sg.Wad, sg.Face, sg.Scam, sg.Celebrity)
resp, err := client.CheckUrl("http://img09.deviantart.net/2bd0/i/2009/276/c/9/magic_forrest_wallpaper_by_goergen.jpg")
```

## Moderate a local image:
```go
// Detect nudity in an image
client := sg.New("api_user", "api_secret", sg.Nudity)
resp, err := client.CheckFile("/full/path/to/image.jpg")

// Detect nudity, weapons, alcohol, drugs and faces in an image, along with image properties and type
client := sg.New("api_user", "api_secret", sg.Nudity, sg.Type, sg.Properties, sg.Wad, sg.Face, sg.Scam, sg.Celebrity)
resp, err := client.CheckFile("/full/path/to/image.jpg")
```
