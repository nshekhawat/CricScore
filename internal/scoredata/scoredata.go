package scoredata

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/nshekhawat/CricScore/internal/scoreapi"
)

// Matches - Cricket Matches
type Matches []struct {
	ID int    `json:"id"`
	T2 string `json:"t2"`
	T1 string `json:"t1"`
}

var matchID int

// GetMatchData - Get live match info and unmarshal them to struct
func GetMatchData(teamName string) (int, error) {
	var matches Matches

	body, err := scoreapi.GetLiveMatches()
	if err != nil {
		return 0, errors.New("Unable to fetch live match details from server")
	}

	if err = json.Unmarshal(body, &matches); err != nil {
		return 0, errors.New("Unable to marshal json data")
	}

	for _, v := range matches {
		if v.T1 == teamName || v.T2 == teamName {
			return v.ID, nil
		}
	}

	return 0, errors.New("Unable to find any match data for " + teamName)
}

// GetLiveScoreData - Get live score info and unmarshal them to struct.
func GetLiveScoreData(teamName string) []byte {
	matchID, err := GetMatchData(teamName)

	if matchID == 0 || err != nil {
		log.Fatalf("%v", err)
	}

	score, err := scoreapi.GetLiveScores(matchID)
	if err != nil {
		log.Fatalf("Unable to fetch latest scores - %v", err)
	}

	return score
}
