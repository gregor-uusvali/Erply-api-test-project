docker build -t erply-api-test-project -f Dockerfile .
docker container run -p 8080:8080 erply-api-test-project