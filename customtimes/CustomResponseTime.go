package customtimes

import(
	"time"
)

type CustomResponseTime struct {
	time.Time
}

const customResponseTimeFormat = "2006-01-02T15:04:05"

func (ct *CustomResponseTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if string(b) == "null" {
		return nil
	}
	ct.Time, err = time.Parse(customResponseTimeFormat, string(b))
	return
}

func (ct *CustomResponseTime) MarshalJSON() ([]byte, error) {
	if(ct == nil){
		return []byte("null"), nil
	}
	return []byte(ct.Time.Format(customResponseTimeFormat)), nil
}