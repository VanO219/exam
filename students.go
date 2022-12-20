package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name           string
	NumOfQuestions int
}

type StudentsList map[int]Student

func NewStudentsList(fileName string) (StudentsList, []int, error) {
	var (
		err error
		stl = StudentsList{}
		sts = []int{}
	)

	fl, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, nil, err
	}

	fileScanner := bufio.NewScanner(fl)

	for fileScanner.Scan() {
		s := fileScanner.Text()
		s = strings.ReplaceAll(s, "\t", " ")
		sl := strings.Split(s, " ")
		studentNum, err := strconv.Atoi(sl[0])
		if err != nil {
			return nil, nil, errors.Wrap(err, "strconv.Atoi(sl[0])")
		}
		studentName := fmt.Sprintf("%s %s %s", sl[1], sl[2], sl[3])
		numOfQuestions, err := strconv.Atoi(sl[4])
		if err != nil {
			return nil, nil, errors.Wrap(err, "strconv.Atoi(sl[4])")
		}
		stl[studentNum] = Student{
			Name:           studentName,
			NumOfQuestions: numOfQuestions,
		}
		sts = append(sts, studentNum)
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		err = errors.Wrap(err, "Error while reading file")
		return nil, nil, err
	}

	return stl, sts, err
}
