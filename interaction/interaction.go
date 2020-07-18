package interaction

import (
	"bytes"
	"fmt"
	"goant/context"
	"goant/log"
	"io"
	"os"
)

var args string = "file,"

var params string ="-D,-h"

func Exec(c *context.Context){

	if c.Logger == nil {

		c.Logger = log.New(os.Stdout)
	}

	c.Exec()
}

func ParseArguments(args []string,c *context.Context){


}

func ReadGoAntBuildXML(c *context.Context) []byte{

	var bts []byte = make([]byte,512)
	var des []byte = make([]byte,2048)
	var n int

	f,err:=os.Open(c.BuildFilePath)

	defer f.Close()

	if err!=nil {

		str:=fmt.Sprintf("%s not exist!",c.BuildFilePath)
		panic(str)
	}

	bf:=bytes.NewBuffer(des)

	for {

		n,err = f.Write(bts)

		if err!=nil {

			str:=fmt.Sprintf("%s parse failed!",c.BuildFilePath)
			panic(str)
		}

		if err==io.EOF{

			break
		}

		temp:=make([]byte,n)

		copy(temp,bts)

		bf.Read(temp)

	}

	return bf.Bytes()
}


func CheckExec(inputTargetName string,c *context.Context) bool{


	return false
}

func CheckExecTargets(inputTargetNames []string,c *context.Context) bool{


	return false
}
