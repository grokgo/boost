// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package migrations

import (
	"fmt"
	"io"
	"math"
	"sort"

	datatransfer "github.com/filecoin-project/boost/datatransfer"
	legacytypes "github.com/filecoin-project/boost/storagemarket/types/legacytypes"
	filestore "github.com/filecoin-project/boost/storagemarket/types/legacytypes/filestore"
	abi "github.com/filecoin-project/go-state-types/abi"
	market "github.com/filecoin-project/specs-actors/actors/builtin/market"
	cid "github.com/ipfs/go-cid"
	peer "github.com/libp2p/go-libp2p/core/peer"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *Proposal1) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{163}); err != nil {
		return err
	}

	// t.Piece (legacytypes.DataRef) (struct)
	if len("Piece") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Piece\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Piece"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Piece")); err != nil {
		return err
	}

	if err := t.Piece.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.DealProposal (market.ClientDealProposal) (struct)
	if len("DealProposal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealProposal\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("DealProposal"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("DealProposal")); err != nil {
		return err
	}

	if err := t.DealProposal.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.FastRetrieval (bool) (bool)
	if len("FastRetrieval") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"FastRetrieval\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("FastRetrieval"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("FastRetrieval")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.FastRetrieval); err != nil {
		return err
	}
	return nil
}

func (t *Proposal1) UnmarshalCBOR(r io.Reader) (err error) {
	*t = Proposal1{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("Proposal1: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Piece (legacytypes.DataRef) (struct)
		case "Piece":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Piece = new(legacytypes.DataRef)
					if err := t.Piece.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Piece pointer: %w", err)
					}
				}

			}
			// t.DealProposal (market.ClientDealProposal) (struct)
		case "DealProposal":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.DealProposal = new(market.ClientDealProposal)
					if err := t.DealProposal.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.DealProposal pointer: %w", err)
					}
				}

			}
			// t.FastRetrieval (bool) (bool)
		case "FastRetrieval":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.FastRetrieval = false
			case 21:
				t.FastRetrieval = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *MinerDeal1) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{180}); err != nil {
		return err
	}

	// t.Ref (legacytypes.DataRef) (struct)
	if len("Ref") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Ref\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Ref"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Ref")); err != nil {
		return err
	}

	if err := t.Ref.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Miner (peer.ID) (string)
	if len("Miner") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Miner\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Miner"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Miner")); err != nil {
		return err
	}

	if len(t.Miner) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Miner was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Miner))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Miner)); err != nil {
		return err
	}

	// t.State (uint64) (uint64)
	if len("State") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"State\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("State"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("State")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.State)); err != nil {
		return err
	}

	// t.Client (peer.ID) (string)
	if len("Client") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Client\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Client"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Client")); err != nil {
		return err
	}

	if len(t.Client) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Client was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Client))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Client)); err != nil {
		return err
	}

	// t.DealID (abi.DealID) (uint64)
	if len("DealID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealID\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("DealID"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("DealID")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.DealID)); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len("Message") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Message\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Message"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Message")); err != nil {
		return err
	}

	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Message)); err != nil {
		return err
	}

	// t.PiecePath (filestore.Path) (string)
	if len("PiecePath") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PiecePath\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PiecePath"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("PiecePath")); err != nil {
		return err
	}

	if len(t.PiecePath) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.PiecePath was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.PiecePath))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.PiecePath)); err != nil {
		return err
	}

	// t.InboundCAR (string) (string)
	if len("InboundCAR") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"InboundCAR\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("InboundCAR"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("InboundCAR")); err != nil {
		return err
	}

	if len(t.InboundCAR) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.InboundCAR was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.InboundCAR))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.InboundCAR)); err != nil {
		return err
	}

	// t.PublishCid (cid.Cid) (struct)
	if len("PublishCid") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PublishCid\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PublishCid"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("PublishCid")); err != nil {
		return err
	}

	if t.PublishCid == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.PublishCid); err != nil {
			return xerrors.Errorf("failed to write cid field t.PublishCid: %w", err)
		}
	}

	// t.SlashEpoch (abi.ChainEpoch) (int64)
	if len("SlashEpoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"SlashEpoch\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("SlashEpoch"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("SlashEpoch")); err != nil {
		return err
	}

	if t.SlashEpoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.SlashEpoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.SlashEpoch-1)); err != nil {
			return err
		}
	}

	// t.AddFundsCid (cid.Cid) (struct)
	if len("AddFundsCid") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"AddFundsCid\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("AddFundsCid"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("AddFundsCid")); err != nil {
		return err
	}

	if t.AddFundsCid == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.AddFundsCid); err != nil {
			return xerrors.Errorf("failed to write cid field t.AddFundsCid: %w", err)
		}
	}

	// t.ProposalCid (cid.Cid) (struct)
	if len("ProposalCid") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ProposalCid\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ProposalCid"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("ProposalCid")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.ProposalCid); err != nil {
		return xerrors.Errorf("failed to write cid field t.ProposalCid: %w", err)
	}

	// t.CreationTime (typegen.CborTime) (struct)
	if len("CreationTime") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CreationTime\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("CreationTime"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("CreationTime")); err != nil {
		return err
	}

	if err := t.CreationTime.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.MetadataPath (filestore.Path) (string)
	if len("MetadataPath") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"MetadataPath\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("MetadataPath"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("MetadataPath")); err != nil {
		return err
	}

	if len(t.MetadataPath) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.MetadataPath was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.MetadataPath))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.MetadataPath)); err != nil {
		return err
	}

	// t.SectorNumber (abi.SectorNumber) (uint64)
	if len("SectorNumber") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"SectorNumber\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("SectorNumber"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("SectorNumber")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.SectorNumber)); err != nil {
		return err
	}

	// t.FastRetrieval (bool) (bool)
	if len("FastRetrieval") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"FastRetrieval\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("FastRetrieval"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("FastRetrieval")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.FastRetrieval); err != nil {
		return err
	}

	// t.FundsReserved (big.Int) (struct)
	if len("FundsReserved") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"FundsReserved\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("FundsReserved"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("FundsReserved")); err != nil {
		return err
	}

	if err := t.FundsReserved.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.TransferChannelId (datatransfer.ChannelID) (struct)
	if len("TransferChannelId") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TransferChannelId\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("TransferChannelId"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("TransferChannelId")); err != nil {
		return err
	}

	if err := t.TransferChannelId.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ClientDealProposal (market.ClientDealProposal) (struct)
	if len("ClientDealProposal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ClientDealProposal\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ClientDealProposal"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("ClientDealProposal")); err != nil {
		return err
	}

	if err := t.ClientDealProposal.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.AvailableForRetrieval (bool) (bool)
	if len("AvailableForRetrieval") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"AvailableForRetrieval\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("AvailableForRetrieval"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("AvailableForRetrieval")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.AvailableForRetrieval); err != nil {
		return err
	}
	return nil
}

