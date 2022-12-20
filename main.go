package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"time"
)

var (
	studentsList   StudentsList
	questionsList  QuestionsList
	questionsSlice []int
	cmd            = exec.Command("clear")
)

func main() {
	rand.Seed(time.Now().Unix())
	var err error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stderr = os.Stderr
	studentsList, _, err = NewStudentsList("studentsInfo.txt")
	if err != nil {
		log.Println(errors.Wrap(err, `NewStudentsList("studentsInfo.txt")`))
		os.Exit(1)
	}
	questionsList, questionsSlice, err = NewQuestionsList("questions.txt")
	if err != nil {
		log.Println(errors.Wrap(err, `NewQuestionsList("questions.txt")`))
		os.Exit(1)
	}

	start()
}

func start() {
	fmt.Println("Добро пожаловать на зачёт!")

	for {
		fmt.Println("Введите номер студента, который будет отвечать: ")
		var studentNumString string
		_, err := fmt.Scanf("%s\n", &studentNumString)
		if err != nil {
			log.Println(errors.Wrap(err, `fmt.Scanf("%s\n", &studentNumString)`))
			return
		}
		studentNum, err := strconv.Atoi(studentNumString)
		if err != nil {
			//clearCommand()
			fmt.Println("Вы ввели некорректный номер. Попробуйте ещё раз.")
			continue
		}
		sc, ok := studentsList[studentNum]
		if !ok {
			//clearCommand()
			fmt.Println("Не найден студент по введенному номеру. Попробуйте ещё раз.")
			continue
		}
		fmt.Printf("Отвечает: \t\t%s\nКоличество вопросов: \t%d\n", sc.Name, sc.NumOfQuestions)
		questionsNums := getQuestionNums(sc.NumOfQuestions, questionsSlice)
		sort.Slice(questionsNums, func(i, j int) bool {
			return questionsNums[i] < questionsNums[j]
		})
		fmt.Println("Вопросы:")
		for _, qn := range questionsNums {
			fmt.Printf("\n%d. \t%s\n", qn, questionsList[qn])
		}
		cont()
	}

}

func cont() {
	fmt.Printf("\nЕсли студент закончил отвечать и готовы перейти к следующему студенту, то нажмите 'y'\n")
	for {
		var oper string
		_, err := fmt.Scanf("%s\n", &oper)
		if err != nil {
			log.Println(errors.Wrap(err, `fmt.Scanf("%s\n", &studentNumString)`))
			return
		}
		switch oper {
		case "y":
			//clearCommand()
			return
		default:
			//clearCommand()
			fmt.Println("Вы ввели неверную команду. Попробуйте ещё раз.")
			continue
		}
	}
}

func clearCommand() {
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func config() {

}
