# command

```shell

go build -o MyService

// register MyService
launchctl load /Library/LaunchDaemons/MyService.plist

// check status
launchctl list | grep My

// check result
tail -f MyService.log

// stop the service
launchctl unload /Library/LaunchDaemons/MyService.plist
```
