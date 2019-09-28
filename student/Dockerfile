FROM golang:1.12

# RUN mkdir /go/src/github.com
# RUN mkdir /go/src/github.com/MohitArora1
WORKDIR /go/src/github.com/MohitArora1/devops/student
COPY . .
#WORKDIR /go/src/github.com/MohitArora1/devops/student

RUN go get -v ./...
EXPOSE 8080

CMD ["student"]
