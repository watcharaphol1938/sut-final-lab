package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestNameMustNotBeBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Check Name must not be blank", func(t *testing.T) {
		c := Customer{
			Name:       "",
			Email:      "T@gmail.com",
			CustomerID: "L1234567",
		}
		ok, err := govalidator.ValidateStruct(c)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Name should not be blank"))
	})
}
