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

//--------WITH_CACHE----------//

//Добавляет запись в кеш если уровень логирования включен
func (level *LogLevel) PrintWithCache(data...interface{}) error{
	if level.TurnedOn{
		level.Cache.SaveToCache(data, level.LevelName)
	} else{
		return errors.New(level.LevelName + " level is off now")
	}
	return nil
}

//Раз в какой-то промежуток времени выводит все записи с кеша и очищает его
func (level *LogLevel)UnloadCache(){
	for level.Cache.TurnedOn{
		logsArray := []LogStruct{}
		for _, v := range level.Cache.Logs{
			logsArray = append(logsArray, v)
		}
		logeArrayJSON, _ := json.Marshal(logsArray)
		level.Stream.Write(logeArrayJSON)
		level.Cache.Logs = []LogStruct{}
		time.Sleep(level.Cache.SleepTime)
	}
}

//Сохраняет в кеш строку которую формирует тут
func (cache *LogCache) SaveToCache(data interface{}, levelName string) error{
	logStruct := NewLogStruct()
	logStruct.Level = levelName
	logStruct.Data = data
	logStruct.Time = time.Now().Format(time.RFC3339)
	cache.Logs = append(cache.Logs, logStruct)
	return nil
}