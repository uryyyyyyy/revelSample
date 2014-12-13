revelSample
===========



## setup

use OS ubuntu14.04

*you need mercurial*


```
VERSION=1.4
OS=linux
ARCH=amd64

wget https://storage.googleapis.com/golang/go$VERSION.$OS-$ARCH.tar.gz
sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

add path

```
# mey be this path is recommended
export GOROOT=/usr/local/go

# you can choice any path for revel-app
export GOPATH=$HOME/revel

# add PATH (why need GOPATH? I dont know)
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```


install revel

```
# to use revel source and samples
go get github.com/revel/revel

# to use revel command
go get github.com/revel/cmd/revel
```

run sample app

```
cd $GOPATH/src
revel run github.com/revel/revel/samples/chat
```

if sample app is running at localhost:9000,

you already install revel!!!


## system
