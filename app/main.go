package main

import (
	"fmt"
	"log"
	"os"
)

const defaultDir = "files"

const (
	answerYes = "yes"
	answerNo  = "no"
)

func main() {
	err := os.MkdirAll(defaultDir, 0770)
	if err != nil {
		log.Fatal(err)
	}

	file, err := enterFileName()
	defer file.Close()

	for {
		fmt.Println("Please enter data for writing into the file")

		var fileData string
		_, err = fmt.Scan(&fileData)
		if err != nil {
			log.Fatal(err)
		}

		isOK := isWriteOK(file, fileData)
		if !isOK {
			break
		}

		fmt.Println("Success! Do you want to write something else? yes/no")

		isContinue, err := checkAnswer()
		if err != nil {
			log.Fatal(err)
		}

		if isContinue {
			continue
		}

		break
	}
}

func enterFileName() (*os.File, error) {
	var file *os.File

	for {
		fmt.Println("Please enter a file name")

		var fileName string
		_, err := fmt.Scan(&fileName)
		if err != nil {
			return nil, err
		}

		fPath := fmt.Sprintf("%s/%s", defaultDir, fileName)

		file, err = os.OpenFile(fPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0770)
		if err != nil {
			fmt.Println("Create file error:", err)

			continue
		}

		break
	}

	return file, nil
}

func checkAnswer() (bool, error) {
	var answer string

	for {
		_, err := fmt.Scan(&answer)
		if err != nil {
			return false, err
		}

		switch answer {
		case answerYes:
			return true, nil
		case answerNo:
			return false, nil
		default:
			fmt.Println("Incorrect answer. Do you want to write something else? yes/no")
			continue
		}
	}
}

func isWriteOK(file *os.File, fileData string) bool {
	_, err := file.WriteString(fileData + "\n")
	if err != nil {
		fmt.Println("Write data error:", err)

		err = os.Remove(file.Name())
		if err != nil {
			fmt.Println("Remove file error:", err)
		}

		return false
	}

	return true
}
