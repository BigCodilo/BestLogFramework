**Example**

```
Loger = blog.BestLog{}
defer Loger.CloseFiles() //close txts files (path to log output)
Loger.OnInfo() //turn on info level logger 
Loger.OnDebug() //turn on debug level logger 
err = Loger.SetDebugPath("ddeebug") //set path to output
Loger.PrintLog("jopa", "chlen", 7324) //prints logs on all incluted level
Loger.Info("params") // print info level (if it included)
```
 If path to files didn't set, logs will print in console