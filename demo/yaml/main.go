package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

type repo struct {
	Charts struct {
		Name string `yaml:"name"`
		Repo string `yaml:"repo"`
		Tags string `yaml:"tags"`
	} `yaml:"Charts"`
}

func init() {
	logrus.Infoln("init...")
}

func tviper() {

	viper.SetConfigName("config")
	viper.AddConfigPath(currentdir())
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatal(err)
	}

	C := repo{}

	err = viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	logrus.Infoln(C)

	// Change value in map and marshal back into yaml
	C.Charts.Tags = "realtag"

	logrus.Infoln(C)

	d, err := yaml.Marshal(&C)
	if err != nil {
		logrus.Errorf("error: %v", err)
	}

	// write to file
	f, err := os.Create("./data2.yaml")
	if err != nil {
		logrus.Errorln(err)
	}

	// err = ioutil.WriteFile("changed.yaml", d, 0777)
	_, err = f.Write(d)
	logrus.Infoln(string(d))
	if err != nil {
		logrus.Errorf("failed to write file, error: %v", err)
	}

	f.Close()

}

func currentdir() (cwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return cwd
}

type File struct {
	TypeVersion string `yaml:"_type-version"`
	Dependency  []Dependency
}

type Dependency struct {
	Name     string
	Type     string
	CWD      string
	Install  []Install
	Requires []Requires
}

type Install map[string]string

func (i Install) name() string {
	return i["name"]
}

func (i Install) group() string {
	return i["group"]
}

type Requires struct {
	Name string
	Type string
}

var data = `
_type-version: "1.0.0"
dependency:
  - name: ui
    type: runner
    cwd: /ui
    install:
       - name: api
         group: test
         grou2p: test22
    requires:
      - name: db
      - type: mongo
      - name: rst
      - name: test
      - name: test2
`

func mapyaml() {
	f := File{}

	err := yaml.Unmarshal([]byte(data), &f)
	if err != nil {
		logrus.Errorf("error: %v", err)
	}
	logrus.Infof("--- t:\n%v\n\n", f)

	d, err := yaml.Marshal(&f)
	if err != nil {
		logrus.Errorf("error: %v", err)
	}
	logrus.Infof("--- t dump:\n%s\n\n", string(d))
}

func main() {
	mapyaml()
}
