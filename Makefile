
docker build -t localhost.io:5005/go-person . \

docker image ls | grep localhost.io:5005/go-person \

docker tag 9c99de864904 localhost.io/go-person \
docker push localhost.io:5005/go-person