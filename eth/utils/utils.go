package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateAddress checks if the given string is a valid Ethereum address.
func ValidateAddress(address string) bool {
	if len(address) != 42 {
		return false
	}
	if !strings.HasPrefix(address, "0x") {
		return false
	}
	hasInvalidChars := regexp.MustCompile(`[^a-fA-F0-9]`).MatchString(address[2:])
	return !hasInvalidChars
}

// DecodeAddress decodes an ethereum address from a topic.
func DecodeAddress(topic string) (string, error) {
	stripped := strings.TrimPrefix(topic, "0x")
	if len(stripped) < 40 {
		return "", fmt.Errorf("unexpected length for Ethereum address topic: %d", len(stripped))
	}
	return "0x" + stripped[len(stripped)-40:], nil
}
