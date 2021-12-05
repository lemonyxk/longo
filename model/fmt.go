/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2021-12-06 01:01
**/

package model

import (
	"os/exec"
)

func goFmt(filePath string) {
	var c = exec.Command("gofmt", "-w", filePath)
	err := c.Start()
	if err != nil {
		panic(err)
	}

	_, err = c.Process.Wait()
	if err != nil {
		panic(err)
	}
}
