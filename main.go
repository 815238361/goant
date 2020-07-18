package main

import (
	"goant/context"
	"goant/interaction"
	"goant/parse"
	"os"
)

const (

	DEFAULT_GOANT_BUILDXML_FILENAME="build.xml"
)
func main(){

	c:= &context.Context{}

	args:=os.Args[1:]

	interaction.ParseArguments(args,c)

	c.SetProject(parse.Parse(c.Logger,interaction.ReadGoAntBuildXML(c)))

	c.InitContext()

	interaction.Exec(c)
}