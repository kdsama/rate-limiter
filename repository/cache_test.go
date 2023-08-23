package repository

import (
	"testing"

	"github.com/kdsama/rate-limiter/entity"
	"github.com/kdsama/rate-limiter/utils"
)

func TestLimiterCacheSet(t *testing.T) {
	var (
		limiter = NewLimiterCache()
	)
	type inpS struct {
		name string
		url  string
		obj  *entity.Limiter
		want error
	}
	testcases := []inpS{
		{
			name: "Valid object value",
			url:  "kdsite/xyz",
			obj:  entity.NewLimiter("longerURLForXyz", "something", "kdsite/xyz", utils.OneWeekFromNow()),
			want: nil,
		},
		{
			name: "invalid object value",
			url:  "kdsite/xyz",
			obj:  nil,
			want: ErrLimiterIsNil,
		},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			got := limiter.Set(test.url, test.obj)
			if got != test.want {
				t.Errorf("Wanted ::%v:: but got ::%v::", test.want, got)
			}
		})
	}

}

func TestLimiterCacheGet(t *testing.T) {
	type inpS struct {
		url  string
		obj  *entity.Limiter
		want error
	}

	_ = []inpS{
		{url: "www.kdsite/xyzxas"},
	}
}
