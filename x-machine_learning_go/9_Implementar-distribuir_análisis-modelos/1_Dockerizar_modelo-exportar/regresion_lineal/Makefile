all: compile docker push

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goregresiontrain

docker:
	sudo docker build --force-rm=true -t zeroidentidad/goregresiontrain:lineal .

push:
	sudo docker push zeroidentidad/goregresiontrain:lineal

clean:
	rm goregresiontrain
