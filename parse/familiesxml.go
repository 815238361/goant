package parse

import (
	"encoding/xml"
	"fmt"
	"goant/families"
	"goant/log"
	"strings"
)

const (
	MKDIR ="mkdir"
	ECHO="echo"
	JAVAC="javac"
)

var taskValidMap map[string]ValidateTaskXML = make(map[string]ValidateTaskXML,0)


func init(){

	//mkdir
	taskValidMap[MKDIR] = &MkDirValidateTaskXML{}
	//echo
	taskValidMap[ECHO] = &EchoValidateTaskXML{}
	//javac
	taskValidMap[JAVAC] = &JavacValidateTaskXML{}
}

type ProjectXML struct {
	Name     string      `xml:"name,attr"`
	FilePath string      `xml:"path,attr"`
	Default string       `xml:"default,attr"`
	Targets  []TargetXML `xml:"target"`
	Properties []PropertyXML `xml:"property"`
}


type TargetXML struct {
	Content   string `xml:",innerxml"`
	Name      string `xml:"name,attr"`
	DependsOn string `xml:"dependson,attr"`
	Execution string `xml:"execution,attr"`
}

type PropertyXML struct {
	Name      string `xml:"name,attr"`
	AttrValue string `xml:"value,attr"`
	Value     string `xml:"Value"`
}

type ValidateTaskXML interface{
	ValidateAndInit(t *families.Target)
	SetLogger(logger log.Logger)
	SetContent(content string)
	SetTargetTokens(tokens map[string]xml.StartElement)
}

type ValidateTaskXMLBase struct {

	logger log.Logger
	content string
	nametagmap map[string]xml.StartElement
}

func(v *ValidateTaskXMLBase) SetLogger(logger log.Logger){
	v.logger = logger
}

func(v *ValidateTaskXMLBase) SetContent(content string){
	v.content = content
}

func(v *ValidateTaskXMLBase) SetTargetTokens(eles map[string]xml.StartElement){
	v.nametagmap = eles
}

//mkdir task start
type MkDirValidateTaskXML struct {

	ValidateTaskXMLBase
}

func (mkDir *MkDirValidateTaskXML) ValidateAndInit(t *families.Target){

	var dirname string

	ttm:=mkDir.nametagmap
	//mkdir StartElement
	mt:=ttm[MKDIR]

	for _,attr:=range mt.Attr{

		if attr.Name.Local == "dirname" {

			dirname = attr.Value

			break
		}
	}

	if dirname==""{

		str:=fmt.Sprintf("%s target %s task must have dirname attribute",t.Name,MKDIR)
		panic(str)
	}

	mkdir:=&families.MkDirTask{}
	mkdir.DirName = dirname
	mkdir.SetLogger(mkDir.logger)
	mkdir.SetTarget(t)

	t.AddTask(mkdir)
}

//mkdir task end

//echo task start
type EchoValidateTaskXML struct {

	ValidateTaskXMLBase
}

func (echo *EchoValidateTaskXML) ValidateAndInit(t *families.Target){

	var tagname string
	var esbs string
	r:=strings.NewReader(echo.content)
	decoder:=xml.NewDecoder(r)

	for token,err:=decoder.Token();err==nil;token,err=decoder.Token(){

		switch token.(type) {

		case xml.StartElement:
			ele,_:=token.(xml.StartElement)
			tagname=ele.Name.Local
		case xml.EndElement:
			ele,_:=token.(xml.EndElement)
			tagname=ele.Name.Local
			if tagname==ECHO {
				break
			}
		case xml.CharData:
			if tagname==ECHO {
				esbs = string([]byte(token.(xml.CharData)))
			}
		}
	}

	if esbs=="" {

		strErr:= fmt.Sprintf("%s target , %s task must have echo content",t.Name,ECHO)
		panic(strErr)

	}

	echoTask:=&families.EchoTask{}
	echoTask.SetLogger(echo.logger)
	echoTask.Es = esbs

	t.AddTask(echoTask)
}
//echo task end

//javac task start
type JavacValidateTaskXML struct {

	ValidateTaskXMLBase
}

func (javac *JavacValidateTaskXML) ValidateAndInit(t *families.Target){

	var src string="src"
	var dest string="dest"
	var cp string="classpath"
	var path string="path"

	var srcpath string
	var destpath string
	var classpath string

	ntm:=javac.nametagmap

	ele:=ntm[JAVAC]

	for _,attr := range ele.Attr{

		if attr.Name.Local == src {

			srcpath = attr.Value
		}

		if attr.Name.Local == dest {

			destpath = attr.Value
		}

		if attr.Name.Local == cp {

			classpath = attr.Value
		}

	}

	//src element
	ele = ntm[src]

	for _,attr:=range ele.Attr{

		if attr.Name.Local == path {

			srcpath = attr.Value
		}
	}

	//dest element
	ele = ntm[dest]

	for _,attr:=range ele.Attr{

		if attr.Name.Local == path {

			destpath = attr.Value
		}
	}

	if ele,ok := ntm[cp]; ok {

		for _,attr:=range ele.Attr{

			if attr.Name.Local == path {

				classpath = attr.Value
			}
		}
	}




	if srcpath==""{

		strErr:= fmt.Sprintf("%s target , %s task must have src",t.Name,JAVAC)
		panic(strErr)

	}

	if destpath==""{

		strErr:= fmt.Sprintf("%s target , %s task must have dest",t.Name,JAVAC)
		panic(strErr)

	}


	javacTask:=&families.JavacTask{

		Srcpath:srcpath,
		Despath:destpath,
		Classpath:classpath,
	}

	javacTask.SetLogger(javac.logger)

	t.AddTask(javacTask)

}
//javac task end