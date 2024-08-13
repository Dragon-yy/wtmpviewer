package wtmp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"time"
)

// ParseWtmp parses the wtmp file and returns a slice of Utmp records.
func ParseWtmp(filePath string) ([]UtmpRecord, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open wtmp file: %w", err)
	}
	defer file.Close()

	var records []UtmpRecord
	for {
		var utmp UtmpRecord
		err = binary.Read(file, binary.BigEndian, &utmp)
		if err != nil {
			if err == io.EOF {
				break // Reached the end of the file
			}
			return nil, fmt.Errorf("failed to read wtmp record: %w", err)
		}

		if utmp.Type == 7 { // Only consider user login entries
			records = append(records, utmp)
		}
	}

	return records, nil
}

// ParseWtmp parses the wtmp file and returns a slice of Utmp records.
func ParseWtmp2(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open wtmp file: %w", err)
	}
	defer file.Close()
	utmp := UtmpRecord{}
	recordSize := binary.Size(utmp)
	buffer := make([]byte, recordSize)

	for {
		_, err := file.Read(buffer)
		if err != nil {
			break
		}

		reader := bytes.NewReader(buffer)
		err = binary.Read(reader, binary.LittleEndian, &utmp)
		if err != nil {
			return fmt.Errorf("Error reading binary data:", err)

		}

		if utmp.Type == 7 { // USER_PROCESS
			fmt.Printf("User: %s, Line: %s, Host: %s, Time: %s\n",
				bytes.Trim(utmp.User[:], "\x00"),
				bytes.Trim(utmp.Line[:], "\x00"),
				bytes.Trim(utmp.Host[:], "\x00"),
				time.Unix(int64(utmp.Tv.Sec), int64(utmp.Tv.Usec)).Format(time.RFC3339))
		}
	}
	return nil
}
