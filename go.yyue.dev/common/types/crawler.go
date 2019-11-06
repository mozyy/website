package types

type HouseInfo struct {
	ID       string `db:"id"`        // id
	URL      string `db:"url"`       // url
	Title    string `db:"title"`     // 标题
	SubTitle string `db:"sub_title"` // 副标题
	Region   string `db:"region"`    // 小区

	Layout        string `db:"layout"`         // 房屋户型
	Floor         string `db:"floor"`          // 所在楼层
	AreaBuild     string `db:"area_build"`     // 建筑面积
	StructHouse   string `db:"struct_house"`   // 户型结构
	AreaInner     string `db:"area_inner"`     // 套内面积
	BuildType     string `db:"build_type"`     // 建筑类型
	Face          string `db:"face"`           // 房屋朝向
	StructBuild   string `db:"struct_build"`   // 建筑结构
	Decoration    string `db:"decoration"`     // 装修情况
	ElevatorRatio string `db:"elevator_ratio"` // 梯户比例
	Elevator      string `db:"elevator"`       // 配备电梯
	Property      string `db:"property"`       // 产权年限

	ListingTime      string `db:"listing_time"`      // 挂牌时间
	TradingAuthority string `db:"trading_authority"` // 交易权属
	LastTransaction  string `db:"last_transaction"`  // 上次交易
	HousingPurposes  string `db:"housing_purposes"`  // 房屋用途
	HouseYear        string `db:"house_year"`        // 房屋年限
	PropertyRights   string `db:"property_rights"`   // 产权所属
	MortgageInfo     string `db:"mortgage_info"`     // 抵押信息
	DocumentPhoto    string `db:"document_photo"`    // 房本备件
}
