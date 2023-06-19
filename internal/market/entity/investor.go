package entity

type Investor struct {
	ID           string
	Name         string
	AssetPostion []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:           id,
		AssetPostion: []*InvestorAssetPosition{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPostion = append(i.AssetPostion, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetID string, qtdShares int) {
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		i.AssetPostion = append(i.AssetPostion, NewInvestorAssetPosition(assetID, qtdShares))
	} else {
		assetPosition.Shares += qtdShares
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, assetPosition := range i.AssetPostion {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil
}

type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}

func NewInvestorAssetPosition(assetId string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetId,
		Shares:  shares,
	}
}
