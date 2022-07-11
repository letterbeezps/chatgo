FROM golang:alpine 
WORKDIR /app
ADD . /app
RUN go build -o chatgo .

CMD ["./chatgo"]