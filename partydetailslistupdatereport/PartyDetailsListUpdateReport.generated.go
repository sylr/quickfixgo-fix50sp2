package partydetailslistupdatereport

import (
	"time"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// PartyDetailsListUpdateReport is the fix50sp2 PartyDetailsListUpdateReport type, MsgType = CK.
type PartyDetailsListUpdateReport struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a PartyDetailsListUpdateReport from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) PartyDetailsListUpdateReport {
	return PartyDetailsListUpdateReport{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m PartyDetailsListUpdateReport) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a PartyDetailsListUpdateReport initialized with the required fields for PartyDetailsListUpdateReport.
func New(partydetailslistreportid field.PartyDetailsListReportIDField) (m PartyDetailsListUpdateReport) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("CK"))
	m.Set(partydetailslistreportid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg PartyDetailsListUpdateReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "CK", r
}

// SetText sets Text, Tag 58.
func (m PartyDetailsListUpdateReport) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m PartyDetailsListUpdateReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m PartyDetailsListUpdateReport) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m PartyDetailsListUpdateReport) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetLastFragment sets LastFragment, Tag 893.
func (m PartyDetailsListUpdateReport) SetLastFragment(v bool) {
	m.Set(field.NewLastFragment(v))
}

// SetApplID sets ApplID, Tag 1180.
func (m PartyDetailsListUpdateReport) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

// SetApplSeqNum sets ApplSeqNum, Tag 1181.
func (m PartyDetailsListUpdateReport) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

// SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350.
func (m PartyDetailsListUpdateReport) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

// SetApplResendFlag sets ApplResendFlag, Tag 1352.
func (m PartyDetailsListUpdateReport) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

// SetPartyDetailsListRequestID sets PartyDetailsListRequestID, Tag 1505.
func (m PartyDetailsListUpdateReport) SetPartyDetailsListRequestID(v string) {
	m.Set(field.NewPartyDetailsListRequestID(v))
}

// SetPartyDetailsListReportID sets PartyDetailsListReportID, Tag 1510.
func (m PartyDetailsListUpdateReport) SetPartyDetailsListReportID(v string) {
	m.Set(field.NewPartyDetailsListReportID(v))
}

// SetTotNoParties sets TotNoParties, Tag 1512.
func (m PartyDetailsListUpdateReport) SetTotNoParties(v int) {
	m.Set(field.NewTotNoParties(v))
}

