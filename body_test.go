package fibererror_test

import (
	"github.com/prongbang/fibererror"
	"testing"
)

type customErr struct {
	fibererror.Body
}

func (u *customErr) Error() string {
	return u.Message
}

func newCustomErr() error {
	return &customErr{
		fibererror.Body{
			Code:    "XXX",
			Message: "YYY",
			Data:    nil,
		},
	}
}

func BenchmarkGetBody(b *testing.B) {
	err := newCustomErr()

	for i := 0; i < b.N; i++ {
		_, _ = fibererror.GetBody(err)
	}
}

func TestGetBody(t *testing.T) {
	err := newCustomErr()

	actual, _ := fibererror.GetBody(err)

	if actual.Code != "XXX" || actual.Message != "YYY" {
		t.Error("Error:", actual)
	}
}
