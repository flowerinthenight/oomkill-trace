```
$ docker run \
    -it \
    --rm \
    -v /lib/modules:/lib/modules:ro \
    -v /sys:/sys:ro \
    -v /usr/src:/usr/src:ro \
    --privileged oomkill-trace
```
