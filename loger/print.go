package loger

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"
)

func (level LogLevel) Print(data...interface{}) error{
	fromPA := reflect.TypeOf(data[0]).String() == "loger.FromPrintAll"
	if level.TurnedOn && !fromPA{
		err := level.println(data)
		if err != nil {
			return err
		}
	} else{
		return errors.New(level.LevelName + " level is off now")
	}
	return nil
}

//Все включенные уровни логирования - запись
func (blog BestLog) PrintAll(data...interface{}) error{
	if blog.Debug.TurnedOn{
		err := blog.Debug.println(data)
		if err != nil {
			return err
		}
	}
	if blog.Info.TurnedOn{
		err := blog.Info.println(data)
		if err != nil {
			return err
		}
	}
	if blog.Warn.TurnedOn{
		err := blog.Warn.println(data)
		if err != nil {
			return err
		}
	}
	if blog.Error.TurnedOn{
		err := blog.Error.println(data)
		if err != nil {
			return err
		}
	}
	if blog.Fatal.TurnedOn{
		err := blog.Fatal.println(data)
		if err != nil {
			return err
		}
	}
	return nil
}

//Запись в уровень
func (level *LogLevel) println(data interface{}) error{

	logStruct := NewLogStruct()
	logStruct.Level = level.LevelName
	logStruct.Data = data
	logStruct.Time = time.Now().Format(time.RFC3339)
	logBinary, err := json.Marshal(logStruct)
	if err != nil{
		return err
	}
	logString := string(logBinary)

	if level.GetFilePath() == ""{
		fmt.Println(logString)
		return nil
	}
	_, err = level.File.WriteString(logString + "\n")
	if err != nil{
		return err
	}
	return nil
}