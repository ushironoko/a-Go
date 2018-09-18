# a-Go

超A&G recording with RTMPDump from Go(exec)

# Use

Get

```
git clone
```

Arrangement

* rtmpdump.exe or rtmpdump install

Make Output Directory

```
mkdir dist
```

Go Command Run

```
go run main.go

or

go build main.go
```

# How to Rec Settings

Switching outputDir, recTime, and connectServer, at current hour

* setting.json
  * startHour(int, Compare with current time) 
  * recTime(sec)
  * outputDir
  * connectServe

# Operation Check

* Windows 10 (Task scheduler)