package ordercancelreject

import (
	"time"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// OrderCancelReject is the fix50sp2 OrderCancelReject type, MsgType = 9.
type OrderCancelReject struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a OrderCancelReject from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) OrderCancelReject {
	return OrderCancelReject{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m OrderCancelReject) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a OrderCancelReject initialized with the required fields for OrderCancelReject.
func New(orderid field.OrderIDField, clordid field.ClOrdIDField, ordstatus field.OrdStatusField, cxlrejresponseto field.CxlRejResponseToField) (m OrderCancelReject) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("9"))
	m.Set(orderid)
	m.Set(clordid)
	m.Set(ordstatus)
	m.Set(cxlrejresponseto)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg OrderCancelReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "9", r
}

// SetAccount sets Account, Tag 1.
func (m OrderCancelReject) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

// SetClOrdID sets ClOrdID, Tag 11.
func (m OrderCancelReject) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

// SetOrderID sets OrderID, Tag 37.
func (m OrderCancelReject) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

// SetOrdStatus sets OrdStatus, Tag 39.
func (m OrderCancelReject) SetOrdStatus(v enum.OrdStatus) {
	m.Set(field.NewOrdStatus(v))
}

// SetOrigClOrdID sets OrigClOrdID, Tag 41.
func (m OrderCancelReject) SetOrigClOrdID(v string) {
	m.Set(field.NewOrigClOrdID(v))
}

// SetText sets Text, Tag 58.
func (m OrderCancelReject) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m OrderCancelReject) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetListID sets ListID, Tag 66.
func (m OrderCancelReject) SetListID(v string) {
	m.Set(field.NewListID(v))
}

// SetTradeDate sets TradeDate, Tag 75.
func (m OrderCancelReject) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

// SetExDestination sets ExDestination, Tag 100.
func (m OrderCancelReject) SetExDestination(v enum.ExDestination) {
	m.Set(field.NewExDestination(v))
}

// SetCxlRejReason sets CxlRejReason, Tag 102.
func (m OrderCancelReject) SetCxlRejReason(v enum.CxlRejReason) {
	m.Set(field.NewCxlRejReason(v))
}

// SetSecondaryOrderID sets SecondaryOrderID, Tag 198.
func (m OrderCancelReject) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

