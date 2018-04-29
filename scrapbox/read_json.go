package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	projectName := getArgs()

	file, err := ioutil.ReadFile("./raw/" + projectName + ".json")
	if err != nil {
		return
	}
	project := Project{}
	json.Unmarshal(file, &project)
	// fmt.Println(file)
	title := sampleTitle(project)
	lineNums := countLines(project)
	log := []string{
		title,
		strconv.Itoa(lineNums)}
	saveLog(log)
}

func getArgs() string {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		return "daiiz"
	}
	return args[0]
}

func countLines(project Project) int {
	lines := 0
	var pages = project.Pages
	for i := 0; i < len(pages); i++ {
		lines += len(pages[i].Lines)
	}
	fmt.Printf("total lines: %d\n", lines)
	return lines
}

func sampleTitle(project Project) string {
	pageLength := len(project.Pages)

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(pageLength)

	// ランダムに記事タイトルを表示する
	randTitle := project.Pages[idx].Lines[0].Text
	fmt.Println(randTitle)
	return randTitle
}
