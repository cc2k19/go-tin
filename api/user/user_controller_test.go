package user

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/cc2k19/go-tin/models"

	"github.com/cc2k19/go-tin/web"

	"net/http/httptest"

	"github.com/cc2k19/go-tin/storage/storagefakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUserController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Controller Suite")
}

var _ = Describe("User Controller tests", func() {
	var userController *controller
	var fakeRepository *storagefakes.FakeRepository
	var request *http.Request
	var responseRecorder *httptest.ResponseRecorder
	var e error

	BeforeEach(func() {
		fakeRepository = &storagefakes.FakeRepository{}
		userController = NewUsersController(fakeRepository, web.CredentialsExtractorFunc(web.BasicCredentialsExtractor))
		responseRecorder = httptest.NewRecorder()
		request, e = http.NewRequest(http.MethodPost, web.UsersURL+"/test", strings.NewReader("fake body"))
		request.SetBasicAuth("test", "test")
		Expect(e).ShouldNot(HaveOccurred())
	})

	Describe("add user", func() {
		Context("when db error occurs", func() {
			It("returns status BadRequest", func() {
				fakeRepository.AddUserReturns(fmt.Errorf("db error"))
				userController.add(responseRecorder, request)

				Expect(responseRecorder.Code).To(Equal(http.StatusBadRequest))
				Expect(fakeRepository.AddUserCallCount()).To(Equal(1))

				context, body := fakeRepository.AddUserArgsForCall(0)
				Expect(context).To(Equal(request.Context()))

				// request body is closed here
				expectedBody, err := ioutil.ReadAll(strings.NewReader("fake body"))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(body).To(Equal(expectedBody))
			})
		})
		Context("when no error have occurred", func() {
			It("returns StatusCreated status", func() {
				userController.add(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusCreated))

				Expect(fakeRepository.AddUserCallCount()).To(Equal(1))

				context, body := fakeRepository.AddUserArgsForCall(0)
				Expect(context).To(Equal(request.Context()))

				// request body is closed here
				expectedBody, err := ioutil.ReadAll(strings.NewReader("fake body"))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(body).To(Equal(expectedBody))
			})
		})
	})

	Describe("getByUsername", func() {
		Context("when db error: 'no such user' occurs", func() {
			It("returns StatusNotFound status", func() {
				fakeRepository.GetUserByUsernameReturns(nil, fmt.Errorf("db error: no such user"))
				userController.getByUsername(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusNotFound))

				Expect(fakeRepository.GetUserByUsernameCallCount()).To(Equal(1))

				context, bodyS := fakeRepository.GetUserByUsernameArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(bodyS).To(Equal("test"))
			})
		})
		Context("when no error have occurred", func() {
			It("returns StatusOK status", func() {
				userController.getByUsername(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusOK))

				Expect(fakeRepository.GetUserByUsernameCallCount()).To(Equal(1))

				context, bodyS := fakeRepository.GetUserByUsernameArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(bodyS).To(Equal("test"))
			})
		})
	})

	Describe("follow", func() {
		Context("when db error: 'could not perform following' occurs", func() {
			It("returns StatusBadRequest status", func() {
				fakeRepository.AddFollowRecordReturns(fmt.Errorf("db error could not perform following"))
				userController.follow(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusBadRequest))

				Expect(fakeRepository.AddFollowRecordCallCount()).To(Equal(1))

				context, follower, target := fakeRepository.AddFollowRecordArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(follower).To(Equal("test"))
				Expect(target).To(Equal("/v1/users/test"))
			})
		})
		Context("when no error have occurred", func() {
			It("returns StatusCreated status", func() {
				userController.follow(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusCreated))

				Expect(fakeRepository.AddFollowRecordCallCount()).To(Equal(1))

				context, follower, target := fakeRepository.AddFollowRecordArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(follower).To(Equal("test"))
				Expect(target).To(Equal("/v1/users/test"))
			})
		})
	})

	Describe("unfollow", func() {
		Context("when db error: 'could not perform unfollowing' occurs", func() {
			It("returns StatusBadRequest status", func() {
				fakeRepository.DeleteFollowRecordReturns(fmt.Errorf("db error could not perform unfollowing"))
				userController.unfollow(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusBadRequest))

				Expect(fakeRepository.DeleteFollowRecordCallCount()).To(Equal(1))

				context, follower, target := fakeRepository.DeleteFollowRecordArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(follower).To(Equal("test"))
				Expect(target).To(Equal("/v1/users/test"))
			})
		})
		Context("when no error have occurred", func() {
			It("returns StatusNoContent status", func() {
				userController.unfollow(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusNoContent))

				Expect(fakeRepository.DeleteFollowRecordCallCount()).To(Equal(1))

				context, follower, target := fakeRepository.DeleteFollowRecordArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(follower).To(Equal("test"))
				Expect(target).To(Equal("/v1/users/test"))
			})
		})
	})

	Describe("getFollowers", func() {
		Context("when db error: 'could not get followers' occurs", func() {
			It("returns StatusNotFound status", func() {
				fakeRepository.GetFollowersReturns(models.UserSlice{}, fmt.Errorf("db error could not get followers"))
				userController.getFollowers(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusNotFound))

				Expect(fakeRepository.GetFollowersCallCount()).To(Equal(1))

				context, user := fakeRepository.GetFollowersArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(user).To(Equal("test"))
			})
		})
		Context("when no error have occurred but no users found", func() {
			It("returns StatusNotFound status", func() {
				fakeRepository.GetFollowersReturns(nil, nil)
				userController.getFollowers(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusNotFound))

				Expect(fakeRepository.GetFollowersCallCount()).To(Equal(1))

				context, user := fakeRepository.GetFollowersArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(user).To(Equal("test"))
			})
		})
		Context("when no error have occurred", func() {
			It("returns StatusOK status", func() {
				fakeRepository.GetFollowersReturns(models.UserSlice{}, nil)
				userController.getFollowers(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusOK))

				Expect(fakeRepository.GetFollowersCallCount()).To(Equal(1))

				context, user := fakeRepository.GetFollowersArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(user).To(Equal("test"))
			})
		})
	})

	Describe("getFollowing", func() {
		Context("when db error: 'could not get following' occurs", func() {
			It("returns StatusNotFound status", func() {
				fakeRepository.GetFollowingReturns(models.UserSlice{}, fmt.Errorf("db error could not get following"))
				userController.getFollowing(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusNotFound))

				Expect(fakeRepository.GetFollowingCallCount()).To(Equal(1))

				context, user := fakeRepository.GetFollowingArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(user).To(Equal("test"))
			})
		})
		Context("when no error have occurred but no users found", func() {
			It("returns StatusNotFound status", func() {
				fakeRepository.GetFollowingReturns(nil, nil)
				userController.getFollowing(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusNotFound))

				Expect(fakeRepository.GetFollowingCallCount()).To(Equal(1))

				context, user := fakeRepository.GetFollowingArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(user).To(Equal("test"))
			})
		})
		Context("when no error have occurred", func() {
			It("returns StatusOK status", func() {
				fakeRepository.GetFollowingReturns(models.UserSlice{}, nil)
				userController.getFollowing(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusOK))

				Expect(fakeRepository.GetFollowingCallCount()).To(Equal(1))

				context, user := fakeRepository.GetFollowingArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(user).To(Equal("test"))
			})
		})
	})
})