// SetTradeOriginationDate sets TradeOriginationDate, Tag 229.
func (m OrderCancelReject) SetTradeOriginationDate(v string) {
	m.Set(field.NewTradeOriginationDate(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m OrderCancelReject) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m OrderCancelReject) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetCxlRejResponseTo sets CxlRejResponseTo, Tag 434.
func (m OrderCancelReject) SetCxlRejResponseTo(v enum.CxlRejResponseTo) {
	m.Set(field.NewCxlRejResponseTo(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m OrderCancelReject) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526.
func (m OrderCancelReject) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

// SetAccountType sets AccountType, Tag 581.
func (m OrderCancelReject) SetAccountType(v enum.AccountType) {
	m.Set(field.NewAccountType(v))
}

// SetClOrdLinkID sets ClOrdLinkID, Tag 583.
func (m OrderCancelReject) SetClOrdLinkID(v string) {
	m.Set(field.NewClOrdLinkID(v))
}

// SetOrigOrdModTime sets OrigOrdModTime, Tag 586.
func (m OrderCancelReject) SetOrigOrdModTime(v time.Time) {
	m.Set(field.NewOrigOrdModTime(v))
}

// SetWorkingIndicator sets WorkingIndicator, Tag 636.
func (m OrderCancelReject) SetWorkingIndicator(v bool) {
	m.Set(field.NewWorkingIndicator(v))
}

// SetAcctIDSource sets AcctIDSource, Tag 660.
func (m OrderCancelReject) SetAcctIDSource(v enum.AcctIDSource) {
	m.Set(field.NewAcctIDSource(v))
}

// SetExDestinationIDSource sets ExDestinationIDSource, Tag 1133.
func (m OrderCancelReject) SetExDestinationIDSource(v enum.ExDestinationIDSource) {
	m.Set(field.NewExDestinationIDSource(v))
}

// SetRejectText sets RejectText, Tag 1328.
func (m OrderCancelReject) SetRejectText(v string) {
	m.Set(field.NewRejectText(v))
}

// SetEncodedRejectTextLen sets EncodedRejectTextLen, Tag 1664.
func (m OrderCancelReject) SetEncodedRejectTextLen(v int) {
	m.Set(field.NewEncodedRejectTextLen(v))
}

// SetEncodedRejectText sets EncodedRejectText, Tag 1665.
func (m OrderCancelReject) SetEncodedRejectText(v string) {
	m.Set(field.NewEncodedRejectText(v))
}

// SetOrderRequestID sets OrderRequestID, Tag 2422.
func (m OrderCancelReject) SetOrderRequestID(v int) {
	m.Set(field.NewOrderRequestID(v))
}

// GetAccount gets Account, Tag 1.
func (m OrderCancelReject) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetClOrdID gets ClOrdID, Tag 11.
func (m OrderCancelReject) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderID gets OrderID, Tag 37.
func (m OrderCancelReject) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrdStatus gets OrdStatus, Tag 39.
func (m OrderCancelReject) GetOrdStatus() (v enum.OrdStatus, err quickfix.MessageRejectError) {
	var f field.OrdStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrigClOrdID gets OrigClOrdID, Tag 41.
func (m OrderCancelReject) GetOrigClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.OrigClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58.
func (m OrderCancelReject) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m OrderCancelReject) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetListID gets ListID, Tag 66.
func (m OrderCancelReject) GetListID() (v string, err quickfix.MessageRejectError) {
	var f field.ListIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeDate gets TradeDate, Tag 75.
func (m OrderCancelReject) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExDestination gets ExDestination, Tag 100.
func (m OrderCancelReject) GetExDestination() (v enum.ExDestination, err quickfix.MessageRejectError) {
	var f field.ExDestinationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCxlRejReason gets CxlRejReason, Tag 102.
func (m OrderCancelReject) GetCxlRejReason() (v enum.CxlRejReason, err quickfix.MessageRejectError) {
	var f field.CxlRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecondaryOrderID gets SecondaryOrderID, Tag 198.
func (m OrderCancelReject) GetSecondaryOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryOrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeOriginationDate gets TradeOriginationDate, Tag 229.
func (m OrderCancelReject) GetTradeOriginationDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeOriginationDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m OrderCancelReject) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m OrderCancelReject) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCxlRejResponseTo gets CxlRejResponseTo, Tag 434.
func (m OrderCancelReject) GetCxlRejResponseTo() (v enum.CxlRejResponseTo, err quickfix.MessageRejectError) {
	var f field.CxlRejResponseToField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m OrderCancelReject) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526.
func (m OrderCancelReject) GetSecondaryClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAccountType gets AccountType, Tag 581.
func (m OrderCancelReject) GetAccountType() (v enum.AccountType, err quickfix.MessageRejectError) {
	var f field.AccountTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetClOrdLinkID gets ClOrdLinkID, Tag 583.
func (m OrderCancelReject) GetClOrdLinkID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdLinkIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrigOrdModTime gets OrigOrdModTime, Tag 586.
func (m OrderCancelReject) GetOrigOrdModTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.OrigOrdModTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetWorkingIndicator gets WorkingIndicator, Tag 636.
func (m OrderCancelReject) GetWorkingIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.WorkingIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAcctIDSource gets AcctIDSource, Tag 660.
func (m OrderCancelReject) GetAcctIDSource() (v enum.AcctIDSource, err quickfix.MessageRejectError) {
	var f field.AcctIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExDestinationIDSource gets ExDestinationIDSource, Tag 1133.
func (m OrderCancelReject) GetExDestinationIDSource() (v enum.ExDestinationIDSource, err quickfix.MessageRejectError) {
	var f field.ExDestinationIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRejectText gets RejectText, Tag 1328.
func (m OrderCancelReject) GetRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.RejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedRejectTextLen gets EncodedRejectTextLen, Tag 1664.
func (m OrderCancelReject) GetEncodedRejectTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedRejectText gets EncodedRejectText, Tag 1665.
func (m OrderCancelReject) GetEncodedRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderRequestID gets OrderRequestID, Tag 2422.
func (m OrderCancelReject) GetOrderRequestID() (v int, err quickfix.MessageRejectError) {
	var f field.OrderRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAccount returns true if Account is present, Tag 1.
func (m OrderCancelReject) HasAccount() bool {
	return m.Has(tag.Account)
}

// HasClOrdID returns true if ClOrdID is present, Tag 11.
func (m OrderCancelReject) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

// HasOrderID returns true if OrderID is present, Tag 37.
func (m OrderCancelReject) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

// HasOrdStatus returns true if OrdStatus is present, Tag 39.
func (m OrderCancelReject) HasOrdStatus() bool {
	return m.Has(tag.OrdStatus)
}

// HasOrigClOrdID returns true if OrigClOrdID is present, Tag 41.
func (m OrderCancelReject) HasOrigClOrdID() bool {
	return m.Has(tag.OrigClOrdID)
}

// HasText returns true if Text is present, Tag 58.
func (m OrderCancelReject) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m OrderCancelReject) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasListID returns true if ListID is present, Tag 66.
func (m OrderCancelReject) HasListID() bool {
	return m.Has(tag.ListID)
}

// HasTradeDate returns true if TradeDate is present, Tag 75.
func (m OrderCancelReject) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

// HasExDestination returns true if ExDestination is present, Tag 100.
func (m OrderCancelReject) HasExDestination() bool {
	return m.Has(tag.ExDestination)
}

// HasCxlRejReason returns true if CxlRejReason is present, Tag 102.
func (m OrderCancelReject) HasCxlRejReason() bool {
	return m.Has(tag.CxlRejReason)
}

// HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198.
func (m OrderCancelReject) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

// HasTradeOriginationDate returns true if TradeOriginationDate is present, Tag 229.
func (m OrderCancelReject) HasTradeOriginationDate() bool {
	return m.Has(tag.TradeOriginationDate)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m OrderCancelReject) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m OrderCancelReject) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasCxlRejResponseTo returns true if CxlRejResponseTo is present, Tag 434.
func (m OrderCancelReject) HasCxlRejResponseTo() bool {
	return m.Has(tag.CxlRejResponseTo)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m OrderCancelReject) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526.
func (m OrderCancelReject) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

// HasAccountType returns true if AccountType is present, Tag 581.
func (m OrderCancelReject) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

// HasClOrdLinkID returns true if ClOrdLinkID is present, Tag 583.
func (m OrderCancelReject) HasClOrdLinkID() bool {
	return m.Has(tag.ClOrdLinkID)
}

// HasOrigOrdModTime returns true if OrigOrdModTime is present, Tag 586.
func (m OrderCancelReject) HasOrigOrdModTime() bool {
	return m.Has(tag.OrigOrdModTime)
}

// HasWorkingIndicator returns true if WorkingIndicator is present, Tag 636.
func (m OrderCancelReject) HasWorkingIndicator() bool {
	return m.Has(tag.WorkingIndicator)
}

// HasAcctIDSource returns true if AcctIDSource is present, Tag 660.
func (m OrderCancelReject) HasAcctIDSource() bool {
	return m.Has(tag.AcctIDSource)
}

// HasExDestinationIDSource returns true if ExDestinationIDSource is present, Tag 1133.
func (m OrderCancelReject) HasExDestinationIDSource() bool {
	return m.Has(tag.ExDestinationIDSource)
}

// HasRejectText returns true if RejectText is present, Tag 1328.
func (m OrderCancelReject) HasRejectText() bool {
	return m.Has(tag.RejectText)
}

// HasEncodedRejectTextLen returns true if EncodedRejectTextLen is present, Tag 1664.
func (m OrderCancelReject) HasEncodedRejectTextLen() bool {
	return m.Has(tag.EncodedRejectTextLen)
}

// HasEncodedRejectText returns true if EncodedRejectText is present, Tag 1665.
func (m OrderCancelReject) HasEncodedRejectText() bool {
	return m.Has(tag.EncodedRejectText)
}

// HasOrderRequestID returns true if OrderRequestID is present, Tag 2422.
func (m OrderCancelReject) HasOrderRequestID() bool {
	return m.Has(tag.OrderRequestID)
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
