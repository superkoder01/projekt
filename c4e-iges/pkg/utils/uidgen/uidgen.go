package uidgenutil

import "github.com/google/uuid"

type UIDGen interface {
	Get() string
}

type uidgen struct{}

func New() UIDGen {
	return &uidgen{}
}

func (u uidgen) Get() string {
	return uuid.New().String()
}
