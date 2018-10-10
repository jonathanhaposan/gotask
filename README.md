# Go Entry Task

### Installation

This app requires [Golang](https://golang.org/dl/)

Copy all data to
```
$GOPATH/src/github.com/jonathanhaposan/
```

Run SQL ```user.sql``

Run ```go get ./...``` inside ```$GOPATH/src/github.com/jonathanhaposan/```

Run these 2 binaries
```sh
$ cd $GOPATH/src/github.com/jonathanhaposan/gotask/cmd/web && go build && ./web

$ cd $GOPATH/src/github.com/jonathanhaposan/gotask/cmd/tcp   && go build && ./tcp
```

### Documentation

**List of endpoint**

| **Endpoints** | **Method** | **Description** |
|-----------------|:------------:|-------------------|
|`/login`| `GET`   | Show login page |
|`/login`| `POST`  | Handle login user Accept form-data : username & password|
|`/profile`| `GET` | Show profile page, require cookie |
|`/profile`| `POST`| Handle to update nickname and picture, accept multipart : nickname & picture, require cookie |


### Performance

**GET /login**
```
    $ bombardier -c 125 -n 100000 -m GET http://localhost:9000/login
    Bombarding http://localhost:9000/login with 100000 request(s) using 125 connection(s)
    100000 / 100000 [=========================================================================================] 100.00% 14s
    Done!
    Statistics        Avg      Stdev        Max
    Reqs/sec      7011.23    1173.75    8662.56
    Latency       17.85ms    25.36ms      1.17s
    HTTP codes:
        1xx - 0, 2xx - 100000, 3xx - 0, 4xx - 0, 5xx - 0
        others - 0
    Throughput:    13.42MB/s
```

**GET /profile**
```
    $ bombardier -c 125 -n 100000 -m GET --header="Cookie: session_cookie=7dccca3f-5ee6-48d6-856f-2e03bdb071a6" http://localhost:9000/profile
    Bombarding http://localhost:9000/profile with 100000 request(s) using 125 connection(s)
    100000 / 100000 [=======================================================================================] 100.00% 1m30s
    Done!
    Statistics        Avg      Stdev        Max
    Reqs/sec      1112.36     395.90    2908.36
    Latency      112.27ms    37.30ms      2.02s
    HTTP codes:
        1xx - 0, 2xx - 100000, 3xx - 0, 4xx - 0, 5xx - 0
        others - 0
    Throughput:     2.67MB/s
```

**POST /login**
```
    $ bombardier -c 125 -n 8000 -m POST --header="Content-Type: application/x-www-form-urlencoded" -f cred_1_login.txt http://localhost:9000/login
    Bombarding http://localhost:9000/login with 8000 request(s) using 125 connection(s)
    8000 / 8000 [==============================================================================================] 100.00% 6sDone!
    Statistics        Avg      Stdev        Max
    Reqs/sec      1197.79     336.90    3140.70
    Latency      104.43ms    17.98ms   275.67ms
    HTTP codes:
        1xx - 0, 2xx - 8000, 3xx - 0, 4xx - 0, 5xx - 0
        others - 0
    Throughput:   505.87KB/s
```

**POST /profile**
```
    $ bombardier -c 125 -n 8000 -m POST --header="Cookie: session_cookie=7dccca3f-5ee6-48d6-856f-2e03bdb071a6" 
      --header="Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW" -f data_1_profile.txt http://localhost:9000/profile
    Bombarding http://localhost:9000/profile with 8000 request(s) using 125 connection(s)
    8000 / 8000 [==========================================================================================================================] 100.00% 14s
    Done!
    Statistics        Avg      Stdev        Max
    Reqs/sec       571.23     309.17    3126.50
    Latency      219.67ms    42.51ms      2.21s
    HTTP codes:
        1xx - 0, 2xx - 8000, 3xx - 0, 4xx - 0, 5xx - 0
        others - 0
    Throughput:   356.65KB/s
```