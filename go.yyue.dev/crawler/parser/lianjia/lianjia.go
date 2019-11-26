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

//<ul class="sellListContent" log-mod="list">
//<li class="clear LOGVIEWDATA LOGCLICKDATA" data-lj_view_evtid="21625" data-lj_evtid="21624"
//data-lj_view_event="ItemExpo" data-lj_click_event="SearchClick"  data-lj_action_source_type="链家_PC_二手列表页卡片"
//data-lj_action_click_position="0" data-lj_action_fb_expo_id='' data-lj_action_fb_query_id=''
//data-lj_action_resblock_id="3011055763863" data-lj_action_housedel_id="106103125055" >
//<a class="noresultRecommend img LOGCLICKDATA" href="https://cd.lianjia.com/ershoufang/106103125055.html"
//  target="_blank" data-log_index="1"data-el="ershoufang" data-housecode="106103125055" data-is_focus=""data-sl="">
//  <!-- 热推标签、埋点 -->
//  <img src="https://s1.ljcdn.com/feroot/pc/asset/img/vr/vrlogo.png?_v=20191009144231" class="vr_item">
//  <img class="lj-lazy" src="https://s1.ljcdn.com/feroot/pc/asset/img/blank.gif?_v=20191009144231"
//  data-original="https://image1.ljcdn.com/510100-inspection/prod-c171a6ef-1680-4609-afdc-df9ed5827d40.jpg.296x216.jpg"
//  alt="成都温江温江大学城">
//</a>
//<div class="info clear">
//<div class="title">
//<a class="" href="https://cd.lianjia.com/ershoufang/106103125055.html" target="_blank" data-log_index="1"  data-el="ershoufang"
//data-housecode="106103125055" data-is_focus="" data-sl="">中铁丽景书香清水套三，可根据自己喜好装修</a>
//<!-- 拆分标签 -->
//</div>
//<div class="address">
//<div class="houseInfo">
//<span class="houseIcon"></span>
//<a href="https://cd.lianjia.com/xiaoqu/3011055763863/" target="_blank" data-log_index="1" data-el="region">中铁丽景书香 </a>
// | 3室2厅 | 89.23平米 | 东南 | 毛坯
//</div></div>
//<div class="flood"><div class="positionInfo"><span class="positionIcon"></span>
//低楼层(共32层)2018年建板塔结合  -  <a href="https://cd.lianjia.com/ershoufang/wenjiangdaxuecheng/" target="_blank">温江大学城</a>
//</div></div><div class="followInfo"><span class="starIcon"></span>117人关注 / 2个月以前发布</div>
//<div class="tag"><span class="good"><img src="https://img.ljcdn.com/beike/bikanhaofang/1559536880827.png"></span>
//  <span class="vr">VR房源</span>
//  <span class="haskey">随时看房</span></div>
//<div class="priceInfo"><div class="totalPrice"><span>90.4</span>万</div>
// <div class="unitPrice" data-hid="106103125055" data-rid="3011055763863" data-price="10132">
//<span>单价10132元/平米</span></div></div></div>
//<div class="listButtonContainer"><div class="btn-follow followBtn" data-hid="106103125055"><span class="follow-text">关注</span></div>
//<div class="compareBtn LOGCLICK" data-hid="106103125055" log-mod="106103125055" data-log_evtid="10230">加入对比</div></div></li>

var domain = "https://cd.lianjia.com"

func City(q *goquery.Document) (engine.Result, error) {
	result := engine.Result{}

	total, err := strconv.Atoi(strings.TrimSpace(q.Find(".resultDes span").Text()))
	if err != nil {
		return result, err
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
			err = fmt.Errorf("parse city error: total is too large: %d", total)
		}
	}

	result.Requests = append(result.Requests, cityItem(q)...)
	morePageRequests, errp := morePage(q)
	if errp != nil {
		err = fmt.Errorf("%s-%s", err, errp)
	}
	result.Requests = append(result.Requests, morePageRequests...)
	return result, err
}

func cityItem(q *goquery.Document) []engine.Request {
	requests := []engine.Request{}
	content := q.Find(".sellListContent>.LOGVIEWDATA")
	content.Each(func(i int, s *goquery.Selection) {
		URL, _ := s.Find(".title a").Attr("href")

		request := engine.Request{
			URL: URL,
			Parser: func(qq *goquery.Document) (engine.Result, error) {
				return House(qq, URL)
			},
		}
		requests = append(requests, request)
	})
	return requests
}

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
	for curPage, totalPage := pageData.CurPage, pageData.TotalPage; curPage < totalPage; curPage++ {
		nextURL := domain + strings.ReplaceAll(pageURL, "{page}", strconv.Itoa(curPage+1))
		requests = append(requests, engine.Request{
			URL:    nextURL,
			Parser: City,
		})
	}
	return requests, nil
}

var idRegexp = regexp.MustCompile(`^https://cd.lianjia.com/ershoufang/(\d+).html`)

func House(q *goquery.Document, URL string) (engine.Result, error) {
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
		if picSmallUrl, ok := s.Find("img").Attr("src"); ok {
			housePic.PicSmallUrl = picSmallUrl
		}
		if picNormalUrl, ok := s.Attr("data-src"); ok {
			housePic.PicNormalUrl = picNormalUrl
		}
		if picLargeUrl, ok := s.Attr("data-pic"); ok {
			housePic.PicLargeUrl = picLargeUrl
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
	return engine.Result{
		Items: []engine.Item{
			house,
		},
	}, nil
}
