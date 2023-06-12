/*
1. Использовать методы и структуры пакетов ioutils и regexp.
2. Программа должна принимать на вход 2 аргумента:
	имя входного файла и имя файла для вывода результатов.
3. Если не найден вывод, создать.
4. Если файл вывода существует, очистить перед записью новых результатов.
5. Использовать буферизированную запись результатов.
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Введите имя входного файла:")
	reader := bufio.NewReader(os.Stdin)

	// файл для чтения
	inputFileName, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadFile(string(inputFileName))
	if err != nil {
		panic(err)
	}

	// файл для записи
	fmt.Println("Введите имя файла для записи результатов:")
	outFileName, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	file, err := os.Create(string(outFileName))
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)

	re := regexp.MustCompile(`([0-9]+)([+-/*/]{1})([0-9]+)=?`)
	subs := re.FindAllStringSubmatch(string(content), -1)

	for i, s := range subs {

		// записывает в файл по 10 записей за раз
		if i%10 == 0 {
			writer.Flush()
		}

		a, _ := strconv.Atoi(s[1])
		b, _ := strconv.Atoi(s[3])
		var c int = 0

		switch s[2] {

		case "+":
			c = a + b
			writer.Write([]byte(s[0] + fmt.Sprint(c) + "\n"))

		case "-":
			c = a - b
			writer.Write([]byte(s[0] + fmt.Sprint(c) + "\n"))

		case "*":
			c = a * b
			writer.Write([]byte(s[0] + fmt.Sprint(c) + "\n"))

		case "/":
			if b == 0 {
				break // в случае деления на 0 пропустит запись
			}
			c = a / b // делит нацело
			writer.Write([]byte(s[0] + fmt.Sprint(c) + "\n"))
		}
	}
	writer.Flush()

}
