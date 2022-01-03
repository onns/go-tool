package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/onns/zip"
)

/*
@Time : 2022/1/3 15:06
@Author : onns
@File : zip-encrypt/onns.go
*/

func Encrypt(zipFile string, src string, password string, encryption zip.EncryptionMethod, withRoot bool) (err error) {
	var (
		baseDir string
	)
	if strings.HasPrefix(src,"./") {
		err = errors.New("src should not has prefix of ./")
		return
	}
	if !strings.HasSuffix(zipFile, ".zip") {
		err = errors.New("zip file name error")
		return
	}
	fzip, err := os.Create(zipFile)
	if err != nil {
		return
	}
	defer fzip.Close()

	zipWriter := zip.NewWriter(fzip)
	defer zipWriter.Close()

	info, err := os.Stat(src)
	if err != nil {
		return
	}

	// 保证文件夹目录规范
	if info.IsDir() && !strings.HasSuffix(src, "/") {
		src += "/"
	}

	if info.IsDir() && withRoot {
		baseDir = filepath.Base(src) + "/"
	}
	log.Printf("src:%s baseDir: %s\n", src, baseDir)

	return filepath.Walk(src, func(path string, fi os.FileInfo, errWalk error) (err error) {
		// WalkFunc规定如果把错误回传就会终止遍历
		if errWalk != nil {
			return errWalk
		}

		var filename string
		filename = strings.TrimPrefix(path, src)
		fmt.Println(filename)

		if fi.IsDir() {
			filename += "/"
		}

		filename = baseDir + filename

		var fh io.Writer
		if encryption != zip.NoEncryption {
			fh, err = zipWriter.Encrypt(filename, password, encryption)
		}
		// TODO 不加密
		// 	fh, err = zipWriter.CreateHeader(header)

		if err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return
		}

		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			return
		}

		ret, err := io.Copy(fh, file)
		if err != nil {
			return
		}

		log.Printf("added: %s, total: %d\n", path, ret)
		return
	})
}
