package command

import (
	"net"
	"strings"
)

// CmdCommand 获取所有的command
const CmdCommand = "command"

func init() {
	Register(CmdCommand, Command)
}

const respStr = `*157
*6
$9
randomkey
:1
*2
+readonly
+random
:0
:0
:0
*6
$7
persist
:2
*2
+write
+fast
:1
:1
:1
*6
$6
bitpos
:-3
*1
+readonly
:1
:1
:1
*6
$10
psubscribe
:-2
*5
+readonly
+pubsub
+noscript
+loading
+stale
:0
:0
:0
*6
$6
config
:-2
*3
+readonly
+admin
+stale
:0
:0
:0
*6
$10
sdiffstore
:-3
*2
+write
+denyoom
:1
:-1
:1
*6
$6
lrange
:4
*1
+readonly
:1
:1
:1
*6
$6
hsetnx
:4
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$4
keys
:2
*2
+readonly
+sort_for_script
:0
:0
:0
*6
$5
blpop
:-3
*2
+write
+noscript
:1
:-2
:1
*6
$4
hdel
:-3
*2
+write
+fast
:1
:1
:1
*6
$6
client
:-2
*2
+readonly
+noscript
:0
:0
:0
*6
$4
echo
:2
*2
+readonly
+fast
:0
:0
:0
*6
$7
zincrby
:4
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$7
hgetall
:2
*1
+readonly
:1
:1
:1
*6
$6
lpushx
:3
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$4
pttl
:2
*2
+readonly
+fast
:1
:1
:1
*6
$12
hincrbyfloat
:4
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$4
hlen
:2
*2
+readonly
+fast
:1
:1
:1
*6
$9
sismember
:3
*2
+readonly
+fast
:1
:1
:1
*6
$7
flushdb
:1
*1
+write
:0
:0
:0
*6
$11
sunionstore
:-3
*2
+write
+denyoom
:1
:-1
:1
*6
$11
zrangebylex
:-4
*1
+readonly
:1
:1
:1
*6
$4
info
:-1
*3
+readonly
+loading
+stale
:0
:0
:0
*6
$4
lrem
:4
*1
+write
:1
:1
:1
*6
$6
sinter
:-2
*2
+readonly
+sort_for_script
:1
:-1
:1
*6
$5
sscan
:-3
*2
+readonly
+random
:1
:1
:1
*6
$6
strlen
:2
*2
+readonly
+fast
:1
:1
:1
*6
$8
shutdown
:-1
*4
+readonly
+admin
+loading
+stale
:0
:0
:0
*6
$6
msetnx
:-3
*2
+write
+denyoom
:1
:-1
:2
*6
$4
rpop
:2
*2
+write
+fast
:1
:1
:1
*6
$8
expireat
:3
*2
+write
+fast
:1
:1
:1
*6
$11
sinterstore
:-3
*2
+write
+denyoom
:1
:-1
:1
*6
$5
hkeys
:2
*2
+readonly
+sort_for_script
:1
:1
:1
*6
$7
evalsha
:-3
*2
+noscript
+movablekeys
:0
:0
:0
*6
$11
unsubscribe
:-1
*5
+readonly
+pubsub
+noscript
+loading
+stale
:0
:0
:0
*6
$8
getrange
:4
*1
+readonly
:1
:1
:1
*6
$5
zcard
:2
*2
+readonly
+fast
:1
:1
:1
*6
$6
script
:-2
*2
+readonly
+noscript
:0
:0
:0
*6
$7
publish
:3
*5
+readonly
+pubsub
+loading
+stale
+fast
:0
:0
:0
*6
$8
replconf
:-1
*5
+readonly
+admin
+noscript
+loading
+stale
:0
:0
:0
*6
$4
sadd
:-3
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$6
select
:2
*3
+readonly
+loading
+fast
:0
:0
:0
*6
$7
linsert
:5
*2
+write
+denyoom
:1
:1
:1
*6
$16
zremrangebyscore
:4
*1
+write
:1
:1
:1
*6
$4
type
:2
*2
+readonly
+fast
:1
:1
:1
*6
$6
zcount
:4
*2
+readonly
+fast
:1
:1
:1
*6
$6
substr
:4
*1
+readonly
:1
:1
:1
*6
$5
brpop
:-3
*2
+write
+noscript
:1
:1
:1
*6
$6
incrby
:3
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$8
bitcount
:-2
*1
+readonly
:1
:1
:1
*6
$7
migrate
:6
*1
+write
:0
:0
:0
*6
$6
setbit
:4
*2
+write
+denyoom
:1
:1
:1
*6
$5
lpush
:-3
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$9
rpoplpush
:3
*2
+write
+denyoom
:1
:2
:1
*6
$7
latency
:-2
*5
+readonly
+admin
+noscript
+loading
+stale
:0
:0
:0
*6
$5
rpush
:-3
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$9
pexpireat
:3
*2
+write
+fast
:1
:1
:1
*6
$8
zrevrank
:3
*2
+readonly
+fast
:1
:1
:1
*6
$9
zrevrange
:-4
*1
+readonly
:1
:1
:1
*6
$12
punsubscribe
:-1
*5
+readonly
+pubsub
+noscript
+loading
+stale
:0
:0
:0
*6
$6
append
:3
*2
+write
+denyoom
:1
:1
:1
*6
$4
lset
:4
*2
+write
+denyoom
:1
:1
:1
*6
$4
zrem
:-3
*2
+write
+fast
:1
:1
:1
*6
$8
smembers
:2
*2
+readonly
+sort_for_script
:1
:1
:1
*6
$8
lastsave
:1
*3
+readonly
+random
+fast
:0
:0
:0
*6
$4
scan
:-2
*2
+readonly
+random
:0
:0
:0
*6
$5
zrank
:3
*2
+readonly
+fast
:1
:1
:1
*6
$7
pfmerge
:-2
*2
+write
+denyoom
:1
:-1
:1
*6
$7
monitor
:1
*3
+readonly
+admin
+noscript
:0
:0
:0
*6
$6
expire
:3
*2
+write
+fast
:1
:1
:1
*6
$4
ping
:-1
*3
+readonly
+stale
+fast
:0
:0
:0
*6
$14
zremrangebylex
:4
*1
+write
:1
:1
:1
*6
$7
hincrby
:4
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$11
srandmember
:-2
*2
+readonly
+random
:1
:1
:1
*6
$6
pubsub
:-2
*5
+readonly
+pubsub
+random
+loading
+stale
:0
:0
:0
*6
$15
zremrangebyrank
:4
*1
+write
:1
:1
:1
*6
$4
role
:1
*3
+noscript
+loading
+stale
:0
:0
:0
*6
$6
object
:3
*1
+readonly
:2
:2
:2
*6
$4
decr
:2
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$5
pfadd
:-2
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$10
pfselftest
:1
*1
+readonly
:0
:0
:0
*6
$4
spop
:2
*4
+write
+noscript
+random
+fast
:1
:1
:1
*6
$5
debug
:-2
*2
+admin
+noscript
:0
:0
:0
*6
$5
smove
:4
*2
+write
+fast
:1
:2
:1
*6
$4
llen
:2
*2
+readonly
+fast
:1
:1
:1
*6
$5
multi
:1
*3
+readonly
+noscript
+fast
:0
:0
:0
*6
$5
sdiff
:-2
*2
+readonly
+sort_for_script
:1
:-1
:1
*6
$6
getset
:3
*2
+write
+denyoom
:1
:1
:1
*6
$5
hscan
:-3
*2
+readonly
+random
:1
:1
:1
*6
$4
save
:1
*3
+readonly
+admin
+noscript
:0
:0
:0
*6
$7
slaveof
:3
*3
+admin
+noscript
+stale
:0
:0
:0
*6
$4
auth
:2
*5
+readonly
+noscript
+loading
+stale
+fast
:0
:0
:0
*6
$6
rename
:3
*1
+write
:1
:2
:1
*6
$6
bgsave
:1
*2
+readonly
+admin
:0
:0
:0
*6
$6
decrby
:3
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$7
discard
:1
*3
+readonly
+noscript
+fast
:0
:0
:0
*6
$6
sunion
:-2
*2
+readonly
+sort_for_script
:1
:-1
:1
*6
$7
pexpire
:3
*2
+write
+fast
:1
:1
:1
*6
$4
sync
:1
*3
+readonly
+admin
+noscript
:0
:0
:0
*6
$7
pfdebug
:-3
*1
+write
:0
:0
:0
*6
$5
hvals
:2
*2
+readonly
+sort_for_script
:1
:1
:1
*6
$5
zscan
:-3
*2
+readonly
+random
:1
:1
:1
*6
$3
get
:2
*2
+readonly
+fast
:1
:1
:1
*6
$6
exists
:2
*2
+readonly
+fast
:1
:1
:1
*6
$6
lindex
:3
*1
+readonly
:1
:1
:1
*6
$7
restore
:4
*2
+write
+denyoom
:1
:1
:1
*6
$4
sort
:-2
*2
+write
+denyoom
:1
:1
:1
*6
$5
setex
:4
*2
+write
+denyoom
:1
:1
:1
*6
$4
incr
:2
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$3
set
:-3
*2
+write
+denyoom
:1
:1
:1
*6
$5
scard
:2
*2
+readonly
+fast
:1
:1
:1
*6
$4
mget
:-2
*1
+readonly
:1
:-1
:1
*6
$10
brpoplpush
:4
*3
+write
+denyoom
+noscript
:1
:2
:1
*6
$6
zscore
:3
*2
+readonly
+fast
:1
:1
:1
*6
$4
srem
:-3
*2
+write
+fast
:1
:1
:1
*6
$14
zrevrangebylex
:-4
*1
+readonly
:1
:1
:1
*6
$8
setrange
:4
*2
+write
+denyoom
:1
:1
:1
*6
$4
mset
:-3
*2
+write
+denyoom
:1
:-1
:2
*6
$7
unwatch
:1
*3
+readonly
+noscript
+fast
:0
:0
:0
*6
$8
flushall
:1
*1
+write
:0
:0
:0
*6
$8
renamenx
:3
*2
+write
+fast
:1
:2
:1
*6
$6
getbit
:3
*2
+readonly
+fast
:1
:1
:1
*6
$12
bgrewriteaof
:1
*2
+readonly
+admin
:0
:0
:0
*6
$9
subscribe
:-2
*5
+readonly
+pubsub
+noscript
+loading
+stale
:0
:0
:0
*6
$6
zrange
:-4
*1
+readonly
:1
:1
:1
*6
$7
slowlog
:-2
*1
+readonly
:0
:0
:0
*6
$4
hget
:3
*2
+readonly
+fast
:1
:1
:1
*6
$5
ltrim
:4
*1
+write
:1
:1
:1
*6
$7
hexists
:3
*2
+readonly
+fast
:1
:1
:1
*6
$6
rpushx
:3
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$16
zrevrangebyscore
:-4
*1
+readonly
:1
:1
:1
*6
$9
zlexcount
:4
*2
+readonly
+fast
:1
:1
:1
*6
$5
psync
:3
*3
+readonly
+admin
+noscript
:0
:0
:0
*6
$4
time
:1
*3
+readonly
+random
+fast
:0
:0
:0
*6
$11
zunionstore
:-4
*3
+write
+denyoom
+movablekeys
:0
:0
:0
*6
$5
setnx
:3
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$4
hset
:4
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$3
ttl
:2
*2
+readonly
+fast
:1
:1
:1
*6
$5
hmget
:-3
*1
+readonly
:1
:1
:1
*6
$3
del
:-2
*1
+write
:1
:-1
:1
*6
$4
dump
:2
*1
+readonly
:1
:1
:1
*6
$4
move
:3
*2
+write
+fast
:1
:1
:1
*6
$5
watch
:-2
*3
+readonly
+noscript
+fast
:1
:-1
:1
*6
$6
psetex
:4
*2
+write
+denyoom
:1
:1
:1
*6
$4
lpop
:2
*2
+write
+fast
:1
:1
:1
*6
$11
incrbyfloat
:3
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$13
zrangebyscore
:-4
*1
+readonly
:1
:1
:1
*6
$5
bitop
:-4
*2
+write
+denyoom
:2
:-1
:1
*6
$7
pfcount
:-2
*1
+readonly
:1
:1
:1
*6
$7
command
:0
*3
+readonly
+loading
+stale
:0
:0
:0
*6
$5
hmset
:-4
*2
+write
+denyoom
:1
:1
:1
*6
$6
dbsize
:1
*2
+readonly
+fast
:0
:0
:0
*6
$4
zadd
:-4
*3
+write
+denyoom
+fast
:1
:1
:1
*6
$4
exec
:1
*2
+noscript
+skip_monitor
:0
:0
:0
*6
$4
eval
:-3
*2
+noscript
+movablekeys
:0
:0
:0
*6
$11
zinterstore
:-4
*3
+write
+denyoom
+movablekeys
:0
:0
:0
`

// Command 获取所有的command
func Command(conn net.Conn, args ...[]byte) {
	lines := strings.Split(respStr, "\n")
	newRespStr := strings.Join(lines, "\r\n")
	conn.Write([]byte(newRespStr))
}
