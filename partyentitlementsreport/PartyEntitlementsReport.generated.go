package partyentitlementsreport

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// PartyEntitlementsReport is the fix50sp2 PartyEntitlementsReport type, MsgType = CV.
type PartyEntitlementsReport struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a PartyEntitlementsReport from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) PartyEntitlementsReport {
	return PartyEntitlementsReport{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m PartyEntitlementsReport) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a PartyEntitlementsReport initialized with the required fields for PartyEntitlementsReport.
func New(entitlementreportid field.EntitlementReportIDField) (m PartyEntitlementsReport) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("CV"))
	m.Set(entitlementreportid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg PartyEntitlementsReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "CV", r
}

// SetText sets Text, Tag 58.
func (m PartyEntitlementsReport) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m PartyEntitlementsReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m PartyEntitlementsReport) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m PartyEntitlementsReport) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetLastFragment sets LastFragment, Tag 893.
func (m PartyEntitlementsReport) SetLastFragment(v bool) {
	m.Set(field.NewLastFragment(v))
}

// SetApplID sets ApplID, Tag 1180.
func (m PartyEntitlementsReport) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

// SetApplSeqNum sets ApplSeqNum, Tag 1181.
func (m PartyEntitlementsReport) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

// SetRejectText sets RejectText, Tag 1328.
func (m PartyEntitlementsReport) SetRejectText(v string) {
	m.Set(field.NewRejectText(v))
}

// SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350.
func (m PartyEntitlementsReport) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

// SetApplResendFlag sets ApplResendFlag, Tag 1352.
func (m PartyEntitlementsReport) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

// SetRequestResult sets RequestResult, Tag 1511.
func (m PartyEntitlementsReport) SetRequestResult(v enum.RequestResult) {
	m.Set(field.NewRequestResult(v))
}

// SetTotNoParties sets TotNoParties, Tag 1512.
func (m PartyEntitlementsReport) SetTotNoParties(v int) {
	m.Set(field.NewTotNoParties(v))
}

// SetEncodedRejectTextLen sets EncodedRejectTextLen, Tag 1664.
func (m PartyEntitlementsReport) SetEncodedRejectTextLen(v int) {
	m.Set(field.NewEncodedRejectTextLen(v))
}

// SetEncodedRejectText sets EncodedRejectText, Tag 1665.
func (m PartyEntitlementsReport) SetEncodedRejectText(v string) {
	m.Set(field.NewEncodedRejectText(v))
}

// SetEntitlementRequestID sets EntitlementRequestID, Tag 1770.
func (m PartyEntitlementsReport) SetEntitlementRequestID(v string) {
	m.Set(field.NewEntitlementRequestID(v))
}

// SetEntitlementReportID sets EntitlementReportID, Tag 1771.
func (m PartyEntitlementsReport) SetEntitlementReportID(v string) {
	m.Set(field.NewEntitlementReportID(v))
}

// SetNoPartyEntitlements sets NoPartyEntitlements, Tag 1772.
func (m PartyEntitlementsReport) SetNoPartyEntitlements(f NoPartyEntitlementsRepeatingGroup) {
	m.SetGroup(f)
}

