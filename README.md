[![main](https://github.com/flowerinthenight/oomkill-trace/actions/workflows/main.yml/badge.svg)](https://github.com/flowerinthenight/oomkill-trace/actions/workflows/main.yml)

**NOTE: Unfinished, don't use. Check out [`oomkill-watch`](https://github.com/flowerinthenight/oomkill-watch) instead.**

```sh
# Build:
$ docker build --rm -t oomkill-trace .

# Run using docker:
$ docker run \
    -it \
    --rm \
    -v /lib/modules:/lib/modules:ro \
    -v /sys:/sys:ro \
    -v /usr/src:/usr/src:ro \
    --privileged oomkill-trace
```
