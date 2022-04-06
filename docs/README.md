```shell
cd webssh
cnpm install
npm run build

sudo docker run -ti --volume="$(pwd)":/webssh --rm node:16 bash
sudo docker run -ti --volume="$(pwd)":/webssh --rm node:14 bash
cd /webssh
npm install -g cnpm --registry=https://registry.npm.taobao.org
cnpm install
npm run build

cd gossh
go build
# ./gossh

cd ../docker
cp ../gossh/gossh webssh

sudo docker-compose build webssh

sudo docker-compose up webssh
sudo docker-compose up -d webssh
sudo docker-compose stop webssh
sudo docker-compose rm webssh

sudo docker exec -it webssh bash
sudo docker logs -f webssh
```

```shell
https://nodejs.org/zh-cn/docs/guides/nodejs-docker-webapp/
```