package creator

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-helios/eos/pkg/util"
	"github.com/go-helios/eos/tpls"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

type Creator struct {
	Project    string
	Type       string
	FilePath   string
	FileName   string
	TitleLower string
	FirstChar  string
	IsFull     bool
}

func (c *Creator) generate() {
	filePath := c.FilePath
	if len(filePath) == 0 {
		if c.Type == "request" || c.Type == "response" {
			filePath = fmt.Sprintf("internal/dto/%s/", c.Type)
		} else {
			filePath = fmt.Sprintf("internal/%s/", c.Type)
		}
	}
	fName := strings.ToLower(strcase.ToSnake(c.FileName)) + ".go"
	f := util.CreateFile(filePath, fName)
	if f == nil {
		fmt.Printf("警告：文件 %s%s %s!\n", filePath, fName, "已存在")
		return
	}
	defer f.Close()
	tpl, err := template.ParseFS(tpls.CreateTemplateFS, fmt.Sprintf("creator/%s.tpl", c.Type))
	if err != nil {
		log.Fatalf("创建 %s 失败: %s\n", c.Type, err.Error())
	}
	err = tpl.Execute(f, c)
	if err != nil {
		log.Fatalf("创建 %s 失败: %s.", c.Type, err.Error())
	}
	log.Printf("创建成功 %s:%s", c.Type, filePath+fName)
}

var CmdCreator = &cobra.Command{
	Use:     "create [type] [name]",
	Short:   "新建一个 handler/service/repository/model/request/response",
	Example: "eos new handler [name]",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var CmdCreatorHandler = &cobra.Command{
	Use:     "handler",
	Short:   "新建一个 handler",
	Example: "eos new handler [name]",
	Args:    cobra.ExactArgs(1),
	Run:     run,
}

var CmdCreatorService = &cobra.Command{
	Use:     "service",
	Short:   "新建一个 service",
	Example: "eos new service [name]",
	Args:    cobra.ExactArgs(1),
	Run:     run,
}

var CmdCreatorRepository = &cobra.Command{
	Use:     "repository",
	Short:   "新建一个 repository",
	Example: "eos new repository [name]",
	Args:    cobra.ExactArgs(1),
	Run:     run,
}

var CmdCreatorModel = &cobra.Command{
	Use:     "model",
	Short:   "新建一个 model",
	Example: "eos new model [name]",
	Args:    cobra.ExactArgs(1),
	Run:     run,
}

var CmdCreatorRequest = &cobra.Command{
	Use:     "request",
	Short:   "新建一个 request",
	Example: "eos new request [name]",
	Args:    cobra.ExactArgs(1),
	Run:     run,
}

var CmdCreatorResponse = &cobra.Command{
	Use:     "response",
	Short:   "新建一个 response",
	Example: "eos new response [name]",
	Args:    cobra.ExactArgs(1),
	Run:     run,
}

var CmdCreatorAll = &cobra.Command{
	Use:     "all",
	Short:   "新建一个 handler & service & repository & model",
	Example: "eos new all [name]",
	Args:    cobra.ExactArgs(1),
	Run:     run,
}

func run(cmd *cobra.Command, args []string) {
	c := &Creator{}
	c.Project = util.ProjectName(".")
	c.Type = cmd.Use
	c.FilePath, c.FileName = filepath.Split(args[0])
	c.FileName = strings.ReplaceAll(strcase.ToCamel(c.FileName), ".go", "")
	c.TitleLower = strings.ToLower(strcase.ToLowerCamel(c.FileName))
	c.FirstChar = string(c.TitleLower[0])

	switch c.Type {
	case "handler", "service", "repository", "model", "request", "response":
		c.generate()
	case "all":
		c.Type = "request"
		c.generate()
		c.Type = "response"
		c.generate()
		c.Type = "handler"
		c.generate()
		c.Type = "service"
		c.generate()
		c.Type = "repository"
		c.generate()
		c.Type = "model"
		c.generate()
	default:
		fmt.Printf("未知的新建类型: %s", c.Type)
	}
}
