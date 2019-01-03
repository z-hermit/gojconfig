package configs

import (
	"encoding/json"
	"github.com/kataras/iris/core/errors"
	"os"
)

type Config interface {
	Section(key string) (s Config, err error)
	Value(key string) (v interface{}, err error)
	UnsafeValue(key string) (v interface{})
	Init(filepath string) error
}

func NewConfig(filePath string) (Config, error) {
	conf := myConfig{}
	err := conf.Init(filePath)
	return &conf, err
}

type myConfig struct {
	config map[string]interface{}
	isInit bool
}

func (f *myConfig) Section(key string) (s Config, err error) {
	if !f.isInit {
		err = errors.New("config not init")
		return
	}
	sectionData := f.config[key]
	section, ok := sectionData.(map[string]interface{})
	if !ok {
		err = errors.New("the value is not a section")
		return
	}
	s = &myConfig{section, true}
	return
}

func (f *myConfig) Value(key string) (v interface{}, err error) {
	if !f.isInit {
		err = errors.New("not init")
		return
	}
	v = f.config[key]
	return
}

func (f *myConfig) UnsafeValue(key string) (v interface{}) {
	v = f.config[key]
	return
}

func (f *myConfig) Init(filepath string) error {
	config_file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	fi, _ := config_file.Stat()
	//if size := fi.Size(); size > (configFileSizeLimit) {
	//	emit("config file (%q) size exceeds reasonable limit (%d) - aborting", path, size)
	//	return &config // REVU: shouldn't this return an error, then?
	//}

	if fi.Size() == 0 {
		return errors.New("config file " + filepath + " is empty, skipping")
	}

	buffer := make([]byte, fi.Size())
	_, err = config_file.Read(buffer)

	//buffer, err = StripComments(buffer) //去掉注释
	//if err != nil {
	//	emit("Failed to strip comments from json: %s\n", err)
	//	return &config
	//}

	buffer = []byte(os.ExpandEnv(string(buffer))) //特殊

	err = json.Unmarshal(buffer, &f.config) //解析json格式数据
	if err != nil {
		return err
	}
	f.isInit = true
	return nil

}
