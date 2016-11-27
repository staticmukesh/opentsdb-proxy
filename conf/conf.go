package conf

import (
	"fmt"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

type Conf struct {
	Port            int
	Servers         []string
	LimitConnection int
}

func ReadConf(flagConf *string) *Conf {
	conf := &Conf{
		Port:            8080,
		Servers:         []string{},
		LimitConnection: 1000,
	}

	filePath := *flagConf

	if filePath == "" {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error: ", err.Error())
			os.Exit(1)
		}

		filePath = path.Join(pwd, "opentsdb.toml")
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	defer file.Close()

	_, err = toml.DecodeReader(file, conf)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	return conf
}
