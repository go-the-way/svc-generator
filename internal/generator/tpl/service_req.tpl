package {{.Struct.Pkg}}

import "github.com/go-the-way/svc"

type (
	GetPageReq struct {
		svc.PageReq
{{.Struct.PkField.Name}} {{.Struct.PkField.Type}} `form:"{{.Struct.PkField.Column.Name}}"` // {{.Struct.PkField.Column.Comment}}
{{range .Struct.Fields}}
  {{.Name}} {{.Type}} `form:"{{.Column.Name}}"` // {{.Column.Comment}}
{{end}}
{{if .Struct.CreateTimeHave}}
  {{.Struct.CreateTimeField.Name}}1 {{.Struct.CreateTimeField.Type}} `form:"{{.Struct.CreateTimeField.Column.Name}}1"` // {{.Struct.CreateTimeField.Column.Comment}}
  {{.Struct.CreateTimeField.Name}}2 {{.Struct.CreateTimeField.Type}} `form:"{{.Struct.CreateTimeField.Column.Name}}2"` // {{.Struct.CreateTimeField.Column.Comment}}
{{end}}
{{if .Struct.UpdateTimeHave}}
  {{.Struct.UpdateTimeField.Name}}1 {{.Struct.UpdateTimeField.Type}} `form:"{{.Struct.UpdateTimeField.Column.Name}}1"` // {{.Struct.UpdateTimeField.Column.Comment}}
  {{.Struct.UpdateTimeField.Name}}2 {{.Struct.UpdateTimeField.Type}} `form:"{{.Struct.UpdateTimeField.Column.Name}}2"` // {{.Struct.UpdateTimeField.Column.Comment}}
{{end}}
{{range .Struct.TimeFields}}
  {{.Name}}1 {{.Type}} `form:"{{.Column.Name}}1"`         // {{.Column.Comment}}
  {{.Name}}2 {{.Type}} `form:"{{.Column.Name}}2"`         // {{.Column.Comment}}
{{end}}
	}

	GetReq struct {
	  {{.Struct.PkField.Name}} {{.Struct.PkField.Type}} `uri:"{{.Struct.PkField.Column.Name}}" form:"{{.Struct.PkField.Column.Name}}" json:"{{.Struct.PkField.Column.Name}}" {{.Struct.PkField.ValidateTag}}` // {{.Struct.PkField.Column.Comment}}
	}

	AddReq struct {
{{range .Struct.Fields}}
		{{.Name}} {{.Type}} `json:"{{.Column.Name}}" {{.ValidateTag}}`  // {{.Column.Comment}}
{{end}}

{{if .OperatorLog}}Callback func(req AddReq){{end}}
	}

	UpdateReq struct {
		GetReq
		AddReq

{{if .OperatorLog}}Callback func(req UpdateReq){{end}}
	}

	DeleteReq struct {
    GetReq

{{if .OperatorLog}}Callback func(req DeleteReq){{end}}
	}
)
