package loger

import (
	"encoding/json"
	"errors"
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
		blog.Debug.println(data)
	}
	if blog.Info.TurnedOn{
		blog.Info.println(data)
	}
	if blog.Warn.TurnedOn{
		blog.Warn.println(data)
	}
	if blog.Error.TurnedOn{
		blog.Error.println(data)
	}
	if blog.Fatal.TurnedOn{
		blog.Fatal.println(data)
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

	_, err = level.Stream.Write([]byte(logString + "\n"))
	if err != nil{
		return err
	}
	return nil
}
