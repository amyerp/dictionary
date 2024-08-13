//////////////////////////////////////////////////////////////////////////////////
// Copyright 2021 Alexey Yanchenko <mail@yanchenko.me>                          //
//                                                                              //
// This file is part of the ERP library.                                        //
//                                                                              //
//  Unauthorized copying of this file, via any media is strictly prohibited     //
//  Proprietary and confidential                                                //
//////////////////////////////////////////////////////////////////////////////////

package model

import (
	"gorm.io/gorm"
)

/*
Categories:
VATs
Languages
Cryptocurrencies
Units
Measures
Weights
Duration
BysinessTypes
Social Networks
Fixed Asset Types
Business Attributes
Person Attributes
Enterprenuer Attributes
Fixed Asset Attributes


Special:
Countries
PhoneCodes
Currency
Cities
States

*/

type DictionaryCategories struct {
	gorm.Model
	UUID       string `gorm:"column:uuid;type:varchar(254);DEFAULT '';" json:"uuid"`
	Category   string `gorm:"column:category;type:varchar(60);DEFAULT '';" json:"category"`
	FilteredBy string `gorm:"column:filtered_by;type:varchar(254);DEFAULT '';" json:"filtered_by"` //any other Category_id
	IsCustom   bool   `gorm:"column:custom;type:bool;DEFAULT true;NOT NULL;" json:"custom"`
}

type DictionaryValue struct {
	gorm.Model
	UUID        string  `gorm:"column:uuid;type:varchar(254);DEFAULT '';" json:"uuid"`
	CategoryID  string  `gorm:"column:catrgoryid;type:varchar(254);DEFAULT '';" json:"categoryid"`
	Value       float64 `gorm:"column:value;type:float;DEFAULT '';" json:"value"` //for example id or vat amount
	Name        string  `gorm:"column:name;type:varchar(254);DEFAULT '';" json:"name"`
	ShortName   string  `gorm:"column:short_name;type:varchar(254);DEFAULT '';" json:"short_name"`
	FilterValue string  `gorm:"column:filter_value;type:varchar(254);DEFAULT '';" json:"filter_value"`
	IsRequired  bool    `gorm:"column:required;type:bool;DEFAULT false;NOT NULL;" json:"required"`
}

type DictionaryValueLoc struct {
	gorm.Model
	ValueID   string `gorm:"column:valueid;type:varchar(254);DEFAULT '';" json:"valueid"`
	Language  string `gorm:"column:language;type:varchar(254);DEFAULT '';" json:"language"`
	Name      string `gorm:"column:value;type:varchar(254);DEFAULT '';" json:"name"`
	ShortName string `gorm:"column:short_value;type:varchar(254);DEFAULT '';" json:"short_name"`
}

type Countries struct {
	gorm.Model
	Name           string `gorm:"column:name;" json:"name"`
	ISO3           string `gorm:"column:iso3;" json:"iso3"`
	NumericCode    string `gorm:"column:numeric_code;" json:"numeric_code"`
	ISO2           string `gorm:"column:iso2;" json:"iso2"`
	Phonecode      string `gorm:"column:phonecode;" json:"phonecode"`
	Capital        string `gorm:"column:capital;" json:"capital"`
	Currency       string `gorm:"column:currency;" json:"currency"`
	CurrencyName   string `gorm:"column:currency_name;" json:"currency_name"`
	CurrencySymbol string `gorm:"column:currency_symbol;" json:"currency_symbol"`
	Tld            string `gorm:"column:tld;" json:"tld"`
	Native         string `gorm:"column:native;" json:"native"`
	Region         string `gorm:"column:region;" json:"region"`
	Subregion      string `gorm:"column:subregion;" json:"subregion"`
	Zone           string `gorm:"column:zone;" json:"zone"`
	Timezones      string `gorm:"column:timezones;" json:"timezones"`
	Translations   string `gorm:"column:translations;" json:"translations"`
	Latitude       string `gorm:"column:latitude;" json:"latitude"`
	Longitude      string `gorm:"column:longitude;" json:"longitude"`
	Emoji          string `gorm:"column:emoji;" json:"emoji"`
	EmojiU         string `gorm:"column:emojiU;" json:"emojiU"`
	Flag           int    `gorm:"column:flag;" json:"flag"`
	WikiDataId     string `gorm:"column:wikiDataId;" json:"wikiDataId"`
}

type Cities struct {
	gorm.Model
	Name         string `gorm:"column:name;" json:"name"`
	StateID      string `gorm:"column:state_id;" json:"state_id"`
	StateCode    string `gorm:"column:state_code;" json:"state_code"`
	CountryID    string `gorm:"column:country_id;" json:"country_id"`
	CountryCode  string `gorm:"column:country_code;" json:"country_code"`
	Translations string `gorm:"column:translations;" json:"translations"`
	Latitude     string `gorm:"column:latitude;" json:"latitude"`
	Longitude    string `gorm:"column:longitude;" json:"longitude"`
	Flag         int    `gorm:"column:flag;" json:"flag"`
	WikiDataId   string `gorm:"column:wikiDataId;" json:"wikiDataId"`
}

type States struct {
	gorm.Model
	Name         string `gorm:"column:name;" json:"name"`
	CountryID    string `gorm:"column:country_id;" json:"country_id"`
	CountryCode  string `gorm:"column:country_code;" json:"country_code"`
	FipsCode     string `gorm:"column:fips_code;" json:"fips_code"`
	ISO2         string `gorm:"column:iso2;" json:"iso2"`
	Type         string `gorm:"column:type;" json:"type"`
	Translations string `gorm:"column:translations;" json:"translations"`
	Latitude     string `gorm:"column:latitude;" json:"latitude"`
	Longitude    string `gorm:"column:longitude;" json:"longitude"`
	Flag         int    `gorm:"column:flag;" json:"flag"`
	WikiDataId   string `gorm:"column:wikiDataId;" json:"wikiDataId"`
}
