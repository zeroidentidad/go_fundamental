all: compile docker push

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goregresionpredict

docker:
	sudo docker build --force-rm=true -t zeroidentidad/goregresionpredict .

push:
	sudo docker push zeroidentidad/goregresionpredict

clean:
	rm goregresionpredict
