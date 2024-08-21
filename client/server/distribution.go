//nolint:wrapcheck // The api server is our server, so we don't need to wrap it.
package server

import (
	"errors"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/gorilla/mux"

	"github.com/storyprotocol/iliad/client/server/utils"
)

func (s *Server) initDistributionRoute() {
	s.httpMux.HandleFunc("/distribution/validators/{validator_address}", utils.SimpleWrap(s.aminoCodec, s.GetDistributionValidatorByValidatorAddress))
	s.httpMux.HandleFunc("/distribution/validators/{validator_address}/commission", utils.SimpleWrap(s.aminoCodec, s.GetValidatorCommissionByValidatorAddress))
	s.httpMux.HandleFunc("/distribution/validators/{validator_address}/outstanding_rewards", utils.SimpleWrap(s.aminoCodec, s.GetValidatorOutstandingRewardsByValidatorAddress))
	s.httpMux.HandleFunc("/distribution/validators/{validator_address}/slashes", utils.AutoWrap(s.aminoCodec, s.GetValidatorSlashesByValidatorAddress))
	s.httpMux.HandleFunc("/distribution/all_validators/outstanding_rewards", utils.AutoWrap(s.aminoCodec, s.GetAllValidatorOutstandingRewards))

	s.httpMux.HandleFunc("/distribution/delegators/{delegator_address}/validators", utils.SimpleWrap(s.aminoCodec, s.GetDistributionValidatorsByDelegatorAddress))
	s.httpMux.HandleFunc("/distribution/delegators/{delegator_address}/rewards", utils.SimpleWrap(s.aminoCodec, s.GetDelegatorRewardsByDelegatorAddress))
	s.httpMux.HandleFunc("/distribution/delegators/{delegator_address}/rewards/{validator_address}", utils.SimpleWrap(s.aminoCodec, s.GetDelegatorRewardsByDelegatorAddressValidatorAddress))
	s.httpMux.HandleFunc("/distribution/delegators/{delegator_address}/withdraw_address", utils.SimpleWrap(s.aminoCodec, s.GetDelegatorWithdrawAddressByDelegatorAddress))
}

// GetDistributionValidatorByValidatorAddress queries validator commission and self-delegation rewards for validator.
func (s *Server) GetDistributionValidatorByValidatorAddress(r *http.Request) (resp any, err error) {
	queryContext, err := s.createQueryContextByHeader(r)
	if err != nil {
		return nil, err
	}

	queryResp, err := keeper.NewQuerier(s.store.GetDistrKeeper()).ValidatorDistributionInfo(queryContext, &distributiontypes.QueryValidatorDistributionInfoRequest{
		ValidatorAddress: mux.Vars(r)["validator_address"],
	})

	if err != nil {
		return nil, err
	}

	return queryResp, nil
}

// GetValidatorCommissionByValidatorAddress queries accumulated commission for a validator.
func (s *Server) GetValidatorCommissionByValidatorAddress(r *http.Request) (resp any, err error) {
	queryContext, err := s.createQueryContextByHeader(r)
	if err != nil {
		return nil, err
	}

	queryResp, err := keeper.NewQuerier(s.store.GetDistrKeeper()).ValidatorCommission(queryContext, &distributiontypes.QueryValidatorCommissionRequest{
		ValidatorAddress: mux.Vars(r)["validator_address"],
	})

	if err != nil {
		return nil, err
	}

	return queryResp, nil
}

// GetValidatorOutstandingRewardsByValidatorAddress queries rewards of a validator address.
func (s *Server) GetValidatorOutstandingRewardsByValidatorAddress(r *http.Request) (resp any, err error) {
	queryContext, err := s.createQueryContextByHeader(r)
	if err != nil {
		return nil, err
	}

	queryResp, err := keeper.NewQuerier(s.store.GetDistrKeeper()).ValidatorOutstandingRewards(queryContext, &distributiontypes.QueryValidatorOutstandingRewardsRequest{
		ValidatorAddress: mux.Vars(r)["validator_address"],
	})

	if err != nil {
		return nil, err
	}

	return queryResp, nil
}

