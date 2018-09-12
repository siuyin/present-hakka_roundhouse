# Hakka Roundhouse -- a microservices ready monolith

Monoliths are simple in go.

1. One static binary to deploy
1. No network latency
1. No network packet loss etc.

What if we could build a monolith which can _easily_ 
be split into independently running microservices
_if and when_ neeeded?

The Hakka Roundhouse monolith tries to do this...

## If running in a docker container
1. docker run -it --name present -v godata:/home/siuyin/go -p 3999:3999 siuyin/go:dev
1. setup .bashrc to have ~/go/bin in PATH or export PATH=~/go/bin:$PATH
1. present -http 0.0.0.0:3999 -orighost 192.168.99.100
