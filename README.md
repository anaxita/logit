# logit

## Information

Logit is a simple logger for writing log data into a file or/and terminal with three levels logging.
Logit use custom date and color views.

## Install
`go get -u github.com/anaxita/logit`

## Using
```
func main() {
	logfilePath := "/var/log/events.log"

	// New func must be used only once in the main func
	_ := logit.New(logfilePath) 
	
	defer logit.Close()

	// use Log func - you can do this func in any place in your program
	logit.Log("Logit!")
}
```
Result: `2021.01.06 02:17:56 [ERROR] Logit!`
## Info

`logit.Info("Logit!")` has green color and using for info data.

`logit.Log("Logit!")` has blue color for not fatal errors.

`logit.Fatal("Logit!")` has red color and using for fatal errors.

![example_log](https://i.ibb.co/YkL9wZN/log.png)

That's all. Enjoy!
