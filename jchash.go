package jchash

import (
	"errors"
	"github.com/spaolacci/murmur3"
)

func JumpConsistentHashStr(key string, num_buckets uint64) (uint64, error) {
	return JumpConsistentHash(murmur3.Sum64([]byte(key)), num_buckets)
}

func JumpConsistentHashInt(key, num_buckets uint64) (uint64, error) {
	var b, j uint64 = 0, 0
	err := errors.New("error invalid bucket size specified")
	for j < num_buckets {
		b = j
		key = key*2862933555777941757 + 1
		j = (b + 1) * uint64(float64(1<<31)/float64((key>>33)+1))
		err = nil
	}
	return b, err
}

func JumpConsistentHash(key interface{}, num_buckets uint64) (uint64, error) {
	switch k := key.(type) {
	case string:
		return JumpConsistentHashStr(k, num_buckets)
	case uint64:
		return JumpConsistentHashInt(k, num_buckets)
	}
	return 0, errors.New("Invalid Type Specified")
}
