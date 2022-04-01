package usernotification

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// UserNotification is the fix50sp2 UserNotification type, MsgType = CB.
type UserNotification struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a UserNotification from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) UserNotification {
	return UserNotification{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m UserNotification) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a UserNotification initialized with the required fields for UserNotification.
func New(userstatus field.UserStatusField) (m UserNotification) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("CB"))
	m.Set(userstatus)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg UserNotification, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "CB", r
}

// SetText sets Text, Tag 58.
func (m UserNotification) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m UserNotification) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m UserNotification) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetNoUsernames sets NoUsernames, Tag 809.
func (m UserNotification) SetNoUsernames(f NoUsernamesRepeatingGroup) {
	m.SetGroup(f)
}

// SetUserStatus sets UserStatus, Tag 926.
func (m UserNotification) SetUserStatus(v enum.UserStatus) {
	m.Set(field.NewUserStatus(v))
}

// SetNoThrottles sets NoThrottles, Tag 1610.
func (m UserNotification) SetNoThrottles(f NoThrottlesRepeatingGroup) {
	m.SetGroup(f)
}

// GetText gets Text, Tag 58.
func (m UserNotification) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m UserNotification) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m UserNotification) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUsernames gets NoUsernames, Tag 809.
func (m UserNotification) GetNoUsernames() (NoUsernamesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUsernamesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUserStatus gets UserStatus, Tag 926.
func (m UserNotification) GetUserStatus() (v enum.UserStatus, err quickfix.MessageRejectError) {
	var f field.UserStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoThrottles gets NoThrottles, Tag 1610.
func (m UserNotification) GetNoThrottles() (NoThrottlesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoThrottlesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasText returns true if Text is present, Tag 58.
func (m UserNotification) HasText() bool {
	return m.Has(tag.Text)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m UserNotification) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m UserNotification) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasNoUsernames returns true if NoUsernames is present, Tag 809.
func (m UserNotification) HasNoUsernames() bool {
	return m.Has(tag.NoUsernames)
}

// HasUserStatus returns true if UserStatus is present, Tag 926.
func (m UserNotification) HasUserStatus() bool {
	return m.Has(tag.UserStatus)
}

// HasNoThrottles returns true if NoThrottles is present, Tag 1610.
func (m UserNotification) HasNoThrottles() bool {
	return m.Has(tag.NoThrottles)
}

// NoUsernames is a repeating group element, Tag 809.
type NoUsernames struct {
	*quickfix.Group
}

// SetUsername sets Username, Tag 553.
func (m NoUsernames) SetUsername(v string) {
	m.Set(field.NewUsername(v))
}

// GetUsername gets Username, Tag 553.
func (m NoUsernames) GetUsername() (v string, err quickfix.MessageRejectError) {
	var f field.UsernameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUsername returns true if Username is present, Tag 553.
func (m NoUsernames) HasUsername() bool {
	return m.Has(tag.Username)
}

// NoUsernamesRepeatingGroup is a repeating group, Tag 809.
type NoUsernamesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUsernamesRepeatingGroup returns an initialized, NoUsernamesRepeatingGroup.
func NewNoUsernamesRepeatingGroup() NoUsernamesRepeatingGroup {
	return NoUsernamesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoUsernames,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.Username),
			},
		),
	}
}

// Add create and append a new NoUsernames to this group.
func (m NoUsernamesRepeatingGroup) Add() NoUsernames {
	g := m.RepeatingGroup.Add()
	return NoUsernames{g}
}

// Get returns the ith NoUsernames in the NoUsernamesRepeatinGroup.
func (m NoUsernamesRepeatingGroup) Get(i int) NoUsernames {
	return NoUsernames{m.RepeatingGroup.Get(i)}
}

// NoThrottles is a repeating group element, Tag 1610.
type NoThrottles struct {
	*quickfix.Group
}

// SetThrottleAction sets ThrottleAction, Tag 1611.
func (m NoThrottles) SetThrottleAction(v enum.ThrottleAction) {
	m.Set(field.NewThrottleAction(v))
}

// SetThrottleType sets ThrottleType, Tag 1612.
func (m NoThrottles) SetThrottleType(v enum.ThrottleType) {
	m.Set(field.NewThrottleType(v))
}

// SetThrottleNoMsgs sets ThrottleNoMsgs, Tag 1613.
func (m NoThrottles) SetThrottleNoMsgs(v int) {
	m.Set(field.NewThrottleNoMsgs(v))
}

// SetThrottleTimeInterval sets ThrottleTimeInterval, Tag 1614.
func (m NoThrottles) SetThrottleTimeInterval(v int) {
	m.Set(field.NewThrottleTimeInterval(v))
}

// SetThrottleTimeUnit sets ThrottleTimeUnit, Tag 1615.
func (m NoThrottles) SetThrottleTimeUnit(v int) {
	m.Set(field.NewThrottleTimeUnit(v))
}

// SetNoThrottleMsgType sets NoThrottleMsgType, Tag 1618.
func (m NoThrottles) SetNoThrottleMsgType(f NoThrottleMsgTypeRepeatingGroup) {
	m.SetGroup(f)
}

