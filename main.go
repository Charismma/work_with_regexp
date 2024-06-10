package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//Чтение из консоли входной файл
	reader_in := bufio.NewReader(os.Stdin)
	fmt.Println("Введите путь к файлу источнику: ")
	str_in, err := reader_in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	str_in = strings.TrimSpace(str_in)
	//считываем файл в переменную content
	content, err := ioutil.ReadFile(str_in)
	if err != nil {
		panic(err)
	}
	//Чтение из консоли выходного файла, если название не указано создается в папке дефолтный файл
	//Если файл будет существовать то функция его перезапишет
	reader_out := bufio.NewReader(os.Stdin)
	fmt.Println("Введите имя файла для вывода: ")
	str_out, err := reader_out.ReadString('\n')
	str_out = strings.TrimSpace(str_out)
	if err != nil {
		panic(err)
	}
	if str_out == "" {
		str_out = "out.txt"
	}
	file, err := os.OpenFile(str_out, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer file.Close()
	writer := bufio.NewWriter(file)
	//регулярное выражение
	re := regexp.MustCompile(`([0-9]+)([+-/*]{1})([0-9]+)([=]{1})([?]{1})`)
	//делим строки на подстроки по нужному шаблону и несколькими группами захвата, чтобы было удобнее работать с ними
	submatch := re.FindAllStringSubmatch(string(content), -1)
	//преобразуем элементы из строк в integer и в конструкции switch
	//выбираем case в зависимости от выражения и если деление не забываем привести
	//к float и буферизированно записываем во writer
	for _, s := range submatch {
		el1, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		el3, err := strconv.Atoi(s[3])
		if err != nil {
			panic(err)
		}
		var res int
		var resdel float64
		switch s[2] {
		case "/":
			resdel = float64(el1) / float64(el3)
			obch := fmt.Sprintf("%v%v%v=%v\n", s[1], s[2], s[3], resdel)
			writer.Write([]byte(obch))
		case "-":
			res = el1 - el3
			obch := fmt.Sprintf("%v%v%v=%v\n", s[1], s[2], s[3], res)
			writer.Write([]byte(obch))
		case "*":
			res = el1 * el3
			obch := fmt.Sprintf("%v%v%v=%v\n", s[1], s[2], s[3], res)
			writer.Write([]byte(obch))
		case "+":
			res = el1 + el3
			obch := fmt.Sprintf("%v%v%v=%v\n", s[1], s[2], s[3], res)
			writer.Write([]byte(obch))

		}

	}
	//записываем все что в буфере в файл
	writer.Flush()
}
