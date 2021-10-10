package src

import (
	"bytes"
	"encoding/binary"
	"math"
	"time"
)

func String(buf *bytes.Buffer) string {

	var strlen int

	if strflag := buf.Next(1); strflag[0] == 0x0b {
		shift := 0

		for { //ULEB128
			b := buf.Next(1)
			strlen |= (int(b[0]) & 0x7f) << shift
			if (b[0] & (1 << 7)) == 0 {
				break
			}
			shift += 7
		}
	}
	str := string(buf.Next(strlen))
	if len(str) < 1 {
		return ""
	}
	return str
}
func Int(buf *bytes.Buffer) int32 {
	i := binary.LittleEndian.Uint32(buf.Next(4))
	return int32(i)
}

func Short(buf *bytes.Buffer) int16 {
	i := binary.LittleEndian.Uint16(buf.Next(2))
	return int16(i)
}
func Long(buf *bytes.Buffer) int64 {
	i := binary.LittleEndian.Uint64(buf.Next(8))
	return int64(i)
}

func IntDouble(buf *bytes.Buffer) uint64 {
	return binary.LittleEndian.Uint64(buf.Next(8))
}

func Bool(buf *bytes.Buffer) bool {
	return buf.Next(1)[0] != 0
}

func DateTime(buf *bytes.Buffer) uint64 {
	return binary.LittleEndian.Uint64(buf.Next(8))
}

func ByteInt(buf *bytes.Buffer) int {
	return int(buf.Next(1)[0])
}

func Single(buf *bytes.Buffer) float32 {
	i := binary.LittleEndian.Uint32(buf.Next(4))
	r := math.Float32frombits(i)
	if math.IsNaN(float64(r)) {
		return 0
	}
	return r

}
func Double(buf *bytes.Buffer) float64 {
	i := binary.LittleEndian.Uint64(buf.Next(8))

	r := math.Float64frombits(i)
	if r == math.Inf(1) {
		return math.MaxFloat64
	}
	if r == math.Inf(-1) {
		return -math.MaxFloat64
	}
	if math.IsNaN(r) {
		return 0
	}
	return r
}

func RankedStatus(i int) (s string) {
	switch i {
	case 0:
		s = "unknown"
	case 1:
		s = "unsubmitted"
	case 2:
		s = "pending/wip/graveyard"
	case 3:
		s = "unused"
	case 4:
		s = "ranked"
	case 5:
		s = "approved"
	case 6:
		s = "qualified"
	case 7:
		s = "loved"
	}
	return
}

func WindowsTick(buf *bytes.Buffer) string {
	tim := int64(binary.LittleEndian.Uint64(buf.Next(8)))
	return time.Unix(tim/10000000-62135596800, tim%10000000).UTC().Format(time.RFC3339)
}
func ModsParser(m int32) (mods string) {
	if m&1 != 0 {
		mods += "NF"
	}
	if m&2 != 0 {
		mods += "EZ"
	}
	if m&4 != 0 {
		mods += "TD"
	}
	if m&8 != 0 {
		mods += "HD"
	}
	if m&16 != 0 {
		mods += "HR"
	}
	if m&32 != 0 {
		mods += "SD"
	}
	if m&64 != 0 {
		mods += "DT"
	}
	if m&128 != 0 {
		mods += "RX"
	}
	if m&256 != 0 {
		mods += "HT"
	}
	if m&512 != 0 {
		mods += "NC"
	}
	if m&1024 != 0 {
		mods += "FL"
	}
	if m&2048 != 0 {
		mods += "AT"
	}
	if m&4096 != 0 {
		mods += "SO"
	}
	if m&8192 != 0 {
		mods += "AP"
	}
	if m&16384 != 0 {
		mods += "PF"
	}
	if m&32768 != 0 {
		mods += "4K"
	}
	if m&65536 != 0 {
		mods += "5K"
	}
	if m&131072 != 0 {
		mods += "6K"
	}
	if m&262144 != 0 {
		mods += "7K"
	}
	if m&524288 != 0 {
		mods += "8K"
	}
	if m&1048576 != 0 {
		mods += "FI"
	}
	if m&2097152 != 0 {
		mods += "RD"
	}
	if m&4194304 != 0 {
		mods += "CN"
	}
	if m&8388608 != 0 {
		mods += "TG"
	}
	if m&16777216 != 0 {
		mods += "9K"
	}
	if m&33554432 != 0 {
		mods += "CO"
	}
	if m&67108864 != 0 {
		mods += "1K"
	}
	if m&134217728 != 0 {
		mods += "3K"
	}
	if m&268435456 != 0 {
		mods += "2K"
	}
	if m&536870912 != 0 {
		mods += "V2"
	}
	if m&1073741824 != 0 {
		mods += "MR"
	}
	return
}