// SetNoRequestingPartyIDs sets NoRequestingPartyIDs, Tag 1657.
func (m PartyDetailsListUpdateReport) SetNoRequestingPartyIDs(f NoRequestingPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoPartyUpdates sets NoPartyUpdates, Tag 1676.
func (m PartyDetailsListUpdateReport) SetNoPartyUpdates(f NoPartyUpdatesRepeatingGroup) {
	m.SetGroup(f)
}

// GetText gets Text, Tag 58.
func (m PartyDetailsListUpdateReport) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m PartyDetailsListUpdateReport) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m PartyDetailsListUpdateReport) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m PartyDetailsListUpdateReport) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLastFragment gets LastFragment, Tag 893.
func (m PartyDetailsListUpdateReport) GetLastFragment() (v bool, err quickfix.MessageRejectError) {
	var f field.LastFragmentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplID gets ApplID, Tag 1180.
func (m PartyDetailsListUpdateReport) GetApplID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplSeqNum gets ApplSeqNum, Tag 1181.
func (m PartyDetailsListUpdateReport) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350.
func (m PartyDetailsListUpdateReport) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplResendFlag gets ApplResendFlag, Tag 1352.
func (m PartyDetailsListUpdateReport) GetApplResendFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplResendFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyDetailsListRequestID gets PartyDetailsListRequestID, Tag 1505.
func (m PartyDetailsListUpdateReport) GetPartyDetailsListRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyDetailsListRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyDetailsListReportID gets PartyDetailsListReportID, Tag 1510.
func (m PartyDetailsListUpdateReport) GetPartyDetailsListReportID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyDetailsListReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotNoParties gets TotNoParties, Tag 1512.
func (m PartyDetailsListUpdateReport) GetTotNoParties() (v int, err quickfix.MessageRejectError) {
	var f field.TotNoPartiesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRequestingPartyIDs gets NoRequestingPartyIDs, Tag 1657.
func (m PartyDetailsListUpdateReport) GetNoRequestingPartyIDs() (NoRequestingPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRequestingPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoPartyUpdates gets NoPartyUpdates, Tag 1676.
func (m PartyDetailsListUpdateReport) GetNoPartyUpdates() (NoPartyUpdatesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyUpdatesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasText returns true if Text is present, Tag 58.
func (m PartyDetailsListUpdateReport) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m PartyDetailsListUpdateReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m PartyDetailsListUpdateReport) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m PartyDetailsListUpdateReport) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasLastFragment returns true if LastFragment is present, Tag 893.
func (m PartyDetailsListUpdateReport) HasLastFragment() bool {
	return m.Has(tag.LastFragment)
}

// HasApplID returns true if ApplID is present, Tag 1180.
func (m PartyDetailsListUpdateReport) HasApplID() bool {
	return m.Has(tag.ApplID)
}

// HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181.
func (m PartyDetailsListUpdateReport) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

// HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350.
func (m PartyDetailsListUpdateReport) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

// HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352.
func (m PartyDetailsListUpdateReport) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

// HasPartyDetailsListRequestID returns true if PartyDetailsListRequestID is present, Tag 1505.
func (m PartyDetailsListUpdateReport) HasPartyDetailsListRequestID() bool {
	return m.Has(tag.PartyDetailsListRequestID)
}

// HasPartyDetailsListReportID returns true if PartyDetailsListReportID is present, Tag 1510.
func (m PartyDetailsListUpdateReport) HasPartyDetailsListReportID() bool {
	return m.Has(tag.PartyDetailsListReportID)
}

// HasTotNoParties returns true if TotNoParties is present, Tag 1512.
func (m PartyDetailsListUpdateReport) HasTotNoParties() bool {
	return m.Has(tag.TotNoParties)
}

// HasNoRequestingPartyIDs returns true if NoRequestingPartyIDs is present, Tag 1657.
func (m PartyDetailsListUpdateReport) HasNoRequestingPartyIDs() bool {
	return m.Has(tag.NoRequestingPartyIDs)
}

// HasNoPartyUpdates returns true if NoPartyUpdates is present, Tag 1676.
func (m PartyDetailsListUpdateReport) HasNoPartyUpdates() bool {
	return m.Has(tag.NoPartyUpdates)
}

// NoRequestingPartyIDs is a repeating group element, Tag 1657.
type NoRequestingPartyIDs struct {
	*quickfix.Group
}

// SetRequestingPartyID sets RequestingPartyID, Tag 1658.
func (m NoRequestingPartyIDs) SetRequestingPartyID(v string) {
	m.Set(field.NewRequestingPartyID(v))
}

// SetRequestingPartyIDSource sets RequestingPartyIDSource, Tag 1659.
func (m NoRequestingPartyIDs) SetRequestingPartyIDSource(v string) {
	m.Set(field.NewRequestingPartyIDSource(v))
}

// SetRequestingPartyRole sets RequestingPartyRole, Tag 1660.
func (m NoRequestingPartyIDs) SetRequestingPartyRole(v int) {
	m.Set(field.NewRequestingPartyRole(v))
}

// SetNoRequestingPartySubIDs sets NoRequestingPartySubIDs, Tag 1661.
func (m NoRequestingPartyIDs) SetNoRequestingPartySubIDs(f NoRequestingPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetRequestingPartyRoleQualifier sets RequestingPartyRoleQualifier, Tag 2338.
func (m NoRequestingPartyIDs) SetRequestingPartyRoleQualifier(v int) {
	m.Set(field.NewRequestingPartyRoleQualifier(v))
}

// GetRequestingPartyID gets RequestingPartyID, Tag 1658.
func (m NoRequestingPartyIDs) GetRequestingPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.RequestingPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRequestingPartyIDSource gets RequestingPartyIDSource, Tag 1659.
func (m NoRequestingPartyIDs) GetRequestingPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RequestingPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRequestingPartyRole gets RequestingPartyRole, Tag 1660.
func (m NoRequestingPartyIDs) GetRequestingPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.RequestingPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRequestingPartySubIDs gets NoRequestingPartySubIDs, Tag 1661.
func (m NoRequestingPartyIDs) GetNoRequestingPartySubIDs() (NoRequestingPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRequestingPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetRequestingPartyRoleQualifier gets RequestingPartyRoleQualifier, Tag 2338.
func (m NoRequestingPartyIDs) GetRequestingPartyRoleQualifier() (v int, err quickfix.MessageRejectError) {
	var f field.RequestingPartyRoleQualifierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRequestingPartyID returns true if RequestingPartyID is present, Tag 1658.
func (m NoRequestingPartyIDs) HasRequestingPartyID() bool {
	return m.Has(tag.RequestingPartyID)
}

// HasRequestingPartyIDSource returns true if RequestingPartyIDSource is present, Tag 1659.
func (m NoRequestingPartyIDs) HasRequestingPartyIDSource() bool {
	return m.Has(tag.RequestingPartyIDSource)
}

// HasRequestingPartyRole returns true if RequestingPartyRole is present, Tag 1660.
func (m NoRequestingPartyIDs) HasRequestingPartyRole() bool {
	return m.Has(tag.RequestingPartyRole)
}

// HasNoRequestingPartySubIDs returns true if NoRequestingPartySubIDs is present, Tag 1661.
func (m NoRequestingPartyIDs) HasNoRequestingPartySubIDs() bool {
	return m.Has(tag.NoRequestingPartySubIDs)
}

// HasRequestingPartyRoleQualifier returns true if RequestingPartyRoleQualifier is present, Tag 2338.
func (m NoRequestingPartyIDs) HasRequestingPartyRoleQualifier() bool {
	return m.Has(tag.RequestingPartyRoleQualifier)
}

// NoRequestingPartySubIDs is a repeating group element, Tag 1661.
type NoRequestingPartySubIDs struct {
	*quickfix.Group
}

// SetRequestingPartySubID sets RequestingPartySubID, Tag 1662.
func (m NoRequestingPartySubIDs) SetRequestingPartySubID(v string) {
	m.Set(field.NewRequestingPartySubID(v))
}

// SetRequestingPartySubIDType sets RequestingPartySubIDType, Tag 1663.
func (m NoRequestingPartySubIDs) SetRequestingPartySubIDType(v int) {
	m.Set(field.NewRequestingPartySubIDType(v))
}

// GetRequestingPartySubID gets RequestingPartySubID, Tag 1662.
func (m NoRequestingPartySubIDs) GetRequestingPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.RequestingPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRequestingPartySubIDType gets RequestingPartySubIDType, Tag 1663.
func (m NoRequestingPartySubIDs) GetRequestingPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.RequestingPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRequestingPartySubID returns true if RequestingPartySubID is present, Tag 1662.
func (m NoRequestingPartySubIDs) HasRequestingPartySubID() bool {
	return m.Has(tag.RequestingPartySubID)
}

// HasRequestingPartySubIDType returns true if RequestingPartySubIDType is present, Tag 1663.
func (m NoRequestingPartySubIDs) HasRequestingPartySubIDType() bool {
	return m.Has(tag.RequestingPartySubIDType)
}

// NoRequestingPartySubIDsRepeatingGroup is a repeating group, Tag 1661.
type NoRequestingPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRequestingPartySubIDsRepeatingGroup returns an initialized, NoRequestingPartySubIDsRepeatingGroup.
func NewNoRequestingPartySubIDsRepeatingGroup() NoRequestingPartySubIDsRepeatingGroup {
	return NoRequestingPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRequestingPartySubIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RequestingPartySubID),
				quickfix.GroupElement(tag.RequestingPartySubIDType),
			},
		),
	}
}

// Add create and append a new NoRequestingPartySubIDs to this group.
func (m NoRequestingPartySubIDsRepeatingGroup) Add() NoRequestingPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoRequestingPartySubIDs{g}
}

// Get returns the ith NoRequestingPartySubIDs in the NoRequestingPartySubIDsRepeatinGroup.
func (m NoRequestingPartySubIDsRepeatingGroup) Get(i int) NoRequestingPartySubIDs {
	return NoRequestingPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoRequestingPartyIDsRepeatingGroup is a repeating group, Tag 1657.
type NoRequestingPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRequestingPartyIDsRepeatingGroup returns an initialized, NoRequestingPartyIDsRepeatingGroup.
func NewNoRequestingPartyIDsRepeatingGroup() NoRequestingPartyIDsRepeatingGroup {
	return NoRequestingPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRequestingPartyIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RequestingPartyID),
				quickfix.GroupElement(tag.RequestingPartyIDSource),
				quickfix.GroupElement(tag.RequestingPartyRole),
				NewNoRequestingPartySubIDsRepeatingGroup(),
				quickfix.GroupElement(tag.RequestingPartyRoleQualifier),
			},
		),
	}
}

// Add create and append a new NoRequestingPartyIDs to this group.
func (m NoRequestingPartyIDsRepeatingGroup) Add() NoRequestingPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoRequestingPartyIDs{g}
}

// Get returns the ith NoRequestingPartyIDs in the NoRequestingPartyIDsRepeatinGroup.
func (m NoRequestingPartyIDsRepeatingGroup) Get(i int) NoRequestingPartyIDs {
	return NoRequestingPartyIDs{m.RepeatingGroup.Get(i)}
}

// NoPartyUpdates is a repeating group element, Tag 1676.
type NoPartyUpdates struct {
	*quickfix.Group
}

// SetListUpdateAction sets ListUpdateAction, Tag 1324.
func (m NoPartyUpdates) SetListUpdateAction(v enum.ListUpdateAction) {
	m.Set(field.NewListUpdateAction(v))
}

// SetNoPartyDetails sets NoPartyDetails, Tag 1671.
func (m NoPartyUpdates) SetNoPartyDetails(f NoPartyDetailsRepeatingGroup) {
	m.SetGroup(f)
}

// GetListUpdateAction gets ListUpdateAction, Tag 1324.
func (m NoPartyUpdates) GetListUpdateAction() (v enum.ListUpdateAction, err quickfix.MessageRejectError) {
	var f field.ListUpdateActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyDetails gets NoPartyDetails, Tag 1671.
func (m NoPartyUpdates) GetNoPartyDetails() (NoPartyDetailsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyDetailsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasListUpdateAction returns true if ListUpdateAction is present, Tag 1324.
func (m NoPartyUpdates) HasListUpdateAction() bool {
	return m.Has(tag.ListUpdateAction)
}

// HasNoPartyDetails returns true if NoPartyDetails is present, Tag 1671.
func (m NoPartyUpdates) HasNoPartyDetails() bool {
	return m.Has(tag.NoPartyDetails)
}

// NoPartyDetails is a repeating group element, Tag 1671.
type NoPartyDetails struct {
	*quickfix.Group
}

// SetPartyDetailID sets PartyDetailID, Tag 1691.
func (m NoPartyDetails) SetPartyDetailID(v string) {
	m.Set(field.NewPartyDetailID(v))
}

// SetPartyDetailIDSource sets PartyDetailIDSource, Tag 1692.
func (m NoPartyDetails) SetPartyDetailIDSource(v string) {
	m.Set(field.NewPartyDetailIDSource(v))
}

// SetPartyDetailRole sets PartyDetailRole, Tag 1693.
func (m NoPartyDetails) SetPartyDetailRole(v int) {
	m.Set(field.NewPartyDetailRole(v))
}

// SetPartyDetailRoleQualifier sets PartyDetailRoleQualifier, Tag 1674.
func (m NoPartyDetails) SetPartyDetailRoleQualifier(v enum.PartyDetailRoleQualifier) {
	m.Set(field.NewPartyDetailRoleQualifier(v))
}

// SetNoPartyDetailSubIDs sets NoPartyDetailSubIDs, Tag 1694.
func (m NoPartyDetails) SetNoPartyDetailSubIDs(f NoPartyDetailSubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoPartyDetailAltID sets NoPartyDetailAltID, Tag 1516.
func (m NoPartyDetails) SetNoPartyDetailAltID(f NoPartyDetailAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoRelatedPartyDetailID sets NoRelatedPartyDetailID, Tag 1562.
func (m NoPartyDetails) SetNoRelatedPartyDetailID(f NoRelatedPartyDetailIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetPartyDetailStatus sets PartyDetailStatus, Tag 1672.
func (m NoPartyDetails) SetPartyDetailStatus(v enum.PartyDetailStatus) {
	m.Set(field.NewPartyDetailStatus(v))
}

// GetPartyDetailID gets PartyDetailID, Tag 1691.
func (m NoPartyDetails) GetPartyDetailID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyDetailIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyDetailIDSource gets PartyDetailIDSource, Tag 1692.
func (m NoPartyDetails) GetPartyDetailIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.PartyDetailIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyDetailRole gets PartyDetailRole, Tag 1693.
func (m NoPartyDetails) GetPartyDetailRole() (v int, err quickfix.MessageRejectError) {
	var f field.PartyDetailRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyDetailRoleQualifier gets PartyDetailRoleQualifier, Tag 1674.
func (m NoPartyDetails) GetPartyDetailRoleQualifier() (v enum.PartyDetailRoleQualifier, err quickfix.MessageRejectError) {
	var f field.PartyDetailRoleQualifierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyDetailSubIDs gets NoPartyDetailSubIDs, Tag 1694.
func (m NoPartyDetails) GetNoPartyDetailSubIDs() (NoPartyDetailSubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyDetailSubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoPartyDetailAltID gets NoPartyDetailAltID, Tag 1516.
func (m NoPartyDetails) GetNoPartyDetailAltID() (NoPartyDetailAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyDetailAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoRelatedPartyDetailID gets NoRelatedPartyDetailID, Tag 1562.
func (m NoPartyDetails) GetNoRelatedPartyDetailID() (NoRelatedPartyDetailIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedPartyDetailIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetPartyDetailStatus gets PartyDetailStatus, Tag 1672.
func (m NoPartyDetails) GetPartyDetailStatus() (v enum.PartyDetailStatus, err quickfix.MessageRejectError) {
	var f field.PartyDetailStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartyDetailID returns true if PartyDetailID is present, Tag 1691.
func (m NoPartyDetails) HasPartyDetailID() bool {
	return m.Has(tag.PartyDetailID)
}

// HasPartyDetailIDSource returns true if PartyDetailIDSource is present, Tag 1692.
func (m NoPartyDetails) HasPartyDetailIDSource() bool {
	return m.Has(tag.PartyDetailIDSource)
}

// HasPartyDetailRole returns true if PartyDetailRole is present, Tag 1693.
func (m NoPartyDetails) HasPartyDetailRole() bool {
	return m.Has(tag.PartyDetailRole)
}

// HasPartyDetailRoleQualifier returns true if PartyDetailRoleQualifier is present, Tag 1674.
func (m NoPartyDetails) HasPartyDetailRoleQualifier() bool {
	return m.Has(tag.PartyDetailRoleQualifier)
}

// HasNoPartyDetailSubIDs returns true if NoPartyDetailSubIDs is present, Tag 1694.
func (m NoPartyDetails) HasNoPartyDetailSubIDs() bool {
	return m.Has(tag.NoPartyDetailSubIDs)
}

// HasNoPartyDetailAltID returns true if NoPartyDetailAltID is present, Tag 1516.
func (m NoPartyDetails) HasNoPartyDetailAltID() bool {
	return m.Has(tag.NoPartyDetailAltID)
}

// HasNoRelatedPartyDetailID returns true if NoRelatedPartyDetailID is present, Tag 1562.
func (m NoPartyDetails) HasNoRelatedPartyDetailID() bool {
	return m.Has(tag.NoRelatedPartyDetailID)
}

// HasPartyDetailStatus returns true if PartyDetailStatus is present, Tag 1672.
func (m NoPartyDetails) HasPartyDetailStatus() bool {
	return m.Has(tag.PartyDetailStatus)
}

// NoPartyDetailSubIDs is a repeating group element, Tag 1694.
type NoPartyDetailSubIDs struct {
	*quickfix.Group
}

// SetPartyDetailSubID sets PartyDetailSubID, Tag 1695.
func (m NoPartyDetailSubIDs) SetPartyDetailSubID(v string) {
	m.Set(field.NewPartyDetailSubID(v))
}

// SetPartyDetailSubIDType sets PartyDetailSubIDType, Tag 1696.
func (m NoPartyDetailSubIDs) SetPartyDetailSubIDType(v int) {
	m.Set(field.NewPartyDetailSubIDType(v))
}

// GetPartyDetailSubID gets PartyDetailSubID, Tag 1695.
func (m NoPartyDetailSubIDs) GetPartyDetailSubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyDetailSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyDetailSubIDType gets PartyDetailSubIDType, Tag 1696.
func (m NoPartyDetailSubIDs) GetPartyDetailSubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.PartyDetailSubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartyDetailSubID returns true if PartyDetailSubID is present, Tag 1695.
func (m NoPartyDetailSubIDs) HasPartyDetailSubID() bool {
	return m.Has(tag.PartyDetailSubID)
}

// HasPartyDetailSubIDType returns true if PartyDetailSubIDType is present, Tag 1696.
func (m NoPartyDetailSubIDs) HasPartyDetailSubIDType() bool {
	return m.Has(tag.PartyDetailSubIDType)
}

// NoPartyDetailSubIDsRepeatingGroup is a repeating group, Tag 1694.
type NoPartyDetailSubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyDetailSubIDsRepeatingGroup returns an initialized, NoPartyDetailSubIDsRepeatingGroup.
func NewNoPartyDetailSubIDsRepeatingGroup() NoPartyDetailSubIDsRepeatingGroup {
	return NoPartyDetailSubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartyDetailSubIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PartyDetailSubID),
				quickfix.GroupElement(tag.PartyDetailSubIDType),
			},
		),
	}
}

// Add create and append a new NoPartyDetailSubIDs to this group.
func (m NoPartyDetailSubIDsRepeatingGroup) Add() NoPartyDetailSubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyDetailSubIDs{g}
}

// Get returns the ith NoPartyDetailSubIDs in the NoPartyDetailSubIDsRepeatinGroup.
func (m NoPartyDetailSubIDsRepeatingGroup) Get(i int) NoPartyDetailSubIDs {
	return NoPartyDetailSubIDs{m.RepeatingGroup.Get(i)}
}

// NoPartyDetailAltID is a repeating group element, Tag 1516.
type NoPartyDetailAltID struct {
	*quickfix.Group
}

// SetPartyDetailAltID sets PartyDetailAltID, Tag 1517.
func (m NoPartyDetailAltID) SetPartyDetailAltID(v string) {
	m.Set(field.NewPartyDetailAltID(v))
}

// SetPartyDetailAltIDSource sets PartyDetailAltIDSource, Tag 1518.
func (m NoPartyDetailAltID) SetPartyDetailAltIDSource(v string) {
	m.Set(field.NewPartyDetailAltIDSource(v))
}

// SetNoPartyDetailAltSubIDs sets NoPartyDetailAltSubIDs, Tag 1519.
func (m NoPartyDetailAltID) SetNoPartyDetailAltSubIDs(f NoPartyDetailAltSubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetPartyDetailAltID gets PartyDetailAltID, Tag 1517.
func (m NoPartyDetailAltID) GetPartyDetailAltID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyDetailAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyDetailAltIDSource gets PartyDetailAltIDSource, Tag 1518.
func (m NoPartyDetailAltID) GetPartyDetailAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.PartyDetailAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyDetailAltSubIDs gets NoPartyDetailAltSubIDs, Tag 1519.
func (m NoPartyDetailAltID) GetNoPartyDetailAltSubIDs() (NoPartyDetailAltSubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyDetailAltSubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasPartyDetailAltID returns true if PartyDetailAltID is present, Tag 1517.
func (m NoPartyDetailAltID) HasPartyDetailAltID() bool {
	return m.Has(tag.PartyDetailAltID)
}

// HasPartyDetailAltIDSource returns true if PartyDetailAltIDSource is present, Tag 1518.
func (m NoPartyDetailAltID) HasPartyDetailAltIDSource() bool {
	return m.Has(tag.PartyDetailAltIDSource)
}

// HasNoPartyDetailAltSubIDs returns true if NoPartyDetailAltSubIDs is present, Tag 1519.
func (m NoPartyDetailAltID) HasNoPartyDetailAltSubIDs() bool {
	return m.Has(tag.NoPartyDetailAltSubIDs)
}

// NoPartyDetailAltSubIDs is a repeating group element, Tag 1519.
type NoPartyDetailAltSubIDs struct {
	*quickfix.Group
}

// SetPartyDetailAltSubID sets PartyDetailAltSubID, Tag 1520.
func (m NoPartyDetailAltSubIDs) SetPartyDetailAltSubID(v string) {
	m.Set(field.NewPartyDetailAltSubID(v))
}

// SetPartyDetailAltSubIDType sets PartyDetailAltSubIDType, Tag 1521.
func (m NoPartyDetailAltSubIDs) SetPartyDetailAltSubIDType(v int) {
	m.Set(field.NewPartyDetailAltSubIDType(v))
}

// GetPartyDetailAltSubID gets PartyDetailAltSubID, Tag 1520.
func (m NoPartyDetailAltSubIDs) GetPartyDetailAltSubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyDetailAltSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyDetailAltSubIDType gets PartyDetailAltSubIDType, Tag 1521.
func (m NoPartyDetailAltSubIDs) GetPartyDetailAltSubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.PartyDetailAltSubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartyDetailAltSubID returns true if PartyDetailAltSubID is present, Tag 1520.
func (m NoPartyDetailAltSubIDs) HasPartyDetailAltSubID() bool {
	return m.Has(tag.PartyDetailAltSubID)
}

// HasPartyDetailAltSubIDType returns true if PartyDetailAltSubIDType is present, Tag 1521.
func (m NoPartyDetailAltSubIDs) HasPartyDetailAltSubIDType() bool {
	return m.Has(tag.PartyDetailAltSubIDType)
}

// NoPartyDetailAltSubIDsRepeatingGroup is a repeating group, Tag 1519.
type NoPartyDetailAltSubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyDetailAltSubIDsRepeatingGroup returns an initialized, NoPartyDetailAltSubIDsRepeatingGroup.
func NewNoPartyDetailAltSubIDsRepeatingGroup() NoPartyDetailAltSubIDsRepeatingGroup {
	return NoPartyDetailAltSubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartyDetailAltSubIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PartyDetailAltSubID),
				quickfix.GroupElement(tag.PartyDetailAltSubIDType),
			},
		),
	}
}

// Add create and append a new NoPartyDetailAltSubIDs to this group.
func (m NoPartyDetailAltSubIDsRepeatingGroup) Add() NoPartyDetailAltSubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyDetailAltSubIDs{g}
}

// Get returns the ith NoPartyDetailAltSubIDs in the NoPartyDetailAltSubIDsRepeatinGroup.
func (m NoPartyDetailAltSubIDsRepeatingGroup) Get(i int) NoPartyDetailAltSubIDs {
	return NoPartyDetailAltSubIDs{m.RepeatingGroup.Get(i)}
}

// NoPartyDetailAltIDRepeatingGroup is a repeating group, Tag 1516.
type NoPartyDetailAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyDetailAltIDRepeatingGroup returns an initialized, NoPartyDetailAltIDRepeatingGroup.
func NewNoPartyDetailAltIDRepeatingGroup() NoPartyDetailAltIDRepeatingGroup {
	return NoPartyDetailAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartyDetailAltID,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PartyDetailAltID),
				quickfix.GroupElement(tag.PartyDetailAltIDSource),
				NewNoPartyDetailAltSubIDsRepeatingGroup(),
			},
		),
	}
}

// Add create and append a new NoPartyDetailAltID to this group.
func (m NoPartyDetailAltIDRepeatingGroup) Add() NoPartyDetailAltID {
	g := m.RepeatingGroup.Add()
	return NoPartyDetailAltID{g}
}

// Get returns the ith NoPartyDetailAltID in the NoPartyDetailAltIDRepeatinGroup.
func (m NoPartyDetailAltIDRepeatingGroup) Get(i int) NoPartyDetailAltID {
	return NoPartyDetailAltID{m.RepeatingGroup.Get(i)}
}

// NoRelatedPartyDetailID is a repeating group element, Tag 1562.
type NoRelatedPartyDetailID struct {
	*quickfix.Group
}

// SetRelatedPartyDetailID sets RelatedPartyDetailID, Tag 1563.
func (m NoRelatedPartyDetailID) SetRelatedPartyDetailID(v string) {
	m.Set(field.NewRelatedPartyDetailID(v))
}

// SetRelatedPartyDetailIDSource sets RelatedPartyDetailIDSource, Tag 1564.
func (m NoRelatedPartyDetailID) SetRelatedPartyDetailIDSource(v string) {
	m.Set(field.NewRelatedPartyDetailIDSource(v))
}

// SetRelatedPartyDetailRole sets RelatedPartyDetailRole, Tag 1565.
func (m NoRelatedPartyDetailID) SetRelatedPartyDetailRole(v int) {
	m.Set(field.NewRelatedPartyDetailRole(v))
}

// SetRelatedPartyDetailRoleQualifier sets RelatedPartyDetailRoleQualifier, Tag 1675.
func (m NoRelatedPartyDetailID) SetRelatedPartyDetailRoleQualifier(v int) {
	m.Set(field.NewRelatedPartyDetailRoleQualifier(v))
}

// SetNoRelatedPartyDetailSubIDs sets NoRelatedPartyDetailSubIDs, Tag 1566.
func (m NoRelatedPartyDetailID) SetNoRelatedPartyDetailSubIDs(f NoRelatedPartyDetailSubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoRelatedPartyDetailAltID sets NoRelatedPartyDetailAltID, Tag 1569.
func (m NoRelatedPartyDetailID) SetNoRelatedPartyDetailAltID(f NoRelatedPartyDetailAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoPartyRelationships sets NoPartyRelationships, Tag 1514.
func (m NoRelatedPartyDetailID) SetNoPartyRelationships(f NoPartyRelationshipsRepeatingGroup) {
	m.SetGroup(f)
}

// GetRelatedPartyDetailID gets RelatedPartyDetailID, Tag 1563.
func (m NoRelatedPartyDetailID) GetRelatedPartyDetailID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRelatedPartyDetailIDSource gets RelatedPartyDetailIDSource, Tag 1564.
func (m NoRelatedPartyDetailID) GetRelatedPartyDetailIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRelatedPartyDetailRole gets RelatedPartyDetailRole, Tag 1565.
func (m NoRelatedPartyDetailID) GetRelatedPartyDetailRole() (v int, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRelatedPartyDetailRoleQualifier gets RelatedPartyDetailRoleQualifier, Tag 1675.
func (m NoRelatedPartyDetailID) GetRelatedPartyDetailRoleQualifier() (v int, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailRoleQualifierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRelatedPartyDetailSubIDs gets NoRelatedPartyDetailSubIDs, Tag 1566.
func (m NoRelatedPartyDetailID) GetNoRelatedPartyDetailSubIDs() (NoRelatedPartyDetailSubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedPartyDetailSubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoRelatedPartyDetailAltID gets NoRelatedPartyDetailAltID, Tag 1569.
func (m NoRelatedPartyDetailID) GetNoRelatedPartyDetailAltID() (NoRelatedPartyDetailAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedPartyDetailAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoPartyRelationships gets NoPartyRelationships, Tag 1514.
func (m NoRelatedPartyDetailID) GetNoPartyRelationships() (NoPartyRelationshipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyRelationshipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasRelatedPartyDetailID returns true if RelatedPartyDetailID is present, Tag 1563.
func (m NoRelatedPartyDetailID) HasRelatedPartyDetailID() bool {
	return m.Has(tag.RelatedPartyDetailID)
}

// HasRelatedPartyDetailIDSource returns true if RelatedPartyDetailIDSource is present, Tag 1564.
func (m NoRelatedPartyDetailID) HasRelatedPartyDetailIDSource() bool {
	return m.Has(tag.RelatedPartyDetailIDSource)
}

// HasRelatedPartyDetailRole returns true if RelatedPartyDetailRole is present, Tag 1565.
func (m NoRelatedPartyDetailID) HasRelatedPartyDetailRole() bool {
	return m.Has(tag.RelatedPartyDetailRole)
}

// HasRelatedPartyDetailRoleQualifier returns true if RelatedPartyDetailRoleQualifier is present, Tag 1675.
func (m NoRelatedPartyDetailID) HasRelatedPartyDetailRoleQualifier() bool {
	return m.Has(tag.RelatedPartyDetailRoleQualifier)
}

// HasNoRelatedPartyDetailSubIDs returns true if NoRelatedPartyDetailSubIDs is present, Tag 1566.
func (m NoRelatedPartyDetailID) HasNoRelatedPartyDetailSubIDs() bool {
	return m.Has(tag.NoRelatedPartyDetailSubIDs)
}

// HasNoRelatedPartyDetailAltID returns true if NoRelatedPartyDetailAltID is present, Tag 1569.
func (m NoRelatedPartyDetailID) HasNoRelatedPartyDetailAltID() bool {
	return m.Has(tag.NoRelatedPartyDetailAltID)
}

// HasNoPartyRelationships returns true if NoPartyRelationships is present, Tag 1514.
func (m NoRelatedPartyDetailID) HasNoPartyRelationships() bool {
	return m.Has(tag.NoPartyRelationships)
}

// NoRelatedPartyDetailSubIDs is a repeating group element, Tag 1566.
type NoRelatedPartyDetailSubIDs struct {
	*quickfix.Group
}

// SetRelatedPartyDetailSubID sets RelatedPartyDetailSubID, Tag 1567.
func (m NoRelatedPartyDetailSubIDs) SetRelatedPartyDetailSubID(v string) {
	m.Set(field.NewRelatedPartyDetailSubID(v))
}

// SetRelatedPartyDetailSubIDType sets RelatedPartyDetailSubIDType, Tag 1568.
func (m NoRelatedPartyDetailSubIDs) SetRelatedPartyDetailSubIDType(v int) {
	m.Set(field.NewRelatedPartyDetailSubIDType(v))
}

// GetRelatedPartyDetailSubID gets RelatedPartyDetailSubID, Tag 1567.
func (m NoRelatedPartyDetailSubIDs) GetRelatedPartyDetailSubID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRelatedPartyDetailSubIDType gets RelatedPartyDetailSubIDType, Tag 1568.
func (m NoRelatedPartyDetailSubIDs) GetRelatedPartyDetailSubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailSubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRelatedPartyDetailSubID returns true if RelatedPartyDetailSubID is present, Tag 1567.
func (m NoRelatedPartyDetailSubIDs) HasRelatedPartyDetailSubID() bool {
	return m.Has(tag.RelatedPartyDetailSubID)
}

// HasRelatedPartyDetailSubIDType returns true if RelatedPartyDetailSubIDType is present, Tag 1568.
func (m NoRelatedPartyDetailSubIDs) HasRelatedPartyDetailSubIDType() bool {
	return m.Has(tag.RelatedPartyDetailSubIDType)
}

// NoRelatedPartyDetailSubIDsRepeatingGroup is a repeating group, Tag 1566.
type NoRelatedPartyDetailSubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedPartyDetailSubIDsRepeatingGroup returns an initialized, NoRelatedPartyDetailSubIDsRepeatingGroup.
func NewNoRelatedPartyDetailSubIDsRepeatingGroup() NoRelatedPartyDetailSubIDsRepeatingGroup {
	return NoRelatedPartyDetailSubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRelatedPartyDetailSubIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RelatedPartyDetailSubID),
				quickfix.GroupElement(tag.RelatedPartyDetailSubIDType),
			},
		),
	}
}

// Add create and append a new NoRelatedPartyDetailSubIDs to this group.
func (m NoRelatedPartyDetailSubIDsRepeatingGroup) Add() NoRelatedPartyDetailSubIDs {
	g := m.RepeatingGroup.Add()
	return NoRelatedPartyDetailSubIDs{g}
}

// Get returns the ith NoRelatedPartyDetailSubIDs in the NoRelatedPartyDetailSubIDsRepeatinGroup.
func (m NoRelatedPartyDetailSubIDsRepeatingGroup) Get(i int) NoRelatedPartyDetailSubIDs {
	return NoRelatedPartyDetailSubIDs{m.RepeatingGroup.Get(i)}
}

// NoRelatedPartyDetailAltID is a repeating group element, Tag 1569.
type NoRelatedPartyDetailAltID struct {
	*quickfix.Group
}

// SetRelatedPartyDetailAltID sets RelatedPartyDetailAltID, Tag 1570.
func (m NoRelatedPartyDetailAltID) SetRelatedPartyDetailAltID(v string) {
	m.Set(field.NewRelatedPartyDetailAltID(v))
}

// SetRelatedPartyDetailAltIDSource sets RelatedPartyDetailAltIDSource, Tag 1571.
func (m NoRelatedPartyDetailAltID) SetRelatedPartyDetailAltIDSource(v string) {
	m.Set(field.NewRelatedPartyDetailAltIDSource(v))
}

// SetNoRelatedPartyDetailAltSubIDs sets NoRelatedPartyDetailAltSubIDs, Tag 1572.
func (m NoRelatedPartyDetailAltID) SetNoRelatedPartyDetailAltSubIDs(f NoRelatedPartyDetailAltSubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetRelatedPartyDetailAltID gets RelatedPartyDetailAltID, Tag 1570.
func (m NoRelatedPartyDetailAltID) GetRelatedPartyDetailAltID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRelatedPartyDetailAltIDSource gets RelatedPartyDetailAltIDSource, Tag 1571.
func (m NoRelatedPartyDetailAltID) GetRelatedPartyDetailAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRelatedPartyDetailAltSubIDs gets NoRelatedPartyDetailAltSubIDs, Tag 1572.
func (m NoRelatedPartyDetailAltID) GetNoRelatedPartyDetailAltSubIDs() (NoRelatedPartyDetailAltSubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedPartyDetailAltSubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasRelatedPartyDetailAltID returns true if RelatedPartyDetailAltID is present, Tag 1570.
func (m NoRelatedPartyDetailAltID) HasRelatedPartyDetailAltID() bool {
	return m.Has(tag.RelatedPartyDetailAltID)
}

// HasRelatedPartyDetailAltIDSource returns true if RelatedPartyDetailAltIDSource is present, Tag 1571.
func (m NoRelatedPartyDetailAltID) HasRelatedPartyDetailAltIDSource() bool {
	return m.Has(tag.RelatedPartyDetailAltIDSource)
}

// HasNoRelatedPartyDetailAltSubIDs returns true if NoRelatedPartyDetailAltSubIDs is present, Tag 1572.
func (m NoRelatedPartyDetailAltID) HasNoRelatedPartyDetailAltSubIDs() bool {
	return m.Has(tag.NoRelatedPartyDetailAltSubIDs)
}

// NoRelatedPartyDetailAltSubIDs is a repeating group element, Tag 1572.
type NoRelatedPartyDetailAltSubIDs struct {
	*quickfix.Group
}

// SetRelatedPartyDetailAltSubID sets RelatedPartyDetailAltSubID, Tag 1573.
func (m NoRelatedPartyDetailAltSubIDs) SetRelatedPartyDetailAltSubID(v string) {
	m.Set(field.NewRelatedPartyDetailAltSubID(v))
}

// SetRelatedPartyDetailAltSubIDType sets RelatedPartyDetailAltSubIDType, Tag 1574.
func (m NoRelatedPartyDetailAltSubIDs) SetRelatedPartyDetailAltSubIDType(v int) {
	m.Set(field.NewRelatedPartyDetailAltSubIDType(v))
}

// GetRelatedPartyDetailAltSubID gets RelatedPartyDetailAltSubID, Tag 1573.
func (m NoRelatedPartyDetailAltSubIDs) GetRelatedPartyDetailAltSubID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailAltSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRelatedPartyDetailAltSubIDType gets RelatedPartyDetailAltSubIDType, Tag 1574.
func (m NoRelatedPartyDetailAltSubIDs) GetRelatedPartyDetailAltSubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailAltSubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRelatedPartyDetailAltSubID returns true if RelatedPartyDetailAltSubID is present, Tag 1573.
func (m NoRelatedPartyDetailAltSubIDs) HasRelatedPartyDetailAltSubID() bool {
	return m.Has(tag.RelatedPartyDetailAltSubID)
}

// HasRelatedPartyDetailAltSubIDType returns true if RelatedPartyDetailAltSubIDType is present, Tag 1574.
func (m NoRelatedPartyDetailAltSubIDs) HasRelatedPartyDetailAltSubIDType() bool {
	return m.Has(tag.RelatedPartyDetailAltSubIDType)
}

// NoRelatedPartyDetailAltSubIDsRepeatingGroup is a repeating group, Tag 1572.
type NoRelatedPartyDetailAltSubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedPartyDetailAltSubIDsRepeatingGroup returns an initialized, NoRelatedPartyDetailAltSubIDsRepeatingGroup.
func NewNoRelatedPartyDetailAltSubIDsRepeatingGroup() NoRelatedPartyDetailAltSubIDsRepeatingGroup {
	return NoRelatedPartyDetailAltSubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRelatedPartyDetailAltSubIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RelatedPartyDetailAltSubID),
				quickfix.GroupElement(tag.RelatedPartyDetailAltSubIDType),
			},
		),
	}
}

// Add create and append a new NoRelatedPartyDetailAltSubIDs to this group.
func (m NoRelatedPartyDetailAltSubIDsRepeatingGroup) Add() NoRelatedPartyDetailAltSubIDs {
	g := m.RepeatingGroup.Add()
	return NoRelatedPartyDetailAltSubIDs{g}
}

// Get returns the ith NoRelatedPartyDetailAltSubIDs in the NoRelatedPartyDetailAltSubIDsRepeatinGroup.
func (m NoRelatedPartyDetailAltSubIDsRepeatingGroup) Get(i int) NoRelatedPartyDetailAltSubIDs {
	return NoRelatedPartyDetailAltSubIDs{m.RepeatingGroup.Get(i)}
}

// NoRelatedPartyDetailAltIDRepeatingGroup is a repeating group, Tag 1569.
type NoRelatedPartyDetailAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedPartyDetailAltIDRepeatingGroup returns an initialized, NoRelatedPartyDetailAltIDRepeatingGroup.
func NewNoRelatedPartyDetailAltIDRepeatingGroup() NoRelatedPartyDetailAltIDRepeatingGroup {
	return NoRelatedPartyDetailAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRelatedPartyDetailAltID,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RelatedPartyDetailAltID),
				quickfix.GroupElement(tag.RelatedPartyDetailAltIDSource),
				NewNoRelatedPartyDetailAltSubIDsRepeatingGroup(),
			},
		),
	}
}

// Add create and append a new NoRelatedPartyDetailAltID to this group.
func (m NoRelatedPartyDetailAltIDRepeatingGroup) Add() NoRelatedPartyDetailAltID {
	g := m.RepeatingGroup.Add()
	return NoRelatedPartyDetailAltID{g}
}

// Get returns the ith NoRelatedPartyDetailAltID in the NoRelatedPartyDetailAltIDRepeatinGroup.
func (m NoRelatedPartyDetailAltIDRepeatingGroup) Get(i int) NoRelatedPartyDetailAltID {
	return NoRelatedPartyDetailAltID{m.RepeatingGroup.Get(i)}
}

// NoPartyRelationships is a repeating group element, Tag 1514.
type NoPartyRelationships struct {
	*quickfix.Group
}

// SetPartyRelationship sets PartyRelationship, Tag 1515.
func (m NoPartyRelationships) SetPartyRelationship(v enum.PartyRelationship) {
	m.Set(field.NewPartyRelationship(v))
}

// GetPartyRelationship gets PartyRelationship, Tag 1515.
func (m NoPartyRelationships) GetPartyRelationship() (v enum.PartyRelationship, err quickfix.MessageRejectError) {
	var f field.PartyRelationshipField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartyRelationship returns true if PartyRelationship is present, Tag 1515.
func (m NoPartyRelationships) HasPartyRelationship() bool {
	return m.Has(tag.PartyRelationship)
}

// NoPartyRelationshipsRepeatingGroup is a repeating group, Tag 1514.
type NoPartyRelationshipsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyRelationshipsRepeatingGroup returns an initialized, NoPartyRelationshipsRepeatingGroup.
func NewNoPartyRelationshipsRepeatingGroup() NoPartyRelationshipsRepeatingGroup {
	return NoPartyRelationshipsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartyRelationships,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PartyRelationship),
			},
		),
	}
}

