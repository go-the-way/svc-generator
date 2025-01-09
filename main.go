// Copyright 2025 svc Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/go-the-way/svc-generator/internal/generator"
	"github.com/go-the-way/svc-generator/internal/logger"
	"github.com/go-the-way/svc-generator/internal/opt"
	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Println(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&opt.DSN, "dsn", "D", "", "The mysql dsn")
	rootCmd.PersistentFlags().StringVarP(&opt.Database, "database", "d", "", "The database name")
	rootCmd.PersistentFlags().StringVarP(&opt.Module, "module", "m", "awesome", "The go module name")
	rootCmd.PersistentFlags().StringVarP(&opt.Table, "table", "t", "", "The table for generating")
	rootCmd.PersistentFlags().StringVar(&opt.TableCommentFile, "table-comment-file", "table-comment.json", "The table comment mapping file")
	rootCmd.PersistentFlags().BoolVar(&opt.TableCommentFileGenerateExample, "table-comment-file-generate-example", false, "The table comment mapping file generated then exit 0")
	rootCmd.PersistentFlags().StringVarP(&opt.SimpleServiceTable, "simple-service-table", "T", "", "The simple service table for generating, only provides Get and Update")
	rootCmd.PersistentFlags().BoolVarP(&opt.Router, "router", "r", true, "Generate gin router pkg?")
	rootCmd.PersistentFlags().StringVar(&opt.RouterRoutePrefix, "router-route-prefix", "/api", "The gin router route prefix")
	rootCmd.PersistentFlags().StringVar(&opt.RouterAppPkg, "router-app-pkg", "awesome/app", "The gin router app pkg")
	rootCmd.PersistentFlags().StringVarP(&opt.RouterOutputDir, "router-output-dir", "R", "router", "The gin router output directory")
	rootCmd.PersistentFlags().BoolVarP(&opt.Service, "service", "s", true, "Generate standard service pkg?")
	rootCmd.PersistentFlags().StringVar(&opt.ServiceModelPkg, "service-model-pkg", "awesome/models", "The standard service model pkg")
	rootCmd.PersistentFlags().StringVarP(&opt.ServiceOutputDir, "service-output-dir", "S", "services", "The standard service output directory")
	rootCmd.PersistentFlags().BoolVarP(&opt.OperatorLog, "operator-log", "o", false, "Generate standard operator log?")
	rootCmd.PersistentFlags().StringVarP(&opt.OperatorLogPkg, "operator-log-pkg", "O", "operatorlog", "The standard operator log pkg")
	rootCmd.PersistentFlags().BoolVar(&opt.GofmtAfterGenerated, "gofmt-after-generated", true, "Go fmt after generated?")
}

var tableCommentExampleJsonFile = `{
  "table_1": "表1",
  "table_2": "表2"
}`

var rootCmd = &cobra.Command{
	Use:   "svc-generator",
	Short: "A standard svc style code generator written in Go.",
	Long:  "A standard svc style code generator written in Go.",
	Run: func(*cobra.Command, []string) {
		if err := opt.Init(); err != nil {
			logger.Println(err.Error())
			return
		}
		if opt.TableCommentFileGenerateExample {
			_ = os.WriteFile(opt.TableCommentFile, []byte(tableCommentExampleJsonFile), 0700)
			return
		}
		generator.Generate()
	},
}
