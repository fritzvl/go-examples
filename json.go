/**
 * Created with IntelliJ IDEA.
 * User: bogdan
 * Date: 1/4/13
 * Time: 3:38 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import "fmt"

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type AERISResponse struct {
	Success bool
	Error []byte "error"
	Response struct {
		Id string "id"
		Loc struct {
			Long float32 "long"
			Lat float32 "lat"
		} "loc"
	} "response"
}

func (fignia AERISResponse) test1(var1 int) (result1 int, err bool, tre string) {

	return result1, err, tre
}

func main() {

	var message AERISResponse
	message.test1()
	fmt.Println("Hello world!")
	const ClientID =  "kmimaLOSUt0k1jA2zW7Kh"
	const ClientSecret = "XKZ2SR0g93s4AMRzNy2aYcL6CS5ET9O7CRqi006m"

	resp, err := http.Get("http://api.aerisapi.com/observations/kyiv,ua?client_id="+ClientID+"&client_secret="+ClientSecret)

	if err !=nil {
		fmt.Println("Weather API error")
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err!=nil{
			fmt.Println("Response read error")
		}
		fmt.Println(string(body))

		parse_err := json.Unmarshal(body, &message)

		if parse_err!=nil{
			fmt.Println("JSON parse error")
			fmt.Println(parse_err)
		} else {
			if message.Success {
				fmt.Println("Success")
				fmt.Println(message.Response.Id)
				fmt.Println(message.Response.Loc.Lat)
				fmt.Println(message.Response.Loc.Long)
			}
		}
	}
}


