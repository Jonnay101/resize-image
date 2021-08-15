package resizeimage

import (
	"image"
	"io"
)

func getImageDimensions(img io.Reader) (int, int, error) {
	imgCfg, _, err := image.DecodeConfig(img)
	if err != nil {
		return 0, 0, err
	}

	return imgCfg.Width, imgCfg.Height, nil
}
