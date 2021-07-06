docker build -t dummy-go2:0.0.1 .

docker tag dummygo2:0.0.1 dmltjq2524/dummy-go2:0.0.1

docker push dmltjq2524/dummy-go2:0.0.1

docker rmi -f [image id]

docker rm [container id]

docker container ls -al $ docker container ls -all
vs
docker container ls --all 


docker run -dp 4000:9998 dmltjq2524/dummy-go2:0.0.1
192.168.0.58:4000
or
http://192.168.0.58:4000

