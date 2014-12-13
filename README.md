revelSample
===========



## setup

use OS ubuntu14.04


```
VERSION=1.4
OS=linux
ARCH=amd64

wget https://storage.googleapis.com/golang/go$VERSION.$OS-$ARCH.tar.gz
sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

add path

```
export GOROOT=/usr/local/go
export GOPATH=$HOME/revel
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
