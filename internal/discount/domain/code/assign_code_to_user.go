package code

import (
	"errors"
	"time"
)

func (d *DiscountCode) AssignToUser(userID string) error {
	if userID == "" {
		return errors.New("empty user id")
	}

	if d.assignedUserID != "" {
		return errors.New("code is already assigned to another user")
	}

	d.assignedUserID = userID
	d.updatedAt = time.Now()

	return nil
}
