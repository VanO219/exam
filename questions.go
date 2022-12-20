package main

import (
	"bufio"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
)

type QuestionsList map[int]string

func NewQuestionsList(fileName string) (QuestionsList, []int, error) {
	var (
		ql  = QuestionsList{}
		qs  = []int{}
		err error
	)

	fl, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, nil, err
	}

	fileScanner := bufio.NewScanner(fl)

	for fileScanner.Scan() {
		s := fileScanner.Text()
		sl := strings.Split(s, "\t")
		questionNum, err := strconv.Atoi(sl[0])
		if err != nil {
			return nil, nil, errors.Wrap(err, "strconv.Atoi(sl[0])")
		}
		question := sl[1]
		ql[questionNum] = question
		qs = append(qs, questionNum)
	}
	return ql, qs, err
}
