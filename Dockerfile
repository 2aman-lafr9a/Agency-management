FROM golang:1.21.5-alpine





ENV GO111MODULE=on
ENV GOFLAGS-mod=vendor

ENV APP_HOME /go/src/agency_management

RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

# copy the dependencies file to the working directory
COPY go.mod ./go.sum ./


RUN go mod download

# copy the source code to the working directory
COPY *.go .

RUN go build -o .

EXPOSE 8080

CMD [ "/Agency_management" ]