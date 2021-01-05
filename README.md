# logit - small package for simple logging in a file and terminal.

## Install
`go get -u github.com/anaxita/logit`

## Using
```
func main() {

    // path to logfile
    logfile := "/var/log/events.log"

    // init logit
	err := logit.New(logfile)

	if err != nil {
		log.Fatal(err)
	}
	
	defer logit.Close()

    // use Log func - you can do this func in any place in your program
	logit.Log("Logit!") // 2021/01/06 02:17:56 Logit!
}
```

## Info
That's all. Enjoy!