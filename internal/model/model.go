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

package model

type (
	Struct struct {
		Table           *Table
		Fields          []*Field
		PkField         *Field
		CreateTimeHave  bool
		CreateTimeField *Field
		UpdateTimeHave  bool
		UpdateTimeField *Field
		TimeFields      []*Field
		Name            string // TableName
		Pkg             string // table_name
		Comment         string // table comment
	}
	Field struct {
		Column      *Column
		Name        string
		Type        string
		OpName      string
		OpVar       string
		ValidateTag string
	}
	Table struct {
		Name    string    `gorm:"column:TABLE_NAME"`
		Columns []*Column `gorm:"-"`
	}
	Column struct {
		Table         string `gorm:"column:TABLE_NAME"`
		Name          string `gorm:"column:COLUMN_NAME"`
		Type          string `gorm:"column:COLUMN_TYPE"`
		DataType      string `gorm:"column:COLUMN_DATA_TYPE"`
		ColumnKey     string `gorm:"column:COLUMN_KEY"`
		Comment       string `gorm:"column:COLUMN_COMMENT"`
		Default       string `gorm:"column:COLUMN_DEFAULT"`
		AutoIncrement int    `gorm:"column:AUTO_INCREMENT" `
		NotNull       int    `gorm:"column:NOT_NULL"`
	}
)
