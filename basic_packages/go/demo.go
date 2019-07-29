package main

import (
	"flag"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	osType = os.Getenv("GOOS") // 获取系统类型
)

type SubFiles struct {
	FilesPath []string
}

func (s *SubFiles) isFilte(path, pattern string) bool {
	ok := strings.HasSuffix(path, "_test.go")
	if ok {
		return true
	}
	ok = strings.HasSuffix(path, ".go")
	if !ok {
		return true
	}
	return false
}
func (s *SubFiles) listFunc(path string, f os.FileInfo, err error) error {

	strRet, err := os.Getwd()
	if err != nil {
		return err
	}

	if osType == "windows" {
		strRet += "\\"
	} else if osType == "linux" {
		strRet += "/"
	}

	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}

	if !s.isFilte(path, "") {
		s.FilesPath = append(s.FilesPath, path)
	}

	return nil
}

func (s *SubFiles) ListPackages(dirPath string) error {
	err := filepath.Walk(dirPath, s.listFunc)
	if err != nil {
		return err
	}
	return err
}

func GetPackages(path string) (s SubFiles, err error) {
	err = s.ListPackages(path)
	if err != nil {
		return s, err
	}
	return s, nil
}

func readPackages() []string {

	return []string{}
}

func readImports(path string) ([]string, error) {
	// 这就是上一章的代码.
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", string(b), 0)
	if err != nil {
		return nil, err
	}

	// Print the AST.
	// ast.Print(fset, f)
	// ast.SortImports(fset, f)
	// ast.Print(fset, f.Imports)
	var imports []string
	for _, v := range f.Imports {
		imports = append(imports, v.Path.Value)
	}
	return imports, nil
}

func main() {
	var path string
	var detail, standard bool
	flag.StringVar(&path, "path", "./", "directory path")
	flag.BoolVar(&detail, "detail", false, "weather print bool value")
	flag.BoolVar(&standard, "standard", false, "weather just print standard library")
	flag.Parse()
	absPath, err := filepath.Abs(path)
	if err != nil {
		logrus.Infoln(1)
		logrus.Errorln(err)
		return
	}
	logrus.Debugln("analyze ", absPath)

	f, err := os.Stat(path)
	if err != nil || !f.IsDir() {
		logrus.Warnf("%s not exist or not a directory!", path)
		return
	}

	s, _ := GetPackages(absPath)

	summary := make(map[string]int, 1)
	for _, v := range s.FilesPath {
		// logrus.Infoln(v)
		res, err := readImports(v)
		if err != nil {
			logrus.Debug(err)
			// logrus.Errorln(err)
		} else {
			if detail {
				logrus.Infoln(v, res)
			}
			for _, i := range res {
				summary[i]++
			}
		}
	}

	for k, v := range summary {
		if standard {
			if !strings.ContainsRune(k, '.') {
				logrus.Infoln(k, v)
			}
		} else {
			logrus.Infoln(k, v)
		}

	}
}
