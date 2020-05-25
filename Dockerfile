############################################################################################################
################## Build                                                      ##############################
############################################################################################################
FROM golang:latest as build

RUN apt-get update && \
    apt-get install lsb-release -y

RUN go version

RUN apt-get update

# k8s CLI
RUN curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
RUN touch /etc/apt/sources.list.d/kubernetes.list
RUN echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" | tee -a /etc/apt/sources.list.d/kubernetes.list
RUN apt-get update && apt-get install -y kubectl

COPY . /go/src/github.com/conplementAG/k8s-semantic-detective

WORKDIR /go/src/github.com/conplementAG/k8s-semantic-detective/cmd/semantic-detective

RUN go build -o k8s-semantic-detective

############################################################################################################
################## Cloudprober                                                ##############################
############################################################################################################
FROM cloudprober/cloudprober:latest as prober

############################################################################################################
################## Finalcontainer                                             ##############################
############################################################################################################
FROM alpine:latest as final

# required for compiled Go binaries to run on alpine (preventing executable not found error)
RUN apk add --no-cache libc6-compat

COPY --from=build /usr/bin/kubectl /usr/bin/kubectl
COPY --from=prober /cloudprober /cloudprober
COPY --from=prober /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/src/github.com/conplementAG/k8s-semantic-detective/cmd/semantic-detective/k8s-semantic-detective /usr/bin/k8s-semantic-detective

ENTRYPOINT ["/usr/bin/k8s-semantic-detective", "probe"]