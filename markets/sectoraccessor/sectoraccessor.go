package sectoraccessor

import (
	"context"
	"fmt"
	"github.com/minio/minio-go"
	"io"
	"os"

	retrievalmarket_types "github.com/filecoin-project/boost/retrievalmarket/types"
	"github.com/filecoin-project/dagstore/mount"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	logging "github.com/ipfs/go-log/v2"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v1api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/markets/dagstore"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/storage/sealer"
	"github.com/filecoin-project/lotus/storage/sectorblocks"
)

var log = logging.Logger("sectoraccessor")

type sectorAccessor struct {
	minio      *minio.Client
	bucketName string
	maddr      address.Address
	secb       sectorblocks.SectorBuilder
	pp         sealer.PieceProvider
	full       v1api.FullNode
}

var _ retrievalmarket_types.SectorAccessor = (*sectorAccessor)(nil)

func NewSectorAccessor(maddr dtypes.MinerAddress, secb sectorblocks.SectorBuilder, pp sealer.PieceProvider, full v1api.FullNode) dagstore.SectorAccessor {
	endpoint, ok := os.LookupEnv("MINIO_ENDPOINT")
	if !ok {
		return nil
	}
	accessKeyID, ok := os.LookupEnv("MINIO_ACCESS_KEY_ID")
	if !ok {
		return nil
	}
	secretAccessKey, ok := os.LookupEnv("MINIO_SECRET_ACCESS_KEY")
	if !ok {
		return nil
	}
	bucketName, ok := os.LookupEnv("MINIO_BUCKET_NAME")
	if !ok {
		return nil
	}

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		log.Errorf("new minio error: %v", err)
	}
	return &sectorAccessor{minioClient, bucketName, address.Address(maddr), secb, pp, full}
}

func (sa *sectorAccessor) UnsealSector(ctx context.Context, sectorID abi.SectorNumber, pieceOffset abi.UnpaddedPieceSize, length abi.UnpaddedPieceSize) (io.ReadCloser, error) {
	return sa.UnsealSectorAt(ctx, sectorID, pieceOffset, length)
}

func (sa *sectorAccessor) UnsealSectorAt(ctx context.Context, sectorID abi.SectorNumber, pieceOffset abi.UnpaddedPieceSize, length abi.UnpaddedPieceSize) (mount.Reader, error) {
	log.Debugf("get sector %d, pieceOffset %d, length %d", sectorID, pieceOffset, length)

	si, err := sa.sectorsStatus(ctx, sectorID, false)
	if err != nil {
		return nil, err
	}

	piece := si.Pieces[0]
	if pieceOffset > 0 && len(si.Pieces) > 1 {
		piece = si.Pieces[1]
	}

	objectName := fmt.Sprintf("%s.car", piece.Piece.PieceCID.String())

	r := MinioMount{
		client:     sa.minio,
		bucketName: sa.bucketName,
		objectName: objectName,
	}
	return r.Fetch(ctx)
}

func (sa *sectorAccessor) IsUnsealed(ctx context.Context, sectorID abi.SectorNumber, offset abi.UnpaddedPieceSize, length abi.UnpaddedPieceSize) (bool, error) {
	return true, nil
}

func (sa *sectorAccessor) sectorsStatus(ctx context.Context, sid abi.SectorNumber, showOnChainInfo bool) (api.SectorInfo, error) {
	sInfo, err := sa.secb.SectorsStatus(ctx, sid, false)
	if err != nil {
		return api.SectorInfo{}, err
	}

	if !showOnChainInfo {
		return sInfo, nil
	}

	onChainInfo, err := sa.full.StateSectorGetInfo(ctx, sa.maddr, sid, types.EmptyTSK)
	if err != nil {
		return sInfo, err
	}
	if onChainInfo == nil {
		return sInfo, nil
	}
	sInfo.SealProof = onChainInfo.SealProof
	sInfo.Activation = onChainInfo.Activation
	sInfo.Expiration = onChainInfo.Expiration
	sInfo.DealWeight = onChainInfo.DealWeight
	sInfo.VerifiedDealWeight = onChainInfo.VerifiedDealWeight
	sInfo.InitialPledge = onChainInfo.InitialPledge

	ex, err := sa.full.StateSectorExpiration(ctx, sa.maddr, sid, types.EmptyTSK)
	if err != nil {
		return sInfo, nil
	}
	sInfo.OnTime = ex.OnTime
	sInfo.Early = ex.Early

	return sInfo, nil
}
