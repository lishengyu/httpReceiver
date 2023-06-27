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

# 帮助信息
./httpReceiver -h
