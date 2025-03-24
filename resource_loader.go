package main

import (
	"bufio"
	"bytes"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

type ResourceLoader struct {
	cache map[string]interface{}
}

func NewResourceLoader() *ResourceLoader {
	return &ResourceLoader{
		cache: map[string]interface{}{},
	}
}

func (r *ResourceLoader) File(path string) (io.Reader, error) {
	if v, ok := r.cache[path]; ok {
		if b, ok := v.(io.Reader); ok {
			return b, nil
		}
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	b := bufio.NewReader(file)
	r.cache[path] = b
	return b, nil
}

func (r *ResourceLoader) Image(path string) (*ebiten.Image, error) {
	if v, ok := r.cache[path]; ok {
		if i, ok := v.(*ebiten.Image); ok {
			return i, nil
		}
	}
	file, err := r.File(path)
	if err != nil {
		return nil, err
	}
	i, _, err := ebitenutil.NewImageFromReader(file)
	r.cache[path] = i
	return i, err
}

func (r *ResourceLoader) FontFace(size int) (text.Face, error) {
	name := "font-" + strconv.Itoa(size)
	if v, ok := r.cache[name]; ok {
		if f, ok := v.(text.Face); ok {
			return f, nil
		}
	}
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		return nil, err
	}
	f := &text.GoTextFace{
		Source: s,
		Size:   float64(size),
	}
	r.cache[name] = f
	return f, nil
}
