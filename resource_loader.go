package main

import (
	"bufio"
	"bytes"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"

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

func (r *ResourceLoader) FontFaceSource(name string, data io.Reader) (*text.GoTextFaceSource, error) {
	if v, ok := r.cache[name]; ok {
		if f, ok := v.(*text.GoTextFaceSource); ok {
			return f, nil
		}
	}
	s, err := text.NewGoTextFaceSource(data)
	if err != nil {
		return nil, err
	}
	r.cache[name] = s
	return s, nil
}

func (r *ResourceLoader) FontFace(size int) (text.Face, error) {
	s, err := r.FontFaceSource("goregular.TTF", bytes.NewReader(goregular.TTF))
	if err != nil {
		return nil, err
	}
	return &text.GoTextFace{
		Source: s,
		Size:   float64(size),
	}, nil
}
