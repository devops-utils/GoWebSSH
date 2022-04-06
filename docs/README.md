```shell
cd webssh
cnpm install
npm run build

sudo docker run -ti --privileged --volume="$(pwd)":/webssh -v $(pwd)/root:/root --rm node:16 bash
# sudo docker run -ti --privileged --volume="$(pwd)":/webssh -v $(pwd)/root:/root --rm node:14 bash
cd /webssh
# npm install -g cnpm --registry=https://registry.npm.taobao.org
# cnpm install
npm install
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
https://www.cnblogs.com/super-lulu/p/11741591.html
http://nginx.org/en/docs/http/ngx_http_auth_request_module.html
```