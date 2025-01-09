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

package generator

var mysqlToGoTypes = map[string]string{
	"bit":             "bool",
	"tinyint":         "byte",
	"smallint":        "int8",
	"mediumint":       "int16",
	"int":             "int",
	"bigint":          "int",  // int64?
	"bigint unsigned": "uint", // uint64?
	"float":           "float32",
	"double":          "float64",
	"decimal":         "float64",
	"date":            "time.Time",
	"time":            "string",
	"year":            "int8",
	"datetime":        "string",
	"timestamp":       "int64",
	"char":            "string",
	"varchar":         "string",
	"tinytext":        "string",
	"mediumtext":      "string",
	"text":            "string",
	"longtext":        "string",
	"tinyblob":        "byte[]",
	"mediumblob":      "byte[]",
	"blob":            "byte[]",
	"longblob":        "byte[]",
}

var goSqlNullTypes = map[string]string{
	"bool":    "sql.NullBool",
	"byte":    "sql.NullByte",
	"int8":    "sql.NullByte",
	"int16":   "sql.NullInt16",
	"int":     "sql.NullInt32",
	"int64":   "sql.NullInt64",
	"float32": "sql.NullFloat64",
	"float64": "sql.NullFloat64",
	"string":  "sql.NullString",
}

var goTypeOps = map[string]string{
	"byte":    ">",
	"int8":    ">",
	"int16":   ">",
	"int32":   ">",
	"int":     ">",
	"int64":   ">",
	"float32": ">",
	"float64": ">",
	"string":  "!=",
}

var goTypeOpValues = map[string]string{
	">":  "0",
	"!=": "\"\"",
}
