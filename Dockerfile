FROM golang:1.6-onbuild

RUN go vet
RUN go test

# CMD "./go/bin/app -location /home 100"
