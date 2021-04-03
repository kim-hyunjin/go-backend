```
type User struct {
	FirstName string	`json:"first_name"`
	LastName string		`json:"last_name"`
	Email string		`json:"email"`
	CreatedAt time.Time	`json:"created_at"`
}
```

구조체에 `json:"first_name"` 이런식으로 태그를 달아두면 구조체의 변수와 json의 key가 달라서 발생하는 문제를 해결할 수 있다.
