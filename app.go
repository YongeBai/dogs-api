package main

import (
	"context"
	"encoding/json"
	"fmt"

	// "fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type RandomImage struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type AllBreeds struct {
	Message map[string][]string `json:"message"`
	Status  string                        `json:"status"`
}

type ImageByBreed struct {
	Message []string `json:"message"`
	Status  string   `json:"status"`
}

func (a *App) GetRandomImageUrl() string {
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var randomImage RandomImage
	err = json.Unmarshal(body, &randomImage)
	if err != nil {
		log.Fatal(err)
	}

	return randomImage.Message
}

func (a *App) GetBreedList() []string {
	response, err := http.Get("https://dog.ceo/api/breeds/list/all")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var allBreeds AllBreeds
	err = json.Unmarshal(body, &allBreeds)
	if err != nil {
		log.Fatal(err)
	}

	breeds := make([]string, 0)
	for breed := range allBreeds.Message {
		breeds = append(breeds, breed)
	}
	sort.Strings(breeds)
	return breeds
}

func (a *App) GetImageUrlsByBreed(breed string) []string {
	url := "https://dog.ceo/api/breed/"+breed+"/images"
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var imageByBreed ImageByBreed
	err = json.Unmarshal(body, &imageByBreed)
	if err != nil {
		log.Fatal(err)
	}

	return imageByBreed.Message	
}
