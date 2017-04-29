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

type Transaction struct {
	Id                  string `json:"txID"`         //Transaction ID from cppe system
	Timestamp           string `json:"EX_TIME"`      //utc timestamp of creation
	TraderA             string `json:"USER_A_ID"`    //UserA ID
	TraderB             string `json:"USER_B_ID"`    //UserB ID
	Seller              string `json:"SELLER_ID"`    //UserA's Seller ID
	Point_Amount        string `json:"POINT_AMOUNT"` //Points owned by UserA after exchange
	Prev_Transaction_id string `json:"PREV_TR_ID"`
}

type AllTx struct {
	TXs []Transaction `json:"tx"`
}
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
type chart struct {
	TDs []AllTx `json:"td"`
}

func main() {
	// tr, er := mainReturnWithCode()
	// 	if er != nil {
	// 		log.Fatal(er)
	// 	}
	mainReturnWithCode()

}

func mainReturnWithCode() {
	var prid string
	q := 0

	res, err := http.Get("http://cyberjon.com/wp-content/tr.json")
	if err != nil {
		// handle error
	}
	defer res.Body.Close()
	bod, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	byteAr := []byte(bod)
	var trans AllTx
	json.Unmarshal(byteAr, &trans)
	rn := len(trans.TXs)
	var founded AllTx
	var foun AllTx
	//var jsonAsByte byte
	for q < rn-1 {
		to := trans.TXs[q].Id
		td := trans.TXs[q+1].Id
		if to == td {
			foun.TXs = append(foun.TXs, trans.TXs[q])

			foun.TXs = append(foun.TXs, trans.TXs[q+1])
		}
		q++
	}
	//vn := len(foun.TXs)
	findIndex := func(str string, trans AllTx) (Transaction, int) {
		var q Transaction
		t := 0
		for i := 0; i < rn; i++ {
			t++
			if t > rn {
				break
			}
			if trans.TXs[i].Prev_Transaction_id == str {
				return trans.TXs[i], i
			}

		}
		return q, -2
	}
	getPrev := func(str string, tid string) (string, int, string) {
		var m, tii string
		var ind, n int
		m = "false"
		tii = ""
		n = -1
		if str == "1" && tid != "" {
			q := 0
			for l := 0; l < rn; l++ {
				q++
				if q > rn {
					break
				}
				if trans.TXs[l].Id == tid {
					ind = l
				}

			}
			return m, ind, tii
		}
		resp, err := http.Get("https://676a3275f6de482aab0e0e9929cccda3-vp0.us.blockchain.ibm.com:5003/transactions/" + str)
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
		rpl := strings.NewReplacer("$", "",
			`%`, "")
		if len(trd) > 0 {
			prid = rpl.Replace(sp[8])
			tn := sp[3]

			t := 0
			for i := 0; i < rn; i++ {
				t++
				if t > rn {
					break
				}
				if trans.TXs[i].Id == tn {
					ind = i
				}

			}
			return prid, ind, tn
		}
		return m, n, tii

	}
	//var inField func(string, AllTx) int
	inField := func(ssd string, trans AllTx) int {
		var ti int

		z := 0
		for z < rn {

			if ssd == "0006" && trans.TXs[z].Prev_Transaction_id == "1" {
				ti = z
				return ti
			}
			z++
		}
		return ti
	}
	var jsonFinal chart
	str := "8600284d-deb7-47f9-9785-42b110e64b14"
ABAR:
	inf := 0
	var getAll func(string, int, AllTx) AllTx
	getAll = func(str string, ff int, prt AllTx) AllTx {
		var at Transaction
		var lst int
		var ttr, tii string
		tii = ""
		count := 0
		ttr = str

	T:
		at, _ = findIndex(str, trans)

		if ttr == "1" {
			prt.TXs = append(prt.TXs, trans.TXs[ff])
			if count < 1 {
				str, _, tii = getPrev(ttr, "")
				count++
				goto T
			}

		} else if at.Prev_Transaction_id != "" {
			str, _, tii = getPrev(str, "")
			prt.TXs = append(prt.TXs, at)
			goto T

		} else {
			lst = inField(tii, trans)
			inf = lst
			prt.TXs = append(prt.TXs, trans.TXs[lst])
			return prt
		}
		inf = 0
		return prt

	}

	//jsonAsTr := getAll(str, 0, founded)

	jsonAsTrs := getAll(str, 0, founded)

	jsonFinal.TDs = append(jsonFinal.TDs, jsonAsTrs)
	q = inf

	if q > 0 {

		//jsonAsTr := getAll(str, 0, founded)
		to := trans.TXs[q].Id
		td := trans.TXs[q-1].Id
		fmt.Println(to)
		fmt.Println(td)
		if to == td {
			foun.TXs = append(foun.TXs, trans.TXs[q])
			jsonAsTr := getAll(trans.TXs[q].Prev_Transaction_id, 1, founded)

			jsonFinal.TDs = append(jsonFinal.TDs, jsonAsTr)

			foun.TXs = append(foun.TXs, trans.TXs[q-1])
			//g := findLast(trans.TXs[q-1].Prev_Transaction_id, trans.TXs[q-1].Id)
			jsonAsTr = getAll(trans.TXs[q-1].Prev_Transaction_id, 4, founded)
			//fmt.Println(g)
			jsonFinal.TDs = append(jsonFinal.TDs, jsonAsTr)
			r := inf
			if r > 0 {
				goto ABAR
			}
		} else {
			goto ONTIM
		}
		q--
	}
ONTIM:
	jsonAsBy, _ := json.Marshal(jsonFinal)
	fmt.Println(string(jsonAsBy))

}
