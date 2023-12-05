package icp

const (
	DFX_VERSION   = "dfx -V"
	DFX_PRINCIPAL = "dfx identity get-principal"
	DFX_DEPLOY    = "dfx deploy"

	DFX_LEDGER_ACCOUNT = "dfx ledger account-id"
	DFX_LEDGER_BALANCE = "dfx ledger balance"
	DFX_LEDGER_CREATE  = "dfx ledger create-canister"

	DFX_CANISTER_ID      = "dfx canister id"
	DFX_CANISTER_CALL    = "dfx canister call"
	DFX_CANISTER_CREATE  = "dfx canister create"
	DFX_CANISTER_STATUS  = "dfx canister status"
	DFX_CANISTER_START   = "dfx canister start"
	DFX_CANISTER_STOP    = "dfx canister stop"
	DFX_CANISTER_DELETE  = "dfx canister delete"
	DFX_CANISTER_INFO    = "dfx canister info"
	DFX_CANISTER_META    = "dfx canister metadata"
	DFX_CANISTER_INSTALL = "dfx canister install"
	DFX_CANISTER_UNINST  = "dfx canister uninstall-code"
	DFX_CANISTER_UPDATE  = "dfx canister update-settings"
	DFX_CANISTER_DEPOSIT = "dfx canister deposit-cycles"

	OPT_NETWORK_IC = "--network ic"
	OPT_PLAYGROUND = "--playground"
	OPT_WITHCYCLES = "--with-cycles"
	OPT_IDENTITY   = "--identity"
	OPT_WALLET     = "--wallet"
	OPT_CANISTER   = "--canister"
	OPT_REINSTALL  = "--mode reinstall"
)
