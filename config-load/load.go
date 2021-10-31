package config_load

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

/*
@Time : 2021/10/16 12:08
@Author : onns
@File : config-load/load.go
*/

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func LoadConfig(dir string) (res []byte, err error) {
	var (
		workDir   string
		configDir string
	)
	configDir = dir
	if !fileExists(configDir) {
		workDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
		log.Println(workDir)
		if err != nil {
			log.Printf("filepath.Abs err(%+v)", err)
			return
		}
		configDir = path.Join(workDir, "config.json")
	}
	if !fileExists(configDir) {
		return
	}
	// log.Printf("%v",configDir)
	res, err = ioutil.ReadFile(configDir)
	if err != nil {
		log.Printf("ioutil.ReadFile err(%+v)", err)
		return
	}
	// log.Printf("%v",string(res))
	return
}
