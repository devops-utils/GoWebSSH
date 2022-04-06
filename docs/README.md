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

sqlite3 .GoWebSSH/GoWebSSH.db
sqlite3 docker/data/.GoWebSSH/GoWebSSH.db
```

```shell
https://nodejs.org/zh-cn/docs/guides/nodejs-docker-webapp/
https://www.cnblogs.com/super-lulu/p/11741591.html
http://nginx.org/en/docs/http/ngx_http_auth_request_module.html
https://www.beekeeperstudio.io/db/sqlite-client/
```

```shell
.backup ?DB? FILE      Backup DB (default "main") to FILE
.bail ON|OFF           Stop after hitting an error.  Default OFF
.databases             List names and files of attached databases
.dump ?TABLE? ...      Dump the database in an SQL text format
                         If TABLE specified, only dump tables matching
                         LIKE pattern TABLE.
.echo ON|OFF           Turn command echo on or off
.exit                  Exit this program
.explain ?ON|OFF?      Turn output mode suitable for EXPLAIN on or off.
                         With no args, it turns EXPLAIN on.
.header(s) ON|OFF      Turn display of headers on or off
.help                  Show this message
.import FILE TABLE     Import data from FILE into TABLE
.indices ?TABLE?       Show names of all indices
                         If TABLE specified, only show indices for tables
                         matching LIKE pattern TABLE.
.load FILE ?ENTRY?     Load an extension library
.log FILE|off          Turn logging on or off.  FILE can be stderr/stdout
.mode MODE ?TABLE?     Set output mode where MODE is one of:
                         csv      Comma-separated values
                         column   Left-aligned columns.  (See .width)
                         html     HTML <table> code
                         insert   SQL insert statements for TABLE
                         line     One value per line
                         list     Values delimited by .separator string
                         tabs     Tab-separated values
                         tcl      TCL list elements
.nullvalue STRING      Use STRING in place of NULL values
.output FILENAME       Send output to FILENAME
.output stdout         Send output to the screen
.print STRING...       Print literal STRING
.prompt MAIN CONTINUE  Replace the standard prompts
.quit                  Exit this program
.read FILENAME         Execute SQL in FILENAME
.restore ?DB? FILE     Restore content of DB (default "main") from FILE
.schema ?TABLE?        Show the CREATE statements
                         If TABLE specified, only show tables matching
                         LIKE pattern TABLE.
.separator STRING      Change separator used by output mode and .import
.show                  Show the current values for various settings
.stats ON|OFF          Turn stats on or off
.tables ?TABLE?        List names of tables
                         If TABLE specified, only list tables matching
                         LIKE pattern TABLE.
.timeout MS            Try opening locked tables for MS milliseconds
.trace FILE|off        Output each SQL statement as it is run
.vfsname ?AUX?         Print the name of the VFS stack
.width NUM1 NUM2 ...   Set column widths for "column" mode
.timer ON|OFF          Turn the CPU timer measurement on or off
```