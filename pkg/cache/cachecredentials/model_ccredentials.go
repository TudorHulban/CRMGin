package cachecredentials

import (
	"os"

	"github.com/TudorHulban/GinCRM/pkg/cache"
	"github.com/TudorHulban/GinCRM/pkg/ostop"
	"github.com/TudorHulban/badgerwrap"
	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

var cacheBadger *badgerwrap.BStore

// GetCache Returns session ID cache object.
// Using varidic for the cases where logger for sure was already created.
func GetCache(l ...*log.LogInfo) (cache.IKV, error) {
	if cacheBadger == nil {
		if len(l) == 0 {
			return nil, errors.New("cannot create cache without logger. please pass logger")
		}

		var errCo error
		cacheBadger, errCo = badgerwrap.NewBStoreInMem(l[0])
		if errCo != nil {
			l[0].Infof("error trying to create credentials cache:%v", errCo)
			os.Exit(ostop.CACHE)
		}
	}

	return cacheBadger, nil
}
