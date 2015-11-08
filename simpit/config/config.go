

package config


import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	//"errors"
)

type Node struct {
	Alias string 	` yaml:"alias" `
	Comment string 	` yaml:"comment" `
	Prop string 	` yaml:"prop" `

}


type Config struct {
	Model string 	` yaml:"model" `
	//DInPins []InputPin	` yaml:"digital_inputs" `
	//AInPins []InputPin	` yaml:"analog_inputs" `
	//DOutPins []OutputPin	` yaml:"digital_outputs" `
	//zxOutputs map[int]string
	Nodes []Node
}

func Load(file_path string) (*Config, error) {



	contents, err_file := ioutil.ReadFile(file_path)
	if err_file != nil {
		return nil, err_file
	}

	conf := new(Config)
	err_yaml := yaml.Unmarshal(contents, &conf)
	if err_yaml != nil {
		return nil, err_yaml
	}

	//conf.LoadFgNodes()
	/*
	conf.LedMap = make(map[string]int)
	for _, led := range conf.Leds {
		conf.LedMap[led.Node] = led.Index
	}
	*/
	//err := conf.Validate()
	//if err != nil {
	//	return conf, err
	//}

	return conf, nil
}

/*
func (me *Config) Validate() error {

	exists := make(map[int]bool)
	mess := ""

	for _, p := range me.DOutPins {
		if p.Pin > 7 {
			mess +=  "OutPin " + p.Prop + " has index > 7\n"
		}
		_, found := exists[p.Pin]
		if found {
			mess +=  "OutPin " + p.Prop + " has duplicate index\n"
		}
		exists[p.Pin] = true
	}
	if mess == "" {
		return nil
	}
	return errors.New(mess)

}
*/