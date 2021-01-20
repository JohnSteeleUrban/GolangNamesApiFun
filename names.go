package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	type Person struct {
		FirstName string `json:"name"`
		LastName  string `json:"surname"`
		Region    string `json:"region"`
	}
	var people []Person

	fmt.Printf("\nPlease enter the number of names to retrieve (up to 500).\n")
	reader := bufio.NewReader(os.Stdin)
	num, _ := reader.ReadString('\n')
	num = strings.TrimSuffix(num, "\n")
	url := "https://uinames.com/api/?amount=" + num

	//fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal([]byte(body), &people)
	len := len(people)
	fmt.Println(num + "Names")
	for i := 0; i <= len; i++ {
		//	fmt.Println(i)
		fmt.Println(people[i].FirstName + " " + people[i].LastName)
	}
	fmt.Printf("People : %+v", people)
}
