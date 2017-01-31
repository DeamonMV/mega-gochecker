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

		if err != nil {

			return false
		}

		if (one1 >= 0) && (one1 <= 6) {

			if (one[1] == "media-err" ) || (one[1] == "other-err") {

				return true
			}
		}
	}
		return false
}

func main() {

	reader1 := bufio.NewReader(os.Stdin)
	text, _ := reader1.ReadString('\n')

	if input(text) == false{
		fmt.Println("Wrong Input")
		os.Exit(1)
	}


	slotkey := string("Slot Number")

	num := strings.Split(text, " ")

	counternum := int(1)



	if strings.Replace(num[1], "\n","", -1) == "media-err" {
		counternum = int(0)
	}

	counter := [2]string{"Media","Other"}


	c1 := exec.Command ("megacli", "-PDList", "-aALL" )
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

		if err != nil {

			if err == io.EOF {
				break
			}

				fmt.Println("Read Error:", err)
			return
		}

		splited := strings.Split(strings.Replace(scanner, "\n", "", -1), ":")

		if (splited[0] == slotkey) && (strings.Replace(splited[1], " ", "", -1) == num[0])  {

			slot = i
		}
		if (slot > 0) && (slot < slot+2)  {

			params := strings.Split(scanner, ":")

			if strings.Split(scanner, " ")[0] == counter[counternum] {

							fmt.Printf("%v", strings.Replace(params[1], " ", "", -1))
			}

			if (i == slot+2) {
				slot = 0
			}
		}

		i = i +1

	}
	if i < 2 {
		fmt.Println("Nothing")

	}
}