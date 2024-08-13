package wtmp

import (
	"fmt"
	"strings"
	"time"
)

// UtmpRecord represents a utmp record in the wtmp file.
type UtmpRecord struct {
	Type    int16
	Padding [2]byte
	Pid     int32
	Line    [32]byte
	ID      [4]byte
	User    [32]byte
	Host    [256]byte
	Exit    struct {
		Termination int16
		Exit        int16
	}
	Session int32
	Tv      struct {
		Sec  int32
		Usec int32
	}
	AddrV6   [4]int32
	Reserved [20]byte
}

// String returns a string representation of a UtmpRecord.
func (u UtmpRecord) String() string {
	username := trimNullBytes(u.User[:])
	line := trimNullBytes(u.Line[:])
	host := trimNullBytes(u.Host[:])
	timestamp := time.Unix(int64(u.Tv.Sec), int64(u.Tv.Usec)*1000)

	return formatUtmpRecord(username, line, host, timestamp)
}

// Helper function to remove null bytes from strings
func trimNullBytes(b []byte) string {
	return strings.TrimRight(string(b), "\x00")
}

// Helper function to format a UtmpRecord for output
func formatUtmpRecord(username, line, host string, timestamp time.Time) string {
	return fmt.Sprintf("User: %s, Line: %s, Host: %s, Time: %s",
		username, line, host, timestamp.Format(time.RFC3339))
}
