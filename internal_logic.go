package main

import (
	"fmt"
	"os"
	"strings"
)

func Start(seq string){

	valid := testValidity(seq)
	if valid {
		fmt.Fprint(os.Stdout, "Validation: true\n")
		fmt.Fprint(os.Stdout, "Average number: ", averageNumber(seq), "\n")
		fmt.Fprint(os.Stdout, "Words: ", wholeStory(seq), "\n")
		stat := storyStats(seq)
		fmt.Fprintf(os.Stdout, "Story Stats:\nShortest word: %s\nLongest word: %s\n"+
			"Averege letter in word number: %d\nAverege words:\n%s\n", stat.shortestWord, stat.longestWord, stat.avgWordLength, strings.Join(stat.avgList, "\n"))
	} else {
		fmt.Fprint(os.Stdout, "Validation: false\n")
	}

}