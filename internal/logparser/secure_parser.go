package logparser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

// SecureRecord represents a successful login record parsed from the secure log.
type SecureRecord struct {
	Username string
	IP       string
	Time     time.Time
}

// ParseSecure parses a secure log file and returns a slice of successful login records.
func ParseSecure(filePath string) ([]SecureRecord, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open secure file: %w", err)
	}
	defer file.Close()

	var records []SecureRecord
	scanner := bufio.NewScanner(file)

	// Regex to match successful SSH login entries
	r := regexp.MustCompile(`(?P<Month>\w{3})\s+(?P<Day>\d+)\s+(?P<Time>\d{2}:\d{2}:\d{2})\s+\S+\s+sshd\[\d+\]:\s+Accepted\s+(?:password|publickey)\s+for\s+(?P<Username>\S+)\s+from\s+(?P<IP>\d+\.\d+\.\d+\.\d+)`)

	for scanner.Scan() {
		line := scanner.Text()
		if matches := r.FindStringSubmatch(line); matches != nil {
			timestamp, err := parseSecureTimestamp(matches[1], matches[2], matches[3])
			if err != nil {
				continue
			}

			record := SecureRecord{
				Username: matches[4],
				IP:       matches[5],
				Time:     timestamp,
			}
			records = append(records, record)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading secure file: %w", err)
	}

	return records, nil
}

// Helper function to parse timestamps in the format found in secure logs
func parseSecureTimestamp(month, day, timeStr string) (time.Time, error) {
	currentYear := time.Now().Year()
	timestampStr := fmt.Sprintf("%s %s %s %d", month, day, timeStr, currentYear)
	return time.Parse("Jan 2 15:04:05 2006", timestampStr)
}

// String returns a string representation of a SecureRecord.
func (r SecureRecord) String() string {
	return fmt.Sprintf("User: %s, IP: %s, Time: %s",
		r.Username, r.IP, r.Time.Format(time.RFC3339))
}
