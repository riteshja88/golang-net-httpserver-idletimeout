# golang-net-httpserver-idletimeout

## no timeout
```
server := &http.Server{Addr: ":3334"}

$ go run p1.go &
$ echo "<<<<start>>>>>"|ts '[%Y-%m-%d %H:%M:%S]';echo -n -e "GET / HTTP/1.1\nHost: 127.0.0.1:3334\n\n"|nc 127.0.0.1 3334| ts '[%Y-%m-%d %H:%M:%S]';echo "<<<<<end>>>>>"|ts '[%Y-%m-%d %H:%M:%S]'
[2022-09-29 13:37:46] <<<<start>>>>>
got / request
[2022-09-29 13:37:46] HTTP/1.1 200 OK
[2022-09-29 13:37:46] Date: Thu, 29 Sep 2022 08:07:46 GMT
[2022-09-29 13:37:46] Content-Length: 20
[2022-09-29 13:37:46] Content-Type: text/plain; charset=utf-8
[2022-09-29 13:37:46]
[2022-09-29 13:37:46] This is my website!
```


## Can lead to too many open idle connections
```
$ for ((i=0;i<4;i++));do echo "Starting infinitely idle connection #$i";echo -n -e
 "GET / HTTP/1.1\nHost: 127.0.0.1:3334\n\n"|nc 127.0.0.1 3334 & done;wait
Starting infinitely idle connection #0
[2] 12814
Starting infinitely idle connection #1
[3] 12816
Starting infinitely idle connection #2
got / request
[4] 12818
Starting infinitely idle connection #3
HTTP/1.1 200 OK
Date: Thu, 29 Sep 2022 08:18:52 GMT
Content-Length: 20
Content-Type: text/plain; charset=utf-8

This is my website!
got / request
HTTP/1.1 200 OK
Date: Thu, 29 Sep 2022 08:18:52 GMT
Content-Length: 20
Content-Type: text/plain; charset=utf-8

This is my website!
[5] 12820
got / request
HTTP/1.1 200 OK
Date: Thu, 29 Sep 2022 08:18:52 GMT
Content-Length: 20
Content-Type: text/plain; charset=utf-8

This is my website!
got / request
HTTP/1.1 200 OK
Date: Thu, 29 Sep 2022 08:18:52 GMT
Content-Length: 20
Content-Type: text/plain; charset=utf-8

This is my website!
```

```
$ netstat -tunp|grep p1
(Not all processes could be identified, non-owned process info
 will not be shown, you would have to be root to see it all.)
tcp6       0      0 127.0.0.1:3334          127.0.0.1:44812         ESTABLISHED 12805/p1
tcp6       0      0 127.0.0.1:3334          127.0.0.1:44814         ESTABLISHED 12805/p1
tcp6       0      0 127.0.0.1:3334          127.0.0.1:44810         ESTABLISHED 12805/p1
tcp6       0      0 127.0.0.1:3334          127.0.0.1:44808         ESTABLISHED 12805/p1
```

## timeout = 10s
```
server := &http.Server{Addr: ":3334", IdleTimeout: time.Duration(10) * time.Second}

$ go run p1.go &
$ echo "<<<<start>>>>>"|ts '[%Y-%m-%d %H:%M:%S]';echo -n -e "GET / HTTP/1.1\nHost: 127.0.0.1:3334\n\n"|nc 127.0.0.1 3334| ts '[%Y-%m-%d %H:%M:%S]';echo "<<<<<end>>>>>"|ts '[%Y-%m-%d %H:%M:%S]'
[2022-09-29 13:28:54] <<<<start>>>>>
got / request
[2022-09-29 13:28:54] HTTP/1.1 200 OK
[2022-09-29 13:28:54] Date: Thu, 29 Sep 2022 07:58:54 GMT
[2022-09-29 13:28:54] Content-Length: 20
[2022-09-29 13:28:54] Content-Type: text/plain; charset=utf-8
[2022-09-29 13:28:54]
[2022-09-29 13:28:54] This is my website!
[2022-09-29 13:29:04] <<<<<end>>>>>
```

## timeout = 20s
```
server := &http.Server{Addr: ":3334", IdleTimeout: time.Duration(20) * time.Second}

$ go run p1.go &
$ echo "<<<<start>>>>>"|ts '[%Y-%m-%d %H:%M:%S]';echo -n -e "GET / HTTP/1.1\nHost: 127.0.0.1:3334\n\n"|nc 127.0.0.1 3334| ts '[%Y-%m-%d %H:%M:%S]';echo "<<<<<end>>>>>"|ts '[%Y-%m-%d %H:%M:%S]'
[2022-09-29 13:29:29] <<<<start>>>>>
got / request
[2022-09-29 13:29:29] HTTP/1.1 200 OK
[2022-09-29 13:29:29] Date: Thu, 29 Sep 2022 07:59:29 GMT
[2022-09-29 13:29:29] Content-Length: 20
[2022-09-29 13:29:29] Content-Type: text/plain; charset=utf-8
[2022-09-29 13:29:29]
[2022-09-29 13:29:29] This is my website!
[2022-09-29 13:29:49] <<<<<end>>>>>
```

## timeout = 30s
```
server := &http.Server{Addr: ":3334", IdleTimeout: time.Duration(30) * time.Second}

$ go run p1.go &
$ echo "<<<<start>>>>>"|ts '[%Y-%m-%d %H:%M:%S]';echo -n -e "GET / HTTP/1.1\nHost: 127.0.0.1:3334\n\n"|nc 127.0.0.1 3334| ts '[%Y-%m-%d %H:%M:%S]';echo "<<<<<end>>>>>"|ts '[%Y-%m-%d %H:%M:%S]'
[2022-09-29 13:30:06] <<<<start>>>>>
got / request
[2022-09-29 13:30:06] HTTP/1.1 200 OK
[2022-09-29 13:30:06] Date: Thu, 29 Sep 2022 08:00:06 GMT
[2022-09-29 13:30:06] Content-Length: 20
[2022-09-29 13:30:06] Content-Type: text/plain; charset=utf-8
[2022-09-29 13:30:06]
[2022-09-29 13:30:06] This is my website!
[2022-09-29 13:30:36] <<<<<end>>>>>
```