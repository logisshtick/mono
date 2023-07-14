package auth

import (
	"errors"
	"github.com/cespare/xxhash"
	"time"

	"github.com/logisshtick/mono/internal/constant"
	"github.com/logisshtick/mono/pkg/cryptograph"
	"github.com/logisshtick/mono/pkg/mu"
)

var (
	ErrAccessTNotExpired               = errors.New("Access token is not expired.")
	ErrRefreshTExpired                 = errors.New("Refresh token is expired.")
	ErrRefreshTNotFound                = errors.New("Refresh token doesnt found.")
	ErrAccessTNotFound                 = errors.New("Access token doesnt found.")
	ErrAccessTExpired                  = errors.New("Access token expired.")
	ErrDeviceIdLenIsBiggerThanExpected = errors.New("Device id len is bigger than expected.")
)

// gen new hash by device id string
func HashDeviceId(deviceId string) (uint64, error) {
	if len(deviceId) > constant.C.MaxDeviceIdLen {
		return 0, ErrDeviceIdLenIsBiggerThanExpected
	}
	// i really dont trust what frontenders send for me)
	return xxhash.Sum64String(deviceId), nil
}

// check is access token correct and not expired
func ValidateAccessToken(tkn uint64) (bool, error) {
	timeNow := time.Now().Unix()

	var (
		t  token
		ok bool
	)
	mu.ExecRWMutex(&maps.accessTsRmu, func() {
		t, ok = maps.accessTs[tkn]
	})
	if !ok {
		return false, ErrAccessTNotFound
	}
	if timeNow-t.date >= accessTLife {
		return false, ErrAccessTExpired
	}
	return true, nil
}

// generate new tokens pair by user account id and device id
func GenTokensPair(id int, deviceId uint64) (uint64, string, error) {
	timeNow := time.Now().Unix()

	hash, err := cryptograph.GenRandHash() // access token
	if err != nil {
		return 0, "", err
	}
	uid, err := cryptograph.GenRandUuid() // refresh token (realization may be changed)
	if err != nil {
		return 0, "", err
	}

	mu.ExecMutex(&maps.accessTsRmu, func() {
		maps.accessTs[hash] = token{
			date:     timeNow,
			id:       id,
			deviceId: deviceId,
		}
	})

	mu.ExecMutex(&maps.refreshTsRmu, func() {
		maps.refreshTs[uid] = token{
			date:     timeNow,
			id:       id,
			deviceId: deviceId,
		}
	})

	return hash, uid, nil
}

// func that generates new tokens pair
// and remove expired and useless tokens
// from db and hashmap cache
func RegenTokensPair(access uint64, refresh string) (uint64, string, error) {
	timeNow := time.Now().Unix()

	var (
		t  token
		ok bool
	)
	mu.ExecRWMutex(&maps.accessTsRmu, func() {
		t, ok = maps.accessTs[access]
	})
	if ok {
		if timeNow-t.date < accessTLife {
			return 0, "", ErrAccessTNotExpired
		}
		mu.ExecMutex(&maps.accessTsRmu, func() {
			delete(maps.accessTs, access)
		})
	}

	mu.ExecRWMutex(&maps.refreshTsRmu, func() {
		t, ok = maps.refreshTs[refresh]
	})
	if !ok {
		return 0, "", ErrRefreshTNotFound
	}
	if timeNow-t.date > refreshTLife {
		return 0, "", ErrRefreshTExpired
	}

	mu.ExecMutex(&maps.refreshTsRmu, func() {
		delete(maps.refreshTs, refresh)
	})

	return GenTokensPair(t.id, t.deviceId)
}
