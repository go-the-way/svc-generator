package generator

import (
	"github.com/go-the-way/svc-generator/internal/logger"
	option "github.com/go-the-way/svc-generator/internal/opt"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/go-the-way/svc-generator/internal/model"
)

type (
	generator struct {
		generatorOption
		tpl, fileName string
	}
	generatorOption struct {
		*Struct

		Module            string
		Router            bool
		RouterRoutePrefix string
		RouterAppPkg      string
		RouterOutputDir   string
		Service           bool
		ServiceModelPkg   string
		ServiceOutputDir  string
		OperatorLog       bool
		OperatorLogPkg    string

		SimpleService   bool
		OutputDirPrefix string
	}
)

func newGenerator(generatorOption generatorOption, tpl, fileName string) *generator {
	return &generator{generatorOption: generatorOption, tpl: tpl, fileName: fileName}
}

func (g *generator) Generate() {
	fileName := filepath.Join(g.OutputDirPrefix, g.fileName)
	dir := filepath.Dir(fileName)
	_ = os.MkdirAll(dir, 0700)
	_ = os.WriteFile(fileName, []byte(g.tpl), 0700)
	if option.GofmtAfterGenerated {
		_ = exec.Command("gofmt", "-w", fileName).Run()
	}
	logger.Println("generated file", fileName)
}
