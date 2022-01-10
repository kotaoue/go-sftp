# go-sftp
Study of SFTP on Go

## Reference
* [【Docker】SFTPサーバコンテナ構築手順と使い方（atmoz/sftp）](https://genchan.net/it/virtualization/docker/13643/)
* [atmoz/sftp](https://github.com/atmoz/sftp) 

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
