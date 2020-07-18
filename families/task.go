package families

import (
	"fmt"
	"goant/log"
	"os"
)

const (
	MKDIR ="mkdir"
	ECHO="echo"
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
