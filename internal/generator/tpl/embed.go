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

package tpl

import (
	"strings"
	"text/template"

	_ "embed"
)

var (
	//go:embed router.tpl
	routerTpl string
	//go:embed service_req.tpl
	serviceReqTpl string
	//go:embed service_req_extra.tpl
	serviceReqExtraTpl string
	//go:embed service_resp.tpl
	serviceRespTpl string
	//go:embed service_service.tpl
	serviceServiceTpl string
	//go:embed service_svc.tpl
	serviceSvcTpl string
	//go:embed service_var.tpl
	serviceVarTpl string
)

func RouterTpl(data any) string          { return execute(routerTpl, data) }
func ServiceReqTpl(data any) string      { return execute(serviceReqTpl, data) }
func ServiceReqExtraTpl(data any) string { return execute(serviceReqExtraTpl, data) }
func ServiceRespTpl(data any) string     { return execute(serviceRespTpl, data) }
func ServiceServiceTpl(data any) string  { return execute(serviceServiceTpl, data) }
func ServiceSvcTpl(data any) string      { return execute(serviceSvcTpl, data) }
func ServiceVarTpl(data any) string      { return execute(serviceVarTpl, data) }

func execute(tpl string, data any) string {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		panic(err)
	}
	var buffer strings.Builder
	err = t.Execute(&buffer, data)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}