// GetThrottleAction gets ThrottleAction, Tag 1611.
func (m NoThrottles) GetThrottleAction() (v enum.ThrottleAction, err quickfix.MessageRejectError) {
	var f field.ThrottleActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetThrottleType gets ThrottleType, Tag 1612.
func (m NoThrottles) GetThrottleType() (v enum.ThrottleType, err quickfix.MessageRejectError) {
	var f field.ThrottleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetThrottleNoMsgs gets ThrottleNoMsgs, Tag 1613.
func (m NoThrottles) GetThrottleNoMsgs() (v int, err quickfix.MessageRejectError) {
	var f field.ThrottleNoMsgsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetThrottleTimeInterval gets ThrottleTimeInterval, Tag 1614.
func (m NoThrottles) GetThrottleTimeInterval() (v int, err quickfix.MessageRejectError) {
	var f field.ThrottleTimeIntervalField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetThrottleTimeUnit gets ThrottleTimeUnit, Tag 1615.
func (m NoThrottles) GetThrottleTimeUnit() (v int, err quickfix.MessageRejectError) {
	var f field.ThrottleTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoThrottleMsgType gets NoThrottleMsgType, Tag 1618.
func (m NoThrottles) GetNoThrottleMsgType() (NoThrottleMsgTypeRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoThrottleMsgTypeRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasThrottleAction returns true if ThrottleAction is present, Tag 1611.
func (m NoThrottles) HasThrottleAction() bool {
	return m.Has(tag.ThrottleAction)
}

// HasThrottleType returns true if ThrottleType is present, Tag 1612.
func (m NoThrottles) HasThrottleType() bool {
	return m.Has(tag.ThrottleType)
}

// HasThrottleNoMsgs returns true if ThrottleNoMsgs is present, Tag 1613.
func (m NoThrottles) HasThrottleNoMsgs() bool {
	return m.Has(tag.ThrottleNoMsgs)
}

// HasThrottleTimeInterval returns true if ThrottleTimeInterval is present, Tag 1614.
func (m NoThrottles) HasThrottleTimeInterval() bool {
	return m.Has(tag.ThrottleTimeInterval)
}

// HasThrottleTimeUnit returns true if ThrottleTimeUnit is present, Tag 1615.
func (m NoThrottles) HasThrottleTimeUnit() bool {
	return m.Has(tag.ThrottleTimeUnit)
}

// HasNoThrottleMsgType returns true if NoThrottleMsgType is present, Tag 1618.
func (m NoThrottles) HasNoThrottleMsgType() bool {
	return m.Has(tag.NoThrottleMsgType)
}

// NoThrottleMsgType is a repeating group element, Tag 1618.
type NoThrottleMsgType struct {
	*quickfix.Group
}

// SetThrottleMsgType sets ThrottleMsgType, Tag 1619.
func (m NoThrottleMsgType) SetThrottleMsgType(v string) {
	m.Set(field.NewThrottleMsgType(v))
}

// GetThrottleMsgType gets ThrottleMsgType, Tag 1619.
func (m NoThrottleMsgType) GetThrottleMsgType() (v string, err quickfix.MessageRejectError) {
	var f field.ThrottleMsgTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasThrottleMsgType returns true if ThrottleMsgType is present, Tag 1619.
func (m NoThrottleMsgType) HasThrottleMsgType() bool {
	return m.Has(tag.ThrottleMsgType)
}

// NoThrottleMsgTypeRepeatingGroup is a repeating group, Tag 1618.
type NoThrottleMsgTypeRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoThrottleMsgTypeRepeatingGroup returns an initialized, NoThrottleMsgTypeRepeatingGroup.
func NewNoThrottleMsgTypeRepeatingGroup() NoThrottleMsgTypeRepeatingGroup {
	return NoThrottleMsgTypeRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoThrottleMsgType,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.ThrottleMsgType),
			},
		),
	}
}

// Add create and append a new NoThrottleMsgType to this group.
func (m NoThrottleMsgTypeRepeatingGroup) Add() NoThrottleMsgType {
	g := m.RepeatingGroup.Add()
	return NoThrottleMsgType{g}
}

// Get returns the ith NoThrottleMsgType in the NoThrottleMsgTypeRepeatinGroup.
func (m NoThrottleMsgTypeRepeatingGroup) Get(i int) NoThrottleMsgType {
	return NoThrottleMsgType{m.RepeatingGroup.Get(i)}
}

// NoThrottlesRepeatingGroup is a repeating group, Tag 1610.
type NoThrottlesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoThrottlesRepeatingGroup returns an initialized, NoThrottlesRepeatingGroup.
func NewNoThrottlesRepeatingGroup() NoThrottlesRepeatingGroup {
	return NoThrottlesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoThrottles,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.ThrottleAction),
				quickfix.GroupElement(tag.ThrottleType),
				quickfix.GroupElement(tag.ThrottleNoMsgs),
				quickfix.GroupElement(tag.ThrottleTimeInterval),
				quickfix.GroupElement(tag.ThrottleTimeUnit),
				NewNoThrottleMsgTypeRepeatingGroup(),
			},
		),
	}
}

// Add create and append a new NoThrottles to this group.
func (m NoThrottlesRepeatingGroup) Add() NoThrottles {
	g := m.RepeatingGroup.Add()
	return NoThrottles{g}
}

// Get returns the ith NoThrottles in the NoThrottlesRepeatinGroup.
func (m NoThrottlesRepeatingGroup) Get(i int) NoThrottles {
	return NoThrottles{m.RepeatingGroup.Get(i)}
}
