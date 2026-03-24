package types

import "github.com/google/uuid"

type OBJECT_ID string

func (id OBJECT_ID) IsEmpty() bool {
	return id == ""
}

func (oid OBJECT_ID) String() string {
	return string(oid)
}

func NewObjectId() OBJECT_ID {
	return OBJECT_ID(uuid.New().String())
}
func NewObjectIdRef() *OBJECT_ID {
	oid := NewObjectId()
	return &oid
}

func StrPtr(s string) *string {
	return &s
}
