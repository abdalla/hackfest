package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	template := fmt.Sprintf("%s-templ.yaml", os.Args[1])
	fmt.Println("TEMPLATE: ", template)

	cmd := exec.Command("sam", "local", "start-api", "-t", template)

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
