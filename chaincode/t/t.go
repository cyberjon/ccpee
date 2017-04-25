package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	//"unicode/utf8"
)

type Transac struct {
	Bid               string `json:"bid"`
	Fun               string `json:"fun"`
	Id                string `json:"id"`
	Timestamp         string `json:"timestamp"`
	TraderA           string `json:"traderA"`
	TraderB           string `json:"traderB"`
	Seller            string `json:"seller"`
	PointAmount       string `json:"pointAmount"`
	PrevTransactionID string `json:"prevTransactionId"`
}

type AllTxs struct {
	TXs []Transac `json:"tx"`
}
type Transact struct {
	Cert        string `json:"cert"`
	ChaincodeID string `json:"chaincodeID"`
	Nonce       string `json:"nonce"`
	Payload     string `json:"payload"`
	Signature   string `json:"signature"`
	Timestamp   string `json:"nanos"`
	Txid        string `json:"txid"`
	Type        int    `json:"type"`
}

func main() {
	tr, er := mainReturnWithCode()
	if er != nil {
		log.Fatal(er)
	}

	fmt.Println(string(tr))

}
func mainReturnWithCode() ([]byte, error) {
	var str []string
	c := 0
	//var x = []byte{}

	var tID = "98d9daad-a66a-414a-9fa3-a1610f08b21c"
M:
	resp, err := http.Get("https://eaf64d13f6fc4d5caeacc5be900d20f0-vp0.us.blockchain.ibm.com:5003/transactions/" + tID)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	byteArray := []byte(body)
	var t Transact
	json.Unmarshal(byteArray, &t)
	st, err := base64.StdEncoding.DecodeString(t.Payload)
	if err != nil {
		log.Fatal(err)
	}
	trd := string(st)
	sp1 := strings.Replace(trd, "\n", " ", -1)
	sp := strings.Split(sp1, " ")
	var se string
	if sp[8] != "1" {
		se = strings.Replace(sp[8], "$", "", 1)
	} else if sp[8] == "1" {
		se = sp[8]
	}

	// trD := `{"bid": "` + sp[1] + `", "fun": "` + sp[2] + `", "id": "` + sp[3] + `", "traderA": "` + sp[4] + `", "traderB": "` + sp[5] + `", "seller": "` + sp[6] + `", "pointAmount": "` + sp[7] + `", "prevTransactionId": "` + sp[8] + `", "timestamp": "` + sp[9] + `"}`
	//
	// 	something := json.RawMessage(trD)
	//
	// 	jsonAsB, _ := something.MarshalJSON()
	// 	var tt AllTxs
	// 	json.Unmarshal(jsonAsB, &tt)

	//jsonAsBy := []byte(sp[3])

	// str = append(str, se)
	// 	tID = se
	// 	c++
	// 	if c < 5 {
	// 		goto M
	// 	}
	if se == "1" {
		str = append(str, se)
	} else {

		str = append(str, se)
		tID = se

		if c < 3 {
			c++
			goto M

		}
	}
	stringByte := "\x00" + strings.Join(str, "\x20\x00")
	jsonAsBy := []byte(stringByte)
	//jsonAsBy, _ := json.Marshal(ids)
	return jsonAsBy, nil
}