func (t *MinerDeal1) UnmarshalCBOR(r io.Reader) (err error) {
	*t = MinerDeal1{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("MinerDeal1: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Ref (legacytypes.DataRef) (struct)
		case "Ref":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Ref = new(legacytypes.DataRef)
					if err := t.Ref.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Ref pointer: %w", err)
					}
				}

			}
			// t.Miner (peer.ID) (string)
		case "Miner":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.Miner = peer.ID(sval)
			}
			// t.State (uint64) (uint64)
		case "State":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.State = uint64(extra)

			}
			// t.Client (peer.ID) (string)
		case "Client":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.Client = peer.ID(sval)
			}
			// t.DealID (abi.DealID) (uint64)
		case "DealID":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.DealID = abi.DealID(extra)

			}
			// t.Message (string) (string)
		case "Message":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.Message = string(sval)
			}
			// t.PiecePath (filestore.Path) (string)
		case "PiecePath":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.PiecePath = filestore.Path(sval)
			}
			// t.InboundCAR (string) (string)
		case "InboundCAR":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.InboundCAR = string(sval)
			}
			// t.PublishCid (cid.Cid) (struct)
		case "PublishCid":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.PublishCid: %w", err)
					}

					t.PublishCid = &c
				}

			}
			// t.SlashEpoch (abi.ChainEpoch) (int64)
		case "SlashEpoch":
			{
				maj, extra, err := cr.ReadHeader()
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.SlashEpoch = abi.ChainEpoch(extraI)
			}
			// t.AddFundsCid (cid.Cid) (struct)
		case "AddFundsCid":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.AddFundsCid: %w", err)
					}

					t.AddFundsCid = &c
				}

			}
			// t.ProposalCid (cid.Cid) (struct)
		case "ProposalCid":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.ProposalCid: %w", err)
				}

				t.ProposalCid = c

			}
			// t.CreationTime (typegen.CborTime) (struct)
		case "CreationTime":

			{

				if err := t.CreationTime.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.CreationTime: %w", err)
				}

			}
			// t.MetadataPath (filestore.Path) (string)
		case "MetadataPath":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.MetadataPath = filestore.Path(sval)
			}
			// t.SectorNumber (abi.SectorNumber) (uint64)
		case "SectorNumber":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.SectorNumber = abi.SectorNumber(extra)

			}
			// t.FastRetrieval (bool) (bool)
		case "FastRetrieval":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.FastRetrieval = false
			case 21:
				t.FastRetrieval = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.FundsReserved (big.Int) (struct)
		case "FundsReserved":

			{

				if err := t.FundsReserved.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.FundsReserved: %w", err)
				}

			}
			// t.TransferChannelId (datatransfer.ChannelID) (struct)
		case "TransferChannelId":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.TransferChannelId = new(datatransfer.ChannelID)
					if err := t.TransferChannelId.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.TransferChannelId pointer: %w", err)
					}
				}

			}
			// t.ClientDealProposal (market.ClientDealProposal) (struct)
		case "ClientDealProposal":

			{

				if err := t.ClientDealProposal.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.ClientDealProposal: %w", err)
				}

			}
			// t.AvailableForRetrieval (bool) (bool)
		case "AvailableForRetrieval":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.AvailableForRetrieval = false
			case 21:
				t.AvailableForRetrieval = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
