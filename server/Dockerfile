FROM golang:1.21

ENV APP_HOME /usr/src/app
WORKDIR $APP_HOME

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/jackc/tern/v2@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY entrypoint.sh $APP_HOME

ADD wait-for-it.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/wait-for-it.sh

CMD ["wait-for-it.sh", "db:5432", "--strict", "--", "./entrypoint.sh"]
