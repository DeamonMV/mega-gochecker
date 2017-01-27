package main

import (
	"bytes"
	"io"
	"os/exec"
	"fmt"
	"strings"
)

func main() {

	c1 := exec.Command ("megacli", "-LDInfo", "-Lall", "-aALL" )
	c2 := exec.Command("grep", "State")

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var outbuf bytes.Buffer
	c2.Stdout = &outbuf

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
	stdout := outbuf.String()

	if len(stdout) > 0 {
		splited := strings.Split(stdout, ":")
		rmspace := strings.Replace(splited[1], " ", "", -1)
		fmt.Printf("%v", rmspace)

		} else {

		fmt.Printf("Nothing\n")
		}
	}


