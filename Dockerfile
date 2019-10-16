FROM golang:latest

RUN apt-get update && \
    apt-get install lsb-release -y

RUN apt-get install apt-transport-https

RUN go version

# dep package manager
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

ADD . /go/src/github.com/conplementAG/k8s-semantic-detective
WORKDIR /go/src/github.com/conplementAG/k8s-semantic-detective

# restore packages
RUN dep ensure

# simple build
WORKDIR /go/src/github.com/conplementAG/k8s-semantic-detective/cmd/semantic-detective
RUN go build -o semantic-detective .

# run the tests
WORKDIR /go/src/github.com/conplementAG/k8s-semantic-detective
RUN go test ./... --cover

CMD [ "/bin/bash" ]