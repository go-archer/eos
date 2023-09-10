package project

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-helios/eos/pkg/util"
	"github.com/spf13/cobra"
)

type Project struct {
	Name string `survey:"name"`
}

func (p *Project) clone() (bool, error) {
	stat, _ := os.Stat(p.Name)
	if stat != nil {
		overwrite := false
		prop := &survey.Confirm{
			Message: fmt.Sprintf("文件夹 %s 已经存在,是否覆盖\n", p.Name),
			Help:    "删除旧项目，同时创建新项目",
		}
		err := survey.AskOne(prop, &overwrite)
		if err != nil {
			return false, err
		}
		if !overwrite {
			return false, nil
		}
		err = os.RemoveAll(p.Name)
		if err != nil {
			fmt.Println("删除旧项目失败", err)
			return false, err
		}
	}
	const repo = "https://github.com/go-helios/eos-layout.git"
	fmt.Printf("git clone %s\n", repo)
	cmd := exec.Command("git", "clone", repo, p.Name)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("git克隆 %s 失败: %s\n", repo, err)
		return false, err
	}
	return true, nil
}

func (p *Project) replace() error {
	pkgName := util.ProjectName(p.Name)
	err := p.replaceFiles(pkgName)
	if err != nil {
		return err
	}
	cmd := exec.Command("go", "mod", "edit", "-module", p.Name)
	cmd.Dir = p.Name
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("go.mod 编辑错误:", err)
		return err
	}
	return nil
}

func (p *Project) replaceFiles(name string) error {
	err := filepath.Walk(p.Name, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		newData := bytes.ReplaceAll(data, []byte(name), []byte(p.Name))
		if err := os.WriteFile(path, newData, 0644); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("遍历文件错误:", err)
		return err
	}
	return nil
}

func (p *Project) tidy() error {
	fmt.Printf("go mod tidy\n")
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = p.Name
	if err := cmd.Run(); err != nil {
		fmt.Println("go mod tidy 错误:", err)
		return err
	}
	return nil
}

func (p *Project) removeGit() {
	os.RemoveAll(filepath.Join(p.Name, "/.git"))
}

func (p *Project) installWire() {
	const wire = "github.com/google/wire/cmd/wire@latest"
	fmt.Printf("go install %s\n", wire)
	cmd := exec.Command("go", "install", wire)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("go install %s 错误\n", err)
	}
}

var CmdNew = &cobra.Command{
	Use:     "new",
	Example: "eos new demo",
	Short:   "创建一个新项目",
	Long:    "使用eos创建一个新的项目",
	Run:     run,
}

func run(cmd *cobra.Command, args []string) {
	p := &Project{}
	if len(args) == 0 {
		err := survey.AskOne(&survey.Input{
			Message: "您的项目名称是什么?",
			Help:    "项目名称",
			Suggest: nil,
		}, &p.Name, survey.WithValidator(survey.Required))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		p.Name = args[0]
	}

	yes, err := p.clone()
	if err != nil || !yes {
		return
	}

	err = p.replace()
	if err != nil {
		return
	}

	err = p.tidy()
	if err != nil {
		return
	}

	p.removeGit()
	p.installWire()
	fmt.Println("项目创建成功!")
}
