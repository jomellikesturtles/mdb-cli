package main

import (
	"bufio"
	"container/list"
	"flag"
	"fmt"
	"io"
	"net/http"

	// "fmt"
	"log"
	"mdb-cli/cmd"
	"os"
	// "rsc.io/quote"
)

func main() {
	cmd.Execute()
	// scan()
}
func main2() {

	cmd.Execute()

	// fmt.Println(quote.Go())
	log.Println("HELLO WORLD")
	versionPtr := flag.String("version", "", "a string")
	log.Println(help())

	flag.Parse()
	// versionPtr
	log.Println("version:", *versionPtr)
	scan()
	os.Exit(0)
}

func help() string {
	return "version 1"
}

func scan() {
	fetchMetadata("tt0120338")
	log.Println("Scanning...")

	// TODO: get library folders

	pathList := list.New()
	defaultPath := "/Users/jommel/Downloads/torrents"
	pathList.PushBack(defaultPath)

	// f, err := os.Open(defaultPath)

	log.Println("Logging..")
	// f.Read()
	// f.Name()

	// for _, num:=range pathList. {

	// }
	// pathList
	entries, err := os.ReadDir(defaultPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		log.Println("Directory name: ", entry.Name())
		// TODO: check if valid file, then check if valid extension
		// add to .db if valid
	}
	log.Println("Scanning completed")
}

func isValid(filePath string) bool {

	return true
}

func fetchMetadata(externalId string) {

	TMDB_URL := "https://api.themoviedb.org/3"
	TMDB_API_KEY := ""
	// TMDB_URL := os.Getenv("TMDB_URL")
	// TMDB_API_KEY := os.Getenv("TMDB_API_KEY")
	// const externalId = "tt1000912"
	external_source := "imdb_id"
	url := TMDB_URL + "/find/" + externalId + "?api_key=" + TMDB_API_KEY + "&language=en-US&external_source=" + external_source
	log.Println("url: " + url)
	response, err := http.Get(url)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	// var apiResponse

	if err != nil {
		panic(err)
	}

	log.Println("Response Status: " + string(body))
	log.Println("Response Status: " + response.Status)
	scanner := bufio.NewScanner(response.Body)
	// scanner.Scan()
	for i := 0; scanner.Scan() && i < 50; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
