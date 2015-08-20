all:
	go build *.go

clean:
	find . -perm +100 -type f -delete
