// Copyright 2020 - 2024 Alexey Yanchenko <mail@yanchenko.me>
//
// This file is part of the Gufo library.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package entrypoint

import (
	. "dictionary/model"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	"github.com/spf13/viper"
)

func CheckDBStructure() {
	//Check DB and table config
	db, err := ConnectDBv2()
	if err != nil {
		SetErrorLog("dbstructure.go:81: " + err.Error())
		//return "error with db"
	}

	dbtype := viper.GetString("database.type")

	if !db.Conn.Migrator().HasTable(&DictionaryCategories{}) {
		if dbtype == "mysql" {
			db.Conn.Set("gorm:table_options", "ENGINE=InnoDB;").Migrator().CreateTable(&DictionaryCategories{})
		} else {
			db.Conn.Migrator().CreateTable(&DictionaryCategories{})
		}
	}

	if !db.Conn.Migrator().HasTable(&DictionaryValue{}) {
		if dbtype == "mysql" {
			db.Conn.Set("gorm:table_options", "ENGINE=InnoDB;").Migrator().CreateTable(&DictionaryValue{})
		} else {
			db.Conn.Migrator().CreateTable(&DictionaryValue{})
		}
	}

	if !db.Conn.Migrator().HasTable(&DictionaryValueLoc{}) {
		if dbtype == "mysql" {
			db.Conn.Set("gorm:table_options", "ENGINE=InnoDB;").Migrator().CreateTable(&DictionaryValueLoc{})
		} else {
			db.Conn.Migrator().CreateTable(&DictionaryValueLoc{})
		}
	}

	if !db.Conn.Migrator().HasTable(&Countries{}) {
		if dbtype == "mysql" {
			db.Conn.Set("gorm:table_options", "ENGINE=InnoDB;").Migrator().CreateTable(&Countries{})
		} else {
			db.Conn.Migrator().CreateTable(&Countries{})
		}
	}

	if !db.Conn.Migrator().HasTable(&Cities{}) {
		if dbtype == "mysql" {
			db.Conn.Set("gorm:table_options", "ENGINE=InnoDB;").Migrator().CreateTable(&Cities{})
		} else {
			db.Conn.Migrator().CreateTable(&Cities{})
		}
	}

	if !db.Conn.Migrator().HasTable(&States{}) {
		if dbtype == "mysql" {
			db.Conn.Set("gorm:table_options", "ENGINE=InnoDB;").Migrator().CreateTable(&States{})
		} else {
			db.Conn.Migrator().CreateTable(&States{})
		}
	}

}
