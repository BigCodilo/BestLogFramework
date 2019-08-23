**Example**

```
    import "blog "github.com/BigCodilo/BestLogFramework/loger"
    Loger = blog.NewBestLog() //Create new loger
    defer Loger.CloseFiles() //Close all opened file for log
    Loger.LDebug.TurnOn() //Turn on level log
    Loger.LDebug.SetFilePath("debugishe") //Set path to save logs
    Loger.Gebug("params") // print debug level (if it included)
    Loger.PrintLog("jopa", "chlen", 7324) // print all turned on level

```
 If path to files didn't set, logs will print in console