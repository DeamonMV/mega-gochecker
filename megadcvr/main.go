package main

import (


	"fmt"
	"strings"
	"strconv"

	"io"
	"os/exec"
	"bytes"
	"bufio"
)

func input(param string) bool {

	param = strings.Replace(param, "\n","", -1)
	if len(param) <= 3 && len(param) > 0 {
		one := strings.Split(param, ";")

		one1, err := strconv.Atoi(one[0])
		//fmt.Printf("one1 %v\n", one1)
		if err != nil {
			fmt.Printf("%v", err)
		}
		one2, err := strconv.Atoi((one[1]))
		//fmt.Printf("one2 %v\n", one2)
		if err != nil {
			fmt.Printf("%v", err)
		}

		if (one1 > 0) && (one1 <= 6) {
			if (one2 == 0) || (one2 == 1) {
				return true
			}
		}
	}

	return false

}

func main() {

	slotkey := string("Slot Number")

	c1 := exec.Command("cat", "data")
	//c1 := exec.Command("megacli", "-PDList", "-aALL")
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
	var buffer bytes.Buffer

	SlotCount := 0

	for {

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

		if (splited[0] == slotkey)  {
			buffer.WriteString(splited[1])
			//fmt.Printf("%v\n", splited[1])
			SlotCount = SlotCount + 1

		}

		//fmt.Println(scanner.Text())
		//fmt.Println(rmspace)
		//fmt.Println(splited[1])
	}

	//fmt.Println(buffer.String())
	//fmt.Println(SlotCount)


	SlotCountSplited := strings.Split((buffer.String())," ")
	i := 1
	if SlotCount > 0 {
			fmt.Println("{\n" +
					"	\"data\":[" )
		for i <= SlotCount {
			fmt.Println( "		{")
			fmt.Printf("			\"{#SLOTNUMBER}\":\"%v\"}", SlotCountSplited[i])
			if i < SlotCount  {
				fmt.Println(",\n")
			} else {
				fmt.Println("]}")
			}
			i ++
		}
			fmt.Println("")
	} else {
		fmt.Println("Empty")
	}
}