package strategy

import (
	"os"

	"go-code-gen/pkg/config"
)

func fileCreate(filePath, tempFile string, c *config.Config) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = c.Temps.ExecuteTemplate(file, tempFile, c)
	if err != nil {
		panic(err)
	}
}

func fileAppend(filePath, tempFile string, c *config.Config) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = c.Temps.ExecuteTemplate(file, tempFile, c)
	if err != nil {
		panic(err)
	}
}

func Run(c *config.Config) {
	handler := new(handlerStrategy)
	service := new(serviceStrategy)
	msg := new(msgStrategy)
	strategy := make([]Strategy, 0, 3)
	strategy = append(strategy, handler, service, msg)
	for _, s := range strategy {
		s.Gen(c)
	}
}

func Runs(configs []*config.Config) {
	for _, c := range configs {
		Run(c)
	}
}
