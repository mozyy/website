package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"go.yyue.dev/common/utils"
	"go.yyue.dev/postgres"
)

type House struct {
	CardType        string      `json:"cardType"`
	CityID          string      `json:"cityId"`
	HouseCode       string      `json:"houseCode"`
	ResblockID      string      `json:"resblockId"`
	DelegateID      string      `json:"delegateId"`
	Title           string      `json:"title"`
	Desc            string      `json:"desc"`
	BangdanTitle    string      `json:"bangdanTitle"`
	RecoDesc        string      `json:"recoDesc"`
	TotalPrice      string      `json:"totalPrice"`
	TotalPriceCount string      `json:"totalPriceCount"`
	UnitPrice       string      `json:"unitPrice"`
	JumpURL         string      `json:"jumpUrl"`
	ListPictureURL  string `json:"listPictureUrl"`
	RecommendReason string `json:"recommendReason"`
	ColorTags       []struct {
		Key       string `json:"key"`
		Title     string `json:"title"`
		Color     string `json:"color"`
		TextColor string `json:"textColor"`
		BgColor   string `json:"bgColor"`
	} `json:"colorTags"`
	HouseStatus  string `json:"houseStatus"`
	IsCtypeHouse bool   `json:"isCtypeHouse"`
	FbExpoID     string `json:"fbExpoId"`
	StatusSwitch struct {
		IsYeZhuTuijian bool `json:"isYeZhuTuijian"`
		IsHaofang      bool `json:"isHaofang"`
		IsYezhuPay     bool `json:"isYezhuPay"`
		IsVr           bool `json:"isVr"`
		IsKey          bool `json:"isKey"`
		IsNew          bool `json:"isNew"`
	} `json:"statusSwitch"`
	BrandTags interface{}   `json:"brandTags"`
	RecomTag  []interface{} `json:"recomTag"`
	HotTop    struct {
		DspAgentUcID string `json:"dspAgentUcId"`
		HotTopDigV   string `json:"hotTopDigV"`
		HotTop       int    `json:"hotTop"`
	} `json:"hotTop"`
	FramePicture []interface{} `json:"framePicture"`
	HousePicture []interface{} `json:"housePicture"`
	MarketBooth  interface{} `json:"marketBooth"`
	PoiDistance  string        `json:"poiDistance"`
	PoiIcon      string        `json:"poiIcon"`
}

type Data struct {
	Code int `json:"code"`
	Data struct { 
		Data struct {
			GetErShouFangList struct {
				adxParams interface{}
				aladinCard interface{}
				breadcrumbs interface{}
				bsRecommendList interface{}
				curPage int
				fbQueryId string
				HasMoreData int `json:"hasMoreData"`
				keyword string
				lessResultReason int
				List []House `json:"list"`
				nearWenan string
				noResultLiuzi interface{}
				noResultReason interface{}
				recommendExtInfo interface{}
				returnCount int
				selected interface{}
				spellCheck interface{}
				tdk interface{}
				totalCount int
			} `json:"getErShouFangList"`
		} `json:"data"`
		errkeys interface{}
		errno int
		msg string
	} `json:"data"`
	msg string
}


func main() {
	
	db, err := postgres.Open("user=postgres password=123456 dbname=beike sslmode=disable")
	utils.PanicErr(err)

	c := colly.NewCollector()

	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1079")
	utils.PanicErr(err)
	c.SetProxyFunc(rp)

	// f, err := os.Create("./beike/result.json")
	// utils.PanicErr(err)
	// defer f.Close()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL, r.Body)
	})

	// _,err =	f.WriteString("[")
	// utils.PanicErr(err)

	c.OnResponse(func(cr *colly.Response) {

		houses, err := getDataHouse(cr.Body)
		if err!= nil {
			var input string
			fmt.Printf("url: %s, error: %s", cr.Request.URL, err)
			for strings.ToLower(input) != "Y" {
				fmt.Println(`input "Y" goon:`)
				fmt.Scanln(&input)
			}
			cr.Request.Retry()
		}		
		// _,err =	f.WriteString(","+s[1:len(s)-1])
		// utils.PanicErr(err)
		err = insert(db, houses)
		utils.PanicErr(err)
	})

	c.OnError(func (cr *colly.Response, err error) {
			var input string
			fmt.Printf("url: %s, error: %s", cr.Request.URL, err)
			for strings.ToLower(input) != "Y" {
				fmt.Println(`input "Y" goon:`)
				fmt.Scanln(&input)
			}
			cr.Request.Retry()
	})

	for i:=1;i < 4579;i++ {
		c.Visit("https://m.ke.com/liverpool/api/ershoufang/getList?cityId=510100&curPage="+strconv.Itoa(i))
	}
	c.Wait()
	//  _,err =	f.WriteString("]")
	//  utils.PanicErr(err)

	// body, err := ioutil.ReadFile("./beike/data.json")
	// utils.PanicErr(err)
	// f, err := os.Create("./beike/result.json")
	// utils.PanicErr(err)
	// defer f.Close()
	// b, err := formatData([]byte(body))
	// houses, err:= getDataHouse(body)
	// utils.PanicErr(err)



	//  _,err =	f.WriteString(b)
	//  utils.PanicErr(err)

	//  f.Sync()
	}

func getDataHouse(body []byte) ([]House, error) {
	var value Data
	err := json.Unmarshal(body, &value)
	return value.Data.Data.GetErShouFangList.List, err
}

func formatData (body []byte) (string,error) {

	house, err := getDataHouse(body)
	result, err := json.Marshal(house)
	if (err != nil) {
		return "", err
	}
	return string(result), nil
}

func insert (db *postgres.DB, houses []House) error {
	var b strings.Builder
	fmt.Fprintf(&b, `
	INSERT INTO %s
	VALUES`, "beike")
	for i, h:= range houses {
		colorTags, _ :=  json.Marshal(h.ColorTags)
		statusSwitch, _ :=  json.Marshal(h.StatusSwitch)
		hotTop, _ :=  json.Marshal(h.HotTop)
		framePicture, _ :=  json.Marshal(h.FramePicture)
		housePicture, _ :=  json.Marshal(h.HousePicture)
		marketBooth, _ :=  json.Marshal(h.MarketBooth)
		brandTags, _ :=  json.Marshal(h.BrandTags)
		fmt.Fprintf(&b, `
		('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %t, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')`,
		h.HouseCode, h.CardType, h.CityID, h.ResblockID, h.DelegateID, h.Title,h.Desc,h.BangdanTitle, 
		h.RecoDesc, h.TotalPrice, h.TotalPriceCount,h.UnitPrice, h.JumpURL,h.ListPictureURL,h.RecommendReason, 
		colorTags, h.HouseStatus, h.IsCtypeHouse, h.FbExpoID, statusSwitch, hotTop,h.PoiDistance, h.PoiIcon, 
		framePicture, housePicture,marketBooth,brandTags)
		if i<len(houses)-1 {
			fmt.Fprintf(&b, ", ")
		}
	}
	_, err := db.DB.Exec(b.String())
	if err!= nil {
		fmt.Println(b.String())
	}
	return err
}