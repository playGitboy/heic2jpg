// https://gophercoding.com/convert-heic-to-jpeg-go/
package main

import (
	"image/jpeg"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"path/filepath"
	"fmt"

	"github.com/adrium/goheif"
)

// Skip Writer for exif writing
type writerSkipper struct {
	w           io.Writer
	bytesToSkip int
}

func main() {
  if len(os.Args) > 1 {
	  heicPaths := getHeicPath(os.Args[1:])
		if len(heicPaths) > 0 {
			for _, heicPath := range heicPaths {
				err := convertHeicToJpg(heicPath, getFileNameNoExt(heicPath) + ".jpg")
				if err != nil {
					log.Fatal(err)
				}	
			}
			fmt.Println("  > 全部文件转换完成！")
		} 
	}else{
		fmt.Println("  > 请指定heic文件路径或所在目录路径...")
	}
}

func convertHeicToJpg(input, output string) error {
	fileInput, err := os.Open(input)
	if err != nil {
		return err
	}
	defer fileInput.Close()

	// Extract exif to add back in after conversion
	exif, err := goheif.ExtractExif(fileInput)
	if err != nil {
		return err
	}

	img, err := goheif.Decode(fileInput)
	if err != nil {
		return err
	}

	fileOutput, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fileOutput.Close()

	// Write both convert file + exif data back
	w, _ := newWriterExif(fileOutput, exif)
	err = jpeg.Encode(w, img, nil)
	if err != nil {
		return err
	}

	return nil
}

func (w *writerSkipper) Write(data []byte) (int, error) {
	if w.bytesToSkip <= 0 {
		return w.w.Write(data)
	}

	if dataLen := len(data); dataLen < w.bytesToSkip {
		w.bytesToSkip -= dataLen
		return dataLen, nil
	}

	if n, err := w.w.Write(data[w.bytesToSkip:]); err == nil {
		n += w.bytesToSkip
		w.bytesToSkip = 0
		return n, nil
	} else {
		return n, err
	}
}

func newWriterExif(w io.Writer, exif []byte) (io.Writer, error) {
	writer := &writerSkipper{w, 2}
	soi := []byte{0xff, 0xd8}
	if _, err := w.Write(soi); err != nil {
		return nil, err
	}

	if exif != nil {
		app1Marker := 0xe1
		markerlen := 2 + len(exif)
		marker := []byte{0xff, uint8(app1Marker), uint8(markerlen >> 8), uint8(markerlen & 0xff)}
		if _, err := w.Write(marker); err != nil {
			return nil, err
		}

		if _, err := w.Write(exif); err != nil {
			return nil, err
		}
	}

	return writer, nil
}

func isFile(szPath string) bool {
	fi, e := os.Stat(szPath)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

func getHeicPath(szPaths []string) (szXlsPath []string) {
	re, _ := regexp.Compile(`(?i)\.heic$`)
	for _, szPath := range szPaths {
		if isFile(szPath) && re.MatchString(szPath) {
			szXlsPath = append(szXlsPath, szPath)
		} else {
			matches, _ := filepath.Glob(filepath.Join(szPath, "*.[hH][eE][iI][cC]"))
			dstMatches := make([]string, 0)
			for _, m := range matches {
				dstMatches = append(dstMatches, m)
			}
			szXlsPath = append(szXlsPath, dstMatches...)
		}
	}
	return
}

func getFileNameNoExt(fileName string) string {
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}