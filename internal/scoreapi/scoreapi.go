package scoreapi

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

const cricscoreurl = "http://cricscore-api.appspot.com/csa"

// GetLiveMatches - Gets the updated list of live matches
func GetLiveMatches() ([]byte, error) {
	resp, err := http.Get(cricscoreurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetLiveScores - Get live score of the selected match.
func GetLiveScores(matchID int) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", cricscoreurl, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("id", strconv.Itoa(matchID))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
