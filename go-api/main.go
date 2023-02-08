package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/go-chi/chi"
)

type PokemonsList struct {
	Next    *string `json:"next"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type Pokemon struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Height int    `json:"height"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Sprites struct {
		BackDefault  string `json:"back_default"`
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
		BackShiny    string `json:"back_shiny"`
	} `json:"sprites"`
}

func GetResponse(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func GetPokemonResponse(url string, pipe chan<- Pokemon) {
	pokemon := Pokemon{}
	err := json.Unmarshal(GetResponse(url), &pokemon)
	if err != nil {
		panic(err)
	}
	pipe <- pokemon
	// return pokemon
}

func main() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	router := chi.NewRouter()

	tempUrl := "https://pokeapi.co/api/v2/pokemon"
	url := &tempUrl
	var pokemons []Pokemon
	for url != nil {
		pokemonsList := PokemonsList{}
		err := json.Unmarshal(GetResponse(*url), &pokemonsList)
		if err != nil {
			panic(err)
		}

		channel := make(chan Pokemon, len(pokemonsList.Results))
		for _, result := range pokemonsList.Results {
			go GetPokemonResponse(result.Url, channel)
		}
		for range pokemonsList.Results {
			pokemons = append(pokemons, <-channel)
		}
		fmt.Println("Pokemon id: ", pokemonsList)
		url = pokemonsList.Next
	}

	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].Id < pokemons[j].Id
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, pokemons); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})
	log.Printf("Starting server on port %s. ", port)
	err1 := http.ListenAndServe(":"+port, router)
	log.Fatal(err1)
}
