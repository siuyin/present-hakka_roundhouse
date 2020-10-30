# Hakka Roundhouse -- a microservices ready monolith

## Biggest challenge in software engineering
## Architectures: Planned City, Slum, Big Ball of Mud, Hakka Roundhouse
## Software Architectures: Old cities adapt or die, planned cities rarely turn out as planned.
## Communications: Local procedure call, remote procedure call, messaging

Monoliths are simple in go.

1. One static binary to deploy
1. No network latency
1. No network packet loss etc.

What if we could build a monolith which can _easily_ 
be split into independently running microservices
_if and when_ neeeded?

The Hakka Roundhouse monolith tries to do this...

```
docker run -it --name present -v godata:/home/siuyin/go -p 3999:3999 siuyin/go:dev
go get golang.org/x/tools/cmd/present
setup .bashrc to have ~/go/bin in PATH or export PATH=~/go/bin:$PATH

present -http 0.0.0.0:3999 -orighost {docker_host_IP}
eg1. present -http 0.0.0.0:3999 -orighost 192.168.99.100
eg2. present -http 0.0.0.0:3999 -orighost 127.0.0.1
```
