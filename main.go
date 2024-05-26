package main

import (
	"bufio"
	// "errors"
	"fmt"
	// "io"
	// "net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const url = "https://cheat.sh/"

func main() {
	curlRequest(prepareReq())
}

func prepareReq() string {
	scanner := bufio.NewScanner(os.Stdin)
	var sb strings.Builder
	langs := []string{"java", "go", "javascript", "lua", "bash", "zsh", "csharp"}
	fmt.Println("[0].Java [1].Go [2].Javascript [3].Lua [4].Bash [5].Zsh [6].Csharp")
	fmt.Println("----------------------------------------------")
	fmt.Println("Pick:")
	scanner.Scan()
	indx := scanner.Text()
	fmt.Println("type your keyword:")
	scanner.Scan()
	keyword := scanner.Text()
	fmt.Println("----------------------------------------------")
	sb.WriteString(url)
	intId, err := strconv.Atoi(indx)
	if err != nil {
		fmt.Println("fjlksjdfljsl")
		return ""
	}

	sb.WriteString(langs[intId])
	sb.WriteString("/")
	sb.WriteString(keyword)
	reqString := sb.String()
	return reqString

}
func curlRequest(reqUrl string) {
	cmd := exec.Command("curl", reqUrl)
	// fmt.Println("ping started..")
	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}

	// Print the command output
	fmt.Println(string(output))
}
