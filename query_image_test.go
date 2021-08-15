package resizeimage

import (
	"io"
	"os"
	"testing"

	_ "image/jpeg"
	_ "image/png"
)

func Test_getImageDimensions(t *testing.T) {
	jpg1280x927, err := os.Open("test-images/1280--927.jpg")
	if jpg1280x927 != nil {
		defer jpg1280x927.Close()
	}
	if err != nil {
		t.Error(err)
		return
	}

	png300x300, err := os.Open("test-images/300--300.png")
	if png300x300 != nil {
		defer png300x300.Close()
	}
	if err != nil {
		t.Error(err)
		return
	}

	invalidImage, err := os.Open("test-images/not--img.jpg")
	if invalidImage != nil {
		defer invalidImage.Close()
	}
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		img io.ReadSeeker
	}
	tests := []struct {
		name       string
		args       args
		wantWidth  int
		wantHeight int
		wantErr    bool
	}{
		{"check jpeg dimensions", args{jpg1280x927}, 1280, 927, false},
		{"check png dimensions", args{png300x300}, 300, 300, false},
		{"invalid image file", args{invalidImage}, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getImageDimensions(tt.args.img)
			if (err != nil) != tt.wantErr {
				t.Errorf("getImageDimensions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantWidth {
				t.Errorf("getImageDimensions() got = %v, want width %v", got, tt.wantWidth)
			}
			if got1 != tt.wantHeight {
				t.Errorf("getImageDimensions() got1 = %v, want height %v", got1, tt.wantHeight)
			}

			// seek files back to first read point fro re-use
			if _, err = tt.args.img.Seek(0, 0); err != nil {
				t.Errorf("Error while resetting file between tests %v", err)
			}
		})
	}
}
