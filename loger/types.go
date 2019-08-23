package loger

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

type FromPrintAll bool

//Struct for logginr (text)
type LogStruct struct{
	Time string
	Level string
	Data interface{}
}

//Decribe logs level
type LogLevel struct{
	//LevelName a name of current level (DEBUG, INFO, WARN, ERROR, FATAL)
	LevelName string
	//TurnedOn true if current level is turned on
	TurnedOn bool
	//Path to outputing file
	FilePath string
	//Stream in wich data will be sent
	Stream io.Writer
}

//Consist of logs level
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
			Stream: os.Stdout,
		},
		Info:  &LogLevel{
			LevelName: "INFO",
			Stream: os.Stdout,
		},
		Warn: &LogLevel{
			LevelName: "WARN",
			Stream: os.Stdout,
		},
		Error: &LogLevel{
			LevelName: "ERROR",
			Stream: os.Stdout,
		},
		Fatal: &LogLevel{
			LevelName: "FATAL",
			Stream: os.Stdout,
		},
	}
}

//Close - close opened files for printing
func (blog *BestLog) CloseFiles(){
	fmt.Println(reflect.TypeOf(blog.Debug.Stream))
	if blog.Debug.Stream != nil && reflect.TypeOf(blog.Debug.Stream).String() == "*os.File"{
		blog.Debug.Stream.(*os.File).Close()
	}
	if blog.Info.Stream != nil && reflect.TypeOf(blog.Info.Stream).String() == "*os.File"{
		blog.Info.Stream.(*os.File).Close()
	}
	if blog.Warn.Stream != nil && reflect.TypeOf(blog.Warn.Stream).String() == "*os.File"{
		blog.Warn.Stream.(*os.File).Close()
	}
	if blog.Error.Stream != nil && reflect.TypeOf(blog.Error.Stream).String() == "*os.File"{
		blog.Error.Stream.(*os.File).Close()
	}
	if blog.Fatal.Stream != nil && reflect.TypeOf(blog.Fatal.Stream).String() == "*os.File"{
		blog.Fatal.Stream.(*os.File).Close()
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
	level.Stream, err = os.OpenFile(level.FilePath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil{
		file, err := os.Create(level.FilePath)
		if err != nil{
			return err
		}
		defer file.Close()
		level.Stream, err = os.OpenFile(level.FilePath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	}
	return nil
}

func (level *LogLevel) SetStreamOutput(stream io.Writer){
	level.Stream = stream
	level.FilePath = ""
}