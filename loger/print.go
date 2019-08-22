package loger

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

//Логирует уровня инфо
func (blog BestLog) Info(data...interface{}) error{
	if blog.IsInfo{
		err := blog.infoPrintln(data)
		if err != nil {
			return err
		}
	} else{
		return errors.New("Info level is off now")
	}
	return nil
}

//Логирует уровень дебаг
func (blog BestLog) Debug(data...interface{}) error{
	if blog.IsDebug{
		err := blog.debugPrintln(data)
		if err != nil {
			return err
		}
	} else{
		return errors.New("Debug level is off now")
	}
	return nil
}

//Все включенные уровни логирования - выводит
func (blog BestLog) PrintLogs(data...interface{}) error{
	if blog.IsDebug{
		err := blog.debugPrintln(data)
		if err != nil {
			return err
		}
	}
	if blog.IsInfo{
		err := blog.infoPrintln(data)
		if err != nil {
			return err
		}
	}
	return nil
}

//Println - print log to file or console
func (blog BestLog) infoPrintln(data interface{}) error{
	dataToPrint, err := MakeString("INFO", data)
	if err != nil{
		return err
	}
	if blog.GetInfoPath() == ""{
		fmt.Println(dataToPrint)
		return nil
	}
	_, err = blog.infoFile.WriteString(dataToPrint + "\n")
	if err != nil{
		return err
	}
	return nil
}

//Println - print log to file or console
func (blog BestLog) debugPrintln(data interface{}) error{
	dataToPrint, err := MakeString("DEBUG", data)
	if err != nil{
		return err
	}
	if blog.GetDebugPath() == ""{
		fmt.Println(dataToPrint)
		return nil
	}
	_, err = blog.debugFile.WriteString(dataToPrint + "\n")
	if err != nil{
		return err
	}
	return nil
}

//Make string with log level, time and data
func MakeString(level string, data interface{}) (string, error){
	dataJSON, err := json.Marshal(data)
	if err != nil{
		return "", err
	}
	currentTime := time.Now().Format(time.RFC3339)
	dataToPrint := currentTime + ": " + level + ": " + string(dataJSON)
	return dataToPrint, nil
}