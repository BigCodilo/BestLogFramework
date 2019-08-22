package loger

import (
	"encoding/json"
	"fmt"
	"time"
)

//Println - print input data to file or console
func (blog BestLog) InfoPrintln(data interface{}) error{
	dataJSON, err := json.Marshal(data)
	if err != nil{
		return err
	}
	currentTime := time.Now().String()
	dataToPrint := currentTime + " - INFO: " + string(dataJSON)
	if blog.GetPath() == ""{
		fmt.Println(string(dataToPrint))
		return nil
	}
	_, err = blog.InfoFile.WriteString(dataToPrint + "\n")
	if err != nil{
		return err
	}
	return nil
}
