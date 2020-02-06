package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	QuerierRoute     = "safetybox"
	QuerySafetyBox   = "safetybox"
	QueryAccountRole = "role"
)

func NewQuerySafetyBoxParams(safetyBoxID string) QuerySafetyBoxParams {
	return QuerySafetyBoxParams{safetyBoxID}
}

type QuerySafetyBoxParams struct {
	SafetyBoxID string `json:"safety_box_id"`
}

func (qsb QuerySafetyBoxParams) String() string {
	return string(codec.MustMarshalJSONIndent(ModuleCdc, qsb))
}

type NodeQuerier interface {
	QueryWithData(path string, data []byte) ([]byte, int64, error)
}

type SafetyBoxRetriever struct {
	querier NodeQuerier
}

func NewSafetyBoxRetriever(querier NodeQuerier) SafetyBoxRetriever {
	return SafetyBoxRetriever{querier: querier}
}

func (sbr SafetyBoxRetriever) GetSafetyBox(safetyBoxID string) (SafetyBox, int64, error) {
	sb, height, err := sbr.GetSafetyBoxWithHeight(safetyBoxID)
	return sb, height, err
}

func (sbr SafetyBoxRetriever) GetSafetyBoxWithHeight(safetyBoxID string) (SafetyBox, int64, error) {
	bs, err := ModuleCdc.MarshalJSON(NewQuerySafetyBoxParams(safetyBoxID))
	if err != nil {
		return SafetyBox{}, 0, err
	}

	res, height, err := sbr.querier.QueryWithData(fmt.Sprintf("custom/%s/%s", QuerierRoute, QuerySafetyBox), bs)
	if err != nil {
		return SafetyBox{}, height, err
	}

	var sb SafetyBox
	if err := ModuleCdc.UnmarshalJSON(res, &sb); err != nil {
		return SafetyBox{}, height, err
	}

	return sb, height, nil
}

type QueryAccountRoleParams struct {
	SafetyBoxID string         `json:"safety_box_id"`
	Role        string         `json:"role"`
	Address     sdk.AccAddress `json:"address"`
}

func (qapp QueryAccountRoleParams) String() string {
	return string(codec.MustMarshalJSONIndent(ModuleCdc, qapp))
}

func NewQueryAccountRoleParams(id, role string, address sdk.AccAddress) QueryAccountRoleParams {
	return QueryAccountRoleParams{id, role, address}
}

type AccountPermissionRetriever struct {
	querier NodeQuerier
}

func NewAccountPermissionRetriever(querier NodeQuerier) AccountPermissionRetriever {
	return AccountPermissionRetriever{querier: querier}
}

func (apr AccountPermissionRetriever) GetAccountPermissions(id, role string, addr sdk.AccAddress) (MsgSafetyBoxRoleResponse, int64, error) {
	pms, height, err := apr.GetAccountPermissionsWithHeight(id, role, addr)
	return pms, height, err
}

func (apr AccountPermissionRetriever) GetAccountPermissionsWithHeight(id, role string, addr sdk.AccAddress) (MsgSafetyBoxRoleResponse, int64, error) {
	bs, err := ModuleCdc.MarshalJSON(NewQueryAccountRoleParams(id, role, addr))
	if err != nil {
		return MsgSafetyBoxRoleResponse{}, 0, err
	}

	res, height, err := apr.querier.QueryWithData(fmt.Sprintf("custom/%s/%s", QuerierRoute, QueryAccountRole), bs)
	if err != nil {
		return MsgSafetyBoxRoleResponse{}, height, err
	}

	var sbpr MsgSafetyBoxRoleResponse
	if err := ModuleCdc.UnmarshalJSON(res, &sbpr); err != nil {
		return MsgSafetyBoxRoleResponse{}, height, err
	}

	return sbpr, height, nil
}
