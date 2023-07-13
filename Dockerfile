FROM golang:latest
COPY . /app
WORKDIR /app
RUN go mod download
# RUN apt-get update && apt-get install -y sqlite3
RUN go build -o Erply-api-test-project .
EXPOSE 8080
CMD [ "/app/Erply-api-test-project" ]