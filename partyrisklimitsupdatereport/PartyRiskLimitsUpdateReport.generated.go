package partyrisklimitsupdatereport

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// PartyRiskLimitsUpdateReport is the fix50sp2 PartyRiskLimitsUpdateReport type, MsgType = CR.
type PartyRiskLimitsUpdateReport struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a PartyRiskLimitsUpdateReport from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) PartyRiskLimitsUpdateReport {
	return PartyRiskLimitsUpdateReport{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m PartyRiskLimitsUpdateReport) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a PartyRiskLimitsUpdateReport initialized with the required fields for PartyRiskLimitsUpdateReport.
func New(risklimitreportid field.RiskLimitReportIDField) (m PartyRiskLimitsUpdateReport) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("CR"))
	m.Set(risklimitreportid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg PartyRiskLimitsUpdateReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "CR", r
}

// SetText sets Text, Tag 58.
func (m PartyRiskLimitsUpdateReport) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m PartyRiskLimitsUpdateReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m PartyRiskLimitsUpdateReport) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m PartyRiskLimitsUpdateReport) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetLastFragment sets LastFragment, Tag 893.
func (m PartyRiskLimitsUpdateReport) SetLastFragment(v bool) {
	m.Set(field.NewLastFragment(v))
}

// SetApplID sets ApplID, Tag 1180.
func (m PartyRiskLimitsUpdateReport) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

// SetApplSeqNum sets ApplSeqNum, Tag 1181.
func (m PartyRiskLimitsUpdateReport) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

// SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350.
func (m PartyRiskLimitsUpdateReport) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

// SetApplResendFlag sets ApplResendFlag, Tag 1352.
func (m PartyRiskLimitsUpdateReport) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

// SetTotNoParties sets TotNoParties, Tag 1512.
func (m PartyRiskLimitsUpdateReport) SetTotNoParties(v int) {
	m.Set(field.NewTotNoParties(v))
}

// SetNoRequestingPartyIDs sets NoRequestingPartyIDs, Tag 1657.
func (m PartyRiskLimitsUpdateReport) SetNoRequestingPartyIDs(f NoRequestingPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetRiskLimitRequestID sets RiskLimitRequestID, Tag 1666.
func (m PartyRiskLimitsUpdateReport) SetRiskLimitRequestID(v string) {
	m.Set(field.NewRiskLimitRequestID(v))
}

// SetRiskLimitReportID sets RiskLimitReportID, Tag 1667.
func (m PartyRiskLimitsUpdateReport) SetRiskLimitReportID(v string) {
	m.Set(field.NewRiskLimitReportID(v))
}

// SetNoPartyRiskLimits sets NoPartyRiskLimits, Tag 1677.
func (m PartyRiskLimitsUpdateReport) SetNoPartyRiskLimits(f NoPartyRiskLimitsRepeatingGroup) {
	m.SetGroup(f)
}

// SetRiskLimitRequestType sets RiskLimitRequestType, Tag 1760.
func (m PartyRiskLimitsUpdateReport) SetRiskLimitRequestType(v enum.RiskLimitRequestType) {
	m.Set(field.NewRiskLimitRequestType(v))
}

// GetText gets Text, Tag 58.
func (m PartyRiskLimitsUpdateReport) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m PartyRiskLimitsUpdateReport) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m PartyRiskLimitsUpdateReport) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m PartyRiskLimitsUpdateReport) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLastFragment gets LastFragment, Tag 893.
func (m PartyRiskLimitsUpdateReport) GetLastFragment() (v bool, err quickfix.MessageRejectError) {
	var f field.LastFragmentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplID gets ApplID, Tag 1180.
func (m PartyRiskLimitsUpdateReport) GetApplID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplSeqNum gets ApplSeqNum, Tag 1181.
func (m PartyRiskLimitsUpdateReport) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350.
func (m PartyRiskLimitsUpdateReport) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplResendFlag gets ApplResendFlag, Tag 1352.
func (m PartyRiskLimitsUpdateReport) GetApplResendFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplResendFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotNoParties gets TotNoParties, Tag 1512.
func (m PartyRiskLimitsUpdateReport) GetTotNoParties() (v int, err quickfix.MessageRejectError) {
	var f field.TotNoPartiesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRequestingPartyIDs gets NoRequestingPartyIDs, Tag 1657.
func (m PartyRiskLimitsUpdateReport) GetNoRequestingPartyIDs() (NoRequestingPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRequestingPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetRiskLimitRequestID gets RiskLimitRequestID, Tag 1666.
func (m PartyRiskLimitsUpdateReport) GetRiskLimitRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.RiskLimitRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRiskLimitReportID gets RiskLimitReportID, Tag 1667.
func (m PartyRiskLimitsUpdateReport) GetRiskLimitReportID() (v string, err quickfix.MessageRejectError) {
	var f field.RiskLimitReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyRiskLimits gets NoPartyRiskLimits, Tag 1677.
func (m PartyRiskLimitsUpdateReport) GetNoPartyRiskLimits() (NoPartyRiskLimitsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyRiskLimitsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetRiskLimitRequestType gets RiskLimitRequestType, Tag 1760.
func (m PartyRiskLimitsUpdateReport) GetRiskLimitRequestType() (v enum.RiskLimitRequestType, err quickfix.MessageRejectError) {
	var f field.RiskLimitRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasText returns true if Text is present, Tag 58.
func (m PartyRiskLimitsUpdateReport) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m PartyRiskLimitsUpdateReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m PartyRiskLimitsUpdateReport) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m PartyRiskLimitsUpdateReport) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasLastFragment returns true if LastFragment is present, Tag 893.
func (m PartyRiskLimitsUpdateReport) HasLastFragment() bool {
	return m.Has(tag.LastFragment)
}

// HasApplID returns true if ApplID is present, Tag 1180.
func (m PartyRiskLimitsUpdateReport) HasApplID() bool {
	return m.Has(tag.ApplID)
}

// HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181.
func (m PartyRiskLimitsUpdateReport) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

// HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350.
func (m PartyRiskLimitsUpdateReport) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

// HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352.
func (m PartyRiskLimitsUpdateReport) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

// HasTotNoParties returns true if TotNoParties is present, Tag 1512.
func (m PartyRiskLimitsUpdateReport) HasTotNoParties() bool {
	return m.Has(tag.TotNoParties)
}

// HasNoRequestingPartyIDs returns true if NoRequestingPartyIDs is present, Tag 1657.
func (m PartyRiskLimitsUpdateReport) HasNoRequestingPartyIDs() bool {
	return m.Has(tag.NoRequestingPartyIDs)
}

// HasRiskLimitRequestID returns true if RiskLimitRequestID is present, Tag 1666.
func (m PartyRiskLimitsUpdateReport) HasRiskLimitRequestID() bool {
	return m.Has(tag.RiskLimitRequestID)
}

// HasRiskLimitReportID returns true if RiskLimitReportID is present, Tag 1667.
func (m PartyRiskLimitsUpdateReport) HasRiskLimitReportID() bool {
	return m.Has(tag.RiskLimitReportID)
}

// HasNoPartyRiskLimits returns true if NoPartyRiskLimits is present, Tag 1677.
func (m PartyRiskLimitsUpdateReport) HasNoPartyRiskLimits() bool {
	return m.Has(tag.NoPartyRiskLimits)
}

// HasRiskLimitRequestType returns true if RiskLimitRequestType is present, Tag 1760.
func (m PartyRiskLimitsUpdateReport) HasRiskLimitRequestType() bool {
	return m.Has(tag.RiskLimitRequestType)
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

// NoPartyRiskLimits is a repeating group element, Tag 1677.
type NoPartyRiskLimits struct {
	*quickfix.Group
}

// SetListUpdateAction sets ListUpdateAction, Tag 1324.
func (m NoPartyRiskLimits) SetListUpdateAction(v enum.ListUpdateAction) {
	m.Set(field.NewListUpdateAction(v))
}

// SetNoPartyDetails sets NoPartyDetails, Tag 1671.
func (m NoPartyRiskLimits) SetNoPartyDetails(f NoPartyDetailsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoRiskLimits sets NoRiskLimits, Tag 1669.
func (m NoPartyRiskLimits) SetNoRiskLimits(f NoRiskLimitsRepeatingGroup) {
	m.SetGroup(f)
}

// SetRiskLimitID sets RiskLimitID, Tag 1670.
func (m NoPartyRiskLimits) SetRiskLimitID(v string) {
	m.Set(field.NewRiskLimitID(v))
}

// SetRiskLimitCheckModelType sets RiskLimitCheckModelType, Tag 2339.
func (m NoPartyRiskLimits) SetRiskLimitCheckModelType(v enum.RiskLimitCheckModelType) {
	m.Set(field.NewRiskLimitCheckModelType(v))
}

// SetPartyRiskLimitStatus sets PartyRiskLimitStatus, Tag 2355.
func (m NoPartyRiskLimits) SetPartyRiskLimitStatus(v enum.PartyRiskLimitStatus) {
	m.Set(field.NewPartyRiskLimitStatus(v))
}

// GetListUpdateAction gets ListUpdateAction, Tag 1324.
func (m NoPartyRiskLimits) GetListUpdateAction() (v enum.ListUpdateAction, err quickfix.MessageRejectError) {
	var f field.ListUpdateActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyDetails gets NoPartyDetails, Tag 1671.
func (m NoPartyRiskLimits) GetNoPartyDetails() (NoPartyDetailsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyDetailsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoRiskLimits gets NoRiskLimits, Tag 1669.
func (m NoPartyRiskLimits) GetNoRiskLimits() (NoRiskLimitsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRiskLimitsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetRiskLimitID gets RiskLimitID, Tag 1670.
func (m NoPartyRiskLimits) GetRiskLimitID() (v string, err quickfix.MessageRejectError) {
	var f field.RiskLimitIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRiskLimitCheckModelType gets RiskLimitCheckModelType, Tag 2339.
func (m NoPartyRiskLimits) GetRiskLimitCheckModelType() (v enum.RiskLimitCheckModelType, err quickfix.MessageRejectError) {
	var f field.RiskLimitCheckModelTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyRiskLimitStatus gets PartyRiskLimitStatus, Tag 2355.
func (m NoPartyRiskLimits) GetPartyRiskLimitStatus() (v enum.PartyRiskLimitStatus, err quickfix.MessageRejectError) {
	var f field.PartyRiskLimitStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasListUpdateAction returns true if ListUpdateAction is present, Tag 1324.
func (m NoPartyRiskLimits) HasListUpdateAction() bool {
	return m.Has(tag.ListUpdateAction)
}

// HasNoPartyDetails returns true if NoPartyDetails is present, Tag 1671.
func (m NoPartyRiskLimits) HasNoPartyDetails() bool {
	return m.Has(tag.NoPartyDetails)
}

// HasNoRiskLimits returns true if NoRiskLimits is present, Tag 1669.
func (m NoPartyRiskLimits) HasNoRiskLimits() bool {
	return m.Has(tag.NoRiskLimits)
}

// HasRiskLimitID returns true if RiskLimitID is present, Tag 1670.
func (m NoPartyRiskLimits) HasRiskLimitID() bool {
	return m.Has(tag.RiskLimitID)
}

// HasRiskLimitCheckModelType returns true if RiskLimitCheckModelType is present, Tag 2339.
func (m NoPartyRiskLimits) HasRiskLimitCheckModelType() bool {
	return m.Has(tag.RiskLimitCheckModelType)
}

// HasPartyRiskLimitStatus returns true if PartyRiskLimitStatus is present, Tag 2355.
func (m NoPartyRiskLimits) HasPartyRiskLimitStatus() bool {
	return m.Has(tag.PartyRiskLimitStatus)
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

// NoRiskLimits is a repeating group element, Tag 1669.
type NoRiskLimits struct {
	*quickfix.Group
}

// SetNoRiskLimitTypes sets NoRiskLimitTypes, Tag 1529.
func (m NoRiskLimits) SetNoRiskLimitTypes(f NoRiskLimitTypesRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoRiskInstrumentScopes sets NoRiskInstrumentScopes, Tag 1534.
func (m NoRiskLimits) SetNoRiskInstrumentScopes(f NoRiskInstrumentScopesRepeatingGroup) {
	m.SetGroup(f)
}

// GetNoRiskLimitTypes gets NoRiskLimitTypes, Tag 1529.
func (m NoRiskLimits) GetNoRiskLimitTypes() (NoRiskLimitTypesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRiskLimitTypesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoRiskInstrumentScopes gets NoRiskInstrumentScopes, Tag 1534.
func (m NoRiskLimits) GetNoRiskInstrumentScopes() (NoRiskInstrumentScopesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRiskInstrumentScopesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasNoRiskLimitTypes returns true if NoRiskLimitTypes is present, Tag 1529.
func (m NoRiskLimits) HasNoRiskLimitTypes() bool {
	return m.Has(tag.NoRiskLimitTypes)
}

// HasNoRiskInstrumentScopes returns true if NoRiskInstrumentScopes is present, Tag 1534.
func (m NoRiskLimits) HasNoRiskInstrumentScopes() bool {
	return m.Has(tag.NoRiskInstrumentScopes)
}

// NoRiskLimitTypes is a repeating group element, Tag 1529.
type NoRiskLimitTypes struct {
	*quickfix.Group
}

// SetRiskLimitType sets RiskLimitType, Tag 1530.
func (m NoRiskLimitTypes) SetRiskLimitType(v enum.RiskLimitType) {
	m.Set(field.NewRiskLimitType(v))
}

// SetRiskLimitAmount sets RiskLimitAmount, Tag 1531.
func (m NoRiskLimitTypes) SetRiskLimitAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewRiskLimitAmount(value, scale))
}

// SetRiskLimitCurrency sets RiskLimitCurrency, Tag 1532.
func (m NoRiskLimitTypes) SetRiskLimitCurrency(v string) {
	m.Set(field.NewRiskLimitCurrency(v))
}

// SetRiskLimitPlatform sets RiskLimitPlatform, Tag 1533.
func (m NoRiskLimitTypes) SetRiskLimitPlatform(v string) {
	m.Set(field.NewRiskLimitPlatform(v))
}

// SetNoRiskWarningLevels sets NoRiskWarningLevels, Tag 1559.
func (m NoRiskLimitTypes) SetNoRiskWarningLevels(f NoRiskWarningLevelsRepeatingGroup) {
	m.SetGroup(f)
}

// SetRiskLimitAction sets RiskLimitAction, Tag 1767.
func (m NoRiskLimitTypes) SetRiskLimitAction(v enum.RiskLimitAction) {
	m.Set(field.NewRiskLimitAction(v))
}

// SetRiskLimitUtilizationAmount sets RiskLimitUtilizationAmount, Tag 1766.
func (m NoRiskLimitTypes) SetRiskLimitUtilizationAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewRiskLimitUtilizationAmount(value, scale))
}

// SetRiskLimitUtilizationPercent sets RiskLimitUtilizationPercent, Tag 1765.
func (m NoRiskLimitTypes) SetRiskLimitUtilizationPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewRiskLimitUtilizationPercent(value, scale))
}

// SetRiskLimitVelocityPeriod sets RiskLimitVelocityPeriod, Tag 2336.
func (m NoRiskLimitTypes) SetRiskLimitVelocityPeriod(v int) {
	m.Set(field.NewRiskLimitVelocityPeriod(v))
}

// SetRiskLimitVelocityUnit sets RiskLimitVelocityUnit, Tag 2337.
func (m NoRiskLimitTypes) SetRiskLimitVelocityUnit(v string) {
	m.Set(field.NewRiskLimitVelocityUnit(v))
}

// GetRiskLimitType gets RiskLimitType, Tag 1530.
func (m NoRiskLimitTypes) GetRiskLimitType() (v enum.RiskLimitType, err quickfix.MessageRejectError) {
	var f field.RiskLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRiskLimitAmount gets RiskLimitAmount, Tag 1531.
func (m NoRiskLimitTypes) GetRiskLimitAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RiskLimitAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRiskLimitCurrency gets RiskLimitCurrency, Tag 1532.
func (m NoRiskLimitTypes) GetRiskLimitCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.RiskLimitCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRiskLimitPlatform gets RiskLimitPlatform, Tag 1533.
func (m NoRiskLimitTypes) GetRiskLimitPlatform() (v string, err quickfix.MessageRejectError) {
	var f field.RiskLimitPlatformField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRiskWarningLevels gets NoRiskWarningLevels, Tag 1559.
func (m NoRiskLimitTypes) GetNoRiskWarningLevels() (NoRiskWarningLevelsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRiskWarningLevelsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetRiskLimitAction gets RiskLimitAction, Tag 1767.
func (m NoRiskLimitTypes) GetRiskLimitAction() (v enum.RiskLimitAction, err quickfix.MessageRejectError) {
	var f field.RiskLimitActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRiskLimitUtilizationAmount gets RiskLimitUtilizationAmount, Tag 1766.
func (m NoRiskLimitTypes) GetRiskLimitUtilizationAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RiskLimitUtilizationAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRiskLimitUtilizationPercent gets RiskLimitUtilizationPercent, Tag 1765.
func (m NoRiskLimitTypes) GetRiskLimitUtilizationPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RiskLimitUtilizationPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRiskLimitVelocityPeriod gets RiskLimitVelocityPeriod, Tag 2336.
func (m NoRiskLimitTypes) GetRiskLimitVelocityPeriod() (v int, err quickfix.MessageRejectError) {
	var f field.RiskLimitVelocityPeriodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRiskLimitVelocityUnit gets RiskLimitVelocityUnit, Tag 2337.
func (m NoRiskLimitTypes) GetRiskLimitVelocityUnit() (v string, err quickfix.MessageRejectError) {
	var f field.RiskLimitVelocityUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRiskLimitType returns true if RiskLimitType is present, Tag 1530.
func (m NoRiskLimitTypes) HasRiskLimitType() bool {
	return m.Has(tag.RiskLimitType)
}

// HasRiskLimitAmount returns true if RiskLimitAmount is present, Tag 1531.
func (m NoRiskLimitTypes) HasRiskLimitAmount() bool {
	return m.Has(tag.RiskLimitAmount)
}

// HasRiskLimitCurrency returns true if RiskLimitCurrency is present, Tag 1532.
func (m NoRiskLimitTypes) HasRiskLimitCurrency() bool {
	return m.Has(tag.RiskLimitCurrency)
}

// HasRiskLimitPlatform returns true if RiskLimitPlatform is present, Tag 1533.
func (m NoRiskLimitTypes) HasRiskLimitPlatform() bool {
	return m.Has(tag.RiskLimitPlatform)
}

// HasNoRiskWarningLevels returns true if NoRiskWarningLevels is present, Tag 1559.
func (m NoRiskLimitTypes) HasNoRiskWarningLevels() bool {
	return m.Has(tag.NoRiskWarningLevels)
}

// HasRiskLimitAction returns true if RiskLimitAction is present, Tag 1767.
func (m NoRiskLimitTypes) HasRiskLimitAction() bool {
	return m.Has(tag.RiskLimitAction)
}

// HasRiskLimitUtilizationAmount returns true if RiskLimitUtilizationAmount is present, Tag 1766.
func (m NoRiskLimitTypes) HasRiskLimitUtilizationAmount() bool {
	return m.Has(tag.RiskLimitUtilizationAmount)
}

// HasRiskLimitUtilizationPercent returns true if RiskLimitUtilizationPercent is present, Tag 1765.
func (m NoRiskLimitTypes) HasRiskLimitUtilizationPercent() bool {
	return m.Has(tag.RiskLimitUtilizationPercent)
}

// HasRiskLimitVelocityPeriod returns true if RiskLimitVelocityPeriod is present, Tag 2336.
func (m NoRiskLimitTypes) HasRiskLimitVelocityPeriod() bool {
	return m.Has(tag.RiskLimitVelocityPeriod)
}

// HasRiskLimitVelocityUnit returns true if RiskLimitVelocityUnit is present, Tag 2337.
func (m NoRiskLimitTypes) HasRiskLimitVelocityUnit() bool {
	return m.Has(tag.RiskLimitVelocityUnit)
}

// NoRiskWarningLevels is a repeating group element, Tag 1559.
type NoRiskWarningLevels struct {
	*quickfix.Group
}

// SetRiskWarningLevelPercent sets RiskWarningLevelPercent, Tag 1560.
func (m NoRiskWarningLevels) SetRiskWarningLevelPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewRiskWarningLevelPercent(value, scale))
}

// SetRiskWarningLevelName sets RiskWarningLevelName, Tag 1561.
func (m NoRiskWarningLevels) SetRiskWarningLevelName(v string) {
	m.Set(field.NewRiskWarningLevelName(v))
}

// SetRiskWarningLevelAmount sets RiskWarningLevelAmount, Tag 1768.
func (m NoRiskWarningLevels) SetRiskWarningLevelAmount(v int) {
	m.Set(field.NewRiskWarningLevelAmount(v))
}

// SetRiskWarningLevelAction sets RiskWarningLevelAction, Tag 1769.
func (m NoRiskWarningLevels) SetRiskWarningLevelAction(v int) {
	m.Set(field.NewRiskWarningLevelAction(v))
}

// GetRiskWarningLevelPercent gets RiskWarningLevelPercent, Tag 1560.
func (m NoRiskWarningLevels) GetRiskWarningLevelPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RiskWarningLevelPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRiskWarningLevelName gets RiskWarningLevelName, Tag 1561.
func (m NoRiskWarningLevels) GetRiskWarningLevelName() (v string, err quickfix.MessageRejectError) {
	var f field.RiskWarningLevelNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRiskWarningLevelAmount gets RiskWarningLevelAmount, Tag 1768.
func (m NoRiskWarningLevels) GetRiskWarningLevelAmount() (v int, err quickfix.MessageRejectError) {
	var f field.RiskWarningLevelAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRiskWarningLevelAction gets RiskWarningLevelAction, Tag 1769.
func (m NoRiskWarningLevels) GetRiskWarningLevelAction() (v int, err quickfix.MessageRejectError) {
	var f field.RiskWarningLevelActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRiskWarningLevelPercent returns true if RiskWarningLevelPercent is present, Tag 1560.
func (m NoRiskWarningLevels) HasRiskWarningLevelPercent() bool {
	return m.Has(tag.RiskWarningLevelPercent)
}

// HasRiskWarningLevelName returns true if RiskWarningLevelName is present, Tag 1561.
func (m NoRiskWarningLevels) HasRiskWarningLevelName() bool {
	return m.Has(tag.RiskWarningLevelName)
}

// HasRiskWarningLevelAmount returns true if RiskWarningLevelAmount is present, Tag 1768.
func (m NoRiskWarningLevels) HasRiskWarningLevelAmount() bool {
	return m.Has(tag.RiskWarningLevelAmount)
}

// HasRiskWarningLevelAction returns true if RiskWarningLevelAction is present, Tag 1769.
func (m NoRiskWarningLevels) HasRiskWarningLevelAction() bool {
	return m.Has(tag.RiskWarningLevelAction)
}

// NoRiskWarningLevelsRepeatingGroup is a repeating group, Tag 1559.
type NoRiskWarningLevelsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRiskWarningLevelsRepeatingGroup returns an initialized, NoRiskWarningLevelsRepeatingGroup.
func NewNoRiskWarningLevelsRepeatingGroup() NoRiskWarningLevelsRepeatingGroup {
	return NoRiskWarningLevelsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRiskWarningLevels,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RiskWarningLevelPercent),
				quickfix.GroupElement(tag.RiskWarningLevelName),
				quickfix.GroupElement(tag.RiskWarningLevelAmount),
				quickfix.GroupElement(tag.RiskWarningLevelAction),
			},
		),
	}
}

// Add create and append a new NoRiskWarningLevels to this group.
func (m NoRiskWarningLevelsRepeatingGroup) Add() NoRiskWarningLevels {
	g := m.RepeatingGroup.Add()
	return NoRiskWarningLevels{g}
}

// Get returns the ith NoRiskWarningLevels in the NoRiskWarningLevelsRepeatinGroup.
func (m NoRiskWarningLevelsRepeatingGroup) Get(i int) NoRiskWarningLevels {
	return NoRiskWarningLevels{m.RepeatingGroup.Get(i)}
}

// NoRiskLimitTypesRepeatingGroup is a repeating group, Tag 1529.
type NoRiskLimitTypesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRiskLimitTypesRepeatingGroup returns an initialized, NoRiskLimitTypesRepeatingGroup.
func NewNoRiskLimitTypesRepeatingGroup() NoRiskLimitTypesRepeatingGroup {
	return NoRiskLimitTypesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRiskLimitTypes,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RiskLimitType),
				quickfix.GroupElement(tag.RiskLimitAmount),
				quickfix.GroupElement(tag.RiskLimitCurrency),
				quickfix.GroupElement(tag.RiskLimitPlatform),
				NewNoRiskWarningLevelsRepeatingGroup(),
				quickfix.GroupElement(tag.RiskLimitAction),
				quickfix.GroupElement(tag.RiskLimitUtilizationAmount),
				quickfix.GroupElement(tag.RiskLimitUtilizationPercent),
				quickfix.GroupElement(tag.RiskLimitVelocityPeriod),
				quickfix.GroupElement(tag.RiskLimitVelocityUnit),
			},
		),
	}
}

// Add create and append a new NoRiskLimitTypes to this group.
func (m NoRiskLimitTypesRepeatingGroup) Add() NoRiskLimitTypes {
	g := m.RepeatingGroup.Add()
	return NoRiskLimitTypes{g}
}

// Get returns the ith NoRiskLimitTypes in the NoRiskLimitTypesRepeatinGroup.
func (m NoRiskLimitTypesRepeatingGroup) Get(i int) NoRiskLimitTypes {
	return NoRiskLimitTypes{m.RepeatingGroup.Get(i)}
}

// NoRiskInstrumentScopes is a repeating group element, Tag 1534.
type NoRiskInstrumentScopes struct {
	*quickfix.Group
}

// SetInstrumentScopeOperator sets InstrumentScopeOperator, Tag 1535.
func (m NoRiskInstrumentScopes) SetInstrumentScopeOperator(v enum.InstrumentScopeOperator) {
	m.Set(field.NewInstrumentScopeOperator(v))
}

// SetInstrumentScopeSymbol sets InstrumentScopeSymbol, Tag 1536.
func (m NoRiskInstrumentScopes) SetInstrumentScopeSymbol(v string) {
	m.Set(field.NewInstrumentScopeSymbol(v))
}

// SetInstrumentScopeSymbolSfx sets InstrumentScopeSymbolSfx, Tag 1537.
func (m NoRiskInstrumentScopes) SetInstrumentScopeSymbolSfx(v string) {
	m.Set(field.NewInstrumentScopeSymbolSfx(v))
}

// SetInstrumentScopeSecurityID sets InstrumentScopeSecurityID, Tag 1538.
func (m NoRiskInstrumentScopes) SetInstrumentScopeSecurityID(v string) {
	m.Set(field.NewInstrumentScopeSecurityID(v))
}

// SetInstrumentScopeSecurityIDSource sets InstrumentScopeSecurityIDSource, Tag 1539.
func (m NoRiskInstrumentScopes) SetInstrumentScopeSecurityIDSource(v string) {
	m.Set(field.NewInstrumentScopeSecurityIDSource(v))
}

// SetNoInstrumentScopeSecurityAltID sets NoInstrumentScopeSecurityAltID, Tag 1540.
func (m NoRiskInstrumentScopes) SetNoInstrumentScopeSecurityAltID(f NoInstrumentScopeSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetInstrumentScopeProduct sets InstrumentScopeProduct, Tag 1543.
func (m NoRiskInstrumentScopes) SetInstrumentScopeProduct(v int) {
	m.Set(field.NewInstrumentScopeProduct(v))
}

// SetInstrumentScopeProductComplex sets InstrumentScopeProductComplex, Tag 1544.
func (m NoRiskInstrumentScopes) SetInstrumentScopeProductComplex(v string) {
	m.Set(field.NewInstrumentScopeProductComplex(v))
}

// SetInstrumentScopeSecurityGroup sets InstrumentScopeSecurityGroup, Tag 1545.
func (m NoRiskInstrumentScopes) SetInstrumentScopeSecurityGroup(v string) {
	m.Set(field.NewInstrumentScopeSecurityGroup(v))
}

// SetInstrumentScopeCFICode sets InstrumentScopeCFICode, Tag 1546.
func (m NoRiskInstrumentScopes) SetInstrumentScopeCFICode(v string) {
	m.Set(field.NewInstrumentScopeCFICode(v))
}

// SetInstrumentScopeSecurityType sets InstrumentScopeSecurityType, Tag 1547.
func (m NoRiskInstrumentScopes) SetInstrumentScopeSecurityType(v string) {
	m.Set(field.NewInstrumentScopeSecurityType(v))
}

// SetInstrumentScopeSecuritySubType sets InstrumentScopeSecuritySubType, Tag 1548.
func (m NoRiskInstrumentScopes) SetInstrumentScopeSecuritySubType(v string) {
	m.Set(field.NewInstrumentScopeSecuritySubType(v))
}

// SetInstrumentScopeMaturityMonthYear sets InstrumentScopeMaturityMonthYear, Tag 1549.
func (m NoRiskInstrumentScopes) SetInstrumentScopeMaturityMonthYear(v string) {
	m.Set(field.NewInstrumentScopeMaturityMonthYear(v))
}

// SetInstrumentScopeMaturityTime sets InstrumentScopeMaturityTime, Tag 1550.
func (m NoRiskInstrumentScopes) SetInstrumentScopeMaturityTime(v string) {
	m.Set(field.NewInstrumentScopeMaturityTime(v))
}

// SetInstrumentScopeRestructuringType sets InstrumentScopeRestructuringType, Tag 1551.
func (m NoRiskInstrumentScopes) SetInstrumentScopeRestructuringType(v string) {
	m.Set(field.NewInstrumentScopeRestructuringType(v))
}

// SetInstrumentScopeSeniority sets InstrumentScopeSeniority, Tag 1552.
func (m NoRiskInstrumentScopes) SetInstrumentScopeSeniority(v string) {
	m.Set(field.NewInstrumentScopeSeniority(v))
}

// SetInstrumentScopePutOrCall sets InstrumentScopePutOrCall, Tag 1553.
func (m NoRiskInstrumentScopes) SetInstrumentScopePutOrCall(v int) {
	m.Set(field.NewInstrumentScopePutOrCall(v))
}

// SetInstrumentScopeFlexibleIndicator sets InstrumentScopeFlexibleIndicator, Tag 1554.
func (m NoRiskInstrumentScopes) SetInstrumentScopeFlexibleIndicator(v bool) {
	m.Set(field.NewInstrumentScopeFlexibleIndicator(v))
}

// SetInstrumentScopeCouponRate sets InstrumentScopeCouponRate, Tag 1555.
func (m NoRiskInstrumentScopes) SetInstrumentScopeCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewInstrumentScopeCouponRate(value, scale))
}

// SetInstrumentScopeSecurityExchange sets InstrumentScopeSecurityExchange, Tag 1616.
func (m NoRiskInstrumentScopes) SetInstrumentScopeSecurityExchange(v string) {
	m.Set(field.NewInstrumentScopeSecurityExchange(v))
}

// SetInstrumentScopeSecurityDesc sets InstrumentScopeSecurityDesc, Tag 1556.
func (m NoRiskInstrumentScopes) SetInstrumentScopeSecurityDesc(v string) {
	m.Set(field.NewInstrumentScopeSecurityDesc(v))
}

// SetInstrumentScopeEncodedSecurityDescLen sets InstrumentScopeEncodedSecurityDescLen, Tag 1620.
func (m NoRiskInstrumentScopes) SetInstrumentScopeEncodedSecurityDescLen(v int) {
	m.Set(field.NewInstrumentScopeEncodedSecurityDescLen(v))
}

// SetInstrumentScopeEncodedSecurityDesc sets InstrumentScopeEncodedSecurityDesc, Tag 1621.
func (m NoRiskInstrumentScopes) SetInstrumentScopeEncodedSecurityDesc(v string) {
	m.Set(field.NewInstrumentScopeEncodedSecurityDesc(v))
}

// SetInstrumentScopeSettlType sets InstrumentScopeSettlType, Tag 1557.
func (m NoRiskInstrumentScopes) SetInstrumentScopeSettlType(v string) {
	m.Set(field.NewInstrumentScopeSettlType(v))
}

// SetInstrumentScopeUPICode sets InstrumentScopeUPICode, Tag 2895.
func (m NoRiskInstrumentScopes) SetInstrumentScopeUPICode(v string) {
	m.Set(field.NewInstrumentScopeUPICode(v))
}

// SetRiskInstrumentMultiplier sets RiskInstrumentMultiplier, Tag 1558.
func (m NoRiskInstrumentScopes) SetRiskInstrumentMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewRiskInstrumentMultiplier(value, scale))
}

// GetInstrumentScopeOperator gets InstrumentScopeOperator, Tag 1535.
func (m NoRiskInstrumentScopes) GetInstrumentScopeOperator() (v enum.InstrumentScopeOperator, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeOperatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSymbol gets InstrumentScopeSymbol, Tag 1536.
func (m NoRiskInstrumentScopes) GetInstrumentScopeSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSymbolSfx gets InstrumentScopeSymbolSfx, Tag 1537.
func (m NoRiskInstrumentScopes) GetInstrumentScopeSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityID gets InstrumentScopeSecurityID, Tag 1538.
func (m NoRiskInstrumentScopes) GetInstrumentScopeSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityIDSource gets InstrumentScopeSecurityIDSource, Tag 1539.
func (m NoRiskInstrumentScopes) GetInstrumentScopeSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentScopeSecurityAltID gets NoInstrumentScopeSecurityAltID, Tag 1540.
func (m NoRiskInstrumentScopes) GetNoInstrumentScopeSecurityAltID() (NoInstrumentScopeSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentScopeSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetInstrumentScopeProduct gets InstrumentScopeProduct, Tag 1543.
func (m NoRiskInstrumentScopes) GetInstrumentScopeProduct() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeProductComplex gets InstrumentScopeProductComplex, Tag 1544.
func (m NoRiskInstrumentScopes) GetInstrumentScopeProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityGroup gets InstrumentScopeSecurityGroup, Tag 1545.
func (m NoRiskInstrumentScopes) GetInstrumentScopeSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeCFICode gets InstrumentScopeCFICode, Tag 1546.
func (m NoRiskInstrumentScopes) GetInstrumentScopeCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityType gets InstrumentScopeSecurityType, Tag 1547.
func (m NoRiskInstrumentScopes) GetInstrumentScopeSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecuritySubType gets InstrumentScopeSecuritySubType, Tag 1548.
func (m NoRiskInstrumentScopes) GetInstrumentScopeSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeMaturityMonthYear gets InstrumentScopeMaturityMonthYear, Tag 1549.
func (m NoRiskInstrumentScopes) GetInstrumentScopeMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeMaturityTime gets InstrumentScopeMaturityTime, Tag 1550.
func (m NoRiskInstrumentScopes) GetInstrumentScopeMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeRestructuringType gets InstrumentScopeRestructuringType, Tag 1551.
func (m NoRiskInstrumentScopes) GetInstrumentScopeRestructuringType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeRestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSeniority gets InstrumentScopeSeniority, Tag 1552.
func (m NoRiskInstrumentScopes) GetInstrumentScopeSeniority() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopePutOrCall gets InstrumentScopePutOrCall, Tag 1553.
func (m NoRiskInstrumentScopes) GetInstrumentScopePutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopePutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeFlexibleIndicator gets InstrumentScopeFlexibleIndicator, Tag 1554.
func (m NoRiskInstrumentScopes) GetInstrumentScopeFlexibleIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeFlexibleIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeCouponRate gets InstrumentScopeCouponRate, Tag 1555.
func (m NoRiskInstrumentScopes) GetInstrumentScopeCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityExchange gets InstrumentScopeSecurityExchange, Tag 1616.
func (m NoRiskInstrumentScopes) GetInstrumentScopeSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityDesc gets InstrumentScopeSecurityDesc, Tag 1556.
func (m NoRiskInstrumentScopes) GetInstrumentScopeSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeEncodedSecurityDescLen gets InstrumentScopeEncodedSecurityDescLen, Tag 1620.
func (m NoRiskInstrumentScopes) GetInstrumentScopeEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeEncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeEncodedSecurityDesc gets InstrumentScopeEncodedSecurityDesc, Tag 1621.
func (m NoRiskInstrumentScopes) GetInstrumentScopeEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeEncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSettlType gets InstrumentScopeSettlType, Tag 1557.
func (m NoRiskInstrumentScopes) GetInstrumentScopeSettlType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSettlTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeUPICode gets InstrumentScopeUPICode, Tag 2895.
func (m NoRiskInstrumentScopes) GetInstrumentScopeUPICode() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeUPICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRiskInstrumentMultiplier gets RiskInstrumentMultiplier, Tag 1558.
func (m NoRiskInstrumentScopes) GetRiskInstrumentMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RiskInstrumentMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasInstrumentScopeOperator returns true if InstrumentScopeOperator is present, Tag 1535.
func (m NoRiskInstrumentScopes) HasInstrumentScopeOperator() bool {
	return m.Has(tag.InstrumentScopeOperator)
}

// HasInstrumentScopeSymbol returns true if InstrumentScopeSymbol is present, Tag 1536.
func (m NoRiskInstrumentScopes) HasInstrumentScopeSymbol() bool {
	return m.Has(tag.InstrumentScopeSymbol)
}

// HasInstrumentScopeSymbolSfx returns true if InstrumentScopeSymbolSfx is present, Tag 1537.
func (m NoRiskInstrumentScopes) HasInstrumentScopeSymbolSfx() bool {
	return m.Has(tag.InstrumentScopeSymbolSfx)
}

// HasInstrumentScopeSecurityID returns true if InstrumentScopeSecurityID is present, Tag 1538.
func (m NoRiskInstrumentScopes) HasInstrumentScopeSecurityID() bool {
	return m.Has(tag.InstrumentScopeSecurityID)
}

// HasInstrumentScopeSecurityIDSource returns true if InstrumentScopeSecurityIDSource is present, Tag 1539.
func (m NoRiskInstrumentScopes) HasInstrumentScopeSecurityIDSource() bool {
	return m.Has(tag.InstrumentScopeSecurityIDSource)
}

// HasNoInstrumentScopeSecurityAltID returns true if NoInstrumentScopeSecurityAltID is present, Tag 1540.
func (m NoRiskInstrumentScopes) HasNoInstrumentScopeSecurityAltID() bool {
	return m.Has(tag.NoInstrumentScopeSecurityAltID)
}

// HasInstrumentScopeProduct returns true if InstrumentScopeProduct is present, Tag 1543.
func (m NoRiskInstrumentScopes) HasInstrumentScopeProduct() bool {
	return m.Has(tag.InstrumentScopeProduct)
}

// HasInstrumentScopeProductComplex returns true if InstrumentScopeProductComplex is present, Tag 1544.
func (m NoRiskInstrumentScopes) HasInstrumentScopeProductComplex() bool {
	return m.Has(tag.InstrumentScopeProductComplex)
}

// HasInstrumentScopeSecurityGroup returns true if InstrumentScopeSecurityGroup is present, Tag 1545.
func (m NoRiskInstrumentScopes) HasInstrumentScopeSecurityGroup() bool {
	return m.Has(tag.InstrumentScopeSecurityGroup)
}

// HasInstrumentScopeCFICode returns true if InstrumentScopeCFICode is present, Tag 1546.
func (m NoRiskInstrumentScopes) HasInstrumentScopeCFICode() bool {
	return m.Has(tag.InstrumentScopeCFICode)
}

// HasInstrumentScopeSecurityType returns true if InstrumentScopeSecurityType is present, Tag 1547.
func (m NoRiskInstrumentScopes) HasInstrumentScopeSecurityType() bool {
	return m.Has(tag.InstrumentScopeSecurityType)
}

// HasInstrumentScopeSecuritySubType returns true if InstrumentScopeSecuritySubType is present, Tag 1548.
func (m NoRiskInstrumentScopes) HasInstrumentScopeSecuritySubType() bool {
	return m.Has(tag.InstrumentScopeSecuritySubType)
}

// HasInstrumentScopeMaturityMonthYear returns true if InstrumentScopeMaturityMonthYear is present, Tag 1549.
func (m NoRiskInstrumentScopes) HasInstrumentScopeMaturityMonthYear() bool {
	return m.Has(tag.InstrumentScopeMaturityMonthYear)
}

// HasInstrumentScopeMaturityTime returns true if InstrumentScopeMaturityTime is present, Tag 1550.
func (m NoRiskInstrumentScopes) HasInstrumentScopeMaturityTime() bool {
	return m.Has(tag.InstrumentScopeMaturityTime)
}

// HasInstrumentScopeRestructuringType returns true if InstrumentScopeRestructuringType is present, Tag 1551.
func (m NoRiskInstrumentScopes) HasInstrumentScopeRestructuringType() bool {
	return m.Has(tag.InstrumentScopeRestructuringType)
}

// HasInstrumentScopeSeniority returns true if InstrumentScopeSeniority is present, Tag 1552.
func (m NoRiskInstrumentScopes) HasInstrumentScopeSeniority() bool {
	return m.Has(tag.InstrumentScopeSeniority)
}

// HasInstrumentScopePutOrCall returns true if InstrumentScopePutOrCall is present, Tag 1553.
func (m NoRiskInstrumentScopes) HasInstrumentScopePutOrCall() bool {
	return m.Has(tag.InstrumentScopePutOrCall)
}

// HasInstrumentScopeFlexibleIndicator returns true if InstrumentScopeFlexibleIndicator is present, Tag 1554.
func (m NoRiskInstrumentScopes) HasInstrumentScopeFlexibleIndicator() bool {
	return m.Has(tag.InstrumentScopeFlexibleIndicator)
}

// HasInstrumentScopeCouponRate returns true if InstrumentScopeCouponRate is present, Tag 1555.
func (m NoRiskInstrumentScopes) HasInstrumentScopeCouponRate() bool {
	return m.Has(tag.InstrumentScopeCouponRate)
}

// HasInstrumentScopeSecurityExchange returns true if InstrumentScopeSecurityExchange is present, Tag 1616.
func (m NoRiskInstrumentScopes) HasInstrumentScopeSecurityExchange() bool {
	return m.Has(tag.InstrumentScopeSecurityExchange)
}

// HasInstrumentScopeSecurityDesc returns true if InstrumentScopeSecurityDesc is present, Tag 1556.
func (m NoRiskInstrumentScopes) HasInstrumentScopeSecurityDesc() bool {
	return m.Has(tag.InstrumentScopeSecurityDesc)
}

// HasInstrumentScopeEncodedSecurityDescLen returns true if InstrumentScopeEncodedSecurityDescLen is present, Tag 1620.
func (m NoRiskInstrumentScopes) HasInstrumentScopeEncodedSecurityDescLen() bool {
	return m.Has(tag.InstrumentScopeEncodedSecurityDescLen)
}

// HasInstrumentScopeEncodedSecurityDesc returns true if InstrumentScopeEncodedSecurityDesc is present, Tag 1621.
func (m NoRiskInstrumentScopes) HasInstrumentScopeEncodedSecurityDesc() bool {
	return m.Has(tag.InstrumentScopeEncodedSecurityDesc)
}

// HasInstrumentScopeSettlType returns true if InstrumentScopeSettlType is present, Tag 1557.
func (m NoRiskInstrumentScopes) HasInstrumentScopeSettlType() bool {
	return m.Has(tag.InstrumentScopeSettlType)
}

// HasInstrumentScopeUPICode returns true if InstrumentScopeUPICode is present, Tag 2895.
func (m NoRiskInstrumentScopes) HasInstrumentScopeUPICode() bool {
	return m.Has(tag.InstrumentScopeUPICode)
}

// HasRiskInstrumentMultiplier returns true if RiskInstrumentMultiplier is present, Tag 1558.
func (m NoRiskInstrumentScopes) HasRiskInstrumentMultiplier() bool {
	return m.Has(tag.RiskInstrumentMultiplier)
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

// NoRiskInstrumentScopesRepeatingGroup is a repeating group, Tag 1534.
type NoRiskInstrumentScopesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRiskInstrumentScopesRepeatingGroup returns an initialized, NoRiskInstrumentScopesRepeatingGroup.
func NewNoRiskInstrumentScopesRepeatingGroup() NoRiskInstrumentScopesRepeatingGroup {
	return NoRiskInstrumentScopesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRiskInstrumentScopes,
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
				quickfix.GroupElement(tag.RiskInstrumentMultiplier),
			},
		),
	}
}

// Add create and append a new NoRiskInstrumentScopes to this group.
func (m NoRiskInstrumentScopesRepeatingGroup) Add() NoRiskInstrumentScopes {
	g := m.RepeatingGroup.Add()
	return NoRiskInstrumentScopes{g}
}

// Get returns the ith NoRiskInstrumentScopes in the NoRiskInstrumentScopesRepeatinGroup.
func (m NoRiskInstrumentScopesRepeatingGroup) Get(i int) NoRiskInstrumentScopes {
	return NoRiskInstrumentScopes{m.RepeatingGroup.Get(i)}
}

// NoRiskLimitsRepeatingGroup is a repeating group, Tag 1669.
type NoRiskLimitsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRiskLimitsRepeatingGroup returns an initialized, NoRiskLimitsRepeatingGroup.
func NewNoRiskLimitsRepeatingGroup() NoRiskLimitsRepeatingGroup {
	return NoRiskLimitsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRiskLimits,
			quickfix.GroupTemplate{
				NewNoRiskLimitTypesRepeatingGroup(),
				NewNoRiskInstrumentScopesRepeatingGroup(),
			},
		),
	}
}

// Add create and append a new NoRiskLimits to this group.
func (m NoRiskLimitsRepeatingGroup) Add() NoRiskLimits {
	g := m.RepeatingGroup.Add()
	return NoRiskLimits{g}
}

// Get returns the ith NoRiskLimits in the NoRiskLimitsRepeatinGroup.
func (m NoRiskLimitsRepeatingGroup) Get(i int) NoRiskLimits {
	return NoRiskLimits{m.RepeatingGroup.Get(i)}
}

// NoPartyRiskLimitsRepeatingGroup is a repeating group, Tag 1677.
type NoPartyRiskLimitsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyRiskLimitsRepeatingGroup returns an initialized, NoPartyRiskLimitsRepeatingGroup.
func NewNoPartyRiskLimitsRepeatingGroup() NoPartyRiskLimitsRepeatingGroup {
	return NoPartyRiskLimitsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartyRiskLimits,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.ListUpdateAction),
				NewNoPartyDetailsRepeatingGroup(),
				NewNoRiskLimitsRepeatingGroup(),
				quickfix.GroupElement(tag.RiskLimitID),
				quickfix.GroupElement(tag.RiskLimitCheckModelType),
				quickfix.GroupElement(tag.PartyRiskLimitStatus),
			},
		),
	}
}

// Add create and append a new NoPartyRiskLimits to this group.
func (m NoPartyRiskLimitsRepeatingGroup) Add() NoPartyRiskLimits {
	g := m.RepeatingGroup.Add()
	return NoPartyRiskLimits{g}
}

// Get returns the ith NoPartyRiskLimits in the NoPartyRiskLimitsRepeatinGroup.
func (m NoPartyRiskLimitsRepeatingGroup) Get(i int) NoPartyRiskLimits {
	return NoPartyRiskLimits{m.RepeatingGroup.Get(i)}
}
