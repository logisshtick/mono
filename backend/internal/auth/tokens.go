package auth

import (
	"errors"
	"github.com/cespare/xxhash"
	"time"
)

const (
	maxDeviceIdLen = 4096
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
	if len(deviceId) > maxDeviceIdLen {
		return 0, ErrDeviceIdLenIsBiggerThanExpected
	}
	// i really dont trust what frontenders send for me)
	return xxhash.Sum64String(deviceId), nil
}

// check is access token correct and not expired
func ValidateAccessToken(token uint64) (bool, error) {
	timeNow := time.Now().Unix()

	maps.accessTsRmu.RLock()
	t, ok := maps.accessTs[token]
	maps.accessTsRmu.RUnlock()
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

	hash, err := genHash() // access token
	if err != nil {
		return 0, "", err
	}
	uid, err := genUuid() // refresh token (realization may be changed)
	if err != nil {
		return 0, "", err
	}

	maps.accessTsRmu.Lock()
	maps.accessTs[hash] = token{
		date:     timeNow,
		id:       id,
		deviceId: deviceId,
	}
	maps.accessTsRmu.Unlock()

	maps.refreshTsRmu.Lock()
	maps.refreshTs[uid] = token{
		date:     timeNow,
		id:       id,
		deviceId: deviceId,
	}
	maps.refreshTsRmu.Unlock()

	return hash, uid, nil
}

// func that generates new tokens pair
// and remove expired and useless tokens
// from db and hashmap cache
func RegenTokensPair(access uint64, refresh string) (uint64, string, error) {
	timeNow := time.Now().Unix()

	maps.accessTsRmu.RLock()
	t, ok := maps.accessTs[access]
	maps.accessTsRmu.RUnlock()
	if ok {
		if timeNow-t.date < accessTLife {
			return 0, "", ErrAccessTNotExpired
		}
		maps.accessTsRmu.Lock()
		delete(maps.accessTs, access)
		maps.accessTsRmu.Unlock()
	}

	maps.refreshTsRmu.RLock()
	t, ok = maps.refreshTs[refresh]
	maps.refreshTsRmu.RUnlock()
	if !ok {
		return 0, "", ErrRefreshTNotFound
	}
	if timeNow-t.date > refreshTLife {
		return 0, "", ErrRefreshTExpired
	}

	id := t.id             // user account id
	deviceId := t.deviceId // uniq session id

	maps.refreshTsRmu.Lock()
	delete(maps.refreshTs, refresh)
	maps.refreshTsRmu.Unlock()

	return GenTokensPair(id, deviceId)
}
