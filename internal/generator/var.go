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

import (
	"path/filepath"
	"strings"

	"github.com/go-the-way/svc-generator/internal/dbload"
	"github.com/go-the-way/svc-generator/internal/generator/tpl"
	"github.com/go-the-way/svc-generator/internal/logger"
	option "github.com/go-the-way/svc-generator/internal/opt"

	. "github.com/go-the-way/svc-generator/internal/model"
)

func Generate() (err error) {
	generators := getGenerators()
	logger.Printf("found %d generator\n", len(generators))
	if len(generators) > 0 {
		for _, gen := range generators {
			gen.Generate()
		}
	}
	return
}

func newGeneratorOption(struct0 *Struct, outputDirPrefix string) generatorOption {
	return generatorOption{
		Struct: struct0,

		Module:            option.Module,
		Router:            option.Router,
		RouterRoutePrefix: option.RouterRoutePrefix,
		RouterAppPkg:      option.RouterAppPkg,
		RouterOutputDir:   option.RouterOutputDir,
		Service:           option.Service,
		ServiceModelPkg:   option.ServiceModelPkg,
		ServiceOutputDir:  option.ServiceOutputDir,
		OperatorLog:       option.OperatorLog,
		OperatorLogPkg:    option.OperatorLogPkg,

		SimpleService:   option.SimpleService(struct0.Table.Name),
		OutputDirPrefix: outputDirPrefix,
	}
}

func getGenerators() (generators []*generator) {
	tableMap := dbload.Loads()
	if len(tableMap) <= 0 {
		return
	}
	for k, v := range tableMap {
		struct0 := transform(v)
		if option.Router {
			opt := newGeneratorOption(struct0, filepath.Join(option.RouterOutputDir, struct0.Pkg))
			generators = append(generators, newGenerator(opt, tpl.RouterTpl(opt), "router.gen.go"))
			logger.Printf("append router 1 generator from `%s`", k)
		}
		if option.Service {
			opt := newGeneratorOption(struct0, filepath.Join(option.ServiceOutputDir, struct0.Pkg))
			generators = append(generators, newGenerator(opt, tpl.ServiceReqTpl(opt), "req.gen.go"))
			// generators = append(generators, newGenerator(opt, tpl.ServiceReqExtraTpl(opt), "req_extra.gen.go"))
			// generators = append(generators, newGenerator(opt, tpl.ServiceRespTpl(opt), "resp.gen.go"))
			// generators = append(generators, newGenerator(opt, tpl.ServiceServiceTpl(opt), "service.gen.go"))
			// generators = append(generators, newGenerator(opt, tpl.ServiceSvcTpl(opt), "svc.gen.go"))
			// generators = append(generators, newGenerator(opt, tpl.ServiceVarTpl(opt), "var.gen.go"))
			logger.Printf("append service 6 generators from `%s`", k)
		}
	}
	return
}

func setFieldValidateTag(field *Field) {
	fieldType := field.Type
	columnType := field.Column.Type
	dataType := field.Column.DataType
	comment := field.Column.Comment
	prefix := "varchar("
	suffix := ")"
	validateTag := ""
	if fieldType == "string" && columnType == "varchar" && strings.HasPrefix(dataType, "varchar(") {
		maxLength := strings.TrimRight(strings.TrimLeft(dataType, prefix), suffix)
		validateTag = `validate:"minlength(1,` + comment + `不能为空) maxlength(` + maxLength + `,` + comment + `长度不能超过` + maxLength + `)"`
	} else {
		validateTag = `validate:"min(1,` + comment + `不能为空)"`
	}
	field.ValidateTag = validateTag
}

func transform(table *Table) *Struct {
	tableName := table.Name
	var (
		fields          []*Field
		pkField         *Field
		createTimeField *Field
		updateTimeField *Field
		timeFields      []*Field
	)
	for _, column := range table.Columns {
		fieldType := map[int]string{0: goSqlNullTypes[mysqlToGoTypes[column.Type]], 1: mysqlToGoTypes[column.Type]}[column.NotNull]
		opName := goTypeOps[fieldType]
		field := &Field{
			Column: column,
			Name:   underlineToUpperCamel(column.Name),
			Type:   fieldType,
			OpName: opName,
			OpVar:  goTypeOpValues[opName],
		}
		setFieldValidateTag(field)
		if column.ColumnKey == "PRI" {
			pkField = field
		} else if column.Name == "create_time" {
			createTimeField = field
		} else if column.Name == "update_time" {
			updateTimeField = field
		} else if strings.HasSuffix(column.Name, "_time") {
			timeFields = append(timeFields, field)
		} else {
			fields = append(fields, field)
		}
	}
	if pkField == nil {
		panic("table of `" + tableName + "` required at least one PRI column")
	}

	return &Struct{
		Table:           table,
		Fields:          fields,
		PkField:         pkField,
		CreateTimeHave:  createTimeField != nil,
		CreateTimeField: createTimeField,
		UpdateTimeHave:  updateTimeField != nil,
		UpdateTimeField: updateTimeField,
		TimeFields:      timeFields,
		Name:            underlineToUpperCamel(tableName),
		Pkg:             tableName,
		Comment:         option.TableComment(tableName),
	}
}

// a_b_c_d => aBCD
func underlineToCamel(str string) string {
	names := strings.Split(str, "_")
	for i, name := range names {
		if i > 0 {
			if len(name) == 1 {
				names[i] = strings.ToUpper(string(name[0]))
			} else if len(name) > 1 {
				names[i] = strings.ToUpper(string(name[0])) + name[1:]
			}
		}
	}
	return strings.Join(names, "")
}

// a_b_c_d => ABCD
func underlineToUpperCamel(str string) string {
	toCamelStr := underlineToCamel(str)
	return strings.ToUpper(string(toCamelStr[0])) + toCamelStr[1:]
}
