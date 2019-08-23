package loger

import (
	"encoding/json"
	"errors"
	"time"
)

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
func (cache *LogCache) SaveToCache(data interface{}, levelName string) error {
	logStruct := NewLogStruct()
	logStruct.Level = levelName
	logStruct.Data = data
	logStruct.Time = time.Now().Format(time.RFC3339)
	cache.Logs = append(cache.Logs, logStruct)
	return nil
}