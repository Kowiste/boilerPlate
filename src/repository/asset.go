package repository

import (
	"boiler/src/model/asset"
	"context"
)

type IAssetRepository interface {
	CreateAsset(c context.Context, asset *asset.Asset) (string, error)
	Assets(c context.Context) (asset.Assets, error)
	AssetByID(c context.Context, id string) (*asset.Asset, error)
	UpdateAsset(c context.Context,asset *asset.Asset) error
	DeleteAsset(c context.Context,id  string) error
}