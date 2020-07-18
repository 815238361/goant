package parse

import (
	"encoding/xml"
	"goant/families"
	"goant/log"
	"strings"
)

const (

	SERIAL = "serial"

	PARALLEL =  "parallel"

	TARGETTAGNAME="target"

	TARGETTAGSTART = "<target>"

	TARGETTAGEND = "</target>"

)


//first normal case
func Parse(logger log.Logger,data []byte) *families.Project {

	ppx:=&ProjectXML{Name: "none",FilePath:"none"}

	err:=xml.Unmarshal(data,ppx)

	if err!=nil{

		logger.Err("parse goant xml failed! "+ err.Error())

		return nil

	}

	if ppx.Name == "" {

		logger.Err("project must have name attribute")

		return nil
	}

	pp:= families.Project{
		Name:ppx.Name,
		FilePath:ppx.FilePath,
		Default:ppx.Default,
		Properties:make([]families.Property,0,100),
		Targets:make([]families.Target,0,100),
	}

	//targetXMLs
	txls:= ppx.Targets

	//propertieXMLS
	pxls:= ppx.Properties

	//pxl propertyXML
	for _,pxl:= range pxls{

		pxlVal:=strings.TrimSpace(pxl.Value)
		pxlAttrVal:=strings.TrimSpace(pxl.AttrValue)
		pxlName:=strings.TrimSpace(pxl.Name)
		//Check property name attribute
		if  pxlName=="" {

			logger.Err("property must have name attribute")

			return nil

		}

		//Check property value subtag ,value attribute
		if pxlVal=="" && pxlAttrVal==""{

			logger.Err("property must have value attribute or value sub tag")

			return nil
		}

		if pxlVal!="" && pxlAttrVal!="" {

			logger.Err("property just must have one of value attribute and value sub tag")

			return nil
		}

		if pxlVal!="" || pxlAttrVal!=""{

			if pxlVal!="" {
				pp.AddProperty(families.Property{Name: pxlName,Value:pxlVal})
			}

			if pxlAttrVal!="" {
				pp.AddProperty(families.Property{Name: pxlName,Value:pxlVal})
			}
		}


	}

	//txl targetXML parse the task
	for _,txl:= range txls{

		var nametagmap map[string]xml.StartElement = make(map[string]xml.StartElement)

		txlName:=strings.TrimSpace(txl.Name)
		txlDp:=strings.TrimSpace(txl.DependsOn)
		txlEx:=strings.TrimSpace(txl.Execution)
		txlContent:=strings.TrimSpace(txl.Content)

		if txlName==""{

			logger.Err("target must have name attribute")

			return nil

		}

		if txlEx!="" && txlEx!=SERIAL && txlEx!=PARALLEL {

			logger.Errf("target execution attribute value must is %s or %s",SERIAL,PARALLEL)

			return nil
		}

		target:= families.Target{}
		target.Name = txlName
		target.Execution = txlEx
		target.DependsOn = txlDp
		target.Tasks = make([]families.Task,0,2)

		//oper txlcontent
		if txlContent!="" {

			txlContent = strings.Join([]string{TARGETTAGSTART,txlContent,TARGETTAGEND},"")

			r :=strings.NewReader(txlContent)

			decoder:=xml.NewDecoder(r)

			for token,err:=decoder.Token();err==nil;token,err=decoder.Token(){

				if err!=nil{

					logger.Err("parse goant xml failed! "+ err.Error())

					return nil

				}

				switch token.(type) {

				case xml.StartElement:

					ele,ok:=token.(xml.StartElement)

					if ok {

						eleName:=ele.Name.Local

						if eleName == TARGETTAGNAME{

							continue
						}

						nametagmap[eleName] = ele
					}
				case xml.EndElement:
				case xml.CharData:
				}
			}


			for k,_:=range nametagmap{

				validateXml:= taskValidMap[k]

				if validateXml!=nil {

					validateXml.SetLogger(logger)
					validateXml.SetContent(txlContent)
					validateXml.SetTargetTokens(nametagmap)
					validateXml.ValidateAndInit(&target)
				}
			}

		}

		pp.AddTarget(target)
	}

	return &pp
}
