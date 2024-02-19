// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package migrations

import (
	"fmt"
	"io"
	"math"
	"sort"

	abi "github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufPieceInfo0 = []byte{130}

func (t *PieceInfo0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufPieceInfo0); err != nil {
		return err
	}

	// t.PieceCID (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.PieceCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
	}

	// t.Deals ([]migrations.DealInfo0) (slice)
	if len(t.Deals) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Deals was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Deals))); err != nil {
		return err
	}
	for _, v := range t.Deals {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}

	}
	return nil
}

func (t *PieceInfo0) UnmarshalCBOR(r io.Reader) (err error) {
	*t = PieceInfo0{}

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

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PieceCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
		}

		t.PieceCID = c

	}
	// t.Deals ([]migrations.DealInfo0) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Deals: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Deals = make([]DealInfo0, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			{

				if err := t.Deals[i].UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.Deals[i]: %w", err)
				}

			}

		}
	}
	return nil
}

var lengthBufDealInfo0 = []byte{132}

func (t *DealInfo0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufDealInfo0); err != nil {
		return err
	}

	// t.DealID (abi.DealID) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.DealID)); err != nil {
		return err
	}

	// t.SectorID (abi.SectorNumber) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.SectorID)); err != nil {
		return err
	}

	// t.Offset (abi.PaddedPieceSize) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Offset)); err != nil {
		return err
	}

	// t.Length (abi.PaddedPieceSize) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Length)); err != nil {
		return err
	}

	return nil
}

func (t *DealInfo0) UnmarshalCBOR(r io.Reader) (err error) {
	*t = DealInfo0{}

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

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.DealID (abi.DealID) (uint64)

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
	// t.SectorID (abi.SectorNumber) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.SectorID = abi.SectorNumber(extra)

	}
	// t.Offset (abi.PaddedPieceSize) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Offset = abi.PaddedPieceSize(extra)

	}
	// t.Length (abi.PaddedPieceSize) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Length = abi.PaddedPieceSize(extra)

	}
	return nil
}

var lengthBufBlockLocation0 = []byte{130}

func (t *BlockLocation0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufBlockLocation0); err != nil {
		return err
	}

	// t.RelOffset (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.RelOffset)); err != nil {
		return err
	}

	// t.BlockSize (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.BlockSize)); err != nil {
		return err
	}

	return nil
}

func (t *BlockLocation0) UnmarshalCBOR(r io.Reader) (err error) {
	*t = BlockLocation0{}

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

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.RelOffset (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.RelOffset = uint64(extra)

	}
	// t.BlockSize (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.BlockSize = uint64(extra)

	}
	return nil
}

var lengthBufPieceBlockLocation0 = []byte{130}

func (t *PieceBlockLocation0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufPieceBlockLocation0); err != nil {
		return err
	}

	// t.BlockLocation0 (migrations.BlockLocation0) (struct)
	if err := t.BlockLocation0.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.PieceCID (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.PieceCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
	}

	return nil
}

func (t *PieceBlockLocation0) UnmarshalCBOR(r io.Reader) (err error) {
	*t = PieceBlockLocation0{}

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

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.BlockLocation0 (migrations.BlockLocation0) (struct)

	{

		if err := t.BlockLocation0.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.BlockLocation0: %w", err)
		}

	}
	// t.PieceCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
		}

		t.PieceCID = c

	}
	return nil
}

var lengthBufCIDInfo0 = []byte{130}

func (t *CIDInfo0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufCIDInfo0); err != nil {
		return err
	}

	// t.CID (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.CID); err != nil {
		return xerrors.Errorf("failed to write cid field t.CID: %w", err)
	}

	// t.PieceBlockLocations ([]migrations.PieceBlockLocation0) (slice)
	if len(t.PieceBlockLocations) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.PieceBlockLocations was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.PieceBlockLocations))); err != nil {
		return err
	}
	for _, v := range t.PieceBlockLocations {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}

	}
	return nil
}

func (t *CIDInfo0) UnmarshalCBOR(r io.Reader) (err error) {
	*t = CIDInfo0{}

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

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.CID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.CID: %w", err)
		}

		t.CID = c

	}
	// t.PieceBlockLocations ([]migrations.PieceBlockLocation0) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.PieceBlockLocations: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.PieceBlockLocations = make([]PieceBlockLocation0, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			{

				if err := t.PieceBlockLocations[i].UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.PieceBlockLocations[i]: %w", err)
				}

			}

		}
	}
	return nil
}
