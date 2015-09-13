# Tcpheartping

This is my tools for test tcp port

It can keep on contact your remote special tcp port and say hi or something to keep your remote port keeping listening

# Compile

You need golang installed.
```
git clone https://github.com/netroby/tcpheartping.git 
cd tcpheartping
go build
```
# Run

```
# tcpheartping 127.0.0.1 8123
tcpheartping www.baidu.com 80
```

# Exit

By default the program will running , but if you want to break, you may press CTRL + c to break it down.
Or you can run kill command to let it down
