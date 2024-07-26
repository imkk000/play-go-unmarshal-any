package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var x struct {
		Status struct {
			Code json.Number `json:"code"`
		} `json:"status"`
	}
	n := 2000
	body := fmt.Sprintf(`{"status":{"code":"%d"}}`, n)
	if err := json.Unmarshal([]byte(body), &x); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("n:", n)
	fmt.Println(body)
	fmt.Printf("%+v\n", x)
	fmt.Println(x.Status.Code.Int64())
	fmt.Println(x.Status.Code.Float64())
	fmt.Println(x.Status.Code.String())
}
