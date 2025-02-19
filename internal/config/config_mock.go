// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package config

import (
	"sync"
)

// Ensure, that ConfigMock does implement Config.
// If this is not the case, regenerate this file with moq.
var _ Config = &ConfigMock{}

// ConfigMock is a mock implementation of Config.
//
//	func TestSomethingThatUsesConfig(t *testing.T) {
//
//		// make and configure a mocked Config
//		mockedConfig := &ConfigMock{
//			AliasesFunc: func() *AliasConfig {
//				panic("mock out the Aliases method")
//			},
//			AuthTokenFunc: func(s string) (string, string) {
//				panic("mock out the AuthToken method")
//			},
//			DefaultHostFunc: func() (string, string) {
//				panic("mock out the DefaultHost method")
//			},
//			GetFunc: func(s1 string, s2 string) (string, error) {
//				panic("mock out the Get method")
//			},
//			GetOrDefaultFunc: func(s1 string, s2 string) (string, error) {
//				panic("mock out the GetOrDefault method")
//			},
//			HostsFunc: func() []string {
//				panic("mock out the Hosts method")
//			},
//			SetFunc: func(s1 string, s2 string, s3 string)  {
//				panic("mock out the Set method")
//			},
//			UnsetHostFunc: func(s string)  {
//				panic("mock out the UnsetHost method")
//			},
//			WriteFunc: func() error {
//				panic("mock out the Write method")
//			},
//		}
//
//		// use mockedConfig in code that requires Config
//		// and then make assertions.
//
//	}
type ConfigMock struct {
	// AliasesFunc mocks the Aliases method.
	AliasesFunc func() *AliasConfig

	// AuthTokenFunc mocks the AuthToken method.
	AuthTokenFunc func(s string) (string, string)

	// DefaultHostFunc mocks the DefaultHost method.
	DefaultHostFunc func() (string, string)

	// GetFunc mocks the Get method.
	GetFunc func(s1 string, s2 string) (string, error)

	// GetOrDefaultFunc mocks the GetOrDefault method.
	GetOrDefaultFunc func(s1 string, s2 string) (string, error)

	// HostsFunc mocks the Hosts method.
	HostsFunc func() []string

	// SetFunc mocks the Set method.
	SetFunc func(s1 string, s2 string, s3 string)

	// UnsetHostFunc mocks the UnsetHost method.
	UnsetHostFunc func(s string)

	// WriteFunc mocks the Write method.
	WriteFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// Aliases holds details about calls to the Aliases method.
		Aliases []struct {
		}
		// AuthToken holds details about calls to the AuthToken method.
		AuthToken []struct {
			// S is the s argument value.
			S string
		}
		// DefaultHost holds details about calls to the DefaultHost method.
		DefaultHost []struct {
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// S1 is the s1 argument value.
			S1 string
			// S2 is the s2 argument value.
			S2 string
		}
		// GetOrDefault holds details about calls to the GetOrDefault method.
		GetOrDefault []struct {
			// S1 is the s1 argument value.
			S1 string
			// S2 is the s2 argument value.
			S2 string
		}
		// Hosts holds details about calls to the Hosts method.
		Hosts []struct {
		}
		// Set holds details about calls to the Set method.
		Set []struct {
			// S1 is the s1 argument value.
			S1 string
			// S2 is the s2 argument value.
			S2 string
			// S3 is the s3 argument value.
			S3 string
		}
		// UnsetHost holds details about calls to the UnsetHost method.
		UnsetHost []struct {
			// S is the s argument value.
			S string
		}
		// Write holds details about calls to the Write method.
		Write []struct {
		}
	}
	lockAliases      sync.RWMutex
	lockAuthToken    sync.RWMutex
	lockDefaultHost  sync.RWMutex
	lockGet          sync.RWMutex
	lockGetOrDefault sync.RWMutex
	lockHosts        sync.RWMutex
	lockSet          sync.RWMutex
	lockUnsetHost    sync.RWMutex
	lockWrite        sync.RWMutex
}

