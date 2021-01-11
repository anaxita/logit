# logit

## Information

Logit is a simple underlogger for write log data into a file and terminal with date and levels.

## Install
`go get -u github.com/anaxita/logit`

## Using
```
func main() {

	// path to logfile
	logfile := "/var/log/events.log"

	// init logit
	err := logit.New(logfile) // ! this func must be used only once in main func

	if err != nil {
		log.Fatal(err)
	}

	defer logit.Close()

	// use Log func - you can do this func in any place in your program
	logit.Log("Logit!")
}
```
Result: `2021/01/06 02:17:56 Logit!`
## Info

`logit.Info("Logit!")` has green color and using for info data.

`logit.Log("Logit!")` has blue color for not fatal errors.

`logit.Fatal("Logit!")` has red color and using for fatal errors.

![example_log](https://i.ibb.co/YkL9wZN/log.png)

That's all. Enjoy!
