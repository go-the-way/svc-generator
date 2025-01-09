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

package opt

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/go-the-way/svc-generator/internal/logger"
)

var (
	DSN                             string
	Database                        string
	Module                          string
	Table                           string
	TableCommentFile                string
	TableCommentFileGenerateExample bool
	SimpleServiceTable              string
	Router                          bool
	RouterRoutePrefix               string
	RouterAppPkg                    string
	RouterOutputDir                 string
	Service                         bool
	ServiceModelPkg                 string
	ServiceOutputDir                string
	OperatorLog                     bool
	OperatorLogPkg                  string
	GofmtAfterGenerated             bool

	simpleServiceMap = map[string]struct{}{}
	tableCommentMap  = map[string]string{}
)

func SimpleService(table string) (yes bool) { _, yes = simpleServiceMap[table]; return }

func TableComment(table string) (comment string) {
	comment = table
	if cmt, ok := tableCommentMap[table]; ok {
		comment = cmt
	}
	return
}

func Init() (err error) {
	if err = check(); err != nil {
		return
	}
	init0()
	return
}

func check() (err error) {
	if DSN == "" {
		return errors.New("the dsn flag is empty")
	}
	if Module == "" {
		return errors.New("the go module flag is empty")
	}
	if Table == "" {
		return errors.New("the table flag is empty")
	}
	return
}

func init0() {
	ss := strings.Split(SimpleServiceTable, ",")
	for _, a := range ss {
		aa := strings.TrimSpace(a)
		if len(aa) > 0 {
			simpleServiceMap[aa] = struct{}{}
		}
	}

	if len(TableCommentFile) > 0 {
		buf, _ := os.ReadFile(TableCommentFile)
		if len(buf) > 0 {
			_ = json.Unmarshal(buf, &tableCommentMap)
			if len(tableCommentMap) > 0 {
				logger.Printf("loaded %d table comment", len(tableCommentMap))
			}
		}
	}
}
