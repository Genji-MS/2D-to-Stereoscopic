package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

func main() {
	dirInput := flag.String("o", "./Input/", "Input Image .path/file.jpg")
	dirOutput := flag.String("i", "./Output/", "Output Image .path/file.jpg")
	flag.Parse()

	var wg sync.WaitGroup //instance of a WaitGroup, which allows for a coroutine
	files, err := ioutil.ReadDir(*dirInput)
	check(err)
	printedFiles := 0
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".jpg" {
			fileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) + ".jpg"
			printedFiles++
			fmt.Println("\033[33mProcessing image " + fileName + "\033[0m")
			wg.Add(1)
			go cropImage(fileName, *dirInput, *dirOutput, &wg) //calls the coroutine function passing the pointer value of wg in
			wg.Wait() //actually makes the waitgroup wait at this point in the code if our waitgroup function is still running. It has allowed other things to process up to this point.
			fmt.Println("\033[32mImage " + fileName + " complete.\033[0m")
		}
	}
	if printedFiles > 0 {
		fmt.Printf("\033[32m\033[1mSuccess!\033[0m Generated \033[1m%s\033[0m image \n", strconv.Itoa(printedFiles))
	} else {
		fmt.Println("\033[31mERROR\033[0m No \033[36m(.png)\033[0m files found in directory")
	}
}

func cropImage(fileName, inputPath, outputPath string, wg *sync.WaitGroup) {
	defer wg.Done()

	outFile := createImg(outputPath + "2D_" + fileName)
	f, err := os.Open(inputPath + fileName)
	check(err)
	imageData, err := jpeg.Decode(f)
	check(err)
	f.Close()

	imageBounds := imageData.Bounds()
	right := imageBounds.Dx()
	width := right/2
	height := imageBounds.Dy()
	mid := right - width
	fmt.Println("full width: " + strconv.Itoa(right) + " width: " + strconv.Itoa(mid) + " height: " + strconv.Itoa(height) )
		// crop the right side image (Keeping the RIGHT)
	rect := image.Rect(mid, 0, right, height)
		// crop the left side image (Keeping the LEFT)
	// rect := image.Rect(0, 0, width, height)

	type subImage interface {
        SubImage(r image.Rectangle) image.Image
    }
	cropImage, _ := imageData.(subImage)
	imageData = cropImage.SubImage(rect)
	err = jpeg.Encode(outFile, imageData, nil)
	check(err)

	outFile.Close()
}

func createImg(filename string) *os.File {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		check(err)
		return file
	} else {
		// trying to write to an existing file gives us errors, delete it instead and create
		err = os.Remove(filename)
		check(err)
		file, err := os.Create(filename)
		check(err)
		return file
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}