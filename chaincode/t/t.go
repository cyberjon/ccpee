package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

	res, err := http.Get("http://127.0.0.1:8080/t.json")
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

		resp, err := http.Get("https://3bdbeca04a864ccd8530ed61cecd741a-vp0.us.blockchain.ibm.com:5003/transactions/" + str)
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
	inField := func(ssd string, spd string, trans AllTx) int {
		var ti int

		z := rn - 1
		for z >= 0 {

			a := []byte(ssd)
			if len(a) > 0 {
				copy(a[0:], a[1:])
				a[len(a)-1] = 0
				a = a[:len(a)-1]

				t, err := strconv.Atoi(string(a))
				if err != nil {
					fmt.Println(err)
				}
				tm, _ := strconv.Atoi(trans.TXs[z].Id)
				if t == tm && spd == trans.TXs[z].Prev_Transaction_id {
					ti = z
					return ti
				}
			}

			z--
		}
		return ti
	}
	var jsonFinal chart
	var jsonAs []byte
	var jsonAsTrs AllTx
	//var jsonAsTr AllTx
	var getBranch func(string, AllTx, int)
	str := "4c1da8f5-4c76-473c-ad86-a08cfb5582d8"
	inf := 0
	var ls, lt int

	getAll := func(str string, ff int, prt AllTx) (AllTx, int) {
		var at Transaction
		var lst int

		var tii string
		tii = ""
		count := 0
		ttr := str

	T:
		at, ls = findIndex(str, trans)
		if ls == ff {
			lt = ls - 2
		}

		if ttr == "1" {
			str, _, tii = getPrev(str, "")

			if count < 1 {
				prt.TXs = append(prt.TXs, trans.TXs[ff])
				count++
				goto T
			}
			return prt, inf

		} else if at.Prev_Transaction_id != "" {

			lst = inField(tii, str, trans)

			q = lst
			if q > 0 {
				to := trans.TXs[q].Id
				td := trans.TXs[q-1].Id

				if to == td {
					getBranch(str, prt, q)
					return prt, inf
				} else {
					str, _, tii = getPrev(str, "")
					prt.TXs = append(prt.TXs, at)
					goto T

				}
				q--
			} else {
				str, _, tii = getPrev(str, "")
				prt.TXs = append(prt.TXs, at)
				goto T

			}

		} else {
			lst = inField(tii, str, trans)

			inf = lst

			prt.TXs = append(prt.TXs, trans.TXs[lst])

			return prt, inf
		}
		return prt, inf

	}
	getBranch = func(str string, jsonAsTr AllTx, q int) {
		if q > 0 {
			to := trans.TXs[q].Id
			td := trans.TXs[q-1].Id
			if to == td {
				foun.TXs = append(foun.TXs, trans.TXs[q])

				jsonAsTr, _ = getAll(trans.TXs[q].Prev_Transaction_id, q, founded)
				jsonFinal.TDs = append(jsonFinal.TDs, jsonAsTr)
				// jsonAs, _ = json.Marshal(jsonFinal)
				// 				fmt.Println(string(jsonAs))
				// 				fmt.Println("----------------------")
				jsonAsTr, _ = getAll(trans.TXs[q-1].Prev_Transaction_id, q-1, founded)
				jsonFinal.TDs = append(jsonFinal.TDs, jsonAsTr)
				jsonAs, _ = json.Marshal(jsonFinal)
				// fmt.Println(string(jsonAs))
				// 				fmt.Println("----------------------")
				return
			}
			q--
		}
		return
	}

	str, _, _ = getPrev(str, "")

	jsonAsTrs, inf = getAll(str, 0, founded)

	jsonFinal.TDs = append(jsonFinal.TDs, jsonAsTrs)
	// q = inf
	jsonAsBy, _ := json.Marshal(jsonFinal)
	fmt.Println(string(jsonAsBy))

}
