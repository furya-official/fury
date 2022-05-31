package bindings

import sdk "github.com/cosmos/cosmos-sdk/types"

type ComdexQuery struct {
	AppData                     *AppData                     `json:"get_app,omitempty"`
	AssetData                   *AssetData                   `json:"get_asset_data,omitempty"`
	State                       *State                       `json:"state,omitempty"`
	MintedToken                 *MintedToken                 `json:"total_supply,omitempty"`
	RemoveWhiteListAssetLocker  *RemoveWhiteListAssetLocker  `json:"remove_white_list_asset,omitempty"`
	WhitelistAppIdVaultInterest *WhitelistAppIdVaultInterest `json:"whitelist_app_id_vault_interest,omitempty"`
	WhitelistAppIdLockerRewards *WhitelistAppIdLockerRewards `json:"whitelist_app_id_locker_rewards,omitempty"`
	ExternalLockerRewards       *ExternalLockerRewards       `json:"external_locker_rewards,omitempty"`
	ExternalVaultRewards        *ExternalVaultRewards        `json:"external_vault_rewards,omitempty"`
}

type AppData struct {
	App_Id uint64 `json:"app_mapping_id"`
}

type AppDataResponse struct {
	MinGovDeposit    int64  `json:"min_gov_deposit"`
	GovTimeInSeconds int64  `json:"gov_time_in_seconds"`
	AssetId          uint64 `json:"gov_token_id"` //only when isGovToken true
}

type AssetData struct {
	Asset_Id uint64 `json:"asset_id"`
}

type AssetDataResponse struct {
	Denom string `json:"denom"`
}

type State struct {
	Address string `json:"address"`
	Denom   string `json:"denom"`
	Height  string `json:"height"`
	Target  string `json:"target"`
}

type StateResponse struct {
	Amount sdk.Coin `json:"amount"`
}

type MintedToken struct {
	App_Id   uint64 `json:"app_id"`
	Asset_Id uint64 `json:"asset_id"`
}

type MintedTokenResponse struct {
	MintedTokens int64 `json:"current_supply"`
}

type RemoveWhiteListAssetLocker struct {
	App_Id   uint64   `json:"app_id"`
	Asset_Id []uint64 `json:"asset_id"`
}

type RemoveWhiteListAssetResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type WhitelistAppIdVaultInterest struct {
	App_Id uint64 `json:"app_id"`
}

type WhitelistAppIdVaultInterestResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type WhitelistAppIdLockerRewards struct {
	App_Id   uint64 `json:"app_id"`
	Asset_Id uint64 `json:"asset_id"`
}

type WhitelistAppIdLockerRewardsResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type ExternalLockerRewards struct {
	App_Id   uint64 `json:"app_id"`
	Asset_Id uint64 `json:"asset_id"`
}

type ExternalLockerRewardsResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type ExternalVaultRewards struct {
	App_Id   uint64 `json:"app_id"`
	Asset_Id uint64 `json:"asset_id"`
}

type ExternalVaultRewardsResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}
