```
type User struct {
	FirstName string	`json:"first_name"`
	LastName string		`json:"last_name"`
	Email string		`json:"email"`
	CreatedAt time.Time	`json:"created_at"`
}
```

구조체에 `json:"first_name"` 이런식으로 태그를 달아두면 구조체의 변수와 json의 key가 달라서 발생하는 문제를 해결할 수 있다.

## 테스트 컨벤션

- 파일명은 끝에 \_test를 붙인다.
- 메소드명은 Test로 시작한다.
- 아규먼트로 \*testing.T 를 받는다.

### 원활한 테스트를 위한 라이브러리

- github.com/smartystreets/goconvey : 파일이 변경될때마다 자동을 테스트를 수행해준다.
- github.com/stretchr/testify : assert 패키지의 메소드를 사용하면 쉽게 테스트할 수 있다.

## Web socket, Event source

https://developer.mozilla.org/ko/docs/Web/API/WebSockets_API/Writing_WebSocket_client_applications

https://developer.mozilla.org/ko/docs/Web/API/EventSource

# 3 tier web

Front - Back - DB

# DB 사용하기

- SQLite3

```
go get github.com/mattn/go-sqlite3
```

- 위 라이브러리는 cgo 패키지다. 그래서 c 표준 컴파일러가 필요하다. sqllite3가 c로 만들어져 있기 때문.
- 윈도우의 경우 https://jmeubank.github.io/tdm-gcc/ 여기서 gcc를 다운받으면 된다.
