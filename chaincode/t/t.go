package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	sp := strings.Split(trd, "\n")
	trD, er := fmt.Printf(`{"tx": [{"bid": "%s", "fun": "%s", "id": "%s", "traderA": "%s", "traderB": "%s", "seller": "%s", "pointAmount": "%s", "prevTransactionId": "%s", "timestamp": "%s"}]}`, sp[1], sp[2], sp[3], sp[4], sp[5], sp[6], sp[7], sp[8], sp[9])
	if err != nil {
		log.Fatal(er)
	}
	trS := string(trD)
	byteArr := []byte(trS)
	var tt AllTxs
	json.Unmarshal(byteArr, &tt)
	var found AllTxs
	for i := range tt.TXs {
		found.TXs = append(found.TXs, tt.TXs[i])
	}
	jsonAsB, _ := json.Marshal(found)
	return jsonAsB
}
