package service

import "errors"

var (
	ErrConstructPVC = errors.New("construct pvc failed")
	ErrCreatePVC    = errors.New("create pvc failed")
	ErrCreatePod    = errors.New("create pod failed")
	ErrDeletePod    = errors.New("delete pod failed")
	ErrDeletePVC    = errors.New("delete pvc failed")
)
