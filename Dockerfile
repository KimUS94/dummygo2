FROM golang:1.16.5

RUN mkdir /echo

COPY dummy2.go /echo

EXPOSE 9998

CMD ["go", "run", "/echo/dummy2.go"]