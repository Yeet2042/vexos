package vexosservice

import "context"

type Service interface {
	V1() V1
}

type V1 interface {
	Start(ctx context.Context) error
}
