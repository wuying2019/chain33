[![pipeline status](https://api.travis-ci.org/bityuan/bityuan.svg?branch=master)](https://travis-ci.org/bityuan/bityuan/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bityuan/bityuan)](https://goreportcard.com/report/github.com/bityuan/bityuan)

# GXL公有链系统

#### 编译

```
git clone https://github.com/wuying2019/gxlchain.git $GOPATH/src/github.com/wuying2019/gxlchain
cd $GOPATH/src/github.com/wuying2019/gxlchain
go build -i -o gxl
go build -i -o gxl-cli github.com/wuying2019/gxlchain/cli
```

#### 运行

拷贝编译好的gxl, gxl-cli, gxlChain.toml这三个文件置于同一个文件夹下，执行：

```
./gxl -f gxlChain.toml
```


