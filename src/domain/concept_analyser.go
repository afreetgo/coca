package domain

import (
	"coca/src/adapter/models"
	"coca/src/domain/support"
	"fmt"
)

type ConceptAnalyser struct {
}

func NewConceptAnalyser() ConceptAnalyser {
	return *&ConceptAnalyser{}
}

func (c ConceptAnalyser) run() {

}

func (c ConceptAnalyser) Analysis(clzs *[]models.JClassNode) {
	buildMethodsFromDeps(*clzs)
}

func buildMethodsFromDeps(clzs []models.JClassNode) {
	var methodsName []string
	var methodStr string
	for _, clz := range clzs {
		for _, method := range clz.Methods {
			methodName := method.Name
			methodsName = append(methodsName, methodName)
			methodStr = methodStr + " " + methodName
		}
	}

	words := support.SegmentConceptCamelcase(methodsName)

	words = removeNormalWords(words)

	wordCounts := support.RankByWordCount(words)
	for _, word := range wordCounts[0:20] {
		fmt.Println(word.Key, word.Value)
	}
}

var normalWords = []string{
	"get",
	"create",
	"update",
	"delete",
	"save",

	"exist",
	"find",
	"new",
	"parse",

	"set",
	"get",

	"type",

	"all",
	"by",
	"id",
	"is",
	"of",
	"not",
	"with",
	"main",
}

func removeNormalWords(words map[string]int) map[string]int {
	var newWords = words
	for _, normalWord := range normalWords {
		if newWords[normalWord] > 0 {
			delete(newWords, normalWord)
		}
	}

	fmt.Println(newWords)
	return newWords
}