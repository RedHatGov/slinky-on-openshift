// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"fmt"
	"strconv"
	"time"
)

func validFirstDigit(str string) bool {
	if len(str) == 0 {
		return false
	}
	return str[0] == '-' || (str[0] == '0' && str == "0") || (str[0] >= '1' && str[0] <= '9')
}

// GetNumberFromAnnotations returns the integer value of annotation.
// Returns 0 if not set or the value is invalid.
func GetNumberFromAnnotations(annotations map[string]string, key string) (int32, error) {
	if value, exist := annotations[key]; exist {
		// values that start with plus sign (e.g, "+10") or leading zeros (e.g., "008") are not valid.
		if !validFirstDigit(value) {
			return 0, fmt.Errorf("invalid value %q", value)
		}

		i, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			// make sure we default to 0 on error.
			return 0, err
		}
		return int32(i), nil
	}
	return 0, nil
}

// GetBoolFromAnnotations returns the value of annotation.
// Returns false if not set or the value is invalid.
func GetBoolFromAnnotations(annotations map[string]string, key string) (bool, error) {
	if value, exist := annotations[key]; exist {
		b, err := strconv.ParseBool(value)
		if err != nil {
			// make sure we default to false on error.
			return false, err
		}
		return b, nil
	}
	return false, nil
}

// GetTimeFromAnnotations returns the integer value of annotation.
// Returns unit Time if not set or the value is invalid.
func GetTimeFromAnnotations(annotations map[string]string, key string) (time.Time, error) {
	if value, ok := annotations[key]; ok {
		t, err := time.Parse(time.RFC3339, value)
		if err != nil {
			// make sure we default to unit Time on error.
			return time.Time{}, err
		}
		return t, nil
	}
	return time.Time{}, nil
}
