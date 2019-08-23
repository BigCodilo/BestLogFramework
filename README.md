**Example**

```
    import "blog "github.com/BigCodilo/BestLogFramework/loger"
    Loger = blog.NewBestLog() //Create new loger
    defer Loger.CloseFiles() //Close all opened file for log
    Loger.Debug.TurnOn() //Turn on level log
    Loger.Debug.SetFilePath("debugishe") //Set path to save logs
    Loger.Debug("params") // print debug level (if it included)
    Loger.PrintLog("jopa", "chlen", 7324) // print all turned on level
    Loger.SetStreamOutput(os.Writer) //set outputb stramm
    
    Loger.Error.TurnOn()
    Loger.Error.TurnOnCache() //Включает кэш
    Loger.Error.SetCachetime(time.Duration) //Устанавливает время через какое веш будет выводить записи и очищаться (по дефолту 10 сек)
    Loger.Error.PrintWithCache("jewdnkeiwnefjneoifnaliuaeflauhfluei") //Добавляет записи в кеш

```
 If path to files didn't set, logs will print in std.Out
 
 **Levels:**
 
 <ol>
 <li>Debug</li>
 <li>Info</li>
 <li>Warn</li>
 <li>Error</li>
 <li>Fatal</li>
 </ol>