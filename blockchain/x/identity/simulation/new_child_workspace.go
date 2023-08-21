package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"gitlab.qredo.com/qrdochain/fusionchain/x/identity/keeper"
	"gitlab.qredo.com/qrdochain/fusionchain/x/identity/types"
)

func SimulateMsgNewChildWorkspace(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgNewChildWorkspace{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the NewChildWorkspace simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "NewChildWorkspace simulation not implemented"), nil, nil
	}
}