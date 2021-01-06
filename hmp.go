package main

import (
	"fmt"
	"os"
	"log"
	"os/exec"
	"path/filepath"
)

func prettyPrintSize(size int64) {
	switch {
	case size > 1024*1024*1024:
		fmt.Printf("Counting penguins for %.1fG packages size\n", float64(size)/(1024*1024*1024))
	case size > 1024*1024:
		fmt.Printf("Counting penguins for %.1fM packages size\n", float64(size)/(1024*1024))
	case size > 1024:
		fmt.Printf("Counting penguins for %.1fK packages size\n", float64(size)/1024)
	default:
		fmt.Printf("Counting penguins for %d packages size", size)
	}
}

func DirSize(path string) (int64, error) {
	var size int64
	adjSize := func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	}
	err := filepath.Walk(path, adjSize)
	return size, err
}

func DownloadPackages() {
	fmt.Println("Downloading packages...")
	yarn := exec.Command("yarn")
	err := yarn.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Packages downloaded.")
}

func Requirements() {
	if _, err := os.Stat("./package.json"); os.IsNotExist(err) {
		log.Fatal("No package.json detected !")
	}
	_, err := exec.LookPath("yarn")
	if err != nil {
		log.Fatal("Please install Yarn cli to continue")
	}
}

func CountPenguins(size int64) (int64) {
	prettyPrintSize(size)

	penguins := size%100
	return penguins
}

func PrintPenguins(penguins int64) {
	fmt.Printf("🔪 ")
	for i := int64(1); i < penguins; i++ {
		if i == 100 {
			fmt.Printf(" You freak... ")	
		} else if i == 200 {
			fmt.Printf(" Really ?! ")	
		} else {
			fmt.Printf("🐧 ")
		}
	}
	fmt.Printf("\n")
}

func main() {
	fmt.Println("Welcome on HowManyPenguin CLI")
	fmt.Println("I'll install all packages listed in your package.json file and calculate how many penguins you just killed...")
	Requirements()
	DownloadPackages()
	size, err := DirSize("./node_modules")
	if err != nil {
		log.Fatal("Please install Yarn cli to continue")
	}
	penguins := CountPenguins(size)
	fmt.Println("Penguins killed: ", penguins)
	PrintPenguins(penguins)
}
