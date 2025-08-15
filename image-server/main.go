package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/image/draw"
)

var validImageName = regexp.MustCompile(`^[a-zA-Z0-9._ -]+\.(jpg|jpeg|png|gif|avif|webp)$`)

func isValidImageName(filename string) bool {
	return validImageName.MatchString(filename)
}

func all_images(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files, err := os.ReadDir("./images")
	if err != nil {
		fmt.Println("Something went wrong", err)
		return
	}
	fmt.Println(files)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>Image Server</h1>"))

	for _, file := range files {
		if !file.IsDir() {
			filename := file.Name()
			thumbnail_path := "/thumb/" + filename
			image_path := "/image/" + filename
			w.Write([]byte(fmt.Sprintf("<a href='%s' target='_blank'><img src='%s'/></a><p>%s</p>", image_path, thumbnail_path, filename)))
		}
	}
}

func serve_image(w http.ResponseWriter, r *http.Request) {
	filename := strings.TrimPrefix(r.URL.Path, "/image/")
	filename = filepath.Base(filename)
	if !isValidImageName(filename) {
		http.Error(w, "Invalid filename", http.StatusBadRequest)
		return
	}
	http.ServeFile(w, r, filepath.Join("./images", filename))
}

func thumbnail(w http.ResponseWriter, r *http.Request) {
	filename := strings.TrimPrefix(r.URL.Path, "/thumb/")
	filename = filepath.Base(filename)
	if !isValidImageName(filename) {
		http.Error(w, "Invalid filename", http.StatusBadRequest)
		return
	}
	file_path := filepath.Join("./images", filename)
	image, err := load_image(file_path)
	if err != nil {
		fmt.Println("Unable to load image", err)
		http.Error(w, "Unable to load image", http.StatusNotFound)
		return
	}
	thumbnail := resize_image(image, 128, 128)
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, thumbnail)
	default:
		w.Header().Set("Content-Type", "image/jpeg")
		jpeg.Encode(w, thumbnail, nil)
	}
}

func load_image(filepath string) (image.Image, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	return image, err
}

func resize_image(img image.Image, maxHeight, maxWidth int) image.Image {
	bounds := img.Bounds()
	origW := bounds.Dx()
	origH := bounds.Dy()

	ratioW := float64(maxWidth) / float64(origW)
	ratioH := float64(maxHeight) / float64(origH)
	ratio := math.Min(ratioW, ratioH)

	newW := int(float64(origW) * ratio)
	newH := int(float64(origH) * ratio)
	resized := image.NewRGBA(image.Rect(0, 0, newW, newH))
	draw.CatmullRom.Scale(resized, resized.Bounds(), img, img.Bounds(), draw.Over, nil)
	return resized
}

func main() {
	fmt.Println("Running")
	http.HandleFunc("/", all_images)
	http.HandleFunc("/image/", serve_image)
	http.HandleFunc("/thumb/", thumbnail)

	port := ":8888"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error running server", err)
	} else {
		fmt.Println("Server running on port", port)
	}
}
