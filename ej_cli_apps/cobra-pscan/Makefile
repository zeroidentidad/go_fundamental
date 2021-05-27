all: win64 default

win64:
	GOOS=windows GOARCH=amd64 go build -o pScan.exe

default:
	GOOS=linux GOARCH=amd64 go build -o pScan

clean:
	rm pScan.exe && rm pScan