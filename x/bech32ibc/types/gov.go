package types

import (
	"fmt"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"strings"
)

const (
	ProposalTypeUpdateHrpIbcChannel = "UpdateHrpIbcChannel"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeUpdateHrpIbcChannel)
}

var _ govtypes.Content = &UpdateHrpIbcChannelProposal{}

func NewUpdateHrpIBCRecordProposal(title, description, hrp, fungibleSourceChannel string, nftSourceChannel string) govtypes.Content {
	return &UpdateHrpIbcChannelProposal{
		Title:                 title,
		Description:           description,
		Hrp:                   hrp,
		FungibleSourceChannel: fungibleSourceChannel,
		NftSourceChannel:      nftSourceChannel,
	}
}

func (p *UpdateHrpIbcChannelProposal) GetTitle() string { return p.Title }

func (p *UpdateHrpIbcChannelProposal) GetDescription() string { return p.Description }

func (p *UpdateHrpIbcChannelProposal) ProposalRoute() string { return RouterKey }

func (p *UpdateHrpIbcChannelProposal) ProposalType() string {
	return ProposalTypeUpdateHrpIbcChannel
}

func (p *UpdateHrpIbcChannelProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}
	return ValidateHrp(p.Hrp)
}

func (p UpdateHrpIbcChannelProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Update HRP IBC Channel Proposal:
  Title:                   %s
  Description:             %s
  HRP:                     %s
  Fungible Source Channel: %s
  NFT Source Chanell:      %s
`, p.Title, p.Description, p.Hrp, p.FungibleSourceChannel, p.NftSourceChannel))
	return b.String()
}
