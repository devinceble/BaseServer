#BaseServerStruct

Version: *** v0.0.1b ***

Base Server Structure for Client Web Applications

My personal Starter Kit for Building Client Server for Web Application and PhaserJS Game

--------
###Change Log
*** v0.0.1b Beta Release ***

  * Routes
  * Models
  * Migrations
  * Tests
  * Sockets

--------

### Building Dependencies and Structure

####Dependencies
```
go 1.3.x
```

####Install Build Dependencies

```bash
#Create Database
  eg.
    MYSQL
    $ mysql -u root -p
    mysql> CREATE SCHEMA databasename;

#Configure Server
  mv Configs/Conf.go.example Configs/Conf.go
  //Change Database Connection and Server Settings
  vim Configs/Conf.go

#Initialize Server
$ go get
$ BaseServer migrate

#Running Debug Server
$ BaseServer start <PORT>

#Running Production Server
$ foreman start

#Deploying Server
$ foreman export FORMAT LOCATION

#Don't Forget to Create test dbase
#and change config at Configs/Conf.go to the test dbase name
#Testing Server
$ go get github.com/smartystreets/goconvey
$ goconvey
OR
$ go test -v ./...

```

#####Access Browser
http://localhost:8000/
