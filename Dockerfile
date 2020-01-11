FROM golang:latest

WORKDIR /go/src

COPY . .

RUN [ "go","build","-o","app" ]


FROM hpprc/mysql_playground:latest
COPY --from=0 /go/src/app .