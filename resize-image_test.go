package resizeimage

import (
	"io"
	"testing"
)

func TestResizeImage(t *testing.T) {
	type args struct {
		img       io.Reader
		newWidth  int
		newHeight int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResizeImage(tt.args.img, tt.args.newWidth, tt.args.newHeight)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResizeImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// get dimensions of resized file
			width, height, err := getImageDimensions(got)
			if err != nil {
				t.Errorf("Error while getting files dimensions: %v", err)
				return
			}

			if width != tt.args.newWidth {
				t.Errorf("ResizeImage() new width = %v, wanted %v", width, tt.args.newWidth)
				return
			}

			if height != tt.args.newHeight {
				t.Errorf("resizeImage() new height = %v, wanted %v", height, tt.args.newHeight)
				return
			}
		})
	}
}
