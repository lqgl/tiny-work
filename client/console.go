package main

import (
	"bufio"
	"github.com/lqgl/tinywork/logger"
	"os"
	"strings"
)

type Console struct {
	chInput chan *InputParam
}

type InputParam struct {
	Command string
	Param   []string
}

func NewConsole() *Console {
	c := &Console{}
	return c
}

func (c *Console) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			logger.Logger.DebugF("input err ,check your input and  try again !!!")
			continue
		}
		strings.TrimSpace(readString)
		readString = strings.Replace(readString, "\n", "", -1)
		readString = strings.Replace(readString, "\r", "", -1)
		split := strings.Split(readString, " ")
		if len(split) == 0 {
			logger.Logger.DebugF("input err, check your input and  try again !!! ")
			continue
		}
		in := &InputParam{
			Command: split[0],
			Param:   split[1:],
		}
		c.chInput <- in // 发送接收到的终端数据
	}
}
