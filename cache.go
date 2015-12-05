package cache

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/gogather/com/log"
	"time"
)

var memCache map[string]map[string]CacheItem

type CacheItem struct {
	expire time.Time
	data   []byte
}

func init() {
	memCache = make(map[string]map[string]CacheItem)
}

func initRegion(region string) {
	memCache[region] = make(map[string]CacheItem)
	return
}

func Set(region, key string, data interface{}) error {
	regionCache, ok := memCache[region]
	if !ok {
		initRegion(region)
		regionCache = memCache[region]
	}

	dataBytes, err := encode(data)
	if err != nil {
		return err
	} else {
		var item CacheItem
		item.data = dataBytes
		regionCache[key] = item
		return nil
	}
}

func Get(region, key string, to interface{}) error {
	regionCache, ok := memCache[region]
	if !ok {
		return errors.New("region not exist")
	} else {
		item, ok := regionCache[key]
		if !ok {
			return errors.New("item not exist")
		} else {
			if !item.expire.IsZero() {
				if time.Now().After(item.expire) {
					log.Yellowln(item)
					delete(regionCache, key)
					return errors.New("item expired")
				}
			}

			data := regionCache[key]
			return decode(data.data, to)

		}
	}
}

func encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
