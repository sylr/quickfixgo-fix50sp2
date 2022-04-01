package partyrisklimitsrequest

import (
	"github.com/shopspring/decimal"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// PartyRiskLimitsRequest is the fix50sp2 PartyRiskLimitsRequest type, MsgType = CL.
type PartyRiskLimitsRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a PartyRiskLimitsRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) PartyRiskLimitsRequest {
	return PartyRiskLimitsRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m PartyRiskLimitsRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a PartyRiskLimitsRequest initialized with the required fields for PartyRiskLimitsRequest.
func New(risklimitrequestid field.RiskLimitRequestIDField) (m PartyRiskLimitsRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("CL"))
	m.Set(risklimitrequestid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg PartyRiskLimitsRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "CL", r
}

// SetText sets Text, Tag 58.
func (m PartyRiskLimitsRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263.
func (m PartyRiskLimitsRequest) SetSubscriptionRequestType(v enum.SubscriptionRequestType) {
	m.Set(field.NewSubscriptionRequestType(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m PartyRiskLimitsRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m PartyRiskLimitsRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m PartyRiskLimitsRequest) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoRequestedPartyRoles sets NoRequestedPartyRoles, Tag 1508.
func (m PartyRiskLimitsRequest) SetNoRequestedPartyRoles(f NoRequestedPartyRolesRepeatingGroup) {
	m.SetGroup(f)
}

// SetRiskLimitPlatform sets RiskLimitPlatform, Tag 1533.
func (m PartyRiskLimitsRequest) SetRiskLimitPlatform(v string) {
	m.Set(field.NewRiskLimitPlatform(v))
}

// SetNoRiskInstrumentScopes sets NoRiskInstrumentScopes, Tag 1534.
func (m PartyRiskLimitsRequest) SetNoRiskInstrumentScopes(f NoRiskInstrumentScopesRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoRequestingPartyIDs sets NoRequestingPartyIDs, Tag 1657.
func (m PartyRiskLimitsRequest) SetNoRequestingPartyIDs(f NoRequestingPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetRiskLimitRequestID sets RiskLimitRequestID, Tag 1666.
func (m PartyRiskLimitsRequest) SetRiskLimitRequestID(v string) {
	m.Set(field.NewRiskLimitRequestID(v))
}

// SetNoRequestedRiskLimitType sets NoRequestedRiskLimitType, Tag 1668.
func (m PartyRiskLimitsRequest) SetNoRequestedRiskLimitType(f NoRequestedRiskLimitTypeRepeatingGroup) {
	m.SetGroup(f)
}

// SetRiskLimitRequestType sets RiskLimitRequestType, Tag 1760.
func (m PartyRiskLimitsRequest) SetRiskLimitRequestType(v enum.RiskLimitRequestType) {
	m.Set(field.NewRiskLimitRequestType(v))
}

// GetText gets Text, Tag 58.
func (m PartyRiskLimitsRequest) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263.
func (m PartyRiskLimitsRequest) GetSubscriptionRequestType() (v enum.SubscriptionRequestType, err quickfix.MessageRejectError) {
	var f field.SubscriptionRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m PartyRiskLimitsRequest) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m PartyRiskLimitsRequest) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m PartyRiskLimitsRequest) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoRequestedPartyRoles gets NoRequestedPartyRoles, Tag 1508.
func (m PartyRiskLimitsRequest) GetNoRequestedPartyRoles() (NoRequestedPartyRolesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRequestedPartyRolesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetRiskLimitPlatform gets RiskLimitPlatform, Tag 1533.
func (m PartyRiskLimitsRequest) GetRiskLimitPlatform() (v string, err quickfix.MessageRejectError) {
	var f field.RiskLimitPlatformField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRiskInstrumentScopes gets NoRiskInstrumentScopes, Tag 1534.
func (m PartyRiskLimitsRequest) GetNoRiskInstrumentScopes() (NoRiskInstrumentScopesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRiskInstrumentScopesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoRequestingPartyIDs gets NoRequestingPartyIDs, Tag 1657.
func (m PartyRiskLimitsRequest) GetNoRequestingPartyIDs() (NoRequestingPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRequestingPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetRiskLimitRequestID gets RiskLimitRequestID, Tag 1666.
func (m PartyRiskLimitsRequest) GetRiskLimitRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.RiskLimitRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRequestedRiskLimitType gets NoRequestedRiskLimitType, Tag 1668.
func (m PartyRiskLimitsRequest) GetNoRequestedRiskLimitType() (NoRequestedRiskLimitTypeRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRequestedRiskLimitTypeRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetRiskLimitRequestType gets RiskLimitRequestType, Tag 1760.
func (m PartyRiskLimitsRequest) GetRiskLimitRequestType() (v enum.RiskLimitRequestType, err quickfix.MessageRejectError) {
	var f field.RiskLimitRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasText returns true if Text is present, Tag 58.
func (m PartyRiskLimitsRequest) HasText() bool {
	return m.Has(tag.Text)
}

// HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263.
func (m PartyRiskLimitsRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m PartyRiskLimitsRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m PartyRiskLimitsRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m PartyRiskLimitsRequest) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasNoRequestedPartyRoles returns true if NoRequestedPartyRoles is present, Tag 1508.
func (m PartyRiskLimitsRequest) HasNoRequestedPartyRoles() bool {
	return m.Has(tag.NoRequestedPartyRoles)
}

// HasRiskLimitPlatform returns true if RiskLimitPlatform is present, Tag 1533.
func (m PartyRiskLimitsRequest) HasRiskLimitPlatform() bool {
	return m.Has(tag.RiskLimitPlatform)
}

// HasNoRiskInstrumentScopes returns true if NoRiskInstrumentScopes is present, Tag 1534.
func (m PartyRiskLimitsRequest) HasNoRiskInstrumentScopes() bool {
	return m.Has(tag.NoRiskInstrumentScopes)
}

// HasNoRequestingPartyIDs returns true if NoRequestingPartyIDs is present, Tag 1657.
func (m PartyRiskLimitsRequest) HasNoRequestingPartyIDs() bool {
	return m.Has(tag.NoRequestingPartyIDs)
}

// HasRiskLimitRequestID returns true if RiskLimitRequestID is present, Tag 1666.
func (m PartyRiskLimitsRequest) HasRiskLimitRequestID() bool {
	return m.Has(tag.RiskLimitRequestID)
}

// HasNoRequestedRiskLimitType returns true if NoRequestedRiskLimitType is present, Tag 1668.
func (m PartyRiskLimitsRequest) HasNoRequestedRiskLimitType() bool {
	return m.Has(tag.NoRequestedRiskLimitType)
}

// HasRiskLimitRequestType returns true if RiskLimitRequestType is present, Tag 1760.
func (m PartyRiskLimitsRequest) HasRiskLimitRequestType() bool {
	return m.Has(tag.RiskLimitRequestType)
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

// NoRequestedPartyRoles is a repeating group element, Tag 1508.
type NoRequestedPartyRoles struct {
	*quickfix.Group
}

// SetRequestedPartyRole sets RequestedPartyRole, Tag 1509.
func (m NoRequestedPartyRoles) SetRequestedPartyRole(v int) {
	m.Set(field.NewRequestedPartyRole(v))
}

// SetRequestedPartyRoleQualifier sets RequestedPartyRoleQualifier, Tag 2386.
func (m NoRequestedPartyRoles) SetRequestedPartyRoleQualifier(v int) {
	m.Set(field.NewRequestedPartyRoleQualifier(v))
}

// GetRequestedPartyRole gets RequestedPartyRole, Tag 1509.
func (m NoRequestedPartyRoles) GetRequestedPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.RequestedPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRequestedPartyRoleQualifier gets RequestedPartyRoleQualifier, Tag 2386.
func (m NoRequestedPartyRoles) GetRequestedPartyRoleQualifier() (v int, err quickfix.MessageRejectError) {
	var f field.RequestedPartyRoleQualifierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRequestedPartyRole returns true if RequestedPartyRole is present, Tag 1509.
func (m NoRequestedPartyRoles) HasRequestedPartyRole() bool {
	return m.Has(tag.RequestedPartyRole)
}

// HasRequestedPartyRoleQualifier returns true if RequestedPartyRoleQualifier is present, Tag 2386.
func (m NoRequestedPartyRoles) HasRequestedPartyRoleQualifier() bool {
	return m.Has(tag.RequestedPartyRoleQualifier)
}

// NoRequestedPartyRolesRepeatingGroup is a repeating group, Tag 1508.
type NoRequestedPartyRolesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRequestedPartyRolesRepeatingGroup returns an initialized, NoRequestedPartyRolesRepeatingGroup.
func NewNoRequestedPartyRolesRepeatingGroup() NoRequestedPartyRolesRepeatingGroup {
	return NoRequestedPartyRolesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRequestedPartyRoles,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RequestedPartyRole),
				quickfix.GroupElement(tag.RequestedPartyRoleQualifier),
			},
		),
	}
}

// Add create and append a new NoRequestedPartyRoles to this group.
func (m NoRequestedPartyRolesRepeatingGroup) Add() NoRequestedPartyRoles {
	g := m.RepeatingGroup.Add()
	return NoRequestedPartyRoles{g}
}

// Get returns the ith NoRequestedPartyRoles in the NoRequestedPartyRolesRepeatinGroup.
func (m NoRequestedPartyRolesRepeatingGroup) Get(i int) NoRequestedPartyRoles {
	return NoRequestedPartyRoles{m.RepeatingGroup.Get(i)}
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

// NoRequestedRiskLimitType is a repeating group element, Tag 1668.
type NoRequestedRiskLimitType struct {
	*quickfix.Group
}

// SetRiskLimitType sets RiskLimitType, Tag 1530.
func (m NoRequestedRiskLimitType) SetRiskLimitType(v enum.RiskLimitType) {
	m.Set(field.NewRiskLimitType(v))
}

// GetRiskLimitType gets RiskLimitType, Tag 1530.
func (m NoRequestedRiskLimitType) GetRiskLimitType() (v enum.RiskLimitType, err quickfix.MessageRejectError) {
	var f field.RiskLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRiskLimitType returns true if RiskLimitType is present, Tag 1530.
func (m NoRequestedRiskLimitType) HasRiskLimitType() bool {
	return m.Has(tag.RiskLimitType)
}

// NoRequestedRiskLimitTypeRepeatingGroup is a repeating group, Tag 1668.
type NoRequestedRiskLimitTypeRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRequestedRiskLimitTypeRepeatingGroup returns an initialized, NoRequestedRiskLimitTypeRepeatingGroup.
func NewNoRequestedRiskLimitTypeRepeatingGroup() NoRequestedRiskLimitTypeRepeatingGroup {
	return NoRequestedRiskLimitTypeRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRequestedRiskLimitType,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RiskLimitType),
			},
		),
	}
}

// Add create and append a new NoRequestedRiskLimitType to this group.
func (m NoRequestedRiskLimitTypeRepeatingGroup) Add() NoRequestedRiskLimitType {
	g := m.RepeatingGroup.Add()
	return NoRequestedRiskLimitType{g}
}

// Get returns the ith NoRequestedRiskLimitType in the NoRequestedRiskLimitTypeRepeatinGroup.
func (m NoRequestedRiskLimitTypeRepeatingGroup) Get(i int) NoRequestedRiskLimitType {
	return NoRequestedRiskLimitType{m.RepeatingGroup.Get(i)}
}
