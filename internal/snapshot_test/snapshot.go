// Package snapshot_test provides visual regression testing utilities for Ebiten games.
// It captures and compares game renderings to detect visual regressions.
//
// from:
//
//	https://zenn.dev/yktakaha4/articles/visual_regression_test_in_ebitengine
package snapshot_test

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	diff "github.com/olegfedoseev/image-diff"
)

const (
	SnapshotErrorThreshold = 0.0
)

// CheckSnapshot performs visual regression testing by comparing the actual image to an expected snapshot.
// It creates or updates snapshot files and reports any visual differences.
func CheckSnapshot(t *testing.T, actualImage *ebiten.Image) error {
	_, callerSourceFileName, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatalf("failed to read filename: %v", t.Name())
	}

	basePath := path.Join(path.Dir(callerSourceFileName), "snapshot")
	baseFileName := strings.ReplaceAll(t.Name(), "/", "_")
	expectedFilePath := path.Join(basePath, fmt.Sprintf("%v.png", baseFileName))
	actualFilePath := path.Join(basePath, fmt.Sprintf("%v_actual.png", baseFileName))
	diffFilePath := path.Join(basePath, fmt.Sprintf("%v_diff.png", baseFileName))

	err := os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	var expectedImage image.Image
	foundExpectedImage := false
	expectedFile, err := os.Open(expectedFilePath)
	if err == nil {
		expectedImage, _, err = image.Decode(expectedFile)
		if err != nil {
			log.Fatal(err)
		}
		foundExpectedImage = true
	} else if !errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}

	_ = os.Remove(diffFilePath)
	_ = os.Remove(actualFilePath)

	updateSnapshot, _ := strconv.ParseBool(os.Getenv("UPDATE_SNAPSHOT"))
	if foundExpectedImage && !updateSnapshot {
		diffImage, percent, err := diff.CompareImages(actualImage, expectedImage)
		if err != nil {
			log.Fatal(err)
		}

		if percent > SnapshotErrorThreshold {
			f, _ := os.Create(diffFilePath)
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(f)

			err = png.Encode(f, diffImage)
			if err != nil {
				log.Fatal(err)
			}

			f, _ = os.Create(actualFilePath)
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(f)

			err = png.Encode(f, actualImage)
			if err != nil {
				log.Fatal(err)
			}

			return fmt.Errorf(
				"snapshot test failed: diff = %v > %v, file = %v",
				percent,
				SnapshotErrorThreshold,
				diffFilePath)
		}
	}

	f, _ := os.Create(expectedFilePath)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	err = png.Encode(f, actualImage)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
