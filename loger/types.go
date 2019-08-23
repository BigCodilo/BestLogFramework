package loger

import (
	"os"
)

type FromPrintAll bool

type LogStruct struct{
	Time string
	Level string
	Data interface{}
}

type LogLevel struct{
	LevelName string
	TurnedOn bool
	FilePath string
	File *os.File
}

type BestLog struct{
	Debug *LogLevel
	Info *LogLevel
	Warn *LogLevel
	Error *LogLevel
	Fatal *LogLevel
}

//------------------------------------END STRUCTS----------------------//

func NewLogStruct() LogStruct{
	return LogStruct{}
}

func NewBestLog() BestLog{
	return BestLog{
		Debug: &LogLevel{
			LevelName: "DEBUG",
		},
		Info:  &LogLevel{
			LevelName: "INFO",
		},
		Warn: &LogLevel{
			LevelName: "WARN",
		},
		Error: &LogLevel{
			LevelName: "ERROR",
		},
		Fatal: &LogLevel{
			LevelName: "FATAL",
		},
	}
}

//Close - close opened files for printing
func (blog *BestLog) CloseFiles(){
	if blog.Debug.File != nil{
		blog.Debug.File.Close()
	}
	if blog.Info.File != nil{
		blog.Info.File.Close()
	}
	if blog.Warn.File != nil{
		blog.Warn.File.Close()
	}
	if blog.Error.File != nil{
		blog.Error.File.Close()
	}
	if blog.Fatal.File != nil{
		blog.Fatal.File.Close()
	}
}


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