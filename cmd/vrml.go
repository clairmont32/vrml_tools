// UNOFFICIAL TOOL IMPLEMENTING THE VRML API
package main

import (
	"encoding/json"
	"fmt"
	"github.com/clairmont32/httplib"
	vrml "vrml/pkg"
)

func main() {
	req := &httplib.FormRequest{
		BaseURL:  "https://api.vrmasterleague.com/",
		Endpoint: "Sponsors",
		Payload:  nil,
		Method:   "GET",
	}

	h := []httplib.Headers{{
		Key:   "User-Agent",
		Value: "Grenade Magnet Finally Scripting Stuff",
	}}

	resp, err := httplib.DefaultRequest(req, h)

	var sponsors vrml.Sponsors
	err = json.Unmarshal(resp, &sponsors)
	if err != nil {
		panic(err)
	}

	fmt.Println(sponsors[0].Url)

}
