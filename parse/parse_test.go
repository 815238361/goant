package parse

import (
	"goant/log"
	"os"
	"testing"
)

func TestParse(t *testing.T) {

	simpleLog:=log.New(os.Stdout)

	genteel :=`<project name="simple">
               	 <property name="flag" value="1"/>
                 <target name="mkdir">
					<mkdir dirname="c://dir"/>
	             </target>
               </project>`
	Parse(simpleLog,[]byte(genteel))

}
