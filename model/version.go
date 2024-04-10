package model

import "time"

type Version struct {
	ItemMajor int64
	ItemMinor int64
	Major     uint16
	Minor     uint16
	Patch     uint16
	Extension uint16
	Created   time.Time
	Updated   time.Time
}

func (ver *Version) ToUint64() uint64 {
	return uint64(ver.Major)<<48 +
		uint64(ver.Minor)<<32 +
		uint64(ver.Patch)<<16 +
		uint64(ver.Extension)
}

func (ver *Version) FromUint64(i uint64) {
	ver.Extension = uint16(i & 0xffff)
	ver.Patch = uint16(i >> 16 & 0xffff)
	ver.Minor = uint16(i >> 32 & 0xffff)
	ver.Major = uint16(i >> 48 & 0xffff)
}
