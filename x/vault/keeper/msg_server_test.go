package keeper_test

import (
	"fmt"

	"github.com/comdex-official/comdex/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/stretchr/testify/suite"
)

func (s *KeeperTestSuite) TestMsgCreate() {

	addr1 := s.addr(1)
	addr2 := s.addr(2)

	appID1 := s.CreateNewApp("appOne")
	appID2 := s.CreateNewApp("appTwo")
	asseOneID := s.CreateNewAsset("ASSET1", "uasset1", 1000000)
	asseTwoID := s.CreateNewAsset("ASSET2", "uasset2", 2000000)
	pairID := s.CreateNewPair(addr1, asseOneID, asseTwoID)
	extendedVaultPairID1 := s.CreateNewExtendedVaultPair("CMDX C", appID1, pairID)
	extendedVaultPairID2 := s.CreateNewExtendedVaultPair("CMDX C", appID2, pairID)

	testCases := []struct {
		Name               string
		Msg                types.MsgCreateRequest
		ExpErr             error
		ExpResp            *types.MsgCreateResponse
		QueryResponseIndex uint64
		QueryResponse      *types.Vault
		AvailableBalance   sdk.Coins
	}{
		{
			Name: "error extended vault pair does not exists",
			Msg: *types.NewMsgCreateRequest(
				addr1, appID1, 123, newInt(10000000), newInt(4000000),
			),
			ExpErr:             types.ErrorExtendedPairVaultDoesNotExists,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "error invalid appID",
			Msg: *types.NewMsgCreateRequest(
				addr1, 12, extendedVaultPairID1, newInt(10000000), newInt(4000000),
			),
			ExpErr:             types.ErrorAppMappingDoesNotExist,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "error appID mismatch",
			Msg: *types.NewMsgCreateRequest(
				addr1, appID2, extendedVaultPairID1, newInt(10000000), newInt(4000000),
			),
			ExpErr:             types.ErrorAppMappingIDMismatch,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "error invalid from address",
			Msg: *types.NewMsgCreateRequest(
				[]byte(""), appID1, extendedVaultPairID1, newInt(10000000), newInt(4000000),
			),
			ExpErr:             fmt.Errorf("empty address string is not allowed"),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "error amount out smaller that debt floor",
			Msg: *types.NewMsgCreateRequest(
				addr1, appID1, extendedVaultPairID1, newInt(10000000), newInt(4000000),
			),
			ExpErr:             types.ErrorAmountOutLessThanDebtFloor,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "error invalid collaterlization ratio",
			Msg: *types.NewMsgCreateRequest(
				addr1, appID1, extendedVaultPairID1, newInt(800000000), newInt(200000000),
			),
			ExpErr:             types.ErrorInvalidCollateralizationRatio,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "error insufficient funds",
			Msg: *types.NewMsgCreateRequest(
				addr1, appID1, extendedVaultPairID1, newInt(1000000000), newInt(200000000),
			),
			ExpErr:             fmt.Errorf(fmt.Sprintf("0uasset1 is smaller than %duasset1: insufficient funds", 1000000000)),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "success valid case app1 user1",
			Msg: *types.NewMsgCreateRequest(
				addr1, appID1, extendedVaultPairID1, newInt(1000000000), newInt(200000000),
			),
			ExpErr:             nil,
			ExpResp:            &types.MsgCreateResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.Vault{
				Id:                  "appOne1",
				AppMappingId:        appID1,
				ExtendedPairVaultID: extendedVaultPairID1,
				Owner:               addr1.String(),
				AmountIn:            newInt(1000000000),
				AmountOut:           newInt(200000000),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset2", newInt(198000000))),
		},
		{
			Name: "success valid case app1 user2",
			Msg: *types.NewMsgCreateRequest(
				addr2, appID1, extendedVaultPairID1, newInt(1000000000), newInt(200000000),
			),
			ExpErr:             nil,
			ExpResp:            &types.MsgCreateResponse{},
			QueryResponseIndex: 1,
			QueryResponse: &types.Vault{
				Id:                  "appOne2",
				AppMappingId:        appID1,
				ExtendedPairVaultID: extendedVaultPairID1,
				Owner:               addr2.String(),
				AmountIn:            newInt(1000000000),
				AmountOut:           newInt(200000000),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset2", newInt(198000000))),
		},
		{
			Name: "error user vault already exists",
			Msg: *types.NewMsgCreateRequest(
				addr1, appID1, extendedVaultPairID1, newInt(1000000000), newInt(200000000),
			),
			ExpErr:             types.ErrorUserVaultAlreadyExists,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "success valid case app2 user1",
			Msg: *types.NewMsgCreateRequest(
				addr1, appID2, extendedVaultPairID2, newInt(1000000000), newInt(200000000),
			),
			ExpErr:             nil,
			ExpResp:            &types.MsgCreateResponse{},
			QueryResponseIndex: 2,
			QueryResponse: &types.Vault{
				Id:                  "appTwo1",
				AppMappingId:        appID2,
				ExtendedPairVaultID: extendedVaultPairID2,
				Owner:               addr1.String(),
				AmountIn:            newInt(1000000000),
				AmountOut:           newInt(200000000),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset2", newInt(198000000*2))),
		},
		{
			Name: "success valid case app2 user2",
			Msg: *types.NewMsgCreateRequest(
				addr2, appID2, extendedVaultPairID2, newInt(1000000000), newInt(200000000),
			),
			ExpErr:             nil,
			ExpResp:            &types.MsgCreateResponse{},
			QueryResponseIndex: 3,
			QueryResponse: &types.Vault{
				Id:                  "appTwo2",
				AppMappingId:        appID2,
				ExtendedPairVaultID: extendedVaultPairID2,
				Owner:               addr2.String(),
				AmountIn:            newInt(1000000000),
				AmountOut:           newInt(200000000),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset2", newInt(198000000*2))),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {

			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.From), sdk.NewCoins(sdk.NewCoin("uasset1", tc.Msg.AmountIn)))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.MsgCreate(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)

				availableBalances := s.getBalances(sdk.MustAccAddressFromBech32(tc.Msg.From))
				s.Require().True(tc.AvailableBalance.IsEqual(availableBalances))

				vaults := s.keeper.GetVaults(s.ctx)
				s.Require().Len(vaults, int(tc.QueryResponseIndex+1))
				s.Require().Equal(tc.QueryResponse.Id, vaults[tc.QueryResponseIndex].Id)
				s.Require().Equal(tc.QueryResponse.Owner, vaults[tc.QueryResponseIndex].Owner)
				s.Require().Equal(tc.QueryResponse.AmountIn, vaults[tc.QueryResponseIndex].AmountIn)
				s.Require().Equal(tc.QueryResponse.AmountOut, vaults[tc.QueryResponseIndex].AmountOut)
				s.Require().Equal(tc.QueryResponse.ExtendedPairVaultID, vaults[tc.QueryResponseIndex].ExtendedPairVaultID)
				s.Require().Equal(tc.QueryResponse.AppMappingId, vaults[tc.QueryResponseIndex].AppMappingId)
			}
		})
	}

}

func (s *KeeperTestSuite) TestMsgDeposit() {

	addr1 := s.addr(1)
	addr2 := s.addr(2)

	appID1 := s.CreateNewApp("appOne")
	appID2 := s.CreateNewApp("appTwo")
	asseOneID := s.CreateNewAsset("ASSET1", "uasset1", 1000000)
	asseTwoID := s.CreateNewAsset("ASSET2", "uasset2", 2000000)
	pairID := s.CreateNewPair(addr1, asseOneID, asseTwoID)
	extendedVaultPairID1 := s.CreateNewExtendedVaultPair("CMDX C", appID1, pairID)
	extendedVaultPairID2 := s.CreateNewExtendedVaultPair("CMDX C", appID2, pairID)

	msg := types.NewMsgCreateRequest(addr1, appID1, extendedVaultPairID1, newInt(1000000000), newInt(200000000))
	s.fundAddr(sdk.MustAccAddressFromBech32(addr1.String()), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(1000000000))))
	s.msgServer.MsgCreate(sdk.WrapSDKContext(s.ctx), msg)

	msg = types.NewMsgCreateRequest(addr2, appID2, extendedVaultPairID2, newInt(1000000000), newInt(200000000))
	s.fundAddr(sdk.MustAccAddressFromBech32(addr2.String()), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(1000000000))))
	s.msgServer.MsgCreate(sdk.WrapSDKContext(s.ctx), msg)

	testCases := []struct {
		Name               string
		Msg                types.MsgDepositRequest
		ExpErr             error
		ExpResp            *types.MsgDepositResponse
		QueryResponseIndex uint64
		QueryResponse      *types.Vault
		AvailableBalance   sdk.Coins
	}{
		{
			Name: "error invalid from address",
			Msg: *types.NewMsgDepositRequest(
				[]byte(""), appID1, extendedVaultPairID1, "appOne1", newInt(69000000),
			),
			ExpErr:             fmt.Errorf("empty address string is not allowed"),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "error extended vault pair does not exists",
			Msg: *types.NewMsgDepositRequest(
				addr1, appID1, 123, "appOne1", newInt(4000000),
			),
			ExpErr:             types.ErrorExtendedPairVaultDoesNotExists,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "error invalid appID",
			Msg: *types.NewMsgDepositRequest(
				addr1, 69, extendedVaultPairID1, "appOne1", newInt(69000000),
			),
			ExpErr:             types.ErrorAppMappingDoesNotExist,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "error appID mismatch",
			Msg: *types.NewMsgDepositRequest(
				addr1, appID2, extendedVaultPairID1, "appOne1", newInt(69000000),
			),
			ExpErr:             types.ErrorAppMappingIDMismatch,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "error vault does not exists",
			Msg: *types.NewMsgDepositRequest(
				addr1, appID1, extendedVaultPairID1, "appOne2", newInt(69000000),
			),
			ExpErr:             types.ErrorVaultDoesNotExist,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "error access unathorized ",
			Msg: *types.NewMsgDepositRequest(
				addr2, appID1, extendedVaultPairID1, "appOne1", newInt(69000000),
			),
			ExpErr:             types.ErrVaultAccessUnauthorised,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "error insufficient funds",
			Msg: *types.NewMsgDepositRequest(
				addr1, appID1, extendedVaultPairID1, "appOne1", newInt(69000000),
			),
			ExpErr:             fmt.Errorf(fmt.Sprintf("0uasset1 is smaller than %duasset1: insufficient funds", 69000000)),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset2", newInt(0))),
		},
		{
			Name: "success valid case app1 user1",
			Msg: *types.NewMsgDepositRequest(
				addr1, appID1, extendedVaultPairID1, "appOne1", newInt(69000000),
			),
			ExpErr:             nil,
			ExpResp:            &types.MsgDepositResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.Vault{
				Id:                  "appOne1",
				AppMappingId:        appID1,
				ExtendedPairVaultID: extendedVaultPairID1,
				Owner:               addr1.String(),
				AmountIn:            newInt(1069000000),
				AmountOut:           newInt(200000000),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset2", newInt(198000000))),
		},
		{
			Name: "success valid case app2 user2",
			Msg: *types.NewMsgDepositRequest(
				addr2, appID2, extendedVaultPairID2, "appTwo1", newInt(69000000),
			),
			ExpErr:             nil,
			ExpResp:            &types.MsgDepositResponse{},
			QueryResponseIndex: 1,
			QueryResponse: &types.Vault{
				Id:                  "appTwo1",
				AppMappingId:        appID2,
				ExtendedPairVaultID: extendedVaultPairID2,
				Owner:               addr2.String(),
				AmountIn:            newInt(1069000000),
				AmountOut:           newInt(200000000),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset2", newInt(198000000))),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {

			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.From), sdk.NewCoins(sdk.NewCoin("uasset1", tc.Msg.Amount)))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.MsgDeposit(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)

				availableBalances := s.getBalances(sdk.MustAccAddressFromBech32(tc.Msg.From))
				s.Require().True(tc.AvailableBalance.IsEqual(availableBalances))

				vaults := s.keeper.GetVaults(s.ctx)
				s.Require().Len(vaults, 2)
				s.Require().Equal(tc.QueryResponse.Id, vaults[tc.QueryResponseIndex].Id)
				s.Require().Equal(tc.QueryResponse.Owner, vaults[tc.QueryResponseIndex].Owner)
				s.Require().Equal(tc.QueryResponse.AmountIn, vaults[tc.QueryResponseIndex].AmountIn)
				s.Require().Equal(tc.QueryResponse.AmountOut, vaults[tc.QueryResponseIndex].AmountOut)
				s.Require().Equal(tc.QueryResponse.ExtendedPairVaultID, vaults[tc.QueryResponseIndex].ExtendedPairVaultID)
				s.Require().Equal(tc.QueryResponse.AppMappingId, vaults[tc.QueryResponseIndex].AppMappingId)
			}
		})
	}
}
