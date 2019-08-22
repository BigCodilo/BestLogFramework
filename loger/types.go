package loger

import (
	"os"
)

func init(){
	os.Setenv("BLOG_INFO_LEVEL", "true")
	os.Setenv("BLOG_DEBUG_LEVEL", "false")
}

type BestLog struct{
	IsInfo bool
	IsDebug bool
	InfoPath string
	DebugPath string
	infoFile *os.File
	debugFile *os.File
}

//NewBestLog - return new logger with path
func NewBestLog(isInfo, isDebug bool, infoPath ,debugPath string) (BestLog, error){
	blog := BestLog{}
	err := blog.SetInfoPath(infoPath)
	if err != nil{
		return BestLog{}, err
	}
	err = blog.SetDebugPath(debugPath)
	if err != nil{
		return BestLog{}, err
	}
	blog.IsInfo = isInfo
	blog.IsDebug = isDebug
	return blog, nil
}

//Close - close opened files for printing
func (blog *BestLog) CloseFiles(){
	if blog.infoFile != nil{
		blog.infoFile.Close()
	}
	if blog.debugFile != nil{
		blog.debugFile.Close()
	}
}


//---------------------------------------------------------------------INFO---------------------------------------------------//

//OnInfo turn on info level
func (blog *BestLog) OnInfo(){
	blog.IsInfo = true
}

//OffInfo turn off info level
func (blog *BestLog) OffInfo(){
	blog.IsInfo = false
}

//Setter for info path
func (blog *BestLog)SetInfoPath(path string) error{
	blog.InfoPath = path
	err := blog.OpenInfoFile()
	return err
}

//Getter for info path
func (blog BestLog)GetInfoPath() string{
	return blog.InfoPath
}

//Open info file by inputing path
func (blog *BestLog) OpenInfoFile() error{
	var err error
	blog.infoFile, err = os.OpenFile(blog.InfoPath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil{
		file, err := os.Create(blog.InfoPath)
		if err != nil{
			return err
		}
		defer file.Close()
		blog.infoFile, err = os.OpenFile(blog.InfoPath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	}
	return nil
}



//-----------------------------------------------------------------DEBUG------------------------------------------------------//

//OnDebug turn on debug level
func (blog *BestLog) OnDebug(){
	blog.IsDebug = true
}

//OffDebug turn on debug level
func (blog *BestLog) OffDebug(){
	blog.IsDebug = false
}

//Setter for info file path
func (blog *BestLog)SetDebugPath(path string) error{
	blog.DebugPath = path
	err := blog.OpenDebugFile()
	return err
}

//Getter for info file path
func (blog BestLog)GetDebugPath() string{
	return blog.DebugPath
}

//Open debug file by inputing path
func (blog *BestLog) OpenDebugFile() error{
	var err error
	blog.debugFile, err = os.OpenFile(blog.DebugPath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil{
		file, err := os.Create(blog.DebugPath)
		if err != nil{
			return err
		}
		defer file.Close()
		blog.debugFile, err = os.OpenFile(blog.DebugPath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	}
	return nil
}
