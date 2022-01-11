# go-sftp
Study of SFTP on Go

## Reference
* [atmoz/sftp](https://github.com/atmoz/sftp) 
* [github.com/pkg/sftp](https://pkg.go.dev/github.com/pkg/sftp)
* [golang.org/x/crypto/ssh](https://pkg.go.dev/golang.org/x/crypto/ssh)

## Install
@local machine
```
mkdir sftp-server
mkdir sftp-server/upload
docker-compose up -d --build
sftp -oPort="2222" foo@localhost
```
If you get an following error when you run the sftp command, remove [localhost]:2222 from known_hosts.
```
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@    WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!     @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
```
```
ssh-keygen -R [localhost]:2222
```

@ sftp container
```
sftp> cd upload
sftp> put test
sftp> exit
```

@local machine
```
ls sftp-server/upload
```
