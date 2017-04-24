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

// type Transac struct {
// 	I                 string `json:"i"`
// 	T                 string `json:"t"`
// 	Tr                string `json:"tr"`
// 	Id                string `json:"id"`
// 	Timestamp         string `json:"timestamp"`
// 	TraderA           string `json:"traderA"`
// 	TraderB           string `json:"traderB"`
// 	Seller            string `json:"seller"`
// 	PointAmount       string `json:"pointAmount"`
// 	PrevTransactionID string `json:"prevTransactionId"`
// }

// type AllTxs struct {
// 	TXs []Transac `json:"tx"`
// }
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
	var tID = "3ec98fdb-d35d-4556-8426-92e538c386ab"
	resp, err := http.Get("http://148.100.4.235:7050/transactions/" + tID)
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

	sp := strings.Split(sp1, "\x20")

	trD := `{"tx": [{"bid": "` + sp[1] + `", "fun": "` + sp[2] + `", "id": "` + sp[3] + `", "traderA": "` + sp[4] + `", "traderB": "` + sp[5] + `", "seller": "` + sp[6] + `", "pointAmount": "` + sp[7] + `", "prevTransactionId": "` + sp[8] + `", "timestamp": "` + sp[9] + `"}]}`

	something := json.RawMessage(trD)

	jsonAsB, _ := something.MarshalJSON()
	return jsonAsB, nil
}
