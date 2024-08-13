package logparser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

// AuthRecord represents a successful login record parsed from the auth.log file.
type AuthRecord struct {
	Username string
	IP       string
	Time     time.Time
}

// ParseAuthLog parses an auth.log file and returns a slice of successful login records.
func ParseAuthLog(filePath string) ([]AuthRecord, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open auth.log file: %w", err)
	}
	defer file.Close()

	var records []AuthRecord
	scanner := bufio.NewScanner(file)

	// Regex to match successful SSH login entries
	r := regexp.MustCompile(`(?P<Month>\w{3})\s+(?P<Day>\d+)\s+(?P<Time>\d{2}:\d{2}:\d{2})\s+\w+\s+sshd\[\d+\]:\s+Accepted\s+(?:password|publickey)\s+for\s+(?P<Username>\w+)\s+from\s+(?P<IP>[^\s]+)`)

	for scanner.Scan() {
		line := scanner.Text()
		if matches := r.FindStringSubmatch(line); matches != nil {
			timestamp, err := parseAuthTimestamp(matches[1], matches[2], matches[3])
			if err != nil {
				continue
			}

			record := AuthRecord{
				Username: matches[4],
				IP:       matches[5],
				Time:     timestamp,
			}
			records = append(records, record)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading auth.log file: %w", err)
	}

	return records, nil
}

// String returns a string representation of an AuthRecord.
func (r AuthRecord) String() string {
	return fmt.Sprintf("User: %s, IP: %s, Time: %s",
		r.Username, r.IP, r.Time.Format(time.RFC3339))
}

// Helper function to parse timestamps in the format found in auth.log and secure logs
func parseAuthTimestamp(month, day, timeStr string) (time.Time, error) {
	currentYear := time.Now().Year()
	timestampStr := fmt.Sprintf("%s %s %s %d", month, day, timeStr, currentYear)
	return time.Parse("Jan 2 15:04:05 2006", timestampStr)
}
