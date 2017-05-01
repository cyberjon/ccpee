package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"strconv"
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
type branch struct {
	TBs []chart `json:"tb"`
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

	for q < rn-1 {
		to := trans.TXs[q].Id
		td := trans.TXs[q+1].Id
		if to == td {
			foun.TXs = append(foun.TXs, trans.TXs[q])

			foun.TXs = append(foun.TXs, trans.TXs[q+1])
		}
		q++
	}
	vn := len(foun.TXs)
	cin := vn
	var getAll func(string, int, AllTx) AllTx

	getPrev := func(str string, tid string) (string, int, string) {
		var m, tii string
		var ind, n int
		m = "false"
		tii = ""
		n = -1
		resp, err := http.Get("https://1c965e4727c24d81ace6787e11ca2b4b-vp0.us.blockchain.ibm.com:5001/transactions/" + str)
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
					if string(a) == trans.TXs[i].Id {
						ind = i
						break
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
				if string(a) == trans.TXs[z].Id && spd == trans.TXs[z].Prev_Transaction_id {
					ti = z
					return ti
				}

			}

			z--
		}
		return ti
	}

	var brFinal = make([]AllTx, vn+1)
	var jsonFinal chart
	var jsonAsTrs AllTx
	var tid, tii, std string
	var getBranch func(string, AllTx, int)
	str := "5066133e-8352-4782-b9e4-6d85bfd16937"

	var n int
	co := 0
	cco := 0
	count := ""
	count2 := ""
	getAll = func(str string, ff int, prt AllTx) AllTx {
		var at Transaction
		var tk int
		tii = ""

		q = ff
		if q > 0 && str != "1" {
			tn := "0"
			td := trans.TXs[q-1].Id
			to := trans.TXs[q].Id
			if vn > 1 {
				tn = trans.TXs[q+1].Id
			}

			if to == td {
				fmt.Println("Loop 1")
				if co == 0 {
					count = trans.TXs[q-1].Prev_Transaction_id
					fmt.Println(count)
					cco++
				}
				count2 = str
				co++
				q--
				getBranch(str, prt, q)
			} else if to == tn {
				fmt.Println("Loop 2")
				if co == 0 {
					count = trans.TXs[q+1].Prev_Transaction_id
					fmt.Println(count)
					cco++
				}
				count2 = str
				co++
				q++
				getBranch(str, prt, q)
			} else {
				fmt.Println("Loop 3")
				str, q, tii = getPrev(str, "")
				tk = inField(tii, str, trans)
				at = trans.TXs[q]

				prt.TXs = append(prt.TXs, at)
				jsonFinal.TDs = append(jsonFinal.TDs, prt)
				jsonAsTrs = getAll(str, tk, founded)

			}

		}

		if str == "1" && co > 0 && lc == 0 {
			fmt.Println("Loop 4")
			at = trans.TXs[ff]
			prt.TXs = append(prt.TXs, at)
			lc++
		}

		if co > 0 {
			fmt.Println("Loop 5")
			co--
			if cco == 0 {
				count = count2
				fmt.Println(count)
			}
			cco--
			str, q, tii = getPrev(count, "")
			tk = inField(tii, str, trans)
			at = trans.TXs[tk]

			prt.TXs = append(prt.TXs, at)
			jsonFinal.TDs = append(jsonFinal.TDs, prt)

			jsonAsTrs = getAll(str, tk, founded)

		}

		return prt

	}
	hk := 0
	getBranch = func(str string, jsonAsTr AllTx, q int) {
		tn := "0"
		td := trans.TXs[q-1].Id
		to := trans.TXs[q].Id
		if vn > 1 {
			tn = trans.TXs[q+1].Id
		}
		if to == td {
			fmt.Println("Branch 3")
			hk++
			if hk < vn {
				q--
				jsonAsTrs = getAll(trans.TXs[q].Prev_Transaction_id, q, founded)
				cin--
			}
		} else if to == tn {
			fmt.Println("Branch 4")
			hk++
			if hk < vn {
				q++
				jsonAsTrs = getAll(trans.TXs[q].Prev_Transaction_id, q, founded)
				cin--
			}

		}
		brFinal[1] = jsonAsTrs
		return
	}

	std, n, tid = getPrev(str, "")
	n = inField("", str, trans)
	fmt.Println(n)
	fmt.Println(std)
	fmt.Println(tid)
	jsonAsTrs = getAll(str, n+1, founded)

	jsonAsBy, _ := json.Marshal(jsonFinal)
	fmt.Println(string(jsonAsBy))

}
