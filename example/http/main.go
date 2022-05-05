package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type result struct {
	Status   string `json:"status"`
	Password string `json:"password"`
}

func main() {
	var (
		passLen     int
		digitsCount int
		withUpper   bool
		allowRepeat bool

		res result
	)

	fmt.Println("Enter pass len:")
	fmt.Scan(&passLen)

	fmt.Println("Enter digits count:")
	fmt.Scan(&digitsCount)

	fmt.Println("Enter with upper:")
	fmt.Scan(&withUpper)

	fmt.Println("Enter allow repeat:")
	fmt.Scan(&allowRepeat)

	resp, err := http.Get(fmt.Sprintf("http://localhost:8888/api/pass?length=%d&digitsCount=%d&withUpper=%t&allowRepeat=%t", passLen, digitsCount, withUpper, allowRepeat))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	btRes, _ := ioutil.ReadAll(resp.Body)
	resp.Close = true

	json.Unmarshal(btRes, &res)

	fmt.Println(res.Password)
}