// GetAllValidatorOutstandingRewards queries rewards of all validators.
func (s *Server) GetAllValidatorOutstandingRewards(req *getAllValidatorOutstandingRewardsRequest, r *http.Request) (resp any, err error) {
	if req.To-req.From > 100 {
		return nil, errors.New("search max 100 blocks")
	}

	curBlock, err := s.cl.Block(r.Context(), nil)
	if err != nil {
		return nil, errors.Join(errors.New("curbock fetch fail"), err)
	}

	querier := keeper.NewQuerier(s.store.GetDistrKeeper())
	result := make([]*getAllValidatorOutstandingRewardsRequestBlockResults, 0)

	for i := req.From; i < min(req.To, curBlock.Block.Height); i++ {
		queryContext, err := s.store.CreateQueryContext(i, false)
		if err != nil {
			return nil, errors.Join(errors.New("create query context fail"), err)
		}

		blockResult := &getAllValidatorOutstandingRewardsRequestBlockResults{
			Height:     i,
			Validators: make(map[string]sdk.DecCoins),
		}
		//nolint: contextcheck // false positive
		querier.IterateValidatorOutstandingRewards(queryContext, func(val sdk.ValAddress, rewards distributiontypes.ValidatorOutstandingRewards) (stop bool) {
			if len(rewards.Rewards) > 0 {
				blockResult.Validators[val.String()] = rewards.Rewards
			}

			return false
		})

		if len(blockResult.Validators) > 0 {
			result = append(result, blockResult)
		}
	}

	return result, nil
}

// GetValidatorSlashesByValidatorAddress queries slash events of a validator.
func (s *Server) GetValidatorSlashesByValidatorAddress(req *getValidatorSlashesByValidatorAddressRequest, r *http.Request) (resp any, err error) {
	queryContext, err := s.createQueryContextByHeader(r)
	if err != nil {
		return nil, err
	}

	queryResp, err := keeper.NewQuerier(s.store.GetDistrKeeper()).ValidatorSlashes(queryContext, &distributiontypes.QueryValidatorSlashesRequest{
		ValidatorAddress: mux.Vars(r)["validator_address"],
		StartingHeight:   req.StartingHeight,
		EndingHeight:     req.EndingHeight,
		Pagination: &query.PageRequest{
			Key:        []byte(req.Pagination.Key),
			Offset:     req.Pagination.Offset,
			Limit:      req.Pagination.Limit,
			CountTotal: req.Pagination.CountTotal,
			Reverse:    req.Pagination.Reverse,
		},
	})

	if err != nil {
		return nil, err
	}

	return queryResp, nil
}

// GetDistributionValidatorsByDelegatorAddress queries the validators of a delegator.
func (s *Server) GetDistributionValidatorsByDelegatorAddress(r *http.Request) (resp any, err error) {
	queryContext, err := s.createQueryContextByHeader(r)
	if err != nil {
		return nil, err
	}

	queryResp, err := keeper.NewQuerier(s.store.GetDistrKeeper()).DelegatorValidators(queryContext, &distributiontypes.QueryDelegatorValidatorsRequest{
		DelegatorAddress: mux.Vars(r)["delegator_address"],
	})

	if err != nil {
		return nil, err
	}

	return queryResp, nil
}

// GetDelegatorRewardsByDelegatorAddress queries the total rewards accrued by each validator.
func (s *Server) GetDelegatorRewardsByDelegatorAddress(r *http.Request) (resp any, err error) {
	queryContext, err := s.createQueryContextByHeader(r)
	if err != nil {
		return nil, err
	}

	queryResp, err := keeper.NewQuerier(s.store.GetDistrKeeper()).DelegationTotalRewards(queryContext, &distributiontypes.QueryDelegationTotalRewardsRequest{
		DelegatorAddress: mux.Vars(r)["delegator_address"],
	})

	if err != nil {
		return nil, err
	}

	return queryResp, nil
}

// GetDelegatorRewardsByDelegatorAddressValidatorAddress queries the total rewards accrued by a delegation.
func (s *Server) GetDelegatorRewardsByDelegatorAddressValidatorAddress(r *http.Request) (resp any, err error) {
	queryContext, err := s.createQueryContextByHeader(r)
	if err != nil {
		return nil, err
	}

	muxVars := mux.Vars(r)
	queryResp, err := keeper.NewQuerier(s.store.GetDistrKeeper()).DelegationRewards(queryContext, &distributiontypes.QueryDelegationRewardsRequest{
		DelegatorAddress: muxVars["delegator_address"],
		ValidatorAddress: muxVars["validator_address"],
	})

	if err != nil {
		return nil, err
	}

	return queryResp, nil
}

// GetDelegatorWithdrawAddressByDelegatorAddress queries withdraw address of a delegator.
func (s *Server) GetDelegatorWithdrawAddressByDelegatorAddress(r *http.Request) (resp any, err error) {
	queryContext, err := s.createQueryContextByHeader(r)
	if err != nil {
		return nil, err
	}

	queryResp, err := keeper.NewQuerier(s.store.GetDistrKeeper()).DelegatorWithdrawAddress(queryContext, &distributiontypes.QueryDelegatorWithdrawAddressRequest{
		DelegatorAddress: mux.Vars(r)["delegator_address"],
	})

	if err != nil {
		return nil, err
	}

	return queryResp, nil
}