// Aliases calls AliasesFunc.
func (mock *ConfigMock) Aliases() *AliasConfig {
	if mock.AliasesFunc == nil {
		panic("ConfigMock.AliasesFunc: method is nil but Config.Aliases was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAliases.Lock()
	mock.calls.Aliases = append(mock.calls.Aliases, callInfo)
	mock.lockAliases.Unlock()
	return mock.AliasesFunc()
}

// AliasesCalls gets all the calls that were made to Aliases.
// Check the length with:
//
//	len(mockedConfig.AliasesCalls())
func (mock *ConfigMock) AliasesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAliases.RLock()
	calls = mock.calls.Aliases
	mock.lockAliases.RUnlock()
	return calls
}

// AuthToken calls AuthTokenFunc.
func (mock *ConfigMock) AuthToken(s string) (string, string) {
	if mock.AuthTokenFunc == nil {
		panic("ConfigMock.AuthTokenFunc: method is nil but Config.AuthToken was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockAuthToken.Lock()
	mock.calls.AuthToken = append(mock.calls.AuthToken, callInfo)
	mock.lockAuthToken.Unlock()
	return mock.AuthTokenFunc(s)
}

// AuthTokenCalls gets all the calls that were made to AuthToken.
// Check the length with:
//
//	len(mockedConfig.AuthTokenCalls())
func (mock *ConfigMock) AuthTokenCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockAuthToken.RLock()
	calls = mock.calls.AuthToken
	mock.lockAuthToken.RUnlock()
	return calls
}

// DefaultHost calls DefaultHostFunc.
func (mock *ConfigMock) DefaultHost() (string, string) {
	if mock.DefaultHostFunc == nil {
		panic("ConfigMock.DefaultHostFunc: method is nil but Config.DefaultHost was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDefaultHost.Lock()
	mock.calls.DefaultHost = append(mock.calls.DefaultHost, callInfo)
	mock.lockDefaultHost.Unlock()
	return mock.DefaultHostFunc()
}

// DefaultHostCalls gets all the calls that were made to DefaultHost.
// Check the length with:
//
//	len(mockedConfig.DefaultHostCalls())
func (mock *ConfigMock) DefaultHostCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDefaultHost.RLock()
	calls = mock.calls.DefaultHost
	mock.lockDefaultHost.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *ConfigMock) Get(s1 string, s2 string) (string, error) {
	if mock.GetFunc == nil {
		panic("ConfigMock.GetFunc: method is nil but Config.Get was just called")
	}
	callInfo := struct {
		S1 string
		S2 string
	}{
		S1: s1,
		S2: s2,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(s1, s2)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedConfig.GetCalls())
func (mock *ConfigMock) GetCalls() []struct {
	S1 string
	S2 string
} {
	var calls []struct {
		S1 string
		S2 string
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// GetOrDefault calls GetOrDefaultFunc.
func (mock *ConfigMock) GetOrDefault(s1 string, s2 string) (string, error) {
	if mock.GetOrDefaultFunc == nil {
		panic("ConfigMock.GetOrDefaultFunc: method is nil but Config.GetOrDefault was just called")
	}
	callInfo := struct {
		S1 string
		S2 string
	}{
		S1: s1,
		S2: s2,
	}
	mock.lockGetOrDefault.Lock()
	mock.calls.GetOrDefault = append(mock.calls.GetOrDefault, callInfo)
	mock.lockGetOrDefault.Unlock()
	return mock.GetOrDefaultFunc(s1, s2)
}

// GetOrDefaultCalls gets all the calls that were made to GetOrDefault.
// Check the length with:
//
//	len(mockedConfig.GetOrDefaultCalls())
func (mock *ConfigMock) GetOrDefaultCalls() []struct {
	S1 string
	S2 string
} {
	var calls []struct {
		S1 string
		S2 string
	}
	mock.lockGetOrDefault.RLock()
	calls = mock.calls.GetOrDefault
	mock.lockGetOrDefault.RUnlock()
	return calls
}

// Hosts calls HostsFunc.
func (mock *ConfigMock) Hosts() []string {
	if mock.HostsFunc == nil {
		panic("ConfigMock.HostsFunc: method is nil but Config.Hosts was just called")
	}
	callInfo := struct {
	}{}
	mock.lockHosts.Lock()
	mock.calls.Hosts = append(mock.calls.Hosts, callInfo)
	mock.lockHosts.Unlock()
	return mock.HostsFunc()
}

// HostsCalls gets all the calls that were made to Hosts.
// Check the length with:
//
//	len(mockedConfig.HostsCalls())
func (mock *ConfigMock) HostsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockHosts.RLock()
	calls = mock.calls.Hosts
	mock.lockHosts.RUnlock()
	return calls
}

// Set calls SetFunc.
func (mock *ConfigMock) Set(s1 string, s2 string, s3 string) {
	if mock.SetFunc == nil {
		panic("ConfigMock.SetFunc: method is nil but Config.Set was just called")
	}
	callInfo := struct {
		S1 string
		S2 string
		S3 string
	}{
		S1: s1,
		S2: s2,
		S3: s3,
	}
	mock.lockSet.Lock()
	mock.calls.Set = append(mock.calls.Set, callInfo)
	mock.lockSet.Unlock()
	mock.SetFunc(s1, s2, s3)
}

// SetCalls gets all the calls that were made to Set.
// Check the length with:
//
//	len(mockedConfig.SetCalls())
func (mock *ConfigMock) SetCalls() []struct {
	S1 string
	S2 string
	S3 string
} {
	var calls []struct {
		S1 string
		S2 string
		S3 string
	}
	mock.lockSet.RLock()
	calls = mock.calls.Set
	mock.lockSet.RUnlock()
	return calls
}

// UnsetHost calls UnsetHostFunc.
func (mock *ConfigMock) UnsetHost(s string) {
	if mock.UnsetHostFunc == nil {
		panic("ConfigMock.UnsetHostFunc: method is nil but Config.UnsetHost was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockUnsetHost.Lock()
	mock.calls.UnsetHost = append(mock.calls.UnsetHost, callInfo)
	mock.lockUnsetHost.Unlock()
	mock.UnsetHostFunc(s)
}

// UnsetHostCalls gets all the calls that were made to UnsetHost.
// Check the length with:
//
//	len(mockedConfig.UnsetHostCalls())
func (mock *ConfigMock) UnsetHostCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockUnsetHost.RLock()
	calls = mock.calls.UnsetHost
	mock.lockUnsetHost.RUnlock()
	return calls
}

// Write calls WriteFunc.
func (mock *ConfigMock) Write() error {
	if mock.WriteFunc == nil {
		panic("ConfigMock.WriteFunc: method is nil but Config.Write was just called")
	}
	callInfo := struct {
	}{}
	mock.lockWrite.Lock()
	mock.calls.Write = append(mock.calls.Write, callInfo)
	mock.lockWrite.Unlock()
	return mock.WriteFunc()
}

// WriteCalls gets all the calls that were made to Write.
// Check the length with:
//
//	len(mockedConfig.WriteCalls())
func (mock *ConfigMock) WriteCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockWrite.RLock()
	calls = mock.calls.Write
	mock.lockWrite.RUnlock()
	return calls
}
