package tokenDictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetIDEXTokenList() (currencies map[string]TokenInformation, err error) {

	const IDEX_url string = "https://api.idex.market/returnCurrencies"
	currencies = make(map[string]TokenInformation)

	client := &http.Client{}

	req, err := http.NewRequest("POST", IDEX_url, nil)
	if err != nil {
		fmt.Println("Error with the http.Request\n")
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error with Client Do \n")
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error with ReadAll ")
		return nil, err
	}

	err = json.Unmarshal([]byte(data), &currencies)
	if err != nil {
		fmt.Printf("Error with Unmarshall %s ", err)
	}

	return currencies, nil
}

type TokenInformation struct {
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	Address  string `json:"address"`
}
