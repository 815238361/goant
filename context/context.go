package context

import (
	"goant/families"
	"goant/log"
)

type Context struct {

	project *families.Project

	targetNums int

	targetNames []string

	defaultTarget string

	Logger log.Logger

	BuildFilePath string
}


func (c *Context) SetProject(p *families.Project){

	c.project = p

}

func (c *Context) InitContext(){

	c.targetNums = len(c.project.Targets)

	if c.targetNames == nil {

		c.targetNames = make([]string,2)
	}

	for  _, t:= range c.project.Targets{

		c.targetNames = append(c.targetNames,t.Name)

	}

	c.defaultTarget = c.project.Default

	c.checkContext()

}



func (c *Context) checkContext(){


}

func (c *Context) checkDefaultTargetExist(){


}

func (c *Context) Exec(){

	p:= c.project
	targets:=p.Targets

	for _,t:=range targets{

		for _,ts:= range t.Tasks{

			ts.Exec()
		}
	}
}