// GetText gets Text, Tag 58.
func (m PartyEntitlementsReport) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m PartyEntitlementsReport) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m PartyEntitlementsReport) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m PartyEntitlementsReport) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLastFragment gets LastFragment, Tag 893.
func (m PartyEntitlementsReport) GetLastFragment() (v bool, err quickfix.MessageRejectError) {
	var f field.LastFragmentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplID gets ApplID, Tag 1180.
func (m PartyEntitlementsReport) GetApplID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplSeqNum gets ApplSeqNum, Tag 1181.
func (m PartyEntitlementsReport) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRejectText gets RejectText, Tag 1328.
func (m PartyEntitlementsReport) GetRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.RejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350.
func (m PartyEntitlementsReport) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplResendFlag gets ApplResendFlag, Tag 1352.
func (m PartyEntitlementsReport) GetApplResendFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplResendFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRequestResult gets RequestResult, Tag 1511.
func (m PartyEntitlementsReport) GetRequestResult() (v enum.RequestResult, err quickfix.MessageRejectError) {
	var f field.RequestResultField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotNoParties gets TotNoParties, Tag 1512.
func (m PartyEntitlementsReport) GetTotNoParties() (v int, err quickfix.MessageRejectError) {
	var f field.TotNoPartiesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedRejectTextLen gets EncodedRejectTextLen, Tag 1664.
func (m PartyEntitlementsReport) GetEncodedRejectTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedRejectText gets EncodedRejectText, Tag 1665.
func (m PartyEntitlementsReport) GetEncodedRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEntitlementRequestID gets EntitlementRequestID, Tag 1770.
func (m PartyEntitlementsReport) GetEntitlementRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.EntitlementRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEntitlementReportID gets EntitlementReportID, Tag 1771.
func (m PartyEntitlementsReport) GetEntitlementReportID() (v string, err quickfix.MessageRejectError) {
	var f field.EntitlementReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyEntitlements gets NoPartyEntitlements, Tag 1772.
func (m PartyEntitlementsReport) GetNoPartyEntitlements() (NoPartyEntitlementsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyEntitlementsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasText returns true if Text is present, Tag 58.
func (m PartyEntitlementsReport) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m PartyEntitlementsReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m PartyEntitlementsReport) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m PartyEntitlementsReport) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasLastFragment returns true if LastFragment is present, Tag 893.
func (m PartyEntitlementsReport) HasLastFragment() bool {
	return m.Has(tag.LastFragment)
}

// HasApplID returns true if ApplID is present, Tag 1180.
func (m PartyEntitlementsReport) HasApplID() bool {
	return m.Has(tag.ApplID)
}

// HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181.
func (m PartyEntitlementsReport) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

// HasRejectText returns true if RejectText is present, Tag 1328.
func (m PartyEntitlementsReport) HasRejectText() bool {
	return m.Has(tag.RejectText)
}

// HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350.
func (m PartyEntitlementsReport) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

// HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352.
func (m PartyEntitlementsReport) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

// HasRequestResult returns true if RequestResult is present, Tag 1511.
func (m PartyEntitlementsReport) HasRequestResult() bool {
	return m.Has(tag.RequestResult)
}

// HasTotNoParties returns true if TotNoParties is present, Tag 1512.
func (m PartyEntitlementsReport) HasTotNoParties() bool {
	return m.Has(tag.TotNoParties)
}

// HasEncodedRejectTextLen returns true if EncodedRejectTextLen is present, Tag 1664.
func (m PartyEntitlementsReport) HasEncodedRejectTextLen() bool {
	return m.Has(tag.EncodedRejectTextLen)
}

// HasEncodedRejectText returns true if EncodedRejectText is present, Tag 1665.
func (m PartyEntitlementsReport) HasEncodedRejectText() bool {
	return m.Has(tag.EncodedRejectText)
}

// HasEntitlementRequestID returns true if EntitlementRequestID is present, Tag 1770.
func (m PartyEntitlementsReport) HasEntitlementRequestID() bool {
	return m.Has(tag.EntitlementRequestID)
}

// HasEntitlementReportID returns true if EntitlementReportID is present, Tag 1771.
func (m PartyEntitlementsReport) HasEntitlementReportID() bool {
	return m.Has(tag.EntitlementReportID)
}

// HasNoPartyEntitlements returns true if NoPartyEntitlements is present, Tag 1772.
func (m PartyEntitlementsReport) HasNoPartyEntitlements() bool {
	return m.Has(tag.NoPartyEntitlements)
}

// NoPartyEntitlements is a repeating group element, Tag 1772.
type NoPartyEntitlements struct {
	*quickfix.Group
}

// SetNoPartyDetails sets NoPartyDetails, Tag 1671.
func (m NoPartyEntitlements) SetNoPartyDetails(f NoPartyDetailsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoEntitlements sets NoEntitlements, Tag 1773.
func (m NoPartyEntitlements) SetNoEntitlements(f NoEntitlementsRepeatingGroup) {
	m.SetGroup(f)
}

// SetEntitlementStatus sets EntitlementStatus, Tag 1883.
func (m NoPartyEntitlements) SetEntitlementStatus(v enum.EntitlementStatus) {
	m.Set(field.NewEntitlementStatus(v))
}

// GetNoPartyDetails gets NoPartyDetails, Tag 1671.
func (m NoPartyEntitlements) GetNoPartyDetails() (NoPartyDetailsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyDetailsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoEntitlements gets NoEntitlements, Tag 1773.
func (m NoPartyEntitlements) GetNoEntitlements() (NoEntitlementsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEntitlementsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetEntitlementStatus gets EntitlementStatus, Tag 1883.
func (m NoPartyEntitlements) GetEntitlementStatus() (v enum.EntitlementStatus, err quickfix.MessageRejectError) {
	var f field.EntitlementStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasNoPartyDetails returns true if NoPartyDetails is present, Tag 1671.
func (m NoPartyEntitlements) HasNoPartyDetails() bool {
	return m.Has(tag.NoPartyDetails)
}

// HasNoEntitlements returns true if NoEntitlements is present, Tag 1773.
func (m NoPartyEntitlements) HasNoEntitlements() bool {
	return m.Has(tag.NoEntitlements)
}

// HasEntitlementStatus returns true if EntitlementStatus is present, Tag 1883.
func (m NoPartyEntitlements) HasEntitlementStatus() bool {
	return m.Has(tag.EntitlementStatus)
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

// NoEntitlements is a repeating group element, Tag 1773.
type NoEntitlements struct {
	*quickfix.Group
}

// SetEntitlementIndicator sets EntitlementIndicator, Tag 1774.
func (m NoEntitlements) SetEntitlementIndicator(v bool) {
	m.Set(field.NewEntitlementIndicator(v))
}

// SetEntitlementType sets EntitlementType, Tag 1775.
func (m NoEntitlements) SetEntitlementType(v enum.EntitlementType) {
	m.Set(field.NewEntitlementType(v))
}

// SetNoEntitlementAttrib sets NoEntitlementAttrib, Tag 1777.
func (m NoEntitlements) SetNoEntitlementAttrib(f NoEntitlementAttribRepeatingGroup) {
	m.SetGroup(f)
}

// SetEntitlementID sets EntitlementID, Tag 1776.
func (m NoEntitlements) SetEntitlementID(v string) {
	m.Set(field.NewEntitlementID(v))
}

// SetEntitlementPlatform sets EntitlementPlatform, Tag 1784.
func (m NoEntitlements) SetEntitlementPlatform(v string) {
	m.Set(field.NewEntitlementPlatform(v))
}

// SetNoInstrumentScopes sets NoInstrumentScopes, Tag 1656.
func (m NoEntitlements) SetNoInstrumentScopes(f NoInstrumentScopesRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoMarketSegments sets NoMarketSegments, Tag 1310.
func (m NoEntitlements) SetNoMarketSegments(f NoMarketSegmentsRepeatingGroup) {
	m.SetGroup(f)
}

// SetEntitlementStartDate sets EntitlementStartDate, Tag 1782.
func (m NoEntitlements) SetEntitlementStartDate(v string) {
	m.Set(field.NewEntitlementStartDate(v))
}

// SetEntitlementEndDate sets EntitlementEndDate, Tag 1783.
func (m NoEntitlements) SetEntitlementEndDate(v string) {
	m.Set(field.NewEntitlementEndDate(v))
}

// SetEntitlementSubType sets EntitlementSubType, Tag 2402.
func (m NoEntitlements) SetEntitlementSubType(v enum.EntitlementSubType) {
	m.Set(field.NewEntitlementSubType(v))
}

// GetEntitlementIndicator gets EntitlementIndicator, Tag 1774.
func (m NoEntitlements) GetEntitlementIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.EntitlementIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEntitlementType gets EntitlementType, Tag 1775.
func (m NoEntitlements) GetEntitlementType() (v enum.EntitlementType, err quickfix.MessageRejectError) {
	var f field.EntitlementTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoEntitlementAttrib gets NoEntitlementAttrib, Tag 1777.
func (m NoEntitlements) GetNoEntitlementAttrib() (NoEntitlementAttribRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEntitlementAttribRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetEntitlementID gets EntitlementID, Tag 1776.
func (m NoEntitlements) GetEntitlementID() (v string, err quickfix.MessageRejectError) {
	var f field.EntitlementIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEntitlementPlatform gets EntitlementPlatform, Tag 1784.
func (m NoEntitlements) GetEntitlementPlatform() (v string, err quickfix.MessageRejectError) {
	var f field.EntitlementPlatformField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentScopes gets NoInstrumentScopes, Tag 1656.
func (m NoEntitlements) GetNoInstrumentScopes() (NoInstrumentScopesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentScopesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoMarketSegments gets NoMarketSegments, Tag 1310.
func (m NoEntitlements) GetNoMarketSegments() (NoMarketSegmentsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMarketSegmentsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetEntitlementStartDate gets EntitlementStartDate, Tag 1782.
func (m NoEntitlements) GetEntitlementStartDate() (v string, err quickfix.MessageRejectError) {
	var f field.EntitlementStartDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEntitlementEndDate gets EntitlementEndDate, Tag 1783.
func (m NoEntitlements) GetEntitlementEndDate() (v string, err quickfix.MessageRejectError) {
	var f field.EntitlementEndDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEntitlementSubType gets EntitlementSubType, Tag 2402.
func (m NoEntitlements) GetEntitlementSubType() (v enum.EntitlementSubType, err quickfix.MessageRejectError) {
	var f field.EntitlementSubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasEntitlementIndicator returns true if EntitlementIndicator is present, Tag 1774.
func (m NoEntitlements) HasEntitlementIndicator() bool {
	return m.Has(tag.EntitlementIndicator)
}

// HasEntitlementType returns true if EntitlementType is present, Tag 1775.
func (m NoEntitlements) HasEntitlementType() bool {
	return m.Has(tag.EntitlementType)
}

// HasNoEntitlementAttrib returns true if NoEntitlementAttrib is present, Tag 1777.
func (m NoEntitlements) HasNoEntitlementAttrib() bool {
	return m.Has(tag.NoEntitlementAttrib)
}

// HasEntitlementID returns true if EntitlementID is present, Tag 1776.
func (m NoEntitlements) HasEntitlementID() bool {
	return m.Has(tag.EntitlementID)
}

// HasEntitlementPlatform returns true if EntitlementPlatform is present, Tag 1784.
func (m NoEntitlements) HasEntitlementPlatform() bool {
	return m.Has(tag.EntitlementPlatform)
}

// HasNoInstrumentScopes returns true if NoInstrumentScopes is present, Tag 1656.
func (m NoEntitlements) HasNoInstrumentScopes() bool {
	return m.Has(tag.NoInstrumentScopes)
}

// HasNoMarketSegments returns true if NoMarketSegments is present, Tag 1310.
func (m NoEntitlements) HasNoMarketSegments() bool {
	return m.Has(tag.NoMarketSegments)
}

// HasEntitlementStartDate returns true if EntitlementStartDate is present, Tag 1782.
func (m NoEntitlements) HasEntitlementStartDate() bool {
	return m.Has(tag.EntitlementStartDate)
}

// HasEntitlementEndDate returns true if EntitlementEndDate is present, Tag 1783.
func (m NoEntitlements) HasEntitlementEndDate() bool {
	return m.Has(tag.EntitlementEndDate)
}

// HasEntitlementSubType returns true if EntitlementSubType is present, Tag 2402.
func (m NoEntitlements) HasEntitlementSubType() bool {
	return m.Has(tag.EntitlementSubType)
}

// NoEntitlementAttrib is a repeating group element, Tag 1777.
type NoEntitlementAttrib struct {
	*quickfix.Group
}

// SetEntitlementAttribType sets EntitlementAttribType, Tag 1778.
func (m NoEntitlementAttrib) SetEntitlementAttribType(v int) {
	m.Set(field.NewEntitlementAttribType(v))
}

// SetEntitlementAttribDatatype sets EntitlementAttribDatatype, Tag 1779.
func (m NoEntitlementAttrib) SetEntitlementAttribDatatype(v enum.EntitlementAttribDatatype) {
	m.Set(field.NewEntitlementAttribDatatype(v))
}

// SetEntitlementAttribValue sets EntitlementAttribValue, Tag 1780.
func (m NoEntitlementAttrib) SetEntitlementAttribValue(v string) {
	m.Set(field.NewEntitlementAttribValue(v))
}

// SetEntitlementAttribCurrency sets EntitlementAttribCurrency, Tag 1781.
func (m NoEntitlementAttrib) SetEntitlementAttribCurrency(v string) {
	m.Set(field.NewEntitlementAttribCurrency(v))
}

// GetEntitlementAttribType gets EntitlementAttribType, Tag 1778.
func (m NoEntitlementAttrib) GetEntitlementAttribType() (v int, err quickfix.MessageRejectError) {
	var f field.EntitlementAttribTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEntitlementAttribDatatype gets EntitlementAttribDatatype, Tag 1779.
func (m NoEntitlementAttrib) GetEntitlementAttribDatatype() (v enum.EntitlementAttribDatatype, err quickfix.MessageRejectError) {
	var f field.EntitlementAttribDatatypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEntitlementAttribValue gets EntitlementAttribValue, Tag 1780.
func (m NoEntitlementAttrib) GetEntitlementAttribValue() (v string, err quickfix.MessageRejectError) {
	var f field.EntitlementAttribValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEntitlementAttribCurrency gets EntitlementAttribCurrency, Tag 1781.
func (m NoEntitlementAttrib) GetEntitlementAttribCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.EntitlementAttribCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasEntitlementAttribType returns true if EntitlementAttribType is present, Tag 1778.
func (m NoEntitlementAttrib) HasEntitlementAttribType() bool {
	return m.Has(tag.EntitlementAttribType)
}

// HasEntitlementAttribDatatype returns true if EntitlementAttribDatatype is present, Tag 1779.
func (m NoEntitlementAttrib) HasEntitlementAttribDatatype() bool {
	return m.Has(tag.EntitlementAttribDatatype)
}

// HasEntitlementAttribValue returns true if EntitlementAttribValue is present, Tag 1780.
func (m NoEntitlementAttrib) HasEntitlementAttribValue() bool {
	return m.Has(tag.EntitlementAttribValue)
}

// HasEntitlementAttribCurrency returns true if EntitlementAttribCurrency is present, Tag 1781.
func (m NoEntitlementAttrib) HasEntitlementAttribCurrency() bool {
	return m.Has(tag.EntitlementAttribCurrency)
}

// NoEntitlementAttribRepeatingGroup is a repeating group, Tag 1777.
type NoEntitlementAttribRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoEntitlementAttribRepeatingGroup returns an initialized, NoEntitlementAttribRepeatingGroup.
func NewNoEntitlementAttribRepeatingGroup() NoEntitlementAttribRepeatingGroup {
	return NoEntitlementAttribRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoEntitlementAttrib,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.EntitlementAttribType),
				quickfix.GroupElement(tag.EntitlementAttribDatatype),
				quickfix.GroupElement(tag.EntitlementAttribValue),
				quickfix.GroupElement(tag.EntitlementAttribCurrency),
			},
		),
	}
}

// Add create and append a new NoEntitlementAttrib to this group.
func (m NoEntitlementAttribRepeatingGroup) Add() NoEntitlementAttrib {
	g := m.RepeatingGroup.Add()
	return NoEntitlementAttrib{g}
}

// Get returns the ith NoEntitlementAttrib in the NoEntitlementAttribRepeatinGroup.
func (m NoEntitlementAttribRepeatingGroup) Get(i int) NoEntitlementAttrib {
	return NoEntitlementAttrib{m.RepeatingGroup.Get(i)}
}

// NoInstrumentScopes is a repeating group element, Tag 1656.
type NoInstrumentScopes struct {
	*quickfix.Group
}

// SetInstrumentScopeOperator sets InstrumentScopeOperator, Tag 1535.
func (m NoInstrumentScopes) SetInstrumentScopeOperator(v enum.InstrumentScopeOperator) {
	m.Set(field.NewInstrumentScopeOperator(v))
}

// SetInstrumentScopeSymbol sets InstrumentScopeSymbol, Tag 1536.
func (m NoInstrumentScopes) SetInstrumentScopeSymbol(v string) {
	m.Set(field.NewInstrumentScopeSymbol(v))
}

// SetInstrumentScopeSymbolSfx sets InstrumentScopeSymbolSfx, Tag 1537.
func (m NoInstrumentScopes) SetInstrumentScopeSymbolSfx(v string) {
	m.Set(field.NewInstrumentScopeSymbolSfx(v))
}

// SetInstrumentScopeSecurityID sets InstrumentScopeSecurityID, Tag 1538.
func (m NoInstrumentScopes) SetInstrumentScopeSecurityID(v string) {
	m.Set(field.NewInstrumentScopeSecurityID(v))
}

// SetInstrumentScopeSecurityIDSource sets InstrumentScopeSecurityIDSource, Tag 1539.
func (m NoInstrumentScopes) SetInstrumentScopeSecurityIDSource(v string) {
	m.Set(field.NewInstrumentScopeSecurityIDSource(v))
}

// SetNoInstrumentScopeSecurityAltID sets NoInstrumentScopeSecurityAltID, Tag 1540.
func (m NoInstrumentScopes) SetNoInstrumentScopeSecurityAltID(f NoInstrumentScopeSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetInstrumentScopeProduct sets InstrumentScopeProduct, Tag 1543.
func (m NoInstrumentScopes) SetInstrumentScopeProduct(v int) {
	m.Set(field.NewInstrumentScopeProduct(v))
}

// SetInstrumentScopeProductComplex sets InstrumentScopeProductComplex, Tag 1544.
func (m NoInstrumentScopes) SetInstrumentScopeProductComplex(v string) {
	m.Set(field.NewInstrumentScopeProductComplex(v))
}

// SetInstrumentScopeSecurityGroup sets InstrumentScopeSecurityGroup, Tag 1545.
func (m NoInstrumentScopes) SetInstrumentScopeSecurityGroup(v string) {
	m.Set(field.NewInstrumentScopeSecurityGroup(v))
}

// SetInstrumentScopeCFICode sets InstrumentScopeCFICode, Tag 1546.
func (m NoInstrumentScopes) SetInstrumentScopeCFICode(v string) {
	m.Set(field.NewInstrumentScopeCFICode(v))
}

// SetInstrumentScopeSecurityType sets InstrumentScopeSecurityType, Tag 1547.
func (m NoInstrumentScopes) SetInstrumentScopeSecurityType(v string) {
	m.Set(field.NewInstrumentScopeSecurityType(v))
}

// SetInstrumentScopeSecuritySubType sets InstrumentScopeSecuritySubType, Tag 1548.
func (m NoInstrumentScopes) SetInstrumentScopeSecuritySubType(v string) {
	m.Set(field.NewInstrumentScopeSecuritySubType(v))
}

// SetInstrumentScopeMaturityMonthYear sets InstrumentScopeMaturityMonthYear, Tag 1549.
func (m NoInstrumentScopes) SetInstrumentScopeMaturityMonthYear(v string) {
	m.Set(field.NewInstrumentScopeMaturityMonthYear(v))
}

// SetInstrumentScopeMaturityTime sets InstrumentScopeMaturityTime, Tag 1550.
func (m NoInstrumentScopes) SetInstrumentScopeMaturityTime(v string) {
	m.Set(field.NewInstrumentScopeMaturityTime(v))
}

// SetInstrumentScopeRestructuringType sets InstrumentScopeRestructuringType, Tag 1551.
func (m NoInstrumentScopes) SetInstrumentScopeRestructuringType(v string) {
	m.Set(field.NewInstrumentScopeRestructuringType(v))
}

// SetInstrumentScopeSeniority sets InstrumentScopeSeniority, Tag 1552.
func (m NoInstrumentScopes) SetInstrumentScopeSeniority(v string) {
	m.Set(field.NewInstrumentScopeSeniority(v))
}

// SetInstrumentScopePutOrCall sets InstrumentScopePutOrCall, Tag 1553.
func (m NoInstrumentScopes) SetInstrumentScopePutOrCall(v int) {
	m.Set(field.NewInstrumentScopePutOrCall(v))
}

// SetInstrumentScopeFlexibleIndicator sets InstrumentScopeFlexibleIndicator, Tag 1554.
func (m NoInstrumentScopes) SetInstrumentScopeFlexibleIndicator(v bool) {
	m.Set(field.NewInstrumentScopeFlexibleIndicator(v))
}

// SetInstrumentScopeCouponRate sets InstrumentScopeCouponRate, Tag 1555.
func (m NoInstrumentScopes) SetInstrumentScopeCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewInstrumentScopeCouponRate(value, scale))
}

// SetInstrumentScopeSecurityExchange sets InstrumentScopeSecurityExchange, Tag 1616.
func (m NoInstrumentScopes) SetInstrumentScopeSecurityExchange(v string) {
	m.Set(field.NewInstrumentScopeSecurityExchange(v))
}

// SetInstrumentScopeSecurityDesc sets InstrumentScopeSecurityDesc, Tag 1556.
func (m NoInstrumentScopes) SetInstrumentScopeSecurityDesc(v string) {
	m.Set(field.NewInstrumentScopeSecurityDesc(v))
}

// SetInstrumentScopeEncodedSecurityDescLen sets InstrumentScopeEncodedSecurityDescLen, Tag 1620.
func (m NoInstrumentScopes) SetInstrumentScopeEncodedSecurityDescLen(v int) {
	m.Set(field.NewInstrumentScopeEncodedSecurityDescLen(v))
}

// SetInstrumentScopeEncodedSecurityDesc sets InstrumentScopeEncodedSecurityDesc, Tag 1621.
func (m NoInstrumentScopes) SetInstrumentScopeEncodedSecurityDesc(v string) {
	m.Set(field.NewInstrumentScopeEncodedSecurityDesc(v))
}

// SetInstrumentScopeSettlType sets InstrumentScopeSettlType, Tag 1557.
func (m NoInstrumentScopes) SetInstrumentScopeSettlType(v string) {
	m.Set(field.NewInstrumentScopeSettlType(v))
}

// SetInstrumentScopeUPICode sets InstrumentScopeUPICode, Tag 2895.
func (m NoInstrumentScopes) SetInstrumentScopeUPICode(v string) {
	m.Set(field.NewInstrumentScopeUPICode(v))
}

// GetInstrumentScopeOperator gets InstrumentScopeOperator, Tag 1535.
func (m NoInstrumentScopes) GetInstrumentScopeOperator() (v enum.InstrumentScopeOperator, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeOperatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSymbol gets InstrumentScopeSymbol, Tag 1536.
func (m NoInstrumentScopes) GetInstrumentScopeSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSymbolSfx gets InstrumentScopeSymbolSfx, Tag 1537.
func (m NoInstrumentScopes) GetInstrumentScopeSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityID gets InstrumentScopeSecurityID, Tag 1538.
func (m NoInstrumentScopes) GetInstrumentScopeSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityIDSource gets InstrumentScopeSecurityIDSource, Tag 1539.
func (m NoInstrumentScopes) GetInstrumentScopeSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentScopeSecurityAltID gets NoInstrumentScopeSecurityAltID, Tag 1540.
func (m NoInstrumentScopes) GetNoInstrumentScopeSecurityAltID() (NoInstrumentScopeSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentScopeSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetInstrumentScopeProduct gets InstrumentScopeProduct, Tag 1543.
func (m NoInstrumentScopes) GetInstrumentScopeProduct() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeProductComplex gets InstrumentScopeProductComplex, Tag 1544.
func (m NoInstrumentScopes) GetInstrumentScopeProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityGroup gets InstrumentScopeSecurityGroup, Tag 1545.
func (m NoInstrumentScopes) GetInstrumentScopeSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeCFICode gets InstrumentScopeCFICode, Tag 1546.
func (m NoInstrumentScopes) GetInstrumentScopeCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityType gets InstrumentScopeSecurityType, Tag 1547.
func (m NoInstrumentScopes) GetInstrumentScopeSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecuritySubType gets InstrumentScopeSecuritySubType, Tag 1548.
func (m NoInstrumentScopes) GetInstrumentScopeSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeMaturityMonthYear gets InstrumentScopeMaturityMonthYear, Tag 1549.
func (m NoInstrumentScopes) GetInstrumentScopeMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeMaturityTime gets InstrumentScopeMaturityTime, Tag 1550.
func (m NoInstrumentScopes) GetInstrumentScopeMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeRestructuringType gets InstrumentScopeRestructuringType, Tag 1551.
func (m NoInstrumentScopes) GetInstrumentScopeRestructuringType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeRestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSeniority gets InstrumentScopeSeniority, Tag 1552.
func (m NoInstrumentScopes) GetInstrumentScopeSeniority() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopePutOrCall gets InstrumentScopePutOrCall, Tag 1553.
func (m NoInstrumentScopes) GetInstrumentScopePutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopePutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeFlexibleIndicator gets InstrumentScopeFlexibleIndicator, Tag 1554.
func (m NoInstrumentScopes) GetInstrumentScopeFlexibleIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeFlexibleIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeCouponRate gets InstrumentScopeCouponRate, Tag 1555.
func (m NoInstrumentScopes) GetInstrumentScopeCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityExchange gets InstrumentScopeSecurityExchange, Tag 1616.
func (m NoInstrumentScopes) GetInstrumentScopeSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityDesc gets InstrumentScopeSecurityDesc, Tag 1556.
func (m NoInstrumentScopes) GetInstrumentScopeSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeEncodedSecurityDescLen gets InstrumentScopeEncodedSecurityDescLen, Tag 1620.
func (m NoInstrumentScopes) GetInstrumentScopeEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeEncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeEncodedSecurityDesc gets InstrumentScopeEncodedSecurityDesc, Tag 1621.
func (m NoInstrumentScopes) GetInstrumentScopeEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeEncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSettlType gets InstrumentScopeSettlType, Tag 1557.
func (m NoInstrumentScopes) GetInstrumentScopeSettlType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSettlTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeUPICode gets InstrumentScopeUPICode, Tag 2895.
func (m NoInstrumentScopes) GetInstrumentScopeUPICode() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeUPICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasInstrumentScopeOperator returns true if InstrumentScopeOperator is present, Tag 1535.
func (m NoInstrumentScopes) HasInstrumentScopeOperator() bool {
	return m.Has(tag.InstrumentScopeOperator)
}

// HasInstrumentScopeSymbol returns true if InstrumentScopeSymbol is present, Tag 1536.
func (m NoInstrumentScopes) HasInstrumentScopeSymbol() bool {
	return m.Has(tag.InstrumentScopeSymbol)
}

// HasInstrumentScopeSymbolSfx returns true if InstrumentScopeSymbolSfx is present, Tag 1537.
func (m NoInstrumentScopes) HasInstrumentScopeSymbolSfx() bool {
	return m.Has(tag.InstrumentScopeSymbolSfx)
}

// HasInstrumentScopeSecurityID returns true if InstrumentScopeSecurityID is present, Tag 1538.
func (m NoInstrumentScopes) HasInstrumentScopeSecurityID() bool {
	return m.Has(tag.InstrumentScopeSecurityID)
}

// HasInstrumentScopeSecurityIDSource returns true if InstrumentScopeSecurityIDSource is present, Tag 1539.
func (m NoInstrumentScopes) HasInstrumentScopeSecurityIDSource() bool {
	return m.Has(tag.InstrumentScopeSecurityIDSource)
}

// HasNoInstrumentScopeSecurityAltID returns true if NoInstrumentScopeSecurityAltID is present, Tag 1540.
func (m NoInstrumentScopes) HasNoInstrumentScopeSecurityAltID() bool {
	return m.Has(tag.NoInstrumentScopeSecurityAltID)
}

// HasInstrumentScopeProduct returns true if InstrumentScopeProduct is present, Tag 1543.
func (m NoInstrumentScopes) HasInstrumentScopeProduct() bool {
	return m.Has(tag.InstrumentScopeProduct)
}

// HasInstrumentScopeProductComplex returns true if InstrumentScopeProductComplex is present, Tag 1544.
func (m NoInstrumentScopes) HasInstrumentScopeProductComplex() bool {
	return m.Has(tag.InstrumentScopeProductComplex)
}

// HasInstrumentScopeSecurityGroup returns true if InstrumentScopeSecurityGroup is present, Tag 1545.
func (m NoInstrumentScopes) HasInstrumentScopeSecurityGroup() bool {
	return m.Has(tag.InstrumentScopeSecurityGroup)
}

// HasInstrumentScopeCFICode returns true if InstrumentScopeCFICode is present, Tag 1546.
func (m NoInstrumentScopes) HasInstrumentScopeCFICode() bool {
	return m.Has(tag.InstrumentScopeCFICode)
}

// HasInstrumentScopeSecurityType returns true if InstrumentScopeSecurityType is present, Tag 1547.
func (m NoInstrumentScopes) HasInstrumentScopeSecurityType() bool {
	return m.Has(tag.InstrumentScopeSecurityType)
}

// HasInstrumentScopeSecuritySubType returns true if InstrumentScopeSecuritySubType is present, Tag 1548.
func (m NoInstrumentScopes) HasInstrumentScopeSecuritySubType() bool {
	return m.Has(tag.InstrumentScopeSecuritySubType)
}

// HasInstrumentScopeMaturityMonthYear returns true if InstrumentScopeMaturityMonthYear is present, Tag 1549.
func (m NoInstrumentScopes) HasInstrumentScopeMaturityMonthYear() bool {
	return m.Has(tag.InstrumentScopeMaturityMonthYear)
}

// HasInstrumentScopeMaturityTime returns true if InstrumentScopeMaturityTime is present, Tag 1550.
func (m NoInstrumentScopes) HasInstrumentScopeMaturityTime() bool {
	return m.Has(tag.InstrumentScopeMaturityTime)
}

// HasInstrumentScopeRestructuringType returns true if InstrumentScopeRestructuringType is present, Tag 1551.
func (m NoInstrumentScopes) HasInstrumentScopeRestructuringType() bool {
	return m.Has(tag.InstrumentScopeRestructuringType)
}

// HasInstrumentScopeSeniority returns true if InstrumentScopeSeniority is present, Tag 1552.
func (m NoInstrumentScopes) HasInstrumentScopeSeniority() bool {
	return m.Has(tag.InstrumentScopeSeniority)
}

// HasInstrumentScopePutOrCall returns true if InstrumentScopePutOrCall is present, Tag 1553.
func (m NoInstrumentScopes) HasInstrumentScopePutOrCall() bool {
	return m.Has(tag.InstrumentScopePutOrCall)
}

// HasInstrumentScopeFlexibleIndicator returns true if InstrumentScopeFlexibleIndicator is present, Tag 1554.
func (m NoInstrumentScopes) HasInstrumentScopeFlexibleIndicator() bool {
	return m.Has(tag.InstrumentScopeFlexibleIndicator)
}

// HasInstrumentScopeCouponRate returns true if InstrumentScopeCouponRate is present, Tag 1555.
func (m NoInstrumentScopes) HasInstrumentScopeCouponRate() bool {
	return m.Has(tag.InstrumentScopeCouponRate)
}

// HasInstrumentScopeSecurityExchange returns true if InstrumentScopeSecurityExchange is present, Tag 1616.
func (m NoInstrumentScopes) HasInstrumentScopeSecurityExchange() bool {
	return m.Has(tag.InstrumentScopeSecurityExchange)
}

// HasInstrumentScopeSecurityDesc returns true if InstrumentScopeSecurityDesc is present, Tag 1556.
func (m NoInstrumentScopes) HasInstrumentScopeSecurityDesc() bool {
	return m.Has(tag.InstrumentScopeSecurityDesc)
}

// HasInstrumentScopeEncodedSecurityDescLen returns true if InstrumentScopeEncodedSecurityDescLen is present, Tag 1620.
func (m NoInstrumentScopes) HasInstrumentScopeEncodedSecurityDescLen() bool {
	return m.Has(tag.InstrumentScopeEncodedSecurityDescLen)
}

// HasInstrumentScopeEncodedSecurityDesc returns true if InstrumentScopeEncodedSecurityDesc is present, Tag 1621.
func (m NoInstrumentScopes) HasInstrumentScopeEncodedSecurityDesc() bool {
	return m.Has(tag.InstrumentScopeEncodedSecurityDesc)
}

// HasInstrumentScopeSettlType returns true if InstrumentScopeSettlType is present, Tag 1557.
func (m NoInstrumentScopes) HasInstrumentScopeSettlType() bool {
	return m.Has(tag.InstrumentScopeSettlType)
}

// HasInstrumentScopeUPICode returns true if InstrumentScopeUPICode is present, Tag 2895.
func (m NoInstrumentScopes) HasInstrumentScopeUPICode() bool {
	return m.Has(tag.InstrumentScopeUPICode)
}

// NoInstrumentScopeSecurityAltID is a repeating group element, Tag 1540.
type NoInstrumentScopeSecurityAltID struct {
	*quickfix.Group
}

// SetInstrumentScopeSecurityAltID sets InstrumentScopeSecurityAltID, Tag 1541.
func (m NoInstrumentScopeSecurityAltID) SetInstrumentScopeSecurityAltID(v string) {
	m.Set(field.NewInstrumentScopeSecurityAltID(v))
}

// SetInstrumentScopeSecurityAltIDSource sets InstrumentScopeSecurityAltIDSource, Tag 1542.
func (m NoInstrumentScopeSecurityAltID) SetInstrumentScopeSecurityAltIDSource(v string) {
	m.Set(field.NewInstrumentScopeSecurityAltIDSource(v))
}

// GetInstrumentScopeSecurityAltID gets InstrumentScopeSecurityAltID, Tag 1541.
func (m NoInstrumentScopeSecurityAltID) GetInstrumentScopeSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityAltIDSource gets InstrumentScopeSecurityAltIDSource, Tag 1542.
func (m NoInstrumentScopeSecurityAltID) GetInstrumentScopeSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasInstrumentScopeSecurityAltID returns true if InstrumentScopeSecurityAltID is present, Tag 1541.
func (m NoInstrumentScopeSecurityAltID) HasInstrumentScopeSecurityAltID() bool {
	return m.Has(tag.InstrumentScopeSecurityAltID)
}

// HasInstrumentScopeSecurityAltIDSource returns true if InstrumentScopeSecurityAltIDSource is present, Tag 1542.
func (m NoInstrumentScopeSecurityAltID) HasInstrumentScopeSecurityAltIDSource() bool {
	return m.Has(tag.InstrumentScopeSecurityAltIDSource)
}

// NoInstrumentScopeSecurityAltIDRepeatingGroup is a repeating group, Tag 1540.
type NoInstrumentScopeSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoInstrumentScopeSecurityAltIDRepeatingGroup returns an initialized, NoInstrumentScopeSecurityAltIDRepeatingGroup.
func NewNoInstrumentScopeSecurityAltIDRepeatingGroup() NoInstrumentScopeSecurityAltIDRepeatingGroup {
	return NoInstrumentScopeSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoInstrumentScopeSecurityAltID,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.InstrumentScopeSecurityAltID),
				quickfix.GroupElement(tag.InstrumentScopeSecurityAltIDSource),
			},
		),
	}
}

// Add create and append a new NoInstrumentScopeSecurityAltID to this group.
func (m NoInstrumentScopeSecurityAltIDRepeatingGroup) Add() NoInstrumentScopeSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoInstrumentScopeSecurityAltID{g}
}

// Get returns the ith NoInstrumentScopeSecurityAltID in the NoInstrumentScopeSecurityAltIDRepeatinGroup.
func (m NoInstrumentScopeSecurityAltIDRepeatingGroup) Get(i int) NoInstrumentScopeSecurityAltID {
	return NoInstrumentScopeSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoInstrumentScopesRepeatingGroup is a repeating group, Tag 1656.
type NoInstrumentScopesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoInstrumentScopesRepeatingGroup returns an initialized, NoInstrumentScopesRepeatingGroup.
func NewNoInstrumentScopesRepeatingGroup() NoInstrumentScopesRepeatingGroup {
	return NoInstrumentScopesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoInstrumentScopes,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.InstrumentScopeOperator),
				quickfix.GroupElement(tag.InstrumentScopeSymbol),
				quickfix.GroupElement(tag.InstrumentScopeSymbolSfx),
				quickfix.GroupElement(tag.InstrumentScopeSecurityID),
				quickfix.GroupElement(tag.InstrumentScopeSecurityIDSource),
				NewNoInstrumentScopeSecurityAltIDRepeatingGroup(),
				quickfix.GroupElement(tag.InstrumentScopeProduct),
				quickfix.GroupElement(tag.InstrumentScopeProductComplex),
				quickfix.GroupElement(tag.InstrumentScopeSecurityGroup),
				quickfix.GroupElement(tag.InstrumentScopeCFICode),
				quickfix.GroupElement(tag.InstrumentScopeSecurityType),
				quickfix.GroupElement(tag.InstrumentScopeSecuritySubType),
				quickfix.GroupElement(tag.InstrumentScopeMaturityMonthYear),
				quickfix.GroupElement(tag.InstrumentScopeMaturityTime),
				quickfix.GroupElement(tag.InstrumentScopeRestructuringType),
				quickfix.GroupElement(tag.InstrumentScopeSeniority),
				quickfix.GroupElement(tag.InstrumentScopePutOrCall),
				quickfix.GroupElement(tag.InstrumentScopeFlexibleIndicator),
				quickfix.GroupElement(tag.InstrumentScopeCouponRate),
				quickfix.GroupElement(tag.InstrumentScopeSecurityExchange),
				quickfix.GroupElement(tag.InstrumentScopeSecurityDesc),
				quickfix.GroupElement(tag.InstrumentScopeEncodedSecurityDescLen),
				quickfix.GroupElement(tag.InstrumentScopeEncodedSecurityDesc),
				quickfix.GroupElement(tag.InstrumentScopeSettlType),
				quickfix.GroupElement(tag.InstrumentScopeUPICode),
			},
		),
	}
}

// Add create and append a new NoInstrumentScopes to this group.
func (m NoInstrumentScopesRepeatingGroup) Add() NoInstrumentScopes {
	g := m.RepeatingGroup.Add()
	return NoInstrumentScopes{g}
}

// Get returns the ith NoInstrumentScopes in the NoInstrumentScopesRepeatinGroup.
func (m NoInstrumentScopesRepeatingGroup) Get(i int) NoInstrumentScopes {
	return NoInstrumentScopes{m.RepeatingGroup.Get(i)}
}

// NoMarketSegments is a repeating group element, Tag 1310.
type NoMarketSegments struct {
	*quickfix.Group
}

// SetMarketID sets MarketID, Tag 1301.
func (m NoMarketSegments) SetMarketID(v string) {
	m.Set(field.NewMarketID(v))
}

// SetMarketSegmentID sets MarketSegmentID, Tag 1300.
func (m NoMarketSegments) SetMarketSegmentID(v string) {
	m.Set(field.NewMarketSegmentID(v))
}

// GetMarketID gets MarketID, Tag 1301.
func (m NoMarketSegments) GetMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketSegmentID gets MarketSegmentID, Tag 1300.
func (m NoMarketSegments) GetMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasMarketID returns true if MarketID is present, Tag 1301.
func (m NoMarketSegments) HasMarketID() bool {
	return m.Has(tag.MarketID)
}

// HasMarketSegmentID returns true if MarketSegmentID is present, Tag 1300.
func (m NoMarketSegments) HasMarketSegmentID() bool {
	return m.Has(tag.MarketSegmentID)
}

// NoMarketSegmentsRepeatingGroup is a repeating group, Tag 1310.
type NoMarketSegmentsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoMarketSegmentsRepeatingGroup returns an initialized, NoMarketSegmentsRepeatingGroup.
func NewNoMarketSegmentsRepeatingGroup() NoMarketSegmentsRepeatingGroup {
	return NoMarketSegmentsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoMarketSegments,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.MarketID),
				quickfix.GroupElement(tag.MarketSegmentID),
			},
		),
	}
}

// Add create and append a new NoMarketSegments to this group.
func (m NoMarketSegmentsRepeatingGroup) Add() NoMarketSegments {
	g := m.RepeatingGroup.Add()
	return NoMarketSegments{g}
}

// Get returns the ith NoMarketSegments in the NoMarketSegmentsRepeatinGroup.
func (m NoMarketSegmentsRepeatingGroup) Get(i int) NoMarketSegments {
	return NoMarketSegments{m.RepeatingGroup.Get(i)}
}

// NoEntitlementsRepeatingGroup is a repeating group, Tag 1773.
type NoEntitlementsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoEntitlementsRepeatingGroup returns an initialized, NoEntitlementsRepeatingGroup.
func NewNoEntitlementsRepeatingGroup() NoEntitlementsRepeatingGroup {
	return NoEntitlementsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoEntitlements,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.EntitlementIndicator),
				quickfix.GroupElement(tag.EntitlementType),
				NewNoEntitlementAttribRepeatingGroup(),
				quickfix.GroupElement(tag.EntitlementID),
				quickfix.GroupElement(tag.EntitlementPlatform),
				NewNoInstrumentScopesRepeatingGroup(),
				NewNoMarketSegmentsRepeatingGroup(),
				quickfix.GroupElement(tag.EntitlementStartDate),
				quickfix.GroupElement(tag.EntitlementEndDate),
				quickfix.GroupElement(tag.EntitlementSubType),
			},
		),
	}
}

// Add create and append a new NoEntitlements to this group.
func (m NoEntitlementsRepeatingGroup) Add() NoEntitlements {
	g := m.RepeatingGroup.Add()
	return NoEntitlements{g}
}

// Get returns the ith NoEntitlements in the NoEntitlementsRepeatinGroup.
func (m NoEntitlementsRepeatingGroup) Get(i int) NoEntitlements {
	return NoEntitlements{m.RepeatingGroup.Get(i)}
}

// NoPartyEntitlementsRepeatingGroup is a repeating group, Tag 1772.
type NoPartyEntitlementsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyEntitlementsRepeatingGroup returns an initialized, NoPartyEntitlementsRepeatingGroup.
func NewNoPartyEntitlementsRepeatingGroup() NoPartyEntitlementsRepeatingGroup {
	return NoPartyEntitlementsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartyEntitlements,
			quickfix.GroupTemplate{
				NewNoPartyDetailsRepeatingGroup(),
				NewNoEntitlementsRepeatingGroup(),
				quickfix.GroupElement(tag.EntitlementStatus),
			},
		),
	}
}

// Add create and append a new NoPartyEntitlements to this group.
func (m NoPartyEntitlementsRepeatingGroup) Add() NoPartyEntitlements {
	g := m.RepeatingGroup.Add()
	return NoPartyEntitlements{g}
}

// Get returns the ith NoPartyEntitlements in the NoPartyEntitlementsRepeatinGroup.
func (m NoPartyEntitlementsRepeatingGroup) Get(i int) NoPartyEntitlements {
	return NoPartyEntitlements{m.RepeatingGroup.Get(i)}
}