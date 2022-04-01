package positiontransferinstructionack

import (
	"time"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// PositionTransferInstructionAck is the fix50sp2 PositionTransferInstructionAck type, MsgType = DM.
type PositionTransferInstructionAck struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a PositionTransferInstructionAck from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) PositionTransferInstructionAck {
	return PositionTransferInstructionAck{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m PositionTransferInstructionAck) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a PositionTransferInstructionAck initialized with the required fields for PositionTransferInstructionAck.
func New(transferinstructionid field.TransferInstructionIDField) (m PositionTransferInstructionAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("DM"))
	m.Set(transferinstructionid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg PositionTransferInstructionAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "DM", r
}

// SetText sets Text, Tag 58.
func (m PositionTransferInstructionAck) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m PositionTransferInstructionAck) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m PositionTransferInstructionAck) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m PositionTransferInstructionAck) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m PositionTransferInstructionAck) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetRejectText sets RejectText, Tag 1328.
func (m PositionTransferInstructionAck) SetRejectText(v string) {
	m.Set(field.NewRejectText(v))
}

// SetNoTargetPartyIDs sets NoTargetPartyIDs, Tag 1461.
func (m PositionTransferInstructionAck) SetNoTargetPartyIDs(f NoTargetPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetEncodedRejectTextLen sets EncodedRejectTextLen, Tag 1664.
func (m PositionTransferInstructionAck) SetEncodedRejectTextLen(v int) {
	m.Set(field.NewEncodedRejectTextLen(v))
}

// SetEncodedRejectText sets EncodedRejectText, Tag 1665.
func (m PositionTransferInstructionAck) SetEncodedRejectText(v string) {
	m.Set(field.NewEncodedRejectText(v))
}

// SetTransferInstructionID sets TransferInstructionID, Tag 2436.
func (m PositionTransferInstructionAck) SetTransferInstructionID(v string) {
	m.Set(field.NewTransferInstructionID(v))
}

// SetTransferID sets TransferID, Tag 2437.
func (m PositionTransferInstructionAck) SetTransferID(v string) {
	m.Set(field.NewTransferID(v))
}

// SetTransferTransType sets TransferTransType, Tag 2439.
func (m PositionTransferInstructionAck) SetTransferTransType(v enum.TransferTransType) {
	m.Set(field.NewTransferTransType(v))
}

// SetTransferType sets TransferType, Tag 2440.
func (m PositionTransferInstructionAck) SetTransferType(v enum.TransferType) {
	m.Set(field.NewTransferType(v))
}

// SetTransferScope sets TransferScope, Tag 2441.
func (m PositionTransferInstructionAck) SetTransferScope(v enum.TransferScope) {
	m.Set(field.NewTransferScope(v))
}

// SetTransferStatus sets TransferStatus, Tag 2442.
func (m PositionTransferInstructionAck) SetTransferStatus(v enum.TransferStatus) {
	m.Set(field.NewTransferStatus(v))
}

// SetTransferRejectReason sets TransferRejectReason, Tag 2443.
func (m PositionTransferInstructionAck) SetTransferRejectReason(v enum.TransferRejectReason) {
	m.Set(field.NewTransferRejectReason(v))
}

// GetText gets Text, Tag 58.
func (m PositionTransferInstructionAck) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m PositionTransferInstructionAck) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m PositionTransferInstructionAck) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m PositionTransferInstructionAck) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m PositionTransferInstructionAck) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetRejectText gets RejectText, Tag 1328.
func (m PositionTransferInstructionAck) GetRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.RejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoTargetPartyIDs gets NoTargetPartyIDs, Tag 1461.
func (m PositionTransferInstructionAck) GetNoTargetPartyIDs() (NoTargetPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTargetPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetEncodedRejectTextLen gets EncodedRejectTextLen, Tag 1664.
func (m PositionTransferInstructionAck) GetEncodedRejectTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedRejectText gets EncodedRejectText, Tag 1665.
func (m PositionTransferInstructionAck) GetEncodedRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransferInstructionID gets TransferInstructionID, Tag 2436.
func (m PositionTransferInstructionAck) GetTransferInstructionID() (v string, err quickfix.MessageRejectError) {
	var f field.TransferInstructionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransferID gets TransferID, Tag 2437.
func (m PositionTransferInstructionAck) GetTransferID() (v string, err quickfix.MessageRejectError) {
	var f field.TransferIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransferTransType gets TransferTransType, Tag 2439.
func (m PositionTransferInstructionAck) GetTransferTransType() (v enum.TransferTransType, err quickfix.MessageRejectError) {
	var f field.TransferTransTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransferType gets TransferType, Tag 2440.
func (m PositionTransferInstructionAck) GetTransferType() (v enum.TransferType, err quickfix.MessageRejectError) {
	var f field.TransferTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransferScope gets TransferScope, Tag 2441.
func (m PositionTransferInstructionAck) GetTransferScope() (v enum.TransferScope, err quickfix.MessageRejectError) {
	var f field.TransferScopeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransferStatus gets TransferStatus, Tag 2442.
func (m PositionTransferInstructionAck) GetTransferStatus() (v enum.TransferStatus, err quickfix.MessageRejectError) {
	var f field.TransferStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransferRejectReason gets TransferRejectReason, Tag 2443.
func (m PositionTransferInstructionAck) GetTransferRejectReason() (v enum.TransferRejectReason, err quickfix.MessageRejectError) {
	var f field.TransferRejectReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasText returns true if Text is present, Tag 58.
func (m PositionTransferInstructionAck) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m PositionTransferInstructionAck) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m PositionTransferInstructionAck) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m PositionTransferInstructionAck) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m PositionTransferInstructionAck) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasRejectText returns true if RejectText is present, Tag 1328.
func (m PositionTransferInstructionAck) HasRejectText() bool {
	return m.Has(tag.RejectText)
}

// HasNoTargetPartyIDs returns true if NoTargetPartyIDs is present, Tag 1461.
func (m PositionTransferInstructionAck) HasNoTargetPartyIDs() bool {
	return m.Has(tag.NoTargetPartyIDs)
}

// HasEncodedRejectTextLen returns true if EncodedRejectTextLen is present, Tag 1664.
func (m PositionTransferInstructionAck) HasEncodedRejectTextLen() bool {
	return m.Has(tag.EncodedRejectTextLen)
}

// HasEncodedRejectText returns true if EncodedRejectText is present, Tag 1665.
func (m PositionTransferInstructionAck) HasEncodedRejectText() bool {
	return m.Has(tag.EncodedRejectText)
}

// HasTransferInstructionID returns true if TransferInstructionID is present, Tag 2436.
func (m PositionTransferInstructionAck) HasTransferInstructionID() bool {
	return m.Has(tag.TransferInstructionID)
}

// HasTransferID returns true if TransferID is present, Tag 2437.
func (m PositionTransferInstructionAck) HasTransferID() bool {
	return m.Has(tag.TransferID)
}

// HasTransferTransType returns true if TransferTransType is present, Tag 2439.
func (m PositionTransferInstructionAck) HasTransferTransType() bool {
	return m.Has(tag.TransferTransType)
}

// HasTransferType returns true if TransferType is present, Tag 2440.
func (m PositionTransferInstructionAck) HasTransferType() bool {
	return m.Has(tag.TransferType)
}

// HasTransferScope returns true if TransferScope is present, Tag 2441.
func (m PositionTransferInstructionAck) HasTransferScope() bool {
	return m.Has(tag.TransferScope)
}

// HasTransferStatus returns true if TransferStatus is present, Tag 2442.
func (m PositionTransferInstructionAck) HasTransferStatus() bool {
	return m.Has(tag.TransferStatus)
}

// HasTransferRejectReason returns true if TransferRejectReason is present, Tag 2443.
func (m PositionTransferInstructionAck) HasTransferRejectReason() bool {
	return m.Has(tag.TransferRejectReason)
}

// NoPartyIDs is a repeating group element, Tag 453.
type NoPartyIDs struct {
	*quickfix.Group
}

// SetPartyID sets PartyID, Tag 448.
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

// SetPartyIDSource sets PartyIDSource, Tag 447.
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

// SetPartyRole sets PartyRole, Tag 452.
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

// SetNoPartySubIDs sets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetPartyRoleQualifier sets PartyRoleQualifier, Tag 2376.
func (m NoPartyIDs) SetPartyRoleQualifier(v int) {
	m.Set(field.NewPartyRoleQualifier(v))
}

// GetPartyID gets PartyID, Tag 448.
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyIDSource gets PartyIDSource, Tag 447.
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyRole gets PartyRole, Tag 452.
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartySubIDs gets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetPartyRoleQualifier gets PartyRoleQualifier, Tag 2376.
func (m NoPartyIDs) GetPartyRoleQualifier() (v int, err quickfix.MessageRejectError) {
	var f field.PartyRoleQualifierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartyID returns true if PartyID is present, Tag 448.
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

// HasPartyIDSource returns true if PartyIDSource is present, Tag 447.
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

// HasPartyRole returns true if PartyRole is present, Tag 452.
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

// HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802.
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

// HasPartyRoleQualifier returns true if PartyRoleQualifier is present, Tag 2376.
func (m NoPartyIDs) HasPartyRoleQualifier() bool {
	return m.Has(tag.PartyRoleQualifier)
}

// NoPartySubIDs is a repeating group element, Tag 802.
type NoPartySubIDs struct {
	*quickfix.Group
}

// SetPartySubID sets PartySubID, Tag 523.
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

// SetPartySubIDType sets PartySubIDType, Tag 803.
func (m NoPartySubIDs) SetPartySubIDType(v enum.PartySubIDType) {
	m.Set(field.NewPartySubIDType(v))
}

// GetPartySubID gets PartySubID, Tag 523.
func (m NoPartySubIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartySubIDType gets PartySubIDType, Tag 803.
func (m NoPartySubIDs) GetPartySubIDType() (v enum.PartySubIDType, err quickfix.MessageRejectError) {
	var f field.PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartySubID returns true if PartySubID is present, Tag 523.
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

// HasPartySubIDType returns true if PartySubIDType is present, Tag 803.
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

// NoPartySubIDsRepeatingGroup is a repeating group, Tag 802.
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup.
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartySubIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PartySubID),
				quickfix.GroupElement(tag.PartySubIDType),
			},
		),
	}
}

// Add create and append a new NoPartySubIDs to this group.
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

// Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup.
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoPartyIDsRepeatingGroup is a repeating group, Tag 453.
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup.
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartyIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PartyID),
				quickfix.GroupElement(tag.PartyIDSource),
				quickfix.GroupElement(tag.PartyRole),
				NewNoPartySubIDsRepeatingGroup(),
				quickfix.GroupElement(tag.PartyRoleQualifier),
			},
		),
	}
}

// Add create and append a new NoPartyIDs to this group.
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

// Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup.
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

// NoTargetPartyIDs is a repeating group element, Tag 1461.
type NoTargetPartyIDs struct {
	*quickfix.Group
}

// SetTargetPartyID sets TargetPartyID, Tag 1462.
func (m NoTargetPartyIDs) SetTargetPartyID(v string) {
	m.Set(field.NewTargetPartyID(v))
}

// SetTargetPartyIDSource sets TargetPartyIDSource, Tag 1463.
func (m NoTargetPartyIDs) SetTargetPartyIDSource(v string) {
	m.Set(field.NewTargetPartyIDSource(v))
}

// SetTargetPartyRole sets TargetPartyRole, Tag 1464.
func (m NoTargetPartyIDs) SetTargetPartyRole(v int) {
	m.Set(field.NewTargetPartyRole(v))
}

// SetTargetPartyRoleQualifier sets TargetPartyRoleQualifier, Tag 1818.
func (m NoTargetPartyIDs) SetTargetPartyRoleQualifier(v int) {
	m.Set(field.NewTargetPartyRoleQualifier(v))
}

// SetNoTargetPartySubIDs sets NoTargetPartySubIDs, Tag 2433.
func (m NoTargetPartyIDs) SetNoTargetPartySubIDs(f NoTargetPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetTargetPartyID gets TargetPartyID, Tag 1462.
func (m NoTargetPartyIDs) GetTargetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.TargetPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTargetPartyIDSource gets TargetPartyIDSource, Tag 1463.
func (m NoTargetPartyIDs) GetTargetPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.TargetPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTargetPartyRole gets TargetPartyRole, Tag 1464.
func (m NoTargetPartyIDs) GetTargetPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.TargetPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTargetPartyRoleQualifier gets TargetPartyRoleQualifier, Tag 1818.
func (m NoTargetPartyIDs) GetTargetPartyRoleQualifier() (v int, err quickfix.MessageRejectError) {
	var f field.TargetPartyRoleQualifierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoTargetPartySubIDs gets NoTargetPartySubIDs, Tag 2433.
func (m NoTargetPartyIDs) GetNoTargetPartySubIDs() (NoTargetPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTargetPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasTargetPartyID returns true if TargetPartyID is present, Tag 1462.
func (m NoTargetPartyIDs) HasTargetPartyID() bool {
	return m.Has(tag.TargetPartyID)
}

// HasTargetPartyIDSource returns true if TargetPartyIDSource is present, Tag 1463.
func (m NoTargetPartyIDs) HasTargetPartyIDSource() bool {
	return m.Has(tag.TargetPartyIDSource)
}

// HasTargetPartyRole returns true if TargetPartyRole is present, Tag 1464.
func (m NoTargetPartyIDs) HasTargetPartyRole() bool {
	return m.Has(tag.TargetPartyRole)
}

// HasTargetPartyRoleQualifier returns true if TargetPartyRoleQualifier is present, Tag 1818.
func (m NoTargetPartyIDs) HasTargetPartyRoleQualifier() bool {
	return m.Has(tag.TargetPartyRoleQualifier)
}

// HasNoTargetPartySubIDs returns true if NoTargetPartySubIDs is present, Tag 2433.
func (m NoTargetPartyIDs) HasNoTargetPartySubIDs() bool {
	return m.Has(tag.NoTargetPartySubIDs)
}

// NoTargetPartySubIDs is a repeating group element, Tag 2433.
type NoTargetPartySubIDs struct {
	*quickfix.Group
}

// SetTargetPartySubID sets TargetPartySubID, Tag 2434.
func (m NoTargetPartySubIDs) SetTargetPartySubID(v string) {
	m.Set(field.NewTargetPartySubID(v))
}

// SetTargetPartySubIDType sets TargetPartySubIDType, Tag 2435.
func (m NoTargetPartySubIDs) SetTargetPartySubIDType(v int) {
	m.Set(field.NewTargetPartySubIDType(v))
}

// GetTargetPartySubID gets TargetPartySubID, Tag 2434.
func (m NoTargetPartySubIDs) GetTargetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.TargetPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTargetPartySubIDType gets TargetPartySubIDType, Tag 2435.
func (m NoTargetPartySubIDs) GetTargetPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.TargetPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasTargetPartySubID returns true if TargetPartySubID is present, Tag 2434.
func (m NoTargetPartySubIDs) HasTargetPartySubID() bool {
	return m.Has(tag.TargetPartySubID)
}

// HasTargetPartySubIDType returns true if TargetPartySubIDType is present, Tag 2435.
func (m NoTargetPartySubIDs) HasTargetPartySubIDType() bool {
	return m.Has(tag.TargetPartySubIDType)
}

// NoTargetPartySubIDsRepeatingGroup is a repeating group, Tag 2433.
type NoTargetPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoTargetPartySubIDsRepeatingGroup returns an initialized, NoTargetPartySubIDsRepeatingGroup.
func NewNoTargetPartySubIDsRepeatingGroup() NoTargetPartySubIDsRepeatingGroup {
	return NoTargetPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoTargetPartySubIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.TargetPartySubID),
				quickfix.GroupElement(tag.TargetPartySubIDType),
			},
		),
	}
}

// Add create and append a new NoTargetPartySubIDs to this group.
func (m NoTargetPartySubIDsRepeatingGroup) Add() NoTargetPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoTargetPartySubIDs{g}
}

// Get returns the ith NoTargetPartySubIDs in the NoTargetPartySubIDsRepeatinGroup.
func (m NoTargetPartySubIDsRepeatingGroup) Get(i int) NoTargetPartySubIDs {
	return NoTargetPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoTargetPartyIDsRepeatingGroup is a repeating group, Tag 1461.
type NoTargetPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoTargetPartyIDsRepeatingGroup returns an initialized, NoTargetPartyIDsRepeatingGroup.
func NewNoTargetPartyIDsRepeatingGroup() NoTargetPartyIDsRepeatingGroup {
	return NoTargetPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoTargetPartyIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.TargetPartyID),
				quickfix.GroupElement(tag.TargetPartyIDSource),
				quickfix.GroupElement(tag.TargetPartyRole),
				quickfix.GroupElement(tag.TargetPartyRoleQualifier),
				NewNoTargetPartySubIDsRepeatingGroup(),
			},
		),
	}
}

// Add create and append a new NoTargetPartyIDs to this group.
func (m NoTargetPartyIDsRepeatingGroup) Add() NoTargetPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoTargetPartyIDs{g}
}

// Get returns the ith NoTargetPartyIDs in the NoTargetPartyIDsRepeatinGroup.
func (m NoTargetPartyIDsRepeatingGroup) Get(i int) NoTargetPartyIDs {
	return NoTargetPartyIDs{m.RepeatingGroup.Get(i)}
}
