FROM golang:1.15

RUN go get -v github.com/influxdata/influxdb1-client github.com/gin-gonic/gin github.com/mitchellh/mapstructure

RUN mkdir /code 
WORKDIR /code/ 
ADD . /code/ 

RUN ls
RUN printenv 

RUN chmod +x /code/dev-entrypoint.sh

EXPOSE 8080