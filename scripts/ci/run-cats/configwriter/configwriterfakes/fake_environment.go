// This file was generated by counterfeiter
package configwriterfakes

import (
	"sync"

	"github.com/cloudfoundry/runtime-ci/scripts/ci/run-cats/configwriter"
)

type FakeEnvironment struct {
	GetSkipSSLValidationStub        func() (bool, error)
	getSkipSSLValidationMutex       sync.RWMutex
	getSkipSSLValidationArgsForCall []struct{}
	getSkipSSLValidationReturns     struct {
		result1 bool
		result2 error
	}
	GetUseHTTPStub        func() (bool, error)
	getUseHTTPMutex       sync.RWMutex
	getUseHTTPArgsForCall []struct{}
	getUseHTTPReturns     struct {
		result1 bool
		result2 error
	}
	GetIncludePrivilegedContainerSupportStub        func() (bool, error)
	getIncludePrivilegedContainerSupportMutex       sync.RWMutex
	getIncludePrivilegedContainerSupportArgsForCall []struct{}
	getIncludePrivilegedContainerSupportReturns     struct {
		result1 bool
		result2 error
	}
	GetDefaultTimeoutInSecondsStub        func() (int, error)
	getDefaultTimeoutInSecondsMutex       sync.RWMutex
	getDefaultTimeoutInSecondsArgsForCall []struct{}
	getDefaultTimeoutInSecondsReturns     struct {
		result1 int
		result2 error
	}
	GetCFPushTimeoutInSecondsStub        func() (int, error)
	getCFPushTimeoutInSecondsMutex       sync.RWMutex
	getCFPushTimeoutInSecondsArgsForCall []struct{}
	getCFPushTimeoutInSecondsReturns     struct {
		result1 int
		result2 error
	}
	GetLongCurlTimeoutInSecondsStub        func() (int, error)
	getLongCurlTimeoutInSecondsMutex       sync.RWMutex
	getLongCurlTimeoutInSecondsArgsForCall []struct{}
	getLongCurlTimeoutInSecondsReturns     struct {
		result1 int
		result2 error
	}
	GetBrokerStartTimeoutInSecondsStub        func() (int, error)
	getBrokerStartTimeoutInSecondsMutex       sync.RWMutex
	getBrokerStartTimeoutInSecondsArgsForCall []struct{}
	getBrokerStartTimeoutInSecondsReturns     struct {
		result1 int
		result2 error
	}
	GetCFAPIStub        func() string
	getCFAPIMutex       sync.RWMutex
	getCFAPIArgsForCall []struct{}
	getCFAPIReturns     struct {
		result1 string
	}
	GetCFAdminUserStub        func() string
	getCFAdminUserMutex       sync.RWMutex
	getCFAdminUserArgsForCall []struct{}
	getCFAdminUserReturns     struct {
		result1 string
	}
	GetCFAdminPasswordStub        func() string
	getCFAdminPasswordMutex       sync.RWMutex
	getCFAdminPasswordArgsForCall []struct{}
	getCFAdminPasswordReturns     struct {
		result1 string
	}
	GetCFAppsDomainStub        func() string
	getCFAppsDomainMutex       sync.RWMutex
	getCFAppsDomainArgsForCall []struct{}
	getCFAppsDomainReturns     struct {
		result1 string
	}
	GetExistingUserStub        func() string
	getExistingUserMutex       sync.RWMutex
	getExistingUserArgsForCall []struct{}
	getExistingUserReturns     struct {
		result1 string
	}
	UseExistingUserStub        func() bool
	useExistingUserMutex       sync.RWMutex
	useExistingUserArgsForCall []struct{}
	useExistingUserReturns     struct {
		result1 bool
	}
	KeepUserAtSuiteEndStub        func() bool
	keepUserAtSuiteEndMutex       sync.RWMutex
	keepUserAtSuiteEndArgsForCall []struct{}
	keepUserAtSuiteEndReturns     struct {
		result1 bool
	}
	GetExistingUserPasswordStub        func() string
	getExistingUserPasswordMutex       sync.RWMutex
	getExistingUserPasswordArgsForCall []struct{}
	getExistingUserPasswordReturns     struct {
		result1 string
	}
	GetStaticBuildpackNameStub        func() string
	getStaticBuildpackNameMutex       sync.RWMutex
	getStaticBuildpackNameArgsForCall []struct{}
	getStaticBuildpackNameReturns     struct {
		result1 string
	}
	GetJavaBuildpackNameStub        func() string
	getJavaBuildpackNameMutex       sync.RWMutex
	getJavaBuildpackNameArgsForCall []struct{}
	getJavaBuildpackNameReturns     struct {
		result1 string
	}
	GetRubyBuildpackNameStub        func() string
	getRubyBuildpackNameMutex       sync.RWMutex
	getRubyBuildpackNameArgsForCall []struct{}
	getRubyBuildpackNameReturns     struct {
		result1 string
	}
	GetNodeJSBuildpackNameStub        func() string
	getNodeJSBuildpackNameMutex       sync.RWMutex
	getNodeJSBuildpackNameArgsForCall []struct{}
	getNodeJSBuildpackNameReturns     struct {
		result1 string
	}
	GetGoBuildpackNameStub        func() string
	getGoBuildpackNameMutex       sync.RWMutex
	getGoBuildpackNameArgsForCall []struct{}
	getGoBuildpackNameReturns     struct {
		result1 string
	}
	GetPythonBuildpackNameStub        func() string
	getPythonBuildpackNameMutex       sync.RWMutex
	getPythonBuildpackNameArgsForCall []struct{}
	getPythonBuildpackNameReturns     struct {
		result1 string
	}
	GetPHPBuildpackNameStub        func() string
	getPHPBuildpackNameMutex       sync.RWMutex
	getPHPBuildpackNameArgsForCall []struct{}
	getPHPBuildpackNameReturns     struct {
		result1 string
	}
	GetBinaryBuildpackNameStub        func() string
	getBinaryBuildpackNameMutex       sync.RWMutex
	getBinaryBuildpackNameArgsForCall []struct{}
	getBinaryBuildpackNameReturns     struct {
		result1 string
	}
	GetPersistentAppHostStub        func() string
	getPersistentAppHostMutex       sync.RWMutex
	getPersistentAppHostArgsForCall []struct{}
	getPersistentAppHostReturns     struct {
		result1 string
	}
	GetPersistentAppSpaceStub        func() string
	getPersistentAppSpaceMutex       sync.RWMutex
	getPersistentAppSpaceArgsForCall []struct{}
	getPersistentAppSpaceReturns     struct {
		result1 string
	}
	GetPersistentAppOrgStub        func() string
	getPersistentAppOrgMutex       sync.RWMutex
	getPersistentAppOrgArgsForCall []struct{}
	getPersistentAppOrgReturns     struct {
		result1 string
	}
	GetPersistentAppQuotaNameStub        func() string
	getPersistentAppQuotaNameMutex       sync.RWMutex
	getPersistentAppQuotaNameArgsForCall []struct{}
	getPersistentAppQuotaNameReturns     struct {
		result1 string
	}
	GetBackendStub        func() (string, error)
	getBackendMutex       sync.RWMutex
	getBackendArgsForCall []struct{}
	getBackendReturns     struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEnvironment) GetSkipSSLValidation() (bool, error) {
	fake.getSkipSSLValidationMutex.Lock()
	fake.getSkipSSLValidationArgsForCall = append(fake.getSkipSSLValidationArgsForCall, struct{}{})
	fake.recordInvocation("GetSkipSSLValidation", []interface{}{})
	fake.getSkipSSLValidationMutex.Unlock()
	if fake.GetSkipSSLValidationStub != nil {
		return fake.GetSkipSSLValidationStub()
	} else {
		return fake.getSkipSSLValidationReturns.result1, fake.getSkipSSLValidationReturns.result2
	}
}

func (fake *FakeEnvironment) GetSkipSSLValidationCallCount() int {
	fake.getSkipSSLValidationMutex.RLock()
	defer fake.getSkipSSLValidationMutex.RUnlock()
	return len(fake.getSkipSSLValidationArgsForCall)
}

func (fake *FakeEnvironment) GetSkipSSLValidationReturns(result1 bool, result2 error) {
	fake.GetSkipSSLValidationStub = nil
	fake.getSkipSSLValidationReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) GetUseHTTP() (bool, error) {
	fake.getUseHTTPMutex.Lock()
	fake.getUseHTTPArgsForCall = append(fake.getUseHTTPArgsForCall, struct{}{})
	fake.recordInvocation("GetUseHTTP", []interface{}{})
	fake.getUseHTTPMutex.Unlock()
	if fake.GetUseHTTPStub != nil {
		return fake.GetUseHTTPStub()
	} else {
		return fake.getUseHTTPReturns.result1, fake.getUseHTTPReturns.result2
	}
}

func (fake *FakeEnvironment) GetUseHTTPCallCount() int {
	fake.getUseHTTPMutex.RLock()
	defer fake.getUseHTTPMutex.RUnlock()
	return len(fake.getUseHTTPArgsForCall)
}

func (fake *FakeEnvironment) GetUseHTTPReturns(result1 bool, result2 error) {
	fake.GetUseHTTPStub = nil
	fake.getUseHTTPReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) GetIncludePrivilegedContainerSupport() (bool, error) {
	fake.getIncludePrivilegedContainerSupportMutex.Lock()
	fake.getIncludePrivilegedContainerSupportArgsForCall = append(fake.getIncludePrivilegedContainerSupportArgsForCall, struct{}{})
	fake.recordInvocation("GetIncludePrivilegedContainerSupport", []interface{}{})
	fake.getIncludePrivilegedContainerSupportMutex.Unlock()
	if fake.GetIncludePrivilegedContainerSupportStub != nil {
		return fake.GetIncludePrivilegedContainerSupportStub()
	} else {
		return fake.getIncludePrivilegedContainerSupportReturns.result1, fake.getIncludePrivilegedContainerSupportReturns.result2
	}
}

func (fake *FakeEnvironment) GetIncludePrivilegedContainerSupportCallCount() int {
	fake.getIncludePrivilegedContainerSupportMutex.RLock()
	defer fake.getIncludePrivilegedContainerSupportMutex.RUnlock()
	return len(fake.getIncludePrivilegedContainerSupportArgsForCall)
}

func (fake *FakeEnvironment) GetIncludePrivilegedContainerSupportReturns(result1 bool, result2 error) {
	fake.GetIncludePrivilegedContainerSupportStub = nil
	fake.getIncludePrivilegedContainerSupportReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) GetDefaultTimeoutInSeconds() (int, error) {
	fake.getDefaultTimeoutInSecondsMutex.Lock()
	fake.getDefaultTimeoutInSecondsArgsForCall = append(fake.getDefaultTimeoutInSecondsArgsForCall, struct{}{})
	fake.recordInvocation("GetDefaultTimeoutInSeconds", []interface{}{})
	fake.getDefaultTimeoutInSecondsMutex.Unlock()
	if fake.GetDefaultTimeoutInSecondsStub != nil {
		return fake.GetDefaultTimeoutInSecondsStub()
	} else {
		return fake.getDefaultTimeoutInSecondsReturns.result1, fake.getDefaultTimeoutInSecondsReturns.result2
	}
}

func (fake *FakeEnvironment) GetDefaultTimeoutInSecondsCallCount() int {
	fake.getDefaultTimeoutInSecondsMutex.RLock()
	defer fake.getDefaultTimeoutInSecondsMutex.RUnlock()
	return len(fake.getDefaultTimeoutInSecondsArgsForCall)
}

func (fake *FakeEnvironment) GetDefaultTimeoutInSecondsReturns(result1 int, result2 error) {
	fake.GetDefaultTimeoutInSecondsStub = nil
	fake.getDefaultTimeoutInSecondsReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) GetCFPushTimeoutInSeconds() (int, error) {
	fake.getCFPushTimeoutInSecondsMutex.Lock()
	fake.getCFPushTimeoutInSecondsArgsForCall = append(fake.getCFPushTimeoutInSecondsArgsForCall, struct{}{})
	fake.recordInvocation("GetCFPushTimeoutInSeconds", []interface{}{})
	fake.getCFPushTimeoutInSecondsMutex.Unlock()
	if fake.GetCFPushTimeoutInSecondsStub != nil {
		return fake.GetCFPushTimeoutInSecondsStub()
	} else {
		return fake.getCFPushTimeoutInSecondsReturns.result1, fake.getCFPushTimeoutInSecondsReturns.result2
	}
}

func (fake *FakeEnvironment) GetCFPushTimeoutInSecondsCallCount() int {
	fake.getCFPushTimeoutInSecondsMutex.RLock()
	defer fake.getCFPushTimeoutInSecondsMutex.RUnlock()
	return len(fake.getCFPushTimeoutInSecondsArgsForCall)
}

func (fake *FakeEnvironment) GetCFPushTimeoutInSecondsReturns(result1 int, result2 error) {
	fake.GetCFPushTimeoutInSecondsStub = nil
	fake.getCFPushTimeoutInSecondsReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) GetLongCurlTimeoutInSeconds() (int, error) {
	fake.getLongCurlTimeoutInSecondsMutex.Lock()
	fake.getLongCurlTimeoutInSecondsArgsForCall = append(fake.getLongCurlTimeoutInSecondsArgsForCall, struct{}{})
	fake.recordInvocation("GetLongCurlTimeoutInSeconds", []interface{}{})
	fake.getLongCurlTimeoutInSecondsMutex.Unlock()
	if fake.GetLongCurlTimeoutInSecondsStub != nil {
		return fake.GetLongCurlTimeoutInSecondsStub()
	} else {
		return fake.getLongCurlTimeoutInSecondsReturns.result1, fake.getLongCurlTimeoutInSecondsReturns.result2
	}
}

func (fake *FakeEnvironment) GetLongCurlTimeoutInSecondsCallCount() int {
	fake.getLongCurlTimeoutInSecondsMutex.RLock()
	defer fake.getLongCurlTimeoutInSecondsMutex.RUnlock()
	return len(fake.getLongCurlTimeoutInSecondsArgsForCall)
}

func (fake *FakeEnvironment) GetLongCurlTimeoutInSecondsReturns(result1 int, result2 error) {
	fake.GetLongCurlTimeoutInSecondsStub = nil
	fake.getLongCurlTimeoutInSecondsReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) GetBrokerStartTimeoutInSeconds() (int, error) {
	fake.getBrokerStartTimeoutInSecondsMutex.Lock()
	fake.getBrokerStartTimeoutInSecondsArgsForCall = append(fake.getBrokerStartTimeoutInSecondsArgsForCall, struct{}{})
	fake.recordInvocation("GetBrokerStartTimeoutInSeconds", []interface{}{})
	fake.getBrokerStartTimeoutInSecondsMutex.Unlock()
	if fake.GetBrokerStartTimeoutInSecondsStub != nil {
		return fake.GetBrokerStartTimeoutInSecondsStub()
	} else {
		return fake.getBrokerStartTimeoutInSecondsReturns.result1, fake.getBrokerStartTimeoutInSecondsReturns.result2
	}
}

func (fake *FakeEnvironment) GetBrokerStartTimeoutInSecondsCallCount() int {
	fake.getBrokerStartTimeoutInSecondsMutex.RLock()
	defer fake.getBrokerStartTimeoutInSecondsMutex.RUnlock()
	return len(fake.getBrokerStartTimeoutInSecondsArgsForCall)
}

func (fake *FakeEnvironment) GetBrokerStartTimeoutInSecondsReturns(result1 int, result2 error) {
	fake.GetBrokerStartTimeoutInSecondsStub = nil
	fake.getBrokerStartTimeoutInSecondsReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) GetCFAPI() string {
	fake.getCFAPIMutex.Lock()
	fake.getCFAPIArgsForCall = append(fake.getCFAPIArgsForCall, struct{}{})
	fake.recordInvocation("GetCFAPI", []interface{}{})
	fake.getCFAPIMutex.Unlock()
	if fake.GetCFAPIStub != nil {
		return fake.GetCFAPIStub()
	} else {
		return fake.getCFAPIReturns.result1
	}
}

func (fake *FakeEnvironment) GetCFAPICallCount() int {
	fake.getCFAPIMutex.RLock()
	defer fake.getCFAPIMutex.RUnlock()
	return len(fake.getCFAPIArgsForCall)
}

func (fake *FakeEnvironment) GetCFAPIReturns(result1 string) {
	fake.GetCFAPIStub = nil
	fake.getCFAPIReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetCFAdminUser() string {
	fake.getCFAdminUserMutex.Lock()
	fake.getCFAdminUserArgsForCall = append(fake.getCFAdminUserArgsForCall, struct{}{})
	fake.recordInvocation("GetCFAdminUser", []interface{}{})
	fake.getCFAdminUserMutex.Unlock()
	if fake.GetCFAdminUserStub != nil {
		return fake.GetCFAdminUserStub()
	} else {
		return fake.getCFAdminUserReturns.result1
	}
}

func (fake *FakeEnvironment) GetCFAdminUserCallCount() int {
	fake.getCFAdminUserMutex.RLock()
	defer fake.getCFAdminUserMutex.RUnlock()
	return len(fake.getCFAdminUserArgsForCall)
}

func (fake *FakeEnvironment) GetCFAdminUserReturns(result1 string) {
	fake.GetCFAdminUserStub = nil
	fake.getCFAdminUserReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetCFAdminPassword() string {
	fake.getCFAdminPasswordMutex.Lock()
	fake.getCFAdminPasswordArgsForCall = append(fake.getCFAdminPasswordArgsForCall, struct{}{})
	fake.recordInvocation("GetCFAdminPassword", []interface{}{})
	fake.getCFAdminPasswordMutex.Unlock()
	if fake.GetCFAdminPasswordStub != nil {
		return fake.GetCFAdminPasswordStub()
	} else {
		return fake.getCFAdminPasswordReturns.result1
	}
}

func (fake *FakeEnvironment) GetCFAdminPasswordCallCount() int {
	fake.getCFAdminPasswordMutex.RLock()
	defer fake.getCFAdminPasswordMutex.RUnlock()
	return len(fake.getCFAdminPasswordArgsForCall)
}

func (fake *FakeEnvironment) GetCFAdminPasswordReturns(result1 string) {
	fake.GetCFAdminPasswordStub = nil
	fake.getCFAdminPasswordReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetCFAppsDomain() string {
	fake.getCFAppsDomainMutex.Lock()
	fake.getCFAppsDomainArgsForCall = append(fake.getCFAppsDomainArgsForCall, struct{}{})
	fake.recordInvocation("GetCFAppsDomain", []interface{}{})
	fake.getCFAppsDomainMutex.Unlock()
	if fake.GetCFAppsDomainStub != nil {
		return fake.GetCFAppsDomainStub()
	} else {
		return fake.getCFAppsDomainReturns.result1
	}
}

func (fake *FakeEnvironment) GetCFAppsDomainCallCount() int {
	fake.getCFAppsDomainMutex.RLock()
	defer fake.getCFAppsDomainMutex.RUnlock()
	return len(fake.getCFAppsDomainArgsForCall)
}

func (fake *FakeEnvironment) GetCFAppsDomainReturns(result1 string) {
	fake.GetCFAppsDomainStub = nil
	fake.getCFAppsDomainReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetExistingUser() string {
	fake.getExistingUserMutex.Lock()
	fake.getExistingUserArgsForCall = append(fake.getExistingUserArgsForCall, struct{}{})
	fake.recordInvocation("GetExistingUser", []interface{}{})
	fake.getExistingUserMutex.Unlock()
	if fake.GetExistingUserStub != nil {
		return fake.GetExistingUserStub()
	} else {
		return fake.getExistingUserReturns.result1
	}
}

func (fake *FakeEnvironment) GetExistingUserCallCount() int {
	fake.getExistingUserMutex.RLock()
	defer fake.getExistingUserMutex.RUnlock()
	return len(fake.getExistingUserArgsForCall)
}

func (fake *FakeEnvironment) GetExistingUserReturns(result1 string) {
	fake.GetExistingUserStub = nil
	fake.getExistingUserReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) UseExistingUser() bool {
	fake.useExistingUserMutex.Lock()
	fake.useExistingUserArgsForCall = append(fake.useExistingUserArgsForCall, struct{}{})
	fake.recordInvocation("UseExistingUser", []interface{}{})
	fake.useExistingUserMutex.Unlock()
	if fake.UseExistingUserStub != nil {
		return fake.UseExistingUserStub()
	} else {
		return fake.useExistingUserReturns.result1
	}
}

func (fake *FakeEnvironment) UseExistingUserCallCount() int {
	fake.useExistingUserMutex.RLock()
	defer fake.useExistingUserMutex.RUnlock()
	return len(fake.useExistingUserArgsForCall)
}

func (fake *FakeEnvironment) UseExistingUserReturns(result1 bool) {
	fake.UseExistingUserStub = nil
	fake.useExistingUserReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeEnvironment) KeepUserAtSuiteEnd() bool {
	fake.keepUserAtSuiteEndMutex.Lock()
	fake.keepUserAtSuiteEndArgsForCall = append(fake.keepUserAtSuiteEndArgsForCall, struct{}{})
	fake.recordInvocation("KeepUserAtSuiteEnd", []interface{}{})
	fake.keepUserAtSuiteEndMutex.Unlock()
	if fake.KeepUserAtSuiteEndStub != nil {
		return fake.KeepUserAtSuiteEndStub()
	} else {
		return fake.keepUserAtSuiteEndReturns.result1
	}
}

func (fake *FakeEnvironment) KeepUserAtSuiteEndCallCount() int {
	fake.keepUserAtSuiteEndMutex.RLock()
	defer fake.keepUserAtSuiteEndMutex.RUnlock()
	return len(fake.keepUserAtSuiteEndArgsForCall)
}

func (fake *FakeEnvironment) KeepUserAtSuiteEndReturns(result1 bool) {
	fake.KeepUserAtSuiteEndStub = nil
	fake.keepUserAtSuiteEndReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeEnvironment) GetExistingUserPassword() string {
	fake.getExistingUserPasswordMutex.Lock()
	fake.getExistingUserPasswordArgsForCall = append(fake.getExistingUserPasswordArgsForCall, struct{}{})
	fake.recordInvocation("GetExistingUserPassword", []interface{}{})
	fake.getExistingUserPasswordMutex.Unlock()
	if fake.GetExistingUserPasswordStub != nil {
		return fake.GetExistingUserPasswordStub()
	} else {
		return fake.getExistingUserPasswordReturns.result1
	}
}

func (fake *FakeEnvironment) GetExistingUserPasswordCallCount() int {
	fake.getExistingUserPasswordMutex.RLock()
	defer fake.getExistingUserPasswordMutex.RUnlock()
	return len(fake.getExistingUserPasswordArgsForCall)
}

func (fake *FakeEnvironment) GetExistingUserPasswordReturns(result1 string) {
	fake.GetExistingUserPasswordStub = nil
	fake.getExistingUserPasswordReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetStaticBuildpackName() string {
	fake.getStaticBuildpackNameMutex.Lock()
	fake.getStaticBuildpackNameArgsForCall = append(fake.getStaticBuildpackNameArgsForCall, struct{}{})
	fake.recordInvocation("GetStaticBuildpackName", []interface{}{})
	fake.getStaticBuildpackNameMutex.Unlock()
	if fake.GetStaticBuildpackNameStub != nil {
		return fake.GetStaticBuildpackNameStub()
	} else {
		return fake.getStaticBuildpackNameReturns.result1
	}
}

func (fake *FakeEnvironment) GetStaticBuildpackNameCallCount() int {
	fake.getStaticBuildpackNameMutex.RLock()
	defer fake.getStaticBuildpackNameMutex.RUnlock()
	return len(fake.getStaticBuildpackNameArgsForCall)
}

func (fake *FakeEnvironment) GetStaticBuildpackNameReturns(result1 string) {
	fake.GetStaticBuildpackNameStub = nil
	fake.getStaticBuildpackNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetJavaBuildpackName() string {
	fake.getJavaBuildpackNameMutex.Lock()
	fake.getJavaBuildpackNameArgsForCall = append(fake.getJavaBuildpackNameArgsForCall, struct{}{})
	fake.recordInvocation("GetJavaBuildpackName", []interface{}{})
	fake.getJavaBuildpackNameMutex.Unlock()
	if fake.GetJavaBuildpackNameStub != nil {
		return fake.GetJavaBuildpackNameStub()
	} else {
		return fake.getJavaBuildpackNameReturns.result1
	}
}

func (fake *FakeEnvironment) GetJavaBuildpackNameCallCount() int {
	fake.getJavaBuildpackNameMutex.RLock()
	defer fake.getJavaBuildpackNameMutex.RUnlock()
	return len(fake.getJavaBuildpackNameArgsForCall)
}

func (fake *FakeEnvironment) GetJavaBuildpackNameReturns(result1 string) {
	fake.GetJavaBuildpackNameStub = nil
	fake.getJavaBuildpackNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetRubyBuildpackName() string {
	fake.getRubyBuildpackNameMutex.Lock()
	fake.getRubyBuildpackNameArgsForCall = append(fake.getRubyBuildpackNameArgsForCall, struct{}{})
	fake.recordInvocation("GetRubyBuildpackName", []interface{}{})
	fake.getRubyBuildpackNameMutex.Unlock()
	if fake.GetRubyBuildpackNameStub != nil {
		return fake.GetRubyBuildpackNameStub()
	} else {
		return fake.getRubyBuildpackNameReturns.result1
	}
}

func (fake *FakeEnvironment) GetRubyBuildpackNameCallCount() int {
	fake.getRubyBuildpackNameMutex.RLock()
	defer fake.getRubyBuildpackNameMutex.RUnlock()
	return len(fake.getRubyBuildpackNameArgsForCall)
}

func (fake *FakeEnvironment) GetRubyBuildpackNameReturns(result1 string) {
	fake.GetRubyBuildpackNameStub = nil
	fake.getRubyBuildpackNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetNodeJSBuildpackName() string {
	fake.getNodeJSBuildpackNameMutex.Lock()
	fake.getNodeJSBuildpackNameArgsForCall = append(fake.getNodeJSBuildpackNameArgsForCall, struct{}{})
	fake.recordInvocation("GetNodeJSBuildpackName", []interface{}{})
	fake.getNodeJSBuildpackNameMutex.Unlock()
	if fake.GetNodeJSBuildpackNameStub != nil {
		return fake.GetNodeJSBuildpackNameStub()
	} else {
		return fake.getNodeJSBuildpackNameReturns.result1
	}
}

func (fake *FakeEnvironment) GetNodeJSBuildpackNameCallCount() int {
	fake.getNodeJSBuildpackNameMutex.RLock()
	defer fake.getNodeJSBuildpackNameMutex.RUnlock()
	return len(fake.getNodeJSBuildpackNameArgsForCall)
}

func (fake *FakeEnvironment) GetNodeJSBuildpackNameReturns(result1 string) {
	fake.GetNodeJSBuildpackNameStub = nil
	fake.getNodeJSBuildpackNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetGoBuildpackName() string {
	fake.getGoBuildpackNameMutex.Lock()
	fake.getGoBuildpackNameArgsForCall = append(fake.getGoBuildpackNameArgsForCall, struct{}{})
	fake.recordInvocation("GetGoBuildpackName", []interface{}{})
	fake.getGoBuildpackNameMutex.Unlock()
	if fake.GetGoBuildpackNameStub != nil {
		return fake.GetGoBuildpackNameStub()
	} else {
		return fake.getGoBuildpackNameReturns.result1
	}
}

func (fake *FakeEnvironment) GetGoBuildpackNameCallCount() int {
	fake.getGoBuildpackNameMutex.RLock()
	defer fake.getGoBuildpackNameMutex.RUnlock()
	return len(fake.getGoBuildpackNameArgsForCall)
}

func (fake *FakeEnvironment) GetGoBuildpackNameReturns(result1 string) {
	fake.GetGoBuildpackNameStub = nil
	fake.getGoBuildpackNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetPythonBuildpackName() string {
	fake.getPythonBuildpackNameMutex.Lock()
	fake.getPythonBuildpackNameArgsForCall = append(fake.getPythonBuildpackNameArgsForCall, struct{}{})
	fake.recordInvocation("GetPythonBuildpackName", []interface{}{})
	fake.getPythonBuildpackNameMutex.Unlock()
	if fake.GetPythonBuildpackNameStub != nil {
		return fake.GetPythonBuildpackNameStub()
	} else {
		return fake.getPythonBuildpackNameReturns.result1
	}
}

func (fake *FakeEnvironment) GetPythonBuildpackNameCallCount() int {
	fake.getPythonBuildpackNameMutex.RLock()
	defer fake.getPythonBuildpackNameMutex.RUnlock()
	return len(fake.getPythonBuildpackNameArgsForCall)
}

func (fake *FakeEnvironment) GetPythonBuildpackNameReturns(result1 string) {
	fake.GetPythonBuildpackNameStub = nil
	fake.getPythonBuildpackNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetPHPBuildpackName() string {
	fake.getPHPBuildpackNameMutex.Lock()
	fake.getPHPBuildpackNameArgsForCall = append(fake.getPHPBuildpackNameArgsForCall, struct{}{})
	fake.recordInvocation("GetPHPBuildpackName", []interface{}{})
	fake.getPHPBuildpackNameMutex.Unlock()
	if fake.GetPHPBuildpackNameStub != nil {
		return fake.GetPHPBuildpackNameStub()
	} else {
		return fake.getPHPBuildpackNameReturns.result1
	}
}

func (fake *FakeEnvironment) GetPHPBuildpackNameCallCount() int {
	fake.getPHPBuildpackNameMutex.RLock()
	defer fake.getPHPBuildpackNameMutex.RUnlock()
	return len(fake.getPHPBuildpackNameArgsForCall)
}

func (fake *FakeEnvironment) GetPHPBuildpackNameReturns(result1 string) {
	fake.GetPHPBuildpackNameStub = nil
	fake.getPHPBuildpackNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetBinaryBuildpackName() string {
	fake.getBinaryBuildpackNameMutex.Lock()
	fake.getBinaryBuildpackNameArgsForCall = append(fake.getBinaryBuildpackNameArgsForCall, struct{}{})
	fake.recordInvocation("GetBinaryBuildpackName", []interface{}{})
	fake.getBinaryBuildpackNameMutex.Unlock()
	if fake.GetBinaryBuildpackNameStub != nil {
		return fake.GetBinaryBuildpackNameStub()
	} else {
		return fake.getBinaryBuildpackNameReturns.result1
	}
}

func (fake *FakeEnvironment) GetBinaryBuildpackNameCallCount() int {
	fake.getBinaryBuildpackNameMutex.RLock()
	defer fake.getBinaryBuildpackNameMutex.RUnlock()
	return len(fake.getBinaryBuildpackNameArgsForCall)
}

func (fake *FakeEnvironment) GetBinaryBuildpackNameReturns(result1 string) {
	fake.GetBinaryBuildpackNameStub = nil
	fake.getBinaryBuildpackNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetPersistentAppHost() string {
	fake.getPersistentAppHostMutex.Lock()
	fake.getPersistentAppHostArgsForCall = append(fake.getPersistentAppHostArgsForCall, struct{}{})
	fake.recordInvocation("GetPersistentAppHost", []interface{}{})
	fake.getPersistentAppHostMutex.Unlock()
	if fake.GetPersistentAppHostStub != nil {
		return fake.GetPersistentAppHostStub()
	} else {
		return fake.getPersistentAppHostReturns.result1
	}
}

func (fake *FakeEnvironment) GetPersistentAppHostCallCount() int {
	fake.getPersistentAppHostMutex.RLock()
	defer fake.getPersistentAppHostMutex.RUnlock()
	return len(fake.getPersistentAppHostArgsForCall)
}

func (fake *FakeEnvironment) GetPersistentAppHostReturns(result1 string) {
	fake.GetPersistentAppHostStub = nil
	fake.getPersistentAppHostReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetPersistentAppSpace() string {
	fake.getPersistentAppSpaceMutex.Lock()
	fake.getPersistentAppSpaceArgsForCall = append(fake.getPersistentAppSpaceArgsForCall, struct{}{})
	fake.recordInvocation("GetPersistentAppSpace", []interface{}{})
	fake.getPersistentAppSpaceMutex.Unlock()
	if fake.GetPersistentAppSpaceStub != nil {
		return fake.GetPersistentAppSpaceStub()
	} else {
		return fake.getPersistentAppSpaceReturns.result1
	}
}

func (fake *FakeEnvironment) GetPersistentAppSpaceCallCount() int {
	fake.getPersistentAppSpaceMutex.RLock()
	defer fake.getPersistentAppSpaceMutex.RUnlock()
	return len(fake.getPersistentAppSpaceArgsForCall)
}

func (fake *FakeEnvironment) GetPersistentAppSpaceReturns(result1 string) {
	fake.GetPersistentAppSpaceStub = nil
	fake.getPersistentAppSpaceReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetPersistentAppOrg() string {
	fake.getPersistentAppOrgMutex.Lock()
	fake.getPersistentAppOrgArgsForCall = append(fake.getPersistentAppOrgArgsForCall, struct{}{})
	fake.recordInvocation("GetPersistentAppOrg", []interface{}{})
	fake.getPersistentAppOrgMutex.Unlock()
	if fake.GetPersistentAppOrgStub != nil {
		return fake.GetPersistentAppOrgStub()
	} else {
		return fake.getPersistentAppOrgReturns.result1
	}
}

func (fake *FakeEnvironment) GetPersistentAppOrgCallCount() int {
	fake.getPersistentAppOrgMutex.RLock()
	defer fake.getPersistentAppOrgMutex.RUnlock()
	return len(fake.getPersistentAppOrgArgsForCall)
}

func (fake *FakeEnvironment) GetPersistentAppOrgReturns(result1 string) {
	fake.GetPersistentAppOrgStub = nil
	fake.getPersistentAppOrgReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetPersistentAppQuotaName() string {
	fake.getPersistentAppQuotaNameMutex.Lock()
	fake.getPersistentAppQuotaNameArgsForCall = append(fake.getPersistentAppQuotaNameArgsForCall, struct{}{})
	fake.recordInvocation("GetPersistentAppQuotaName", []interface{}{})
	fake.getPersistentAppQuotaNameMutex.Unlock()
	if fake.GetPersistentAppQuotaNameStub != nil {
		return fake.GetPersistentAppQuotaNameStub()
	} else {
		return fake.getPersistentAppQuotaNameReturns.result1
	}
}

func (fake *FakeEnvironment) GetPersistentAppQuotaNameCallCount() int {
	fake.getPersistentAppQuotaNameMutex.RLock()
	defer fake.getPersistentAppQuotaNameMutex.RUnlock()
	return len(fake.getPersistentAppQuotaNameArgsForCall)
}

func (fake *FakeEnvironment) GetPersistentAppQuotaNameReturns(result1 string) {
	fake.GetPersistentAppQuotaNameStub = nil
	fake.getPersistentAppQuotaNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetBackend() (string, error) {
	fake.getBackendMutex.Lock()
	fake.getBackendArgsForCall = append(fake.getBackendArgsForCall, struct{}{})
	fake.recordInvocation("GetBackend", []interface{}{})
	fake.getBackendMutex.Unlock()
	if fake.GetBackendStub != nil {
		return fake.GetBackendStub()
	} else {
		return fake.getBackendReturns.result1, fake.getBackendReturns.result2
	}
}

func (fake *FakeEnvironment) GetBackendCallCount() int {
	fake.getBackendMutex.RLock()
	defer fake.getBackendMutex.RUnlock()
	return len(fake.getBackendArgsForCall)
}

func (fake *FakeEnvironment) GetBackendReturns(result1 string, result2 error) {
	fake.GetBackendStub = nil
	fake.getBackendReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSkipSSLValidationMutex.RLock()
	defer fake.getSkipSSLValidationMutex.RUnlock()
	fake.getUseHTTPMutex.RLock()
	defer fake.getUseHTTPMutex.RUnlock()
	fake.getIncludePrivilegedContainerSupportMutex.RLock()
	defer fake.getIncludePrivilegedContainerSupportMutex.RUnlock()
	fake.getDefaultTimeoutInSecondsMutex.RLock()
	defer fake.getDefaultTimeoutInSecondsMutex.RUnlock()
	fake.getCFPushTimeoutInSecondsMutex.RLock()
	defer fake.getCFPushTimeoutInSecondsMutex.RUnlock()
	fake.getLongCurlTimeoutInSecondsMutex.RLock()
	defer fake.getLongCurlTimeoutInSecondsMutex.RUnlock()
	fake.getBrokerStartTimeoutInSecondsMutex.RLock()
	defer fake.getBrokerStartTimeoutInSecondsMutex.RUnlock()
	fake.getCFAPIMutex.RLock()
	defer fake.getCFAPIMutex.RUnlock()
	fake.getCFAdminUserMutex.RLock()
	defer fake.getCFAdminUserMutex.RUnlock()
	fake.getCFAdminPasswordMutex.RLock()
	defer fake.getCFAdminPasswordMutex.RUnlock()
	fake.getCFAppsDomainMutex.RLock()
	defer fake.getCFAppsDomainMutex.RUnlock()
	fake.getExistingUserMutex.RLock()
	defer fake.getExistingUserMutex.RUnlock()
	fake.useExistingUserMutex.RLock()
	defer fake.useExistingUserMutex.RUnlock()
	fake.keepUserAtSuiteEndMutex.RLock()
	defer fake.keepUserAtSuiteEndMutex.RUnlock()
	fake.getExistingUserPasswordMutex.RLock()
	defer fake.getExistingUserPasswordMutex.RUnlock()
	fake.getStaticBuildpackNameMutex.RLock()
	defer fake.getStaticBuildpackNameMutex.RUnlock()
	fake.getJavaBuildpackNameMutex.RLock()
	defer fake.getJavaBuildpackNameMutex.RUnlock()
	fake.getRubyBuildpackNameMutex.RLock()
	defer fake.getRubyBuildpackNameMutex.RUnlock()
	fake.getNodeJSBuildpackNameMutex.RLock()
	defer fake.getNodeJSBuildpackNameMutex.RUnlock()
	fake.getGoBuildpackNameMutex.RLock()
	defer fake.getGoBuildpackNameMutex.RUnlock()
	fake.getPythonBuildpackNameMutex.RLock()
	defer fake.getPythonBuildpackNameMutex.RUnlock()
	fake.getPHPBuildpackNameMutex.RLock()
	defer fake.getPHPBuildpackNameMutex.RUnlock()
	fake.getBinaryBuildpackNameMutex.RLock()
	defer fake.getBinaryBuildpackNameMutex.RUnlock()
	fake.getPersistentAppHostMutex.RLock()
	defer fake.getPersistentAppHostMutex.RUnlock()
	fake.getPersistentAppSpaceMutex.RLock()
	defer fake.getPersistentAppSpaceMutex.RUnlock()
	fake.getPersistentAppOrgMutex.RLock()
	defer fake.getPersistentAppOrgMutex.RUnlock()
	fake.getPersistentAppQuotaNameMutex.RLock()
	defer fake.getPersistentAppQuotaNameMutex.RUnlock()
	fake.getBackendMutex.RLock()
	defer fake.getBackendMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeEnvironment) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ configwriter.Environment = new(FakeEnvironment)
