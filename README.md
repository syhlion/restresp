# rest response


## Install

`go get -u github.com/syhlion/restresp`

## Usege

```
type customError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (c *customError) Error() string {
	return c.Msg
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	d := &customError{5, "wtf"}
	restresp.Write(w, d, 200)
}
```

