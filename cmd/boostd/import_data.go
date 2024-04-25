package main

import (
	"errors"
	"fmt"
	"github.com/filecoin-project/boost/storagemarket/types/dealcheckpoints"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
	"os"
	"path/filepath"
	"strings"

	bcli "github.com/filecoin-project/boost/cli"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
)

var importDataCmd = &cli.Command{
	Name:      "import-data",
	Usage:     "Import data for offline deal made with Boost",
	ArgsUsage: "<proposal CID> <file> or <deal UUID> <file>",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "delete-after-import",
			Usage: "whether to delete the data for the offline deal after the deal has been added to a sector",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "remote",
			Usage: "is it a remote file",
			Value: true,
		},
		&cli.StringFlag{
			Name:  "remote-path",
			Usage: "remote file download path",
		},
		&cli.StringFlag{
			Name:     "local-path",
			Usage:    "local file path",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() < 2 {
			return fmt.Errorf("must specify proposal CID / deal UUID and file path")
		}

		id := cctx.Args().Get(0)
		fileName := cctx.Args().Get(1)
		localPath := cctx.String("local-path")
		if strings.HasSuffix(localPath, "/") {
			localPath = fmt.Sprintf("%s%s", localPath, fileName)
		} else {
			localPath = fmt.Sprintf("%s/%s", localPath, fileName)
		}

		// Parse the first parameter as a deal UUID or a proposal CID
		var proposalCid *cid.Cid
		dealUuid, err := uuid.Parse(id)
		if err != nil {
			propCid, err := cid.Decode(id)
			if err != nil {
				return fmt.Errorf("could not parse '%s' as deal uuid or proposal cid", id)
			}
			proposalCid = &propCid
		}

		napi, closer, err := bcli.GetBoostAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		pds, err := napi.BoostDeal(cctx.Context, dealUuid)
		if err != nil {
			return err
		}

		if pds.Checkpoint != dealcheckpoints.Accepted {
			return xerrors.Errorf("the order %s has been imported", dealUuid.String())
		}

		remote := cctx.Bool("remote")
		localFileExists, _ := pathExists(localPath)
		if !localFileExists {
			if remote {
				remotePath := cctx.String("remote-path")
				if remotePath == "" {
					return errors.New("remote-path not empty")
				}
				if strings.HasSuffix(remotePath, "/") {
					remotePath = fmt.Sprintf("%s%s", remotePath, fileName)
				} else {
					remotePath = fmt.Sprintf("%s/%s", remotePath, fileName)
				}
				if err := downloadFile(localPath, remotePath); err != nil {
					_ = os.RemoveAll(localPath)
					return err
				}
			} else {
				return errors.New("local file does not exist")
			}
		}

		path, err := homedir.Expand(localPath)
		if err != nil {
			return fmt.Errorf("expanding file path: %w", err)
		}

		filePath, err := filepath.Abs(path)
		if err != nil {
			return fmt.Errorf("failed get absolute path for file: %w", err)
		}

		_, err = os.Stat(filePath)
		if err != nil {
			return fmt.Errorf("opening file %s: %w", filePath, err)
		}

		// If the user has supplied a signed proposal cid
		deleteAfterImport := cctx.Bool("delete-after-import")
		if proposalCid != nil {

			// Look up the deal in the boost database
			deal, err := napi.BoostDealBySignedProposalCid(cctx.Context, *proposalCid)
			if err != nil {
				// If the error is anything other than a Not Found error,
				// return the error
				if !strings.Contains(err.Error(), "not found") {
					return err
				}

				return fmt.Errorf("cannot find boost deal with proposal cid %s and legacy deals are no olnger supported", proposalCid)
			}

			// Get the deal UUID from the deal
			dealUuid = deal.DealUuid
		}

		// Deal proposal by deal uuid (v1.2.0 deal)
		rej, err := napi.BoostOfflineDealWithData(cctx.Context, dealUuid, filePath, deleteAfterImport)
		if err != nil {
			return fmt.Errorf("failed to execute offline deal: %w", err)
		}
		if rej != nil && rej.Reason != "" {
			return fmt.Errorf("offline deal %s rejected: %s", dealUuid, rej.Reason)
		}
		fmt.Printf("Offline deal import for v1.2.0 deal %s scheduled for execution\n", dealUuid)
		return nil
	},
}
