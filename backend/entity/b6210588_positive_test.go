package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestCustomerPass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Check Customer Passes", func(t *testing.T) {
		c := Customer{
			Name:       "Teerasil",
			Email:      "T@gmail.com",
			CustomerID: "L1234567",
		}
		ok, err := govalidator.ValidateStruct(c)
		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())
		g.Expect(err)
	})
}
