package github

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"

	"github.com/ziyaaktas/sa-backend/gosrc/helper"
)

// GetPopularRepos get popular repos from github.com
func GetPopularRepos(w http.ResponseWriter, r *http.Request) {
	payload := get()
	language := r.URL.Query().Get("language")
	fmt.Println(language)
	// w.Write([]byte("welcome"))
	helper.RespondwithJSON(w, http.StatusOK, payload)
}

func get() helper.Repos {
    fmt.Println("Get function")
    resp, err := http.Get("https://api.github.com/search/repositories?q=created:>2019-01-10&sort=stars&order=desc")
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    // bodyString := string(bodyBytes)
    // fmt.Println("API Response as String:\n" + bodyString)

    // Convert response body to Todo struct
    var repoStruct helper.Repos
	json.Unmarshal(bodyBytes, &repoStruct)
	return repoStruct
    // fmt.Printf("API Response as struct %+v\n", repoStruct)
}
