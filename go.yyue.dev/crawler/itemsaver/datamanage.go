package itemsaver

import (
	"database/sql"
	"fmt"

	// sql driver
	_ "github.com/go-sql-driver/mysql"
	"go.yyue.dev/crawler/proto"
)

// InsertHouseSummary impl Query insert
func InsertHouseSummary(db *sql.DB, houseSummary proto.HouseSummary) error {

	if err := db.Ping(); err != nil {
		return err
	}
	sql := fmt.Sprintf(`INSERT INTO %s (
		house_no, url, title, total_price, unit_price, plot, region, layout, area, face, decoration, floor, house_year, struct_build, image, follow, release_time, tags
	) VALUES (%s, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %d, %q, %d);`,
		"house_summary",
		houseSummary.GetHouseNo(),
		houseSummary.GetUrl(),
		houseSummary.GetTitle(),
		houseSummary.GetTotalPrice(),
		houseSummary.GetUnitPrice(),
		houseSummary.GetPlot(),
		houseSummary.GetRegion(),
		houseSummary.GetLayout(),
		houseSummary.GetArea(),
		houseSummary.GetFace(),
		houseSummary.GetDecoration(),
		houseSummary.GetFloor(),
		houseSummary.GetHouseYear(),
		houseSummary.GetStructBuild(),
		houseSummary.GetImage(),
		houseSummary.GetFollow(),
		houseSummary.GetReleaseTime(),
		houseSummary.GetTags(),
	)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

// InsertHouse insert house
func InsertHouse(db *sql.DB, house proto.House) error {

	if err := db.Ping(); err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	houseInfo := house.GetHouseInfo()
	sql := fmt.Sprintf(`INSERT INTO %s (
		house_no, url, title, sub_title, region, total_price, unit_price, room_info, room_sub, type_info, type_sub, area_info, area_sub
	) VALUES (%s, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q);`,
		"house_info",
		houseInfo.GetHouseNo(),
		houseInfo.GetUrl(),
		houseInfo.GetTitle(),
		houseInfo.GetSubTitle(),
		houseInfo.GetRegion(),
		houseInfo.GetTotalPrice(),
		houseInfo.GetUnitPrice(),
		houseInfo.GetRoomInfo(),
		houseInfo.GetRoomSub(),
		houseInfo.GetTypeInfo(),
		houseInfo.GetTypeSub(),
		houseInfo.GetAreaInfo(),
		houseInfo.GetAreaSub(),
	)
	_, err = tx.Exec(sql)
	if err != nil {
		return err
	}
	houseBaseInfo := house.GetHouseBaseInfo()
	_, err = tx.Exec(
		fmt.Sprintf(`INSERT INTO %s (
			house_no, layout, floor, area_build, struct_house, area_inner, build_type, face, struct_build, decoration, elevator_ratio,
			elevator, property
		) VALUES (%s, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q);`,
			"house_base_info",
			houseBaseInfo.GetHouseNo(),
			houseBaseInfo.GetLayout(),
			houseBaseInfo.GetFloor(),
			houseBaseInfo.GetAreaBuild(),
			houseBaseInfo.GetStructHouse(),
			houseBaseInfo.GetAreaInner(),
			houseBaseInfo.GetBuildType(),
			houseBaseInfo.GetFace(),
			houseBaseInfo.GetStructBuild(),
			houseBaseInfo.GetDecoration(),
			houseBaseInfo.GetElevatorRatio(),
			houseBaseInfo.GetElevator(),
			houseBaseInfo.GetProperty(),
		))
	if err != nil {
		return err
	}
	houseTransactionInfo := house.GetHouseTransactionInfo()
	_, err = tx.Exec(
		fmt.Sprintf(`INSERT INTO %s (
			house_no, listing_time, trading_authority, last_transaction, housing_purposes, house_year, property_rights, mortgage_info, document_photo
		) VALUES (%s, %q, %q, %q, %q, %q, %q, %q, %q);`,
			"house_transaction_info",
			houseTransactionInfo.GetHouseNo(),
			houseTransactionInfo.GetListingTime(),
			houseTransactionInfo.GetTradingAuthority(),
			houseTransactionInfo.GetLastTransaction(),
			houseTransactionInfo.GetHousingPurposes(),
			houseTransactionInfo.GetHouseYear(),
			houseTransactionInfo.GetPropertyRights(),
			houseTransactionInfo.GetMortgageInfo(),
			houseTransactionInfo.GetDocumentPhoto(),
		))
	if err != nil {
		return err
	}
	housePics := house.GetHousePics()
	for _, housePic := range housePics {
		_, err = tx.Exec(
			fmt.Sprintf(`INSERT INTO %s (
				house_no, description, pic_small_url, pic_normal_url, pic_large_url
			) VALUES (%s, %q, %q, %q, %q);`,
				"house_pic",
				housePic.GetHouseNo(),
				housePic.GetDescription(),
				housePic.GetPicSmallUrl(),
				housePic.GetPicNormalUrl(),
				housePic.GetPicLargeUrl(),
			))
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}
