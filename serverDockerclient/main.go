package main

import "net/http"
import "encoding/json"
import op "./callops"

func main() {
	var containers []op.container
	resp, _ := http.Get("http://localhost:8080/api/continaers/json")
	jd := json.NewDecoder(resp.Body)
	err := jd.Decode(containers)
	resp.close()
}
