package helper

import (
	"regexp"

	"github.com/huydq189/bbb-storage/domain"
)

// IsValidRecordID return if record id is numberic and single dash format
func IsValidRecordID(id string) (b bool, err error) {
	if m, _ := regexp.MatchString("^[a-z0-9]*-[0-9]*$", id); !m {
		return false, domain.ErrBadParamInput
	}
	return true, nil
}