// Add create and append a new NoPartyRelationships to this group.
func (m NoPartyRelationshipsRepeatingGroup) Add() NoPartyRelationships {
	g := m.RepeatingGroup.Add()
	return NoPartyRelationships{g}
}

// Get returns the ith NoPartyRelationships in the NoPartyRelationshipsRepeatinGroup.
func (m NoPartyRelationshipsRepeatingGroup) Get(i int) NoPartyRelationships {
	return NoPartyRelationships{m.RepeatingGroup.Get(i)}
}

// NoRelatedPartyDetailIDRepeatingGroup is a repeating group, Tag 1562.
type NoRelatedPartyDetailIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedPartyDetailIDRepeatingGroup returns an initialized, NoRelatedPartyDetailIDRepeatingGroup.
func NewNoRelatedPartyDetailIDRepeatingGroup() NoRelatedPartyDetailIDRepeatingGroup {
	return NoRelatedPartyDetailIDRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRelatedPartyDetailID,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RelatedPartyDetailID),
				quickfix.GroupElement(tag.RelatedPartyDetailIDSource),
				quickfix.GroupElement(tag.RelatedPartyDetailRole),
				quickfix.GroupElement(tag.RelatedPartyDetailRoleQualifier),
				NewNoRelatedPartyDetailSubIDsRepeatingGroup(),
				NewNoRelatedPartyDetailAltIDRepeatingGroup(),
				NewNoPartyRelationshipsRepeatingGroup(),
			},
		),
	}
}

// Add create and append a new NoRelatedPartyDetailID to this group.
func (m NoRelatedPartyDetailIDRepeatingGroup) Add() NoRelatedPartyDetailID {
	g := m.RepeatingGroup.Add()
	return NoRelatedPartyDetailID{g}
}

// Get returns the ith NoRelatedPartyDetailID in the NoRelatedPartyDetailIDRepeatinGroup.
func (m NoRelatedPartyDetailIDRepeatingGroup) Get(i int) NoRelatedPartyDetailID {
	return NoRelatedPartyDetailID{m.RepeatingGroup.Get(i)}
}

// NoPartyDetailsRepeatingGroup is a repeating group, Tag 1671.
type NoPartyDetailsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyDetailsRepeatingGroup returns an initialized, NoPartyDetailsRepeatingGroup.
func NewNoPartyDetailsRepeatingGroup() NoPartyDetailsRepeatingGroup {
	return NoPartyDetailsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartyDetails,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PartyDetailID),
				quickfix.GroupElement(tag.PartyDetailIDSource),
				quickfix.GroupElement(tag.PartyDetailRole),
				quickfix.GroupElement(tag.PartyDetailRoleQualifier),
				NewNoPartyDetailSubIDsRepeatingGroup(),
				NewNoPartyDetailAltIDRepeatingGroup(),
				NewNoRelatedPartyDetailIDRepeatingGroup(),
				quickfix.GroupElement(tag.PartyDetailStatus),
			},
		),
	}
}

// Add create and append a new NoPartyDetails to this group.
func (m NoPartyDetailsRepeatingGroup) Add() NoPartyDetails {
	g := m.RepeatingGroup.Add()
	return NoPartyDetails{g}
}

// Get returns the ith NoPartyDetails in the NoPartyDetailsRepeatinGroup.
func (m NoPartyDetailsRepeatingGroup) Get(i int) NoPartyDetails {
	return NoPartyDetails{m.RepeatingGroup.Get(i)}
}

// NoPartyUpdatesRepeatingGroup is a repeating group, Tag 1676.
type NoPartyUpdatesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyUpdatesRepeatingGroup returns an initialized, NoPartyUpdatesRepeatingGroup.
func NewNoPartyUpdatesRepeatingGroup() NoPartyUpdatesRepeatingGroup {
	return NoPartyUpdatesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartyUpdates,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.ListUpdateAction),
				NewNoPartyDetailsRepeatingGroup(),
			},
		),
	}
}

// Add create and append a new NoPartyUpdates to this group.
func (m NoPartyUpdatesRepeatingGroup) Add() NoPartyUpdates {
	g := m.RepeatingGroup.Add()
	return NoPartyUpdates{g}
}

// Get returns the ith NoPartyUpdates in the NoPartyUpdatesRepeatinGroup.
func (m NoPartyUpdatesRepeatingGroup) Get(i int) NoPartyUpdates {
	return NoPartyUpdates{m.RepeatingGroup.Get(i)}
}
