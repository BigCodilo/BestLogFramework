package loger

import (
	"os"
)


type LogLevel struct{
	Level string
	TurnedOn bool
	FilePath string
	File *os.File
}

type BestLog struct{
	LInfo *LogLevel
	LDebug *LogLevel
}

//------------------------------------END STRUCTS----------------------//

func NewBestLog() (BestLog){
	return BestLog{
		LInfo:  &LogLevel{
			Level: "INFO",
		},
		LDebug: &LogLevel{
			Level: "DEBUG",
		},
	}
}

//Close - close opened files for printing
func (blog *BestLog) CloseFiles(){
	if blog.LInfo.File != nil{
		blog.LInfo.File.Close()
	}
	if blog.LDebug.File != nil{
		blog.LDebug.File.Close()
	}
}


//---------------------------------------------------------------------INFO---------------------------------------------------//

func (level *LogLevel) TurnOn(){
	level.TurnedOn = true
}

func (level *LogLevel) TurnOff(){
	level.TurnedOn = false
}

////OnInfo turn on info level
//func (blog *BestLog) OnInfo(){
//	blog.LInfo.TurnedOn = true
//}
//
////OffInfo turn off info level
//func (blog *BestLog) OffInfo(){
//	blog.LInfo.TurnedOn = false
//}

func (level *LogLevel) SetFilePath(path string) error{
	level.FilePath = path
	err := level.OpenFile()
	return err
}

//Setter for info path
//func (blog *BestLog)SetInfoPath(path string) error{
//	blog.LInfo.FilePath = path
//	err := blog.OpenInfoFile()
//	return err
//}

//Getter for info path
func (blog BestLog)GetInfoPath() string{
	return blog.LInfo.FilePath
}

func (level *LogLevel) OpenFile() error{
	var err error
	level.File, err = os.OpenFile(level.FilePath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil{
		file, err := os.Create(level.FilePath)
		if err != nil{
			return err
		}
		defer file.Close()
		level.File, err = os.OpenFile(level.FilePath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	}
	return nil
}

////Open info file by inputing path
//func (blog *BestLog) OpenInfoFile() error{
//	var err error
//	blog.LInfo.File, err = os.OpenFile(blog.LInfo.FilePath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
//	if err != nil{
//		file, err := os.Create(blog.LInfo.FilePath)
//		if err != nil{
//			return err
//		}
//		defer file.Close()
//		blog.LInfo.File, err = os.OpenFile(blog.LInfo.FilePath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
//	}
//	return nil
//}



//-----------------------------------------------------------------DEBUG------------------------------------------------------//

//OnDebug turn on debug level
func (blog *BestLog) OnDebug(){
	blog.LDebug.TurnedOn = true
}

//OffDebug turn on debug level
func (blog *BestLog) OffDebug(){
	blog.LDebug.TurnedOn = false
}

//Setter for info file path
func (blog *BestLog)SetDebugPath(path string) error{
	blog.LDebug.FilePath = path
	err := blog.OpenDebugFile()
	return err
}

//Getter for info file path
func (blog BestLog)GetDebugPath() string{
	return blog.LDebug.FilePath
}

//Open debug file by inputing path
func (blog *BestLog) OpenDebugFile() error{
	var err error
	blog.LDebug.File, err = os.OpenFile(blog.LDebug.FilePath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil{
		file, err := os.Create(blog.LDebug.FilePath)
		if err != nil{
			return err
		}
		defer file.Close()
		blog.LDebug.File, err = os.OpenFile(blog.LDebug.FilePath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	}
	return nil
}
