# Verze A pro úlohu 05
Dosadil jsem vlastní importy, ale gRPC container mi při startu končí na:
**standard_init_linux.go:228: exec user process caused: no such file or directory** 
Pokoušel jsem o Debug přes úpravy v Dockerfilu i spouštění containeru samostatně, ale na řešení jsem nepřišel.



# Etcd and GRPC

### Setup

1. Github [repository](https://github.com/etcd-io/etcd)
2. Download 3.5.1 from etcd releases [page](https://github.com/etcd-io/etcd/releases)
    * [Windows](https://github.com/etcd-io/etcd/releases/download/v3.5.1/etcd-v3.5.1-windows-amd64.zip)
    * [Linux](https://github.com/etcd-io/etcd/releases/download/v3.5.1/etcd-v3.5.1-linux-amd64.tar.gz)
3. Add `etcdctl.exe` to your PATH


### Assignment

1. Create GRPC server and client application
2. GRPC server communicates with Etcd server using go etcd library
3. GRPC client communicates with GRPC server
4. GRPC server API should support at least GET/POST/DELETE operations defined using Protobuf
5. Bundle GRPC server together with Etcd as docker compose file