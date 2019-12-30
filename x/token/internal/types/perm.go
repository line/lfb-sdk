package types

import (
	"fmt"
)

var _ PermissionI = (*Permission)(nil)

const (
	MintAction  = "mint"
	IssueAction = "issue"
)

type PermissionI interface {
	GetResource() string
	GetAction() string
	Equal(string, string) bool
}

type Permissions []PermissionI

func (pms Permissions) String() string {
	return fmt.Sprintf("%#v", pms)
}

type Permission struct {
	Action   string `json:"action"`
	Resource string `json:"resource"`
}

func (p Permission) Validate() bool {
	if len(p.GetResource()) == 0 || len(p.GetAction()) == 0 {
		return false
	}
	return true
}

func (p Permission) GetResource() string {
	return p.Resource
}

func (p Permission) GetAction() string {
	return p.Action
}

func (p Permission) Equal(res, act string) bool {
	if p.GetResource() == res && p.GetAction() == act {
		return true
	}
	return false
}

func NewMintPermission(resource string) PermissionI {
	return &Permission{
		Action:   MintAction,
		Resource: resource,
	}
}

func NewIssuePermission(resource string) PermissionI {
	return &Permission{
		Action:   IssueAction,
		Resource: resource,
	}
}
