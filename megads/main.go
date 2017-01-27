package main

import (


	"fmt"
	"strings"
	"strconv"
	"io"
	"os/exec"
	"bytes"
	"bufio"
	"os"
)

func input(param string) bool {

	param = strings.Replace(param, "\n","", -1)
	if len(param) == 11 {
		one := strings.Split(param, " ")

		one1, err := strconv.Atoi(one[0])
		//fmt.Printf("one1 %v\n", one1)
		if err != nil {
			//fmt.Printf("%v", err)
			return false
		}
		/*one2, err := strconv.Atoi((one[1]))
		//fmt.Printf("one2 %v\n", one2)
		if err != nil {
			//fmt.Printf("%v", err)
			return false

		}
		*/
		if (one1 >= 0) && (one1 <= 6) {
			if (one[1] == "media-err" ) || (one[1] == "other-err") {
				return true
			}
		}
	}
		return false

}

func main() {

	//reader1 := bufio.NewReader(os.Stdin)
		//fmt.Print("set Slot number and Mediad Error = 1 or Other error = 2 : ")
	//text, _ := reader1.ReadString('\n')
		//fmt.Println(text)

	text := string("0 other-err")
	if input(text) == false{
		fmt.Println("Wrong Input")
		os.Exit(1)
	}


	slotkey := string("Slot Number")
	num := strings.Split(text, " ")
	fmt.Printf("text   %v\n", text)
	counternum := int(1)

	if num[1] == "media-err" {
		counternum = int(0)
	}

	fmt.Printf("counternum  %v\n", counternum)
	//if err != nil {
	//	//fmt.Printf("%v", err)
	//}
	counter := [2]string{"Media","Other"}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	c1 := exec.Command("cat", "data")
	//c1 := exec.Command ("megacli", "-PDList", "-aALL" )
	c2 := exec.Command("egrep", "Enclosure Device ID:|Slot Number:|Inquiry Data:|Error Count:|state")

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	outbuf := bytes.NewBuffer([]byte{})
	c2.Stdout = outbuf

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()

	reader := bufio.NewReader(outbuf)

	i := 0
	slot := 0
	for  {

		scanner, err := reader.ReadString('\n')
			//	fmt.Println(scanner)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Read Error:", err)
			return
		}
		splited := strings.Split(strings.Replace(scanner, "\n", "", -1), ":")
			//fmt.Printf("%v\n", num)
			//fmt.Printf("%v\n", counternum)

		// находим в строке и сравниваем сразу с двумя условаия - во втором условии перед сравнением удаляем лишние пробелы
		if (splited[0] == slotkey) && (strings.Replace(splited[1], " ", "", -1) == num[0])  {
			//fmt.Printf("%v", splited[1])
			// запоминаем строку
			slot = i
		}
		// заходим в условие если номер строки не нулевой, и меньше чем номер_строки + 2 - тут мы будем видеть только две строки с нужными значениями
		if (slot > 0) && (slot < slot+2)  {
			// разделяем строку по ":" это будет параметр и значение
			params := strings.Split(scanner, ":")
			//fmt.Printf("%v\n", strings.Replace(params[1], " ", "", -1))
			//fmt.Printf("%v\n", strings.Split(scanner, ":")[0])
			// перед стравнением какой нам нужен счетчик - разделяем его через пробел и берем только первый элимент
			if strings.Split(scanner, " ")[0] == counter[counternum] {
				// выводим значение, в выводе убираем пробелы
				fmt.Printf("%v", strings.Replace(params[1], " ", "", -1))
			}
			// сбрасываем счетчик
			if (i == slot+2) {
				slot = 0
			}
		}
		//fmt.Println(scanner.Text())
		//fmt.Println(rmspace)
		//fmt.Println(splited[1])
		i = i +1

	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	}
}