```shell
cd gossh
go build
# ./gossh

cd ../docker
cp ../gossh/gossh webssh

sudo docker-compose up webssh
sudo docker-compose up -d webssh
sudo docker-compose stop webssh
sudo docker-compose rm webssh
```