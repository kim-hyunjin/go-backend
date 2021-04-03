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
