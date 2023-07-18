package models

import (
	"encoding/base32"

	"github.com/pborman/uuid"
)

var encoding = base32.NewEncoding("gyudij7pobxsfqn52hzeat01836wrcm4").WithPadding(base32.NoPadding)

func NewId() string {
	return encoding.EncodeToString(uuid.NewRandom())
}
