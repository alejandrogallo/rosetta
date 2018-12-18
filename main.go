package main

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
)

func GetSnippetsPaths() []string {
	var paths []string
	files, err := ioutil.ReadDir("./snippets/")
	if err != nil {
		log.Fatal(err)
	}
	for _, f:= range files {
		paths = append(paths, path.Join("./snippets", f.Name()))
	}
	return paths
}

func main() {
	fmt.Println("hello world")

	viper.SetConfigName("config.yaml")
	viper.AddConfigPath("/etc/rosetta/")
	viper.AddConfigPath("$HOME/.rosetta")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No config file found")
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	paths := GetSnippetsPaths()
	fmt.Println(paths)
}
