package {{.Struct.Pkg}}

import (
{{if .OperatorLog}}	"fmt"

{{end}}	"github.com/gin-gonic/gin"
	"github.com/go-the-way/svc"

	app "{{.RouterAppPkg}}"{{if .OperatorLog}}

	{{.OperatorLogPkg}} "{{.Module}}/services/{{.OperatorLogPkg}}"{{end}}

	. "{{.Module}}/services/{{.Struct.Pkg}}"
)
{{if .SimpleService}}
func init() {
	g := app.GetAppWithGroup("{{.RouterRoutePrefix}}/{{.Struct.Pkg}}")
	g.GET("", get)
	g.PUT("", put)
}

func get(ctx *gin.Context) {
	svc.QueryResp(ctx, Get}
}

func put(ctx *gin.Context) {
	svc.BodyReq(ctx, UpdateReq{{"{"}}{{if .OperatorLog}}
		Callback: func(req UpdateReq) {
			_ = {{.OperatorLogPkg}}.AddFromCtx(ctx, fmt.Sprintf("修改{{.Struct.Comment}}：{{.Struct.PkField.Column.Name}}[%v]{{range .Struct.Fields}} {{.Column.Name}}[%v]{{end}}", req.{{.Struct.PkField.Name}}{{range .Struct.Fields}}, req.{{.Name}}{{end}}))
		},
	{{end}}}, Update)
}
{{else}}
func init() {
	g := app.GetAppWithGroup("{{.RouterRoutePrefix}}/{{.Struct.Pkg}}")
	g.GET("", get)
	g.POST("", post)
	g.PUT("", put)
	g.DELETE("", delete0)
}

func get(ctx *gin.Context) {
	svc.QueryReqResp(ctx, GetPageReq{}, GetPage)
}

func post(ctx *gin.Context) {
	svc.BodyReq(ctx, AddReq{{"{"}}{{if .OperatorLog}}
		Callback: func(req AddReq) {
			_ = {{.OperatorLogPkg}}.AddFromCtx(ctx, fmt.Sprintf("添加{{.Struct.Comment}}：{{.Struct.PkField.Column.Name}}[%v]{{range .Struct.Fields}} {{.Column.Name}}[%v]{{end}}", req.{{.Struct.PkField.Name}}{{range .Struct.Fields}}, req.{{.Name}}{{end}}))
		},
	{{end}}}, Add)
}

func put(ctx *gin.Context) {
	svc.BodyReq(ctx, UpdateReq{{"{"}}{{if .OperatorLog}}
		Callback: func(req UpdateReq) {
			_ = {{.OperatorLogPkg}}.AddFromCtx(ctx, fmt.Sprintf("修改{{.Struct.Comment}}：{{.Struct.PkField.Column.Name}}[%v]{{range .Struct.Fields}} {{.Column.Name}}[%v]{{end}}", req.{{.Struct.PkField.Name}}{{range .Struct.Fields}}, req.{{.Name}}{{end}}))
		},
	{{end}}}, Update)
}

func delete0(ctx *gin.Context) {
	svc.BodyReq(ctx, DeleteReq{{"{"}}{{if .OperatorLog}}
		Callback: func(req DeleteReq) {
			_ = {{.OperatorLogPkg}}.AddFromCtx(ctx, fmt.Sprintf("删除{{.Struct.Comment}}：{{.Struct.PkField.Column.Name}}[%v]", req.{{.Struct.PkField.Name}}))
		},
	{{end}}}, Delete)
}
{{end}}