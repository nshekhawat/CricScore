package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nshekhawat/CricScore/internal/scoredata"
)

func main() {

	teamName := flag.String("t", "", "Provide team name")
	flag.Parse()

	if *teamName == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	score := scoredata.GetLiveScoreData(*teamName)

	fmt.Println(string(score))

}
