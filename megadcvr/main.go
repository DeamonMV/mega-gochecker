package main

import (


	"fmt"
	"strings"
	"io"
	"os/exec"
	"bytes"
	"bufio"
)


func main() {

	slotkey := string("Slot Number")

	//c1 := exec.Command("cat", "data")
	c1 := exec.Command("megacli", "-PDList", "-aALL")
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
			SlotCount = SlotCount + 1
		}
	}

	SlotCountSplited := strings.Split((buffer.String())," ")
	i := 1
	if SlotCount > 0 {
			fmt.Println("{\n" +
					"	\"data\":[" )
		for i <= SlotCount {
			fmt.Println( "		{")
			fmt.Printf("			\"{#SLOTNUMBER}\":\"%v\"}", SlotCountSplited[i])
			if i < SlotCount  {
				fmt.Println(",")
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