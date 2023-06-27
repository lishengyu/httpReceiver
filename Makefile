all:
	go build -o httpReceiver *.go

clean:
	rm -rf httpReceiver 
