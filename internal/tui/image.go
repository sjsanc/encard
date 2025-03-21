package tui

import (
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/sjsanc/encard/internal/log"
)

// TODO: support Kitty flags using structs
// TODO: images need to be cleared up when the program ends
// TODO: images rendered as history need to be darkened to match the Fade filter
// TODO: images that extend past the screen need to be truncated
// TODO: support GIFs

// Note: images are currently using the local file transmission flag (t=f) so won't work over SSH

type Image struct {
	ext  string
	path string
}

func NewImage(path string) *Image {
	return &Image{
		ext:  filepath.Ext(path),
		path: path,
	}
}

func (i *Image) Print() string {
	info, err := os.Stat(i.path)
	if info == nil {
		log.Warn("image not found: %s", i.path)
		return ""
	}
	if err != nil {
		log.Warn("error reading image: %s", err)
		return ""
	}

	if i.ext == ".png" {
		return i.printPNG()
	} else if i.ext == ".jpg" {
		return i.printJPG()
	} else {
		panic("unsupported image format")
	}
}

// i= id
// t= transmission medium (f=file)
// q= suppress responses
// f= format (100=png, 32=jpg)
// a= action (T=transmit & display)

func (i *Image) printPNG() string {
	encoded := base64.StdEncoding.EncodeToString([]byte(i.path))
	return fmt.Sprintf("\033_Gi=1,t=f,q=1,f=100,a=T;%s\033\\", encoded)
}

func (i *Image) printJPG() string {
	file, err := os.Open(i.path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	outfile, err := os.CreateTemp("", "encard-*.png")
	if err != nil {
		panic(err)
	}
	defer func() {
		outfile.Close()
		// os.Remove(outfile.Name())
	}()

	err = png.Encode(outfile, img)
	if err != nil {
		panic(err)
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(outfile.Name()))
	return fmt.Sprintf("\033_Gi=2,t=f,q=1,f=100,a=T;%s\033\\", encoded)
}
