package util

import (
    "strings"
    "errors"
    "image"
    "os"
    "image/jpeg"
    "image/png"
    "sync"
)

var syncLock sync.Once

// method to register the supported image formats; only run this method once
// by applying sync.Once.Do(func())
func registerImageFormats() {
    syncLock.Do(func() {
        image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
        image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
    })
}

// decode the image information (e.g. width, height, colorModel - for color translation)
func DecodeImage(file string) (imgCfg image.Config, imgFormat string, err error)  {
    // validation
    if strings.Compare(file, "") == 0 {
        err = errors.New("file location given is empty~")
        return
    }
    // load the file
    imgFilePtr, err := os.Open(file)

    defer imgFilePtr.Close()
    if err != nil {
        return
    }
    // decode
    registerImageFormats()
    imgCfg, imgFormat, err = image.DecodeConfig(imgFilePtr)

    return
}
