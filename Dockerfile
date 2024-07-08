FROM golang
WORKDIR go-app/
COPY . .
RUN go mod download
RUN go build -o /nexus_server
EXPOSE "8080"
CMD ["/nexus_server"]