package main

import (
    "fmt"
    "os"
    "time"
    "encoding/json"
    "github.com/xtls/libxray/share"
)

func handleError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

func main() {

    // mute the stdout
    originalStdout := os.Stdout
    os.Stdout = nil

    url := os.Args[1]

    config, err := share.ConvertShareLinksToXrayJson(url)
    handleError(err)

    config.OutboundConfigs[0].SendThrough = nil

    json, err := json.MarshalIndent(config, "", "    ")
    handleError(err)

    // unmute the stdout
    time.Sleep(50 * time.Millisecond)
    os.Stdout = originalStdout

    os.Stdout.Write(json)

}

