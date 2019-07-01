// Code generated by counterfeiter. DO NOT EDIT.
package storagefakes

import (
	"context"
	"sync"

	"github.com/cc2k19/go-tin/models"
	"github.com/cc2k19/go-tin/storage"
)

type FakeRepository struct {
	AddFollowRecordStub        func(context.Context, string, string) error
	addFollowRecordMutex       sync.RWMutex
	addFollowRecordArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	addFollowRecordReturns struct {
		result1 error
	}
	addFollowRecordReturnsOnCall map[int]struct {
		result1 error
	}
	AddPostStub        func(context.Context, string, []byte) error
	addPostMutex       sync.RWMutex
	addPostArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []byte
	}
	addPostReturns struct {
		result1 error
	}
	addPostReturnsOnCall map[int]struct {
		result1 error
	}
	AddUserStub        func(context.Context, []byte) error
	addUserMutex       sync.RWMutex
	addUserArgsForCall []struct {
		arg1 context.Context
		arg2 []byte
	}
	addUserReturns struct {
		result1 error
	}
	addUserReturnsOnCall map[int]struct {
		result1 error
	}
	AssertCredentialsStub        func(context.Context, []byte, []byte) error
	assertCredentialsMutex       sync.RWMutex
	assertCredentialsArgsForCall []struct {
		arg1 context.Context
		arg2 []byte
		arg3 []byte
	}
	assertCredentialsReturns struct {
		result1 error
	}
	assertCredentialsReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteFollowRecordStub        func(context.Context, string, string) error
	deleteFollowRecordMutex       sync.RWMutex
	deleteFollowRecordArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	deleteFollowRecordReturns struct {
		result1 error
	}
	deleteFollowRecordReturnsOnCall map[int]struct {
		result1 error
	}
	GetFollowersStub        func(context.Context, string) (models.UserSlice, error)
	getFollowersMutex       sync.RWMutex
	getFollowersArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getFollowersReturns struct {
		result1 models.UserSlice
		result2 error
	}
	getFollowersReturnsOnCall map[int]struct {
		result1 models.UserSlice
		result2 error
	}
	GetFollowingStub        func(context.Context, string) (models.UserSlice, error)
	getFollowingMutex       sync.RWMutex
	getFollowingArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getFollowingReturns struct {
		result1 models.UserSlice
		result2 error
	}
	getFollowingReturnsOnCall map[int]struct {
		result1 models.UserSlice
		result2 error
	}
	GetTargetsPostsStub        func(context.Context, string) (models.PostSlice, error)
	getTargetsPostsMutex       sync.RWMutex
	getTargetsPostsArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getTargetsPostsReturns struct {
		result1 models.PostSlice
		result2 error
	}
	getTargetsPostsReturnsOnCall map[int]struct {
		result1 models.PostSlice
		result2 error
	}
	GetUserByUsernameStub        func(context.Context, string) (*models.User, error)
	getUserByUsernameMutex       sync.RWMutex
	getUserByUsernameArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getUserByUsernameReturns struct {
		result1 *models.User
		result2 error
	}
	getUserByUsernameReturnsOnCall map[int]struct {
		result1 *models.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepository) AddFollowRecord(arg1 context.Context, arg2 string, arg3 string) error {
	fake.addFollowRecordMutex.Lock()
	ret, specificReturn := fake.addFollowRecordReturnsOnCall[len(fake.addFollowRecordArgsForCall)]
	fake.addFollowRecordArgsForCall = append(fake.addFollowRecordArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("AddFollowRecord", []interface{}{arg1, arg2, arg3})
	fake.addFollowRecordMutex.Unlock()
	if fake.AddFollowRecordStub != nil {
		return fake.AddFollowRecordStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addFollowRecordReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) AddFollowRecordCallCount() int {
	fake.addFollowRecordMutex.RLock()
	defer fake.addFollowRecordMutex.RUnlock()
	return len(fake.addFollowRecordArgsForCall)
}

func (fake *FakeRepository) AddFollowRecordCalls(stub func(context.Context, string, string) error) {
	fake.addFollowRecordMutex.Lock()
	defer fake.addFollowRecordMutex.Unlock()
	fake.AddFollowRecordStub = stub
}

func (fake *FakeRepository) AddFollowRecordArgsForCall(i int) (context.Context, string, string) {
	fake.addFollowRecordMutex.RLock()
	defer fake.addFollowRecordMutex.RUnlock()
	argsForCall := fake.addFollowRecordArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRepository) AddFollowRecordReturns(result1 error) {
	fake.addFollowRecordMutex.Lock()
	defer fake.addFollowRecordMutex.Unlock()
	fake.AddFollowRecordStub = nil
	fake.addFollowRecordReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) AddFollowRecordReturnsOnCall(i int, result1 error) {
	fake.addFollowRecordMutex.Lock()
	defer fake.addFollowRecordMutex.Unlock()
	fake.AddFollowRecordStub = nil
	if fake.addFollowRecordReturnsOnCall == nil {
		fake.addFollowRecordReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addFollowRecordReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) AddPost(arg1 context.Context, arg2 string, arg3 []byte) error {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.addPostMutex.Lock()
	ret, specificReturn := fake.addPostReturnsOnCall[len(fake.addPostArgsForCall)]
	fake.addPostArgsForCall = append(fake.addPostArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []byte
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("AddPost", []interface{}{arg1, arg2, arg3Copy})
	fake.addPostMutex.Unlock()
	if fake.AddPostStub != nil {
		return fake.AddPostStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addPostReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) AddPostCallCount() int {
	fake.addPostMutex.RLock()
	defer fake.addPostMutex.RUnlock()
	return len(fake.addPostArgsForCall)
}

func (fake *FakeRepository) AddPostCalls(stub func(context.Context, string, []byte) error) {
	fake.addPostMutex.Lock()
	defer fake.addPostMutex.Unlock()
	fake.AddPostStub = stub
}

func (fake *FakeRepository) AddPostArgsForCall(i int) (context.Context, string, []byte) {
	fake.addPostMutex.RLock()
	defer fake.addPostMutex.RUnlock()
	argsForCall := fake.addPostArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRepository) AddPostReturns(result1 error) {
	fake.addPostMutex.Lock()
	defer fake.addPostMutex.Unlock()
	fake.AddPostStub = nil
	fake.addPostReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) AddPostReturnsOnCall(i int, result1 error) {
	fake.addPostMutex.Lock()
	defer fake.addPostMutex.Unlock()
	fake.AddPostStub = nil
	if fake.addPostReturnsOnCall == nil {
		fake.addPostReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addPostReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) AddUser(arg1 context.Context, arg2 []byte) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.addUserMutex.Lock()
	ret, specificReturn := fake.addUserReturnsOnCall[len(fake.addUserArgsForCall)]
	fake.addUserArgsForCall = append(fake.addUserArgsForCall, struct {
		arg1 context.Context
		arg2 []byte
	}{arg1, arg2Copy})
	fake.recordInvocation("AddUser", []interface{}{arg1, arg2Copy})
	fake.addUserMutex.Unlock()
	if fake.AddUserStub != nil {
		return fake.AddUserStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addUserReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) AddUserCallCount() int {
	fake.addUserMutex.RLock()
	defer fake.addUserMutex.RUnlock()
	return len(fake.addUserArgsForCall)
}

func (fake *FakeRepository) AddUserCalls(stub func(context.Context, []byte) error) {
	fake.addUserMutex.Lock()
	defer fake.addUserMutex.Unlock()
	fake.AddUserStub = stub
}

func (fake *FakeRepository) AddUserArgsForCall(i int) (context.Context, []byte) {
	fake.addUserMutex.RLock()
	defer fake.addUserMutex.RUnlock()
	argsForCall := fake.addUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) AddUserReturns(result1 error) {
	fake.addUserMutex.Lock()
	defer fake.addUserMutex.Unlock()
	fake.AddUserStub = nil
	fake.addUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) AddUserReturnsOnCall(i int, result1 error) {
	fake.addUserMutex.Lock()
	defer fake.addUserMutex.Unlock()
	fake.AddUserStub = nil
	if fake.addUserReturnsOnCall == nil {
		fake.addUserReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addUserReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) AssertCredentials(arg1 context.Context, arg2 []byte, arg3 []byte) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.assertCredentialsMutex.Lock()
	ret, specificReturn := fake.assertCredentialsReturnsOnCall[len(fake.assertCredentialsArgsForCall)]
	fake.assertCredentialsArgsForCall = append(fake.assertCredentialsArgsForCall, struct {
		arg1 context.Context
		arg2 []byte
		arg3 []byte
	}{arg1, arg2Copy, arg3Copy})
	fake.recordInvocation("AssertCredentials", []interface{}{arg1, arg2Copy, arg3Copy})
	fake.assertCredentialsMutex.Unlock()
	if fake.AssertCredentialsStub != nil {
		return fake.AssertCredentialsStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.assertCredentialsReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) AssertCredentialsCallCount() int {
	fake.assertCredentialsMutex.RLock()
	defer fake.assertCredentialsMutex.RUnlock()
	return len(fake.assertCredentialsArgsForCall)
}

func (fake *FakeRepository) AssertCredentialsCalls(stub func(context.Context, []byte, []byte) error) {
	fake.assertCredentialsMutex.Lock()
	defer fake.assertCredentialsMutex.Unlock()
	fake.AssertCredentialsStub = stub
}

func (fake *FakeRepository) AssertCredentialsArgsForCall(i int) (context.Context, []byte, []byte) {
	fake.assertCredentialsMutex.RLock()
	defer fake.assertCredentialsMutex.RUnlock()
	argsForCall := fake.assertCredentialsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRepository) AssertCredentialsReturns(result1 error) {
	fake.assertCredentialsMutex.Lock()
	defer fake.assertCredentialsMutex.Unlock()
	fake.AssertCredentialsStub = nil
	fake.assertCredentialsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) AssertCredentialsReturnsOnCall(i int, result1 error) {
	fake.assertCredentialsMutex.Lock()
	defer fake.assertCredentialsMutex.Unlock()
	fake.AssertCredentialsStub = nil
	if fake.assertCredentialsReturnsOnCall == nil {
		fake.assertCredentialsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.assertCredentialsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) DeleteFollowRecord(arg1 context.Context, arg2 string, arg3 string) error {
	fake.deleteFollowRecordMutex.Lock()
	ret, specificReturn := fake.deleteFollowRecordReturnsOnCall[len(fake.deleteFollowRecordArgsForCall)]
	fake.deleteFollowRecordArgsForCall = append(fake.deleteFollowRecordArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("DeleteFollowRecord", []interface{}{arg1, arg2, arg3})
	fake.deleteFollowRecordMutex.Unlock()
	if fake.DeleteFollowRecordStub != nil {
		return fake.DeleteFollowRecordStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteFollowRecordReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) DeleteFollowRecordCallCount() int {
	fake.deleteFollowRecordMutex.RLock()
	defer fake.deleteFollowRecordMutex.RUnlock()
	return len(fake.deleteFollowRecordArgsForCall)
}

func (fake *FakeRepository) DeleteFollowRecordCalls(stub func(context.Context, string, string) error) {
	fake.deleteFollowRecordMutex.Lock()
	defer fake.deleteFollowRecordMutex.Unlock()
	fake.DeleteFollowRecordStub = stub
}

func (fake *FakeRepository) DeleteFollowRecordArgsForCall(i int) (context.Context, string, string) {
	fake.deleteFollowRecordMutex.RLock()
	defer fake.deleteFollowRecordMutex.RUnlock()
	argsForCall := fake.deleteFollowRecordArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRepository) DeleteFollowRecordReturns(result1 error) {
	fake.deleteFollowRecordMutex.Lock()
	defer fake.deleteFollowRecordMutex.Unlock()
	fake.DeleteFollowRecordStub = nil
	fake.deleteFollowRecordReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) DeleteFollowRecordReturnsOnCall(i int, result1 error) {
	fake.deleteFollowRecordMutex.Lock()
	defer fake.deleteFollowRecordMutex.Unlock()
	fake.DeleteFollowRecordStub = nil
	if fake.deleteFollowRecordReturnsOnCall == nil {
		fake.deleteFollowRecordReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteFollowRecordReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) GetFollowers(arg1 context.Context, arg2 string) (models.UserSlice, error) {
	fake.getFollowersMutex.Lock()
	ret, specificReturn := fake.getFollowersReturnsOnCall[len(fake.getFollowersArgsForCall)]
	fake.getFollowersArgsForCall = append(fake.getFollowersArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetFollowers", []interface{}{arg1, arg2})
	fake.getFollowersMutex.Unlock()
	if fake.GetFollowersStub != nil {
		return fake.GetFollowersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getFollowersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) GetFollowersCallCount() int {
	fake.getFollowersMutex.RLock()
	defer fake.getFollowersMutex.RUnlock()
	return len(fake.getFollowersArgsForCall)
}

func (fake *FakeRepository) GetFollowersCalls(stub func(context.Context, string) (models.UserSlice, error)) {
	fake.getFollowersMutex.Lock()
	defer fake.getFollowersMutex.Unlock()
	fake.GetFollowersStub = stub
}

func (fake *FakeRepository) GetFollowersArgsForCall(i int) (context.Context, string) {
	fake.getFollowersMutex.RLock()
	defer fake.getFollowersMutex.RUnlock()
	argsForCall := fake.getFollowersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) GetFollowersReturns(result1 models.UserSlice, result2 error) {
	fake.getFollowersMutex.Lock()
	defer fake.getFollowersMutex.Unlock()
	fake.GetFollowersStub = nil
	fake.getFollowersReturns = struct {
		result1 models.UserSlice
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetFollowersReturnsOnCall(i int, result1 models.UserSlice, result2 error) {
	fake.getFollowersMutex.Lock()
	defer fake.getFollowersMutex.Unlock()
	fake.GetFollowersStub = nil
	if fake.getFollowersReturnsOnCall == nil {
		fake.getFollowersReturnsOnCall = make(map[int]struct {
			result1 models.UserSlice
			result2 error
		})
	}
	fake.getFollowersReturnsOnCall[i] = struct {
		result1 models.UserSlice
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetFollowing(arg1 context.Context, arg2 string) (models.UserSlice, error) {
	fake.getFollowingMutex.Lock()
	ret, specificReturn := fake.getFollowingReturnsOnCall[len(fake.getFollowingArgsForCall)]
	fake.getFollowingArgsForCall = append(fake.getFollowingArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetFollowing", []interface{}{arg1, arg2})
	fake.getFollowingMutex.Unlock()
	if fake.GetFollowingStub != nil {
		return fake.GetFollowingStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getFollowingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) GetFollowingCallCount() int {
	fake.getFollowingMutex.RLock()
	defer fake.getFollowingMutex.RUnlock()
	return len(fake.getFollowingArgsForCall)
}

func (fake *FakeRepository) GetFollowingCalls(stub func(context.Context, string) (models.UserSlice, error)) {
	fake.getFollowingMutex.Lock()
	defer fake.getFollowingMutex.Unlock()
	fake.GetFollowingStub = stub
}

func (fake *FakeRepository) GetFollowingArgsForCall(i int) (context.Context, string) {
	fake.getFollowingMutex.RLock()
	defer fake.getFollowingMutex.RUnlock()
	argsForCall := fake.getFollowingArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) GetFollowingReturns(result1 models.UserSlice, result2 error) {
	fake.getFollowingMutex.Lock()
	defer fake.getFollowingMutex.Unlock()
	fake.GetFollowingStub = nil
	fake.getFollowingReturns = struct {
		result1 models.UserSlice
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetFollowingReturnsOnCall(i int, result1 models.UserSlice, result2 error) {
	fake.getFollowingMutex.Lock()
	defer fake.getFollowingMutex.Unlock()
	fake.GetFollowingStub = nil
	if fake.getFollowingReturnsOnCall == nil {
		fake.getFollowingReturnsOnCall = make(map[int]struct {
			result1 models.UserSlice
			result2 error
		})
	}
	fake.getFollowingReturnsOnCall[i] = struct {
		result1 models.UserSlice
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetTargetsPosts(arg1 context.Context, arg2 string) (models.PostSlice, error) {
	fake.getTargetsPostsMutex.Lock()
	ret, specificReturn := fake.getTargetsPostsReturnsOnCall[len(fake.getTargetsPostsArgsForCall)]
	fake.getTargetsPostsArgsForCall = append(fake.getTargetsPostsArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetTargetsPosts", []interface{}{arg1, arg2})
	fake.getTargetsPostsMutex.Unlock()
	if fake.GetTargetsPostsStub != nil {
		return fake.GetTargetsPostsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getTargetsPostsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) GetTargetsPostsCallCount() int {
	fake.getTargetsPostsMutex.RLock()
	defer fake.getTargetsPostsMutex.RUnlock()
	return len(fake.getTargetsPostsArgsForCall)
}

func (fake *FakeRepository) GetTargetsPostsCalls(stub func(context.Context, string) (models.PostSlice, error)) {
	fake.getTargetsPostsMutex.Lock()
	defer fake.getTargetsPostsMutex.Unlock()
	fake.GetTargetsPostsStub = stub
}

func (fake *FakeRepository) GetTargetsPostsArgsForCall(i int) (context.Context, string) {
	fake.getTargetsPostsMutex.RLock()
	defer fake.getTargetsPostsMutex.RUnlock()
	argsForCall := fake.getTargetsPostsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) GetTargetsPostsReturns(result1 models.PostSlice, result2 error) {
	fake.getTargetsPostsMutex.Lock()
	defer fake.getTargetsPostsMutex.Unlock()
	fake.GetTargetsPostsStub = nil
	fake.getTargetsPostsReturns = struct {
		result1 models.PostSlice
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetTargetsPostsReturnsOnCall(i int, result1 models.PostSlice, result2 error) {
	fake.getTargetsPostsMutex.Lock()
	defer fake.getTargetsPostsMutex.Unlock()
	fake.GetTargetsPostsStub = nil
	if fake.getTargetsPostsReturnsOnCall == nil {
		fake.getTargetsPostsReturnsOnCall = make(map[int]struct {
			result1 models.PostSlice
			result2 error
		})
	}
	fake.getTargetsPostsReturnsOnCall[i] = struct {
		result1 models.PostSlice
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetUserByUsername(arg1 context.Context, arg2 string) (*models.User, error) {
	fake.getUserByUsernameMutex.Lock()
	ret, specificReturn := fake.getUserByUsernameReturnsOnCall[len(fake.getUserByUsernameArgsForCall)]
	fake.getUserByUsernameArgsForCall = append(fake.getUserByUsernameArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetUserByUsername", []interface{}{arg1, arg2})
	fake.getUserByUsernameMutex.Unlock()
	if fake.GetUserByUsernameStub != nil {
		return fake.GetUserByUsernameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getUserByUsernameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) GetUserByUsernameCallCount() int {
	fake.getUserByUsernameMutex.RLock()
	defer fake.getUserByUsernameMutex.RUnlock()
	return len(fake.getUserByUsernameArgsForCall)
}

func (fake *FakeRepository) GetUserByUsernameCalls(stub func(context.Context, string) (*models.User, error)) {
	fake.getUserByUsernameMutex.Lock()
	defer fake.getUserByUsernameMutex.Unlock()
	fake.GetUserByUsernameStub = stub
}

func (fake *FakeRepository) GetUserByUsernameArgsForCall(i int) (context.Context, string) {
	fake.getUserByUsernameMutex.RLock()
	defer fake.getUserByUsernameMutex.RUnlock()
	argsForCall := fake.getUserByUsernameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) GetUserByUsernameReturns(result1 *models.User, result2 error) {
	fake.getUserByUsernameMutex.Lock()
	defer fake.getUserByUsernameMutex.Unlock()
	fake.GetUserByUsernameStub = nil
	fake.getUserByUsernameReturns = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetUserByUsernameReturnsOnCall(i int, result1 *models.User, result2 error) {
	fake.getUserByUsernameMutex.Lock()
	defer fake.getUserByUsernameMutex.Unlock()
	fake.GetUserByUsernameStub = nil
	if fake.getUserByUsernameReturnsOnCall == nil {
		fake.getUserByUsernameReturnsOnCall = make(map[int]struct {
			result1 *models.User
			result2 error
		})
	}
	fake.getUserByUsernameReturnsOnCall[i] = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addFollowRecordMutex.RLock()
	defer fake.addFollowRecordMutex.RUnlock()
	fake.addPostMutex.RLock()
	defer fake.addPostMutex.RUnlock()
	fake.addUserMutex.RLock()
	defer fake.addUserMutex.RUnlock()
	fake.assertCredentialsMutex.RLock()
	defer fake.assertCredentialsMutex.RUnlock()
	fake.deleteFollowRecordMutex.RLock()
	defer fake.deleteFollowRecordMutex.RUnlock()
	fake.getFollowersMutex.RLock()
	defer fake.getFollowersMutex.RUnlock()
	fake.getFollowingMutex.RLock()
	defer fake.getFollowingMutex.RUnlock()
	fake.getTargetsPostsMutex.RLock()
	defer fake.getTargetsPostsMutex.RUnlock()
	fake.getUserByUsernameMutex.RLock()
	defer fake.getUserByUsernameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRepository) recordInvocation(key string, args []interface{}) {
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

var _ storage.Repository = new(FakeRepository)
