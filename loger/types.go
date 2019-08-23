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

func (level *LogLevel) SetFilePath(path string) error{
	level.FilePath = path
	err := level.OpenFile()
	return err
}

func (level LogLevel) GetFilePath() string{
	return level.FilePath
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