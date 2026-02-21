package utils

import (
	"math"
	"sort"
	"strings"
)

// ContainsAny returns true if s contains any of the substrings (case-insensitive).
func ContainsAny(s string, substrings ...string) bool {
	upper := strings.ToUpper(s)
	for _, sub := range substrings {
		if strings.Contains(upper, strings.ToUpper(sub)) {
			return true
		}
	}
	return false
}

// Min returns the smaller of two float64 values.
func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of two float64 values.
func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Round rounds a float64 to the given number of decimal places.
func Round(val float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(val*shift) / shift
}

// Mean calculates the arithmetic mean of a slice of float64 values.
func Mean(vals []float64) float64 {
	if len(vals) == 0 {
		return 0
	}
	var sum float64
	for _, v := range vals {
		sum += v
	}
	return sum / float64(len(vals))
}

// Median calculates the median of a slice of float64 values.
// The input slice is not modified.
func Median(vals []float64) float64 {
	if len(vals) == 0 {
		return 0
	}
	sorted := make([]float64, len(vals))
	copy(sorted, vals)
	sort.Float64s(sorted)
	n := len(sorted)
	if n%2 == 0 {
		return (sorted[n/2-1] + sorted[n/2]) / 2
	}
	return sorted[n/2]
}

// StdDev calculates the population standard deviation.
func StdDev(vals []float64) float64 {
	if len(vals) < 2 {
		return 0
	}
	m := Mean(vals)
	var sumSq float64
	for _, v := range vals {
		diff := v - m
		sumSq += diff * diff
	}
	return math.Sqrt(sumSq / float64(len(vals)))
}

// Truncate shortens a string to maxLen characters, appending "..." if trimmed.
func Truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}

// NormalizeString trims whitespace and normalises line endings.
func NormalizeString(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	return strings.TrimSpace(s)
}

// StringSliceContains returns true if needle is in haystack (case-insensitive).
func StringSliceContains(haystack []string, needle string) bool {
	needle = strings.ToLower(needle)
	for _, h := range haystack {
		if strings.ToLower(h) == needle {
			return true
		}
	}
	return false
}

// ParseBool parses common boolean string representations.
func ParseBool(s string) bool {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "true", "yes", "1", "t", "y":
		return true
	default:
		return false
	}
}

// SafePercentage returns 0 when total is 0 to avoid division by zero.
func SafePercentage(part, total int) float64 {
	if total == 0 {
		return 0
	}
	return float64(part) / float64(total) * 100
}

// UniqueStrings removes duplicates from a string slice while preserving order.
func UniqueStrings(ss []string) []string {
	seen := make(map[string]struct{}, len(ss))
	out := make([]string, 0, len(ss))
	for _, s := range ss {
		if _, ok := seen[s]; !ok {
			seen[s] = struct{}{}
			out = append(out, s)
		}
	}
	return out
}
