package main

import "math/rand"

func newCopySlice(sl []int) []int {
	newSl := make([]int, len(sl))
	copy(newSl, sl)
	return newSl
}

func del(in []int, index int) []int {
	return append(in[:index], in[index+1:]...)
}

func getQuestionNums(numOfQuestion int, allQuestions []int) []int {
	questions := newCopySlice(allQuestions)
	outQ := make([]int, 0, numOfQuestion)

	for i := 0; i < numOfQuestion; i++ {
		indexOfQ := rand.Intn(len(questions))
		outQ = append(outQ, questions[indexOfQ])
		questions = del(questions, indexOfQ)
	}
	return outQ
}
