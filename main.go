package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

)


func main() {

	logFile, fileError := os.Create("history.log")

	if fileError != nil {
		log.Fatal(fileError)
	}

	defer logFile.Close()

	log.SetOutput(logFile)

	if len(os.Args) != 4 {
		fmt.Println("Использование:  <путь_к_каталогу> <текст_для_замены> <новый_текст>")
		return
	}


	dirPath := os.Args[1]
	searchText := os.Args[2]
	replaceText := os.Args[3]


	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() {

			return replaceInFile(path,searchText,replaceText,info)
			
		}
		return nil

	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Замена выполнена успешно!")
}

func replaceInFile(path string, searchText string, replaceText string, info os.FileInfo) error {
	var result strings.Builder
	start := 0
	resultIndex:=0

	content, err := os.ReadFile(path)

	if err != nil {
		 return err
	 }
	var text = string(content)

	for {

		index := strings.Index(text[start:], searchText)

		if index == -1 {
			 result.WriteString(text[start:])
			 resultIndex+=len(text[start:])
			break
		}

		result.WriteString(text[start : start+index])
		result.WriteString(replaceText)
		resultIndex+=len(text[start : start+index])
		position := start + index

		var logString = fmt.Sprintf("\n Имя файла: %s, Замена на позиции:  %d, %s -> %s",
		info.Name(),
		position,
		text[position:position+len(searchText)],
		result.String()[resultIndex:resultIndex+len(replaceText)])

		log.Println(logString);
		fmt.Println(logString)

		resultIndex+=len(replaceText)

		start += index + len(searchText)
	}
	
	err = os.WriteFile(path, []byte(result.String()), info.Mode())
	if err != nil {
		return err
	}
	
	return nil
}