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
	lc := 0
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
	//var jsonAsTr AllTx
	var getAll func(string, int, AllTx) AllTx

	getPrev := func(str string, tid string) (string, int, string) {
		var m, tii string
		var ind, n int
		m = "false"
		tii = ""
		n = -1
		resp, err := http.Get("https://1bb5c3cf7def48bcaef058393604b7e7-vp0.us.blockchain.ibm.com:5003/transactions/" + str)
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
			if tid != "" {
				//tn = tid

			}
			t := 0
			for i := 0; i < rn; i++ {
				t++
				if t > rn {
					break
				}
				a := []byte(tn)
				if len(a) > 0 {
					copy(a[0:], a[1:])
					a[len(a)-1] = 0
					a = a[:len(a)-1]

					t, err := strconv.Atoi(string(a))
					if err != nil {
						fmt.Println(err)
					}
					tm, _ := strconv.Atoi(trans.TXs[i].Id)
					if t == tm {
						ind = i
						break
						//return prid, ind, tn
					}

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
	var jsonAsTrs AllTx
	var tid, tii, std string
	var getBranch func(string, AllTx, int)
	str := "f972064f-4ac1-45eb-8347-dff93ecc265d"

	var n int
	//co := 0
	count := ""
	getAll = func(str string, ff int, prt AllTx) AllTx {
		var at Transaction
		var tk int
		tii = ""

		q = ff
		if q > 0 && str != "1" {
			td := trans.TXs[q-1].Id
			to := trans.TXs[q].Id
			tn := trans.TXs[q+1].Id
			if to == td {
				count = str
				getBranch(str, prt, q-1)
			} else if to == tn {
				count = str
				getBranch(str, prt, q+1)
			} else {
				str, q, tii = getPrev(str, "")
				tk = inField(tii, str, trans)
				at = trans.TXs[tk]

				prt.TXs = append(prt.TXs, at)
				jsonFinal.TDs = append(jsonFinal.TDs, prt)
				jsonAsTrs = getAll(str, tk, founded)

			}

		}
		if str == "1" && len(count) > 0 && lc == 0 {
			at = trans.TXs[ff]
			prt.TXs = append(prt.TXs, at)
			lc++
		}

		if len(count) > 0 {
			str, q, tii = getPrev(count, "")
			tk = inField(tii, str, trans)
			at = trans.TXs[tk]

			prt.TXs = append(prt.TXs, at)
			count = ""
			jsonFinal.TDs = append(jsonFinal.TDs, prt)
			jsonAsTrs = getAll(str, tk, founded)

		}

		return prt

	}

	getBranch = func(str string, jsonAsTr AllTx, q int) {
		td := trans.TXs[q-1].Id
		to := trans.TXs[q].Id
		tn := trans.TXs[q+1].Id
		if to == td {
			jsonAsTrs = getAll(trans.TXs[q].Prev_Transaction_id, q, founded)
		} else if to == tn {
			jsonAsTrs = getAll(trans.TXs[q].Prev_Transaction_id, q, founded)
		}
		return
	}

	std, n, tid = getPrev(str, "")
	n = inField("", str, trans)
	fmt.Println(n)
	fmt.Println(std)
	fmt.Println(tid)
	jsonAsTrs = getAll(str, 1, founded)
	jsonAsBy, _ := json.Marshal(jsonFinal)
	fmt.Println(string(jsonAsBy))

}
