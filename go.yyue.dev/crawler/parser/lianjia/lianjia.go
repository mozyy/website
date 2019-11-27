package lianjia

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"go.yyue.dev/crawler/engine"
	"go.yyue.dev/crawler/proto"
)

var domain = "https://cd.lianjia.com"

// ParseError is parse error
type ParseError struct {
	Func   string
	Detail string
	Err    error
}

// Unwarp is impl errors Unwarp
func (e *ParseError) Unwarp() error {
	return e.Err
}
func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error, %s: %s : %s", e.Func, e.Detail, e.Err)
}

// City parse city
func City(q *goquery.Document) (engine.Result, error) {
	result := engine.Result{}

	total, err := strconv.Atoi(strings.TrimSpace(q.Find(".resultDes span").Text()))
	if err != nil {
		return result, &ParseError{"City", "get total", err}
	}
	if total == 0 {
		return result, nil
	}
	err = nil
	// 其他区域 (链家二手房列表最多100页, 所以需要按区域爬取数据)
	if total > 3000 {
		areas := q.Find("div[data-role=ershoufang] a[href]")
		selectedLen := q.Find("div[data-role=ershoufang] .selected").Length()
		switch selectedLen {
		case 0:
			// 爬大行政区页面
			areas.Each(func(i int, s *goquery.Selection) {
				href, _ := s.Attr("href")
				result.Requests = append(result.Requests, engine.Request{URL: domain + href, Parser: City})
			})
			return result, nil
		case 1:
			// 爬小区域页面
			links := q.Find("div[data-role=ershoufang]>div").Last().Find("a[href]")
			links.Each(func(i int, s *goquery.Selection) {
				href, _ := s.Attr("href")
				result.Requests = append(result.Requests, engine.Request{URL: domain + href, Parser: City})
			})
			return result, nil
		default:
			// TODO: 如果还有多余数据, 再细化搜索条件
			// log.Println("total: ", total, selectedLen, q.Find("div[data-role=ershoufang] .selected").Text())
			// result.Items = append(result.Items, engine.Item{Payload: total})
			err = &ParseError{"City", fmt.Sprintf("large total[%d]", total), engine.ErrListLarge}
		}
	}

	// get house item request
	// result.Requests = append(result.Requests, cityItem(q)...)

	// get house summary item request
	result.Items = append(result.Items, houseSummary(q)...)

	// get house page house length
	// result.Items = []engine.Item{q.Find(".sellListContent>.LOGVIEWDATA").Length()}

	morePageRequests, errp := morePage(q)
	if errp != nil {
		err = &ParseError{"City", "more page", fmt.Errorf("%v-%v", err, errp)}
	}
	result.Requests = append(result.Requests, morePageRequests...)
	return result, err
}

var tagsClass = []string{".goodhouse_tag", ".taxfree", ".five", ".vr", ".subway", ".haskey"}

// houseSummary get house summary
func houseSummary(q *goquery.Document) []engine.Item {
	items := []engine.Item{}
	content := q.Find(".sellListContent>.LOGVIEWDATA")
	content.Each(func(i int, s *goquery.Selection) {
		URL, ok := s.Find(".title a").Attr("href")
		if !ok {
			return
		}
		image, _ := s.Find(".lj-lazy").Attr("src")

		HouseNo := idRegexp.FindStringSubmatch(URL)[1]

		followInfo := strings.Split(s.Find(".followInfo").Text(), " / ")
		follow, err := strconv.Atoi(strings.Replace(followInfo[0], "人关注", "", 1))
		if err != nil {
			follow = -1
		}

		tags := 0
		for i, v := range tagsClass {
			if s.Find(v).Length() > 0 {
				tags |= 1 << i
			}
		}

		item := proto.HouseSummary{
			HouseNo:    HouseNo,
			Url:        URL,
			Title:      s.Find(".title a").Text(),
			TotalPrice: s.Find(".totalPrice span").Text(),
			UnitPrice:  s.Find(".unitPrice span").Text(),
			Plot:       s.Find(".positionInfo a").Eq(0).Text(),
			Region:     s.Find(".positionInfo a").Eq(1).Text(),
			// Layout:      infos[0],
			// Area:        infos[1],
			// Face:        infos[2],
			// Decoration:  infos[3],
			// Floor:       infos[4],
			// HouseYear:   infos[5],
			// StructBuild: infos[6],
			Image:       image,
			Follow:      int32(follow),
			ReleaseTime: followInfo[1],
			Tags:        int32(tags),
		}
		infos := strings.Split(s.Find(".houseInfo").Text(), "|")
		for i, info := range infos {
			info = strings.TrimSpace(info)
			if info == "" || info == "暂无数据" {
				continue
			}
			switch i {
			case 0:
				item.Layout = info
			case 1:
				item.Area = info
			case 2:
				item.Face = info
			case 3:
				item.Decoration = info
			case 4:
				item.Floor = info
			case 5:
				item.HouseYear = info
			case 6:
				item.StructBuild = info
			}
		}
		items = append(items, item)
	})
	return items
}

// cityItem get house item request
func cityItem(q *goquery.Document) []engine.Request {
	requests := []engine.Request{}
	content := q.Find(".sellListContent>.LOGVIEWDATA")
	content.Each(func(i int, s *goquery.Selection) {
		URL, _ := s.Find(".title a").Attr("href")

		request := engine.Request{
			URL: URL,
			Parser: func(qq *goquery.Document) (engine.Result, error) {
				result := engine.Result{}
				house, err := House(qq, URL)
				if err == nil {
					result.Items = []engine.Item{house}
				}
				return result, err
			},
		}
		requests = append(requests, request)
	})
	return requests
}

