// Code generated by mockery v2.50.2. DO NOT EDIT.

package mocks

import (
	context "context"

	internalclient "github.com/shellhub-io/shellhub/pkg/api/internalclient"
	mock "github.com/stretchr/testify/mock"

	models "github.com/shellhub-io/shellhub/pkg/models"

	requests "github.com/shellhub-io/shellhub/pkg/api/requests"

	websocket "github.com/gorilla/websocket"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// BillingEvaluate provides a mock function with given fields: tenantID
func (_m *Client) BillingEvaluate(tenantID string) (*models.BillingEvaluation, int, error) {
	ret := _m.Called(tenantID)

	if len(ret) == 0 {
		panic("no return value specified for BillingEvaluate")
	}

	var r0 *models.BillingEvaluation
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (*models.BillingEvaluation, int, error)); ok {
		return rf(tenantID)
	}
	if rf, ok := ret.Get(0).(func(string) *models.BillingEvaluation); ok {
		r0 = rf(tenantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.BillingEvaluation)
		}
	}

	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(tenantID)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(tenantID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// BillingReport provides a mock function with given fields: tenant, action
func (_m *Client) BillingReport(tenant string, action string) (int, error) {
	ret := _m.Called(tenant, action)

	if len(ret) == 0 {
		panic("no return value specified for BillingReport")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (int, error)); ok {
		return rf(tenant, action)
	}
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(tenant, action)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(tenant, action)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePrivateKey provides a mock function with no fields
func (_m *Client) CreatePrivateKey() (*models.PrivateKey, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CreatePrivateKey")
	}

	var r0 *models.PrivateKey
	var r1 error
	if rf, ok := ret.Get(0).(func() (*models.PrivateKey, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *models.PrivateKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PrivateKey)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceLookup provides a mock function with given fields: tenantID, name
func (_m *Client) DeviceLookup(tenantID, name string) (*models.Device, error) {
	ret := _m.Called(tenantID, name)

	if len(ret) == 0 {
		panic("no return value specified for DeviceLookup")
	}

	var r0 *models.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*models.Device, error)); ok {
		return rf(tenantID, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) *models.Device); ok {
		r0 = rf(tenantID, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(tenantID, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DevicesHeartbeat provides a mock function with given fields: uid
func (_m *Client) DevicesHeartbeat(uid string) error {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for DevicesHeartbeat")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DevicesOffline provides a mock function with given fields: uid
func (_m *Client) DevicesOffline(uid string) error {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for DevicesOffline")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EvaluateKey provides a mock function with given fields: fingerprint, dev, username
func (_m *Client) EvaluateKey(fingerprint string, dev *models.Device, username string) (bool, error) {
	ret := _m.Called(fingerprint, dev, username)

	if len(ret) == 0 {
		panic("no return value specified for EvaluateKey")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *models.Device, string) (bool, error)); ok {
		return rf(fingerprint, dev, username)
	}
	if rf, ok := ret.Get(0).(func(string, *models.Device, string) bool); ok {
		r0 = rf(fingerprint, dev, username)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, *models.Device, string) error); ok {
		r1 = rf(fingerprint, dev, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventSessionStream provides a mock function with given fields: ctx, uid
func (_m *Client) EventSessionStream(ctx context.Context, uid string) (*websocket.Conn, error) {
	ret := _m.Called(ctx, uid)

	if len(ret) == 0 {
		panic("no return value specified for EventSessionStream")
	}

	var r0 *websocket.Conn
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*websocket.Conn, error)); ok {
		return rf(ctx, uid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *websocket.Conn); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*websocket.Conn)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FinishSession provides a mock function with given fields: uid
func (_m *Client) FinishSession(uid string) []error {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for FinishSession")
	}

	var r0 []error
	if rf, ok := ret.Get(0).(func(string) []error); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}

// FirewallEvaluate provides a mock function with given fields: lookup
func (_m *Client) FirewallEvaluate(lookup map[string]string) error {
	ret := _m.Called(lookup)

	if len(ret) == 0 {
		panic("no return value specified for FirewallEvaluate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string) error); ok {
		r0 = rf(lookup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDevice provides a mock function with given fields: uid
func (_m *Client) GetDevice(uid string) (*models.Device, error) {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for GetDevice")
	}

	var r0 *models.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Device, error)); ok {
		return rf(uid)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Device); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPublicKey provides a mock function with given fields: fingerprint, tenant
func (_m *Client) GetPublicKey(fingerprint string, tenant string) (*models.PublicKey, error) {
	ret := _m.Called(fingerprint, tenant)

	if len(ret) == 0 {
		panic("no return value specified for GetPublicKey")
	}

	var r0 *models.PublicKey
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*models.PublicKey, error)); ok {
		return rf(fingerprint, tenant)
	}
	if rf, ok := ret.Get(0).(func(string, string) *models.PublicKey); ok {
		r0 = rf(fingerprint, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PublicKey)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(fingerprint, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InviteMember provides a mock function with given fields: ctx, tenantID, userID, forwardedHost
func (_m *Client) InviteMember(ctx context.Context, tenantID string, userID string, forwardedHost string) error {
	ret := _m.Called(ctx, tenantID, userID, forwardedHost)

	if len(ret) == 0 {
		panic("no return value specified for InviteMember")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, tenantID, userID, forwardedHost)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KeepAliveSession provides a mock function with given fields: uid
func (_m *Client) KeepAliveSession(uid string) []error {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for KeepAliveSession")
	}

	var r0 []error
	if rf, ok := ret.Get(0).(func(string) []error); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}

// ListDevices provides a mock function with no fields
func (_m *Client) ListDevices() ([]models.Device, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListDevices")
	}

	var r0 []models.Device
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Device, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Lookup provides a mock function with given fields: lookup
func (_m *Client) Lookup(lookup map[string]string) (string, []error) {
	ret := _m.Called(lookup)

	if len(ret) == 0 {
		panic("no return value specified for Lookup")
	}

	var r0 string
	var r1 []error
	if rf, ok := ret.Get(0).(func(map[string]string) (string, []error)); ok {
		return rf(lookup)
	}
	if rf, ok := ret.Get(0).(func(map[string]string) string); ok {
		r0 = rf(lookup)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(map[string]string) []error); ok {
		r1 = rf(lookup)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]error)
		}
	}

	return r0, r1
}

// LookupTunnel provides a mock function with given fields: address
func (_m *Client) LookupTunnel(address string) (*internalclient.Tunnel, error) {
	ret := _m.Called(address)

	if len(ret) == 0 {
		panic("no return value specified for LookupTunnel")
	}

	var r0 *internalclient.Tunnel
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*internalclient.Tunnel, error)); ok {
		return rf(address)
	}
	if rf, ok := ret.Get(0).(func(string) *internalclient.Tunnel); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalclient.Tunnel)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceLookup provides a mock function with given fields: tenant
func (_m *Client) NamespaceLookup(tenant string) (*models.Namespace, []error) {
	ret := _m.Called(tenant)

	if len(ret) == 0 {
		panic("no return value specified for NamespaceLookup")
	}

	var r0 *models.Namespace
	var r1 []error
	if rf, ok := ret.Get(0).(func(string) (*models.Namespace, []error)); ok {
		return rf(tenant)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Namespace); ok {
		r0 = rf(tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(string) []error); ok {
		r1 = rf(tenant)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]error)
		}
	}

	return r0, r1
}

// SaveSession provides a mock function with given fields: uid, seat
func (_m *Client) SaveSession(uid string, seat int) error {
	ret := _m.Called(uid, seat)

	if len(ret) == 0 {
		panic("no return value specified for SaveSession")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(uid, seat)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionAsAuthenticated provides a mock function with given fields: uid
func (_m *Client) SessionAsAuthenticated(uid string) []error {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for SessionAsAuthenticated")
	}

	var r0 []error
	if rf, ok := ret.Get(0).(func(string) []error); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}

// SessionCreate provides a mock function with given fields: session
func (_m *Client) SessionCreate(session requests.SessionCreate) error {
	ret := _m.Called(session)

	if len(ret) == 0 {
		panic("no return value specified for SessionCreate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(requests.SessionCreate) error); ok {
		r0 = rf(session)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSession provides a mock function with given fields: uid, model
func (_m *Client) UpdateSession(uid string, model *models.SessionUpdate) error {
	ret := _m.Called(uid, model)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSession")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *models.SessionUpdate) error); ok {
		r0 = rf(uid, model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
