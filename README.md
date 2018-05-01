# cmd-tools
cmd-tools是一个命令行工具，旨在解决现在命令行工具不方便解决或不能解决的问题。

目前包含的功能包括：

1. 检查某端口是否已打开，因为在某些情况下nc或nmap不太好用。
1. 查询某主机的SRV列表，并格式化输出。

未来可能还会加入其它功能。

## Download
```
mkdir -p $GOPATH/src/github.com/sycki
git clone https://github.com/sycki/cmd-tools.git $GOPATH/src/github.com/sycki/
```

## Build
```
cd $GOPATH/src/github.com/sycki/cmd-tools/cmd
go install *
```

## Usage
```
$GOPATH/bin/net
```