// morePage get other page
func morePage(q *goquery.Document) ([]engine.Request, error) {
	var requests []engine.Request
	listNode := q.Find(".house-lst-page-box[page-url][page-data]")
	pageURL, _ := listNode.Attr("page-url")
	pageDataStr, _ := listNode.Attr("page-data")
	var pageData struct {
		TotalPage int
		CurPage   int
	}
	err := json.Unmarshal([]byte(pageDataStr), &pageData)
	if err != nil {
		return requests, err
	}
	if pageData.CurPage == 0 {
		return requests, nil
	}
	for cur, total := pageData.CurPage, pageData.TotalPage; cur < total; cur++ {
		nextURL := domain + strings.ReplaceAll(pageURL, "{page}", strconv.Itoa(cur+1))
		requests = append(requests, engine.Request{
			URL:    nextURL,
			Parser: City,
		})
	}
	return requests, nil
}

var idRegexp = regexp.MustCompile(`^https://cd.lianjia.com/ershoufang/(\d+).html`)

// House get house detail
func House(q *goquery.Document, URL string) (proto.House, error) {
	HouseNo := idRegexp.FindStringSubmatch(URL)[1]

	houseInfo := proto.HouseInfo{
		HouseNo:    HouseNo,
		Url:        URL,
		Title:      q.Find(".title-wrapper .title .main").Text(),
		SubTitle:   q.Find(".title-wrapper .title .sub").Text(),
		Region:     q.Find(".communityName .info").Text(),
		TotalPrice: q.Find(".price .total").Text(),
		UnitPrice:  strings.Replace(q.Find(".unitPriceValue").Text(), "元/平米", "", 1),
		RoomInfo:   q.Find(".room .mainInfo").Text(),
		RoomSub:    q.Find(".room .subInfo").Text(),
		TypeInfo:   q.Find(".type .mainInfo").Text(),
		TypeSub:    q.Find(".type .subInfo").Text(),
		AreaInfo:   q.Find(".area .mainInfo").Text(),
		AreaSub:    q.Find(".area .subInfo").Text(),
	}

	introduction := q.Find("#introduction")
	baseInfo := proto.HouseBaseInfo{HouseNo: HouseNo}
	introduction.Find(".base .content ul li").Each(func(i int, s *goquery.Selection) {
		s.Contents().Each(func(_ int, s *goquery.Selection) {
			if goquery.NodeName(s) == "#text" {
				value := strings.TrimSpace(s.Text())
				switch i {
				case 0:
					baseInfo.Layout = value // 房屋户型
				case 1:
					baseInfo.Floor = value // 所在楼层
				case 2:
					baseInfo.AreaBuild = value // 建筑面积
				case 3:
					baseInfo.StructHouse = value // 户型结构
				case 4:
					baseInfo.AreaInner = value // 套内面积
				case 5:
					baseInfo.BuildType = value // 建筑类型
				case 6:
					baseInfo.Face = value // 房屋朝向
				case 7:
					baseInfo.StructBuild = value // 建筑结构
				case 8:
					baseInfo.Decoration = value // 装修情况
				case 9:
					baseInfo.ElevatorRatio = value // 梯户比例
				case 10:
					baseInfo.Elevator = value // 配备电梯
				case 11:
					baseInfo.Property = value // 产权年限
				}
			}
		})
	})
	transactionInfo := proto.HouseTransactionInfo{HouseNo: HouseNo}
	introduction.Find(".transaction .content ul li span:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		value := strings.TrimSpace(s.Text())
		switch i {
		case 0:
			transactionInfo.ListingTime = value // 挂牌时间
		case 1:
			transactionInfo.TradingAuthority = value // 交易权属
		case 2:
			transactionInfo.LastTransaction = value // 上次交易
		case 3:
			transactionInfo.HousingPurposes = value // 房屋用途
		case 4:
			transactionInfo.HouseYear = value // 房屋年限
		case 5:
			transactionInfo.PropertyRights = value // 产权所属
		case 6:
			transactionInfo.MortgageInfo = value // 抵押信息
		case 7:
			transactionInfo.DocumentPhoto = value // 房本备件
		}
	})
	housePics := []*proto.HousePic{}
	q.Find(".smallpic li").Each(func(i int, s *goquery.Selection) {
		housePic := proto.HousePic{HouseNo: HouseNo}
		if desc, ok := s.Attr("data-desc"); ok {
			housePic.Description = desc
		} else {
			housePic.Description = "vr"
		}
		if picSmallURL, ok := s.Find("img").Attr("src"); ok {
			housePic.PicSmallUrl = picSmallURL
		}
		if picNormalURL, ok := s.Attr("data-src"); ok {
			housePic.PicNormalUrl = picNormalURL
		}
		if picLargeURL, ok := s.Attr("data-pic"); ok {
			housePic.PicLargeUrl = picLargeURL
		}
		housePics = append(housePics, &housePic)
	})
	house := proto.House{
		HouseNo:              HouseNo,
		HouseInfo:            &houseInfo,
		HouseBaseInfo:        &baseInfo,
		HouseTransactionInfo: &transactionInfo,
		HousePics:            housePics,
	}
	return house, nil
}
