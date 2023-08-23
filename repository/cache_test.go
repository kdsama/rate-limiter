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
	var (
		limiter = NewLimiterCache()
		key     = "kdsite/xyz"
	)
	type inpS struct {
		name string
		url  string
		want error
	}

	testcases := []inpS{
		{
			name: "For keys that are not present",
			url:  "www.kdsite/xyzxas",
			want: ErrKeyNotFound,
		},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			_, got := limiter.Get(test.url)
			if got != test.want {
				t.Errorf("Wanted ::%v:: but got ::%v::", test.want, got)
			}
		})
	}

	err := limiter.Set(key, entity.NewLimiter("longerURLForXyz", "something", "kdsite/xyz", utils.OneWeekFromNow()))
	if err != nil {
		t.Errorf("Did not expect any error but got %v", err)
	}
	obj, err := limiter.Get(key)
	if err != nil {
		t.Errorf("Did not expect any error but got %v", err)
	}
	if obj.ShortUrl != key {
		t.Errorf("Wanted %v but got %v ", key, obj.ShortUrl)
	}
}
