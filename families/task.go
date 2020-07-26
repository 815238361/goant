package families

import (
	"fmt"
	"goant/log"
	"os"
	"os/exec"
)

const (
	MKDIR ="mkdir"
	ECHO="echo"
	JAVAC="javac"
)


type Task interface {
	Exec()
	SetTarget(t *Target)
	SetLogger(log log.Logger)
}

type TaskBase struct {

	logger log.Logger
	target *Target
}

func (v *TaskBase) SetTarget(t *Target){
	v.target = t
}

func (v *TaskBase) SetLogger(log log.Logger ){
	v.logger = log
}

//mkdir task start
type MkDirTask struct {

	TaskBase
	DirName string
}

func (mkDir *MkDirTask) Exec(){

	dirName:=mkDir.DirName

	err:=os.Mkdir(dirName,os.ModePerm)

	if err!=nil {

		strErr:= fmt.Sprintf("%s target , %s task make dir(%s) failed.",mkDir.target.Name,MKDIR,mkDir.DirName)
		panic(strErr)
	}

}

//mkdir task end

//echo task start
type EchoTask struct {

	TaskBase
	Es string
}

func (echo *EchoTask) Exec(){

	echo.logger.Info(echo.Es)

}
//echo task end


//javac task start
type JavacTask struct {

	TaskBase
	Srcpath string
	Despath string
	Classpath string
}

func (javac *JavacTask) Exec(){

	var cmd *exec.Cmd

	var strCmd string = "javac"
	var args1 string = javac.Srcpath
	var args2 string = " -d "+javac.Despath

	if javac.Classpath==""{

		cmd= exec.Command(strCmd,args1,args2)

	}else{

		var args3 string = " -classpath "+javac.Classpath
		cmd= exec.Command(strCmd,args1,args3,args2)

	}

	err:=cmd.Run()

	if err!=nil {

		panic(err)
	}
}
//javac task end