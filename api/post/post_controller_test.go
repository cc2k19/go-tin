package post

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cc2k19/go-tin/storage/storagefakes"
	"github.com/cc2k19/go-tin/web"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPostController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Post Controller Suite")
}

var _ = Describe("Post Controller tests", func() {
	var postController *controller
	var fakeRepository *storagefakes.FakeRepository
	var request *http.Request
	var responseRecorder *httptest.ResponseRecorder
	var e error

	BeforeEach(func() {
		fakeRepository = &storagefakes.FakeRepository{}
		postController = NewPostsController(fakeRepository, web.CredentialsExtractorFunc(web.BasicCredentialsExtractor))
		responseRecorder = httptest.NewRecorder()
		request, e = http.NewRequest(http.MethodPost, web.PostsURL, strings.NewReader("fake body"))
		request.SetBasicAuth("test", "test")
		Expect(e).ShouldNot(HaveOccurred())
	})

	Describe("add", func() {
		Context("when db error occurs", func() {
			It("returns status BadRequest", func() {
				fakeRepository.AddPostReturns(fmt.Errorf("db error"))
				postController.add(responseRecorder, request)

				Expect(responseRecorder.Code).To(Equal(http.StatusBadRequest))
				Expect(fakeRepository.AddPostCallCount()).To(Equal(1))

				context, username, post := fakeRepository.AddPostArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(username).To(Equal("test"))

				// request body is closed here
				expectedBody, err := ioutil.ReadAll(strings.NewReader("fake body"))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(post).To(Equal(expectedBody))
			})
		})

		Context("when no error have occurred", func() {
			It("returns StatusCreated status", func() {
				postController.add(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusCreated))
				Expect(fakeRepository.AddPostCallCount()).To(Equal(1))

				context, username, post := fakeRepository.AddPostArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(username).To(Equal("test"))

				// request body is closed here
				expectedBody, err := ioutil.ReadAll(strings.NewReader("fake body"))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(post).To(Equal(expectedBody))
			})
		})
	})

	Describe("get", func() {
		Context("when db error: 'no posts found' occurs", func() {
			It("returns StatusNotFound status", func() {
				fakeRepository.GetTargetsPostsReturns(nil, fmt.Errorf("db error: no posts found"))
				postController.get(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusNotFound))

				Expect(fakeRepository.GetTargetsPostsCallCount()).To(Equal(1))

				context, bodyS := fakeRepository.GetTargetsPostsArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(bodyS).To(Equal("test"))
			})
		})
		Context("when no error have occurred", func() {
			It("returns StatusOK status", func() {
				postController.get(responseRecorder, request)
				Expect(responseRecorder.Code).To(Equal(http.StatusOK))

				Expect(fakeRepository.GetTargetsPostsCallCount()).To(Equal(1))

				context, bodyS := fakeRepository.GetTargetsPostsArgsForCall(0)
				Expect(context).To(Equal(request.Context()))
				Expect(bodyS).To(Equal("test"))
			})
		})
	})
})
