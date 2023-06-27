# 程序运行
## https
./httpReceiver -port=9999 -url=/test -https
## http
./httpReceiver -port=9999 -url=/test

# 模拟访问
## https
curl --location -k --request POST 'https://localhost:9999/test' -H "Content-Disposition:form-data;filename=mtx_all_234221124203.rule;md5=15b42e0262015b95d15af71f75a850d6;batch=20230620" -F 'file=@/home/syli/work/go/src/gin/Makefile'
## http
curl --location -k --request POST 'http://localhost:9999/test' -H "Content-Disposition:form-data;filename=mtx_all_234221124203.rule;md5=15b42e0262015b95d15af71f75a850d6;batch=20230620" -F 'file=@/home/syli/work/go/src/gin/Makefile'

# pprof
## 内存分析
### 采集
curl http://localhost:8080/debug/pprof/heap?seconds=10 > heap.out
### 分析
命令行分析: go tool pprof heap.out
            常用命令: top10/list

web分析:go tool pprof -http=:8000 heap.out

# 帮助信息
./httpReceiver -h
