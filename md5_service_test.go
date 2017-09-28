package main_test

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	. "github.com/Fipul/md5-service"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func randomString(length int) string {
	result := ""
	for i := 0; i < length; i++ {
		result += string(48 + rand.Intn(43))
	}
	return result
}

var s *httptest.Server

func postMd5(input string) (int, string, error) {
	resp, err := http.Post(s.URL+"/md5", "", strings.NewReader(input))
	if err != nil {
		return 0, "", err
	}
	var respBody []byte
	if resp.ContentLength > 0 {
		respBody, err = ioutil.ReadAll(resp.Body)
	}
	return resp.StatusCode, string(respBody), err
}

var _ = Describe("Md5Service", func() {
	BeforeSuite(func() {
		rand.Seed(time.Now().UnixNano())
		gin.SetMode(gin.ReleaseMode)
		s = httptest.NewServer(NewService())
	})
	Context("when valid data", func() {
		It("valid result", func() {
			statusCode, result, err := postMd5(`{"id":111,"text":"test text"}`)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(statusCode).To(Equal(http.StatusOK))
			Expect(result).To(Equal("06df1217b6c8e9ad6e119b65107fbb251"))
		})
	})
	Context("when wrong ID", func() {
		It("should status code 400", func() {
			statusCode, _, err := postMd5(`{"id":aaa,"text":"test text"}`)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(statusCode).To(Equal(http.StatusBadRequest))
		})
	})
	Context("when invalid text", func() {
		It("status code 400 bad request", func() {
			statusCode, _, err := postMd5(`{"id":111,"text":1.001}`)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(statusCode).To(Equal(http.StatusBadRequest))
		})
	})
	Context("when text is too long", func() {
		It("should status code 400", func() {
			statusCode, _, err := postMd5(`{"id":111,"text":"` + randomString(101) + `"}`)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(statusCode).To(Equal(http.StatusBadRequest))
		})
	})
})
