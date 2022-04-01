package userresponse

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// UserResponse is the fix50sp2 UserResponse type, MsgType = BF.
type UserResponse struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a UserResponse from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) UserResponse {
	return UserResponse{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m UserResponse) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a UserResponse initialized with the required fields for UserResponse.
func New(userrequestid field.UserRequestIDField, username field.UsernameField) (m UserResponse) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BF"))
	m.Set(userrequestid)
	m.Set(username)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg UserResponse, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "BF", r
}

// SetUsername sets Username, Tag 553.
func (m UserResponse) SetUsername(v string) {
	m.Set(field.NewUsername(v))
}

// SetUserRequestID sets UserRequestID, Tag 923.
func (m UserResponse) SetUserRequestID(v string) {
	m.Set(field.NewUserRequestID(v))
}

// SetUserStatus sets UserStatus, Tag 926.
func (m UserResponse) SetUserStatus(v enum.UserStatus) {
	m.Set(field.NewUserStatus(v))
}

// SetUserStatusText sets UserStatusText, Tag 927.
func (m UserResponse) SetUserStatusText(v string) {
	m.Set(field.NewUserStatusText(v))
}

// SetNoThrottles sets NoThrottles, Tag 1610.
func (m UserResponse) SetNoThrottles(f NoThrottlesRepeatingGroup) {
	m.SetGroup(f)
}

// GetUsername gets Username, Tag 553.
func (m UserResponse) GetUsername() (v string, err quickfix.MessageRejectError) {
	var f field.UsernameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUserRequestID gets UserRequestID, Tag 923.
func (m UserResponse) GetUserRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.UserRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUserStatus gets UserStatus, Tag 926.
func (m UserResponse) GetUserStatus() (v enum.UserStatus, err quickfix.MessageRejectError) {
	var f field.UserStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUserStatusText gets UserStatusText, Tag 927.
func (m UserResponse) GetUserStatusText() (v string, err quickfix.MessageRejectError) {
	var f field.UserStatusTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoThrottles gets NoThrottles, Tag 1610.
func (m UserResponse) GetNoThrottles() (NoThrottlesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoThrottlesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasUsername returns true if Username is present, Tag 553.
func (m UserResponse) HasUsername() bool {
	return m.Has(tag.Username)
}

// HasUserRequestID returns true if UserRequestID is present, Tag 923.
func (m UserResponse) HasUserRequestID() bool {
	return m.Has(tag.UserRequestID)
}

// HasUserStatus returns true if UserStatus is present, Tag 926.
func (m UserResponse) HasUserStatus() bool {
	return m.Has(tag.UserStatus)
}

// HasUserStatusText returns true if UserStatusText is present, Tag 927.
func (m UserResponse) HasUserStatusText() bool {
	return m.Has(tag.UserStatusText)
}

// HasNoThrottles returns true if NoThrottles is present, Tag 1610.
func (m UserResponse) HasNoThrottles() bool {
	return m.Has(tag.NoThrottles)
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
