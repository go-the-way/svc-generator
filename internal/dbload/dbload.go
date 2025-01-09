// Copyright 2025 svc Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dbload

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"

	"github.com/go-the-way/svc-generator/internal/opt"

	. "github.com/go-the-way/svc-generator/internal/model"
)

func Loads() map[string]*Table {
	tables, err := loadTables()
	if err != nil {
		panic(err)
	}
	columns, err := loadColumns()
	if err != nil {
		panic(err)
	}
	tableMap := transformTables(tables)
	setTableColumns(tableMap, transformColumns(columns))
	return tableMap
}

func getDB() *gorm.DB {
	gDB, err := gorm.Open(mysql.Open(opt.DSN), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic(err)
	}
	return gDB
}

func loadTables() (tables []*Table, err error) {
	table := strings.ReplaceAll(opt.Table, ",", "','")
	sqlStr := "SELECT t.`TABLE_NAME`, IFNULL(t.TABLE_COMMENT, '') AS TABLE_COMMENT FROM information_schema.`TABLES` AS t WHERE t.`TABLE_SCHEMA` = ? AND t.`TABLE_NAME` IN('" + table + "')"
	err = getDB().Raw(sqlStr, opt.Database).Find(&tables).Error
	return
}

func loadColumns() (columns []*Column, err error) {
	table := strings.ReplaceAll(opt.Table, ",", "','")
	sqlStr := `SELECT
		  t.TABLE_NAME,
		  t.COLUMN_KEY,
		  t.COLUMN_NAME,
		  t.IS_NULLABLE = 'NO' AS NOT_NULL,
		  IFNULL(t.COLUMN_DEFAULT, '__NULL__') AS COLUMN_DEFAULT,
		  IF(
			LOCATE('(', t.COLUMN_TYPE) = 0,
			t.COLUMN_TYPE,
			SUBSTRING(
			  t.COLUMN_TYPE,
			  1,
			  LOCATE('(', t.COLUMN_TYPE) - 1
			)
		  ) AS COLUMN_TYPE,
			t.COLUMN_TYPE as COLUMN_DATA_TYPE,
		  IFNULL(t.COLUMN_COMMENT, '') AS COLUMN_COMMENT,
		  t.EXTRA = 'auto_increment' as AUTO_INCREMENT
		FROM
		  information_schema.COLUMNS AS t
		WHERE t.TABLE_SCHEMA = ? AND t.TABLE_NAME IN ('` + table + `')
		ORDER BY t.TABLE_NAME ASC,
		  t.ORDINAL_POSITION ASC`
	err = getDB().Raw(sqlStr, opt.Database).Find(&columns).Error
	return
}

func transformTables(tables []*Table) map[string]*Table {
	tableMap := make(map[string]*Table, len(tables))
	for _, t := range tables {
		tableMap[t.Name] = t
	}
	return tableMap
}

func transformColumns(columns []*Column) map[string]*[]*Column {
	columnMap := make(map[string]*[]*Column)
	for _, c := range columns {
		cs, have := columnMap[c.Table]
		if have {
			*cs = append(*cs, c)
		} else {
			csp := make([]*Column, 1)
			csp[0] = c
			columnMap[c.Table] = &csp
		}
	}
	return columnMap
}

func setTableColumns(tableMap map[string]*Table, columnMap map[string]*[]*Column) {
	for k, v := range tableMap {
		if cc, have := columnMap[k]; have {
			v.Columns = append(v.Columns, *cc...)
		}
	}
}
