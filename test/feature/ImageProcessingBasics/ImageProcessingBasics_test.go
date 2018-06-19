package ImageProcessingBasics

import (
    "github.com/DATA-DOG/godog"
    "fmt"
    "QPongServer/util"
    "image"
    "strings"
    "strconv"
)

// struct to encapsulate the image for the feature test
type imageModel struct {
    Location string
    Config image.Config
}
var imgModel imageModel = imageModel{}

// set the image file location to the imageModel struct
func setImageFileLocation(file string) error {
    imgModel.Location = file
    return nil
}

// decode the image file (if valid)
func decodeImage(imageFormat string) error {
    imgCfg, imgFormat, err := util.DecodeImage(imgModel.Location)
    if err != nil {
        return fmt.Errorf("could not decode the given image %v => %v", imgModel.Location, err)
    }
    imgModel.Config = imgCfg

    switch imageFormat {
    case "jpg":
        if strings.Compare("jpeg", imgFormat) == 0 {
            return nil
        }
    case "jpeg":
        if strings.Compare("jpeg", imgFormat) == 0 {
            return nil
        }
    case "png":
        if strings.Compare("png", imgFormat) == 0 {
            return nil
        }
    default:
        return fmt.Errorf("non supported format %v", imgFormat)
    }
    // should not be reaching here though
    return godog.ErrPending
}

// method to check the dimensions of a decoded image
func checkDimensions(width, height string) error {
    iW, _ := strconv.ParseInt(width, 10, 64)
    iH, _ := strconv.ParseInt(height, 10, 64)

    if imgModel.Config.Width != int(iW) || imgModel.Config.Height != int(iH) {
        return fmt.Errorf("un-matched dimensions~ got [ %v x %v ] instead of the expected {%v x %v}",
            imgModel.Config.Width, imgModel.Config.Height, width, height)
    }
    return nil
}

func FeatureContext(s *godog.Suite) {
    s.Step(`^there is an image at "([^"]*)"$`, setImageFileLocation)
    s.Step(`^after loading the image file; the format should be decoded into "([^"]*)"$`, decodeImage)
    s.Step(`^the width is "([^"]*)" and height is "([^"]*)" pixels$`, checkDimensions)
}
