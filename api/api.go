package api

import (
	"io/ioutil"
	"net/http"
)

const APIURL = "https://api.sdkman.io/2"

func getAPI(segment string) []byte {
	resp, err := http.Get(APIURL + segment)
	if err != nil {
		return []byte(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(err.Error())
	}
	return body
}

func GetBroadcast() []byte {
	return getAPI("/broadcast/latest")
}

func GetCandidatesList() []byte {
	return getAPI("/candidates/list")
}
