package main

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)

type Response struct {
	Messages []struct {
		Name string
	}
}

func main() {
	resp, err := http.Get("http://demo8948446.mockable.io/test")

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	// _, err = io.Copy(os.Stdout, resp.Body)

	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		log.Fatal(err)
	}

	for _, child := range r.Messages {
		fmt.Println(child.Name)
	}

}
