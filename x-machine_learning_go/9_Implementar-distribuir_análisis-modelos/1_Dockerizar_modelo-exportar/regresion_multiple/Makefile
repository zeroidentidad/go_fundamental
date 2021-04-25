all: compile docker push

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goregresiontrain

docker:
	sudo docker build --force-rm=true -t zeroidentidad/goregresiontrain:multiple .

push:
	sudo docker push zeroidentidad/goregresiontrain:multiple

clean:
	rm goregresiontrain
