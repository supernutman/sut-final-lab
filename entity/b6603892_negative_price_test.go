package entity 

import (
	"testing"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
)

func Test2(t *testing.T) {
    g := NewGomegaWithT(t)
	t.Run("fetches the correct count", func(t *testing.T){
		book := Book{
			Title : "Name",
			Price : 49 ,
			Code : "BK123456",
		}
	ok,err := govalidator.ValidateStruct(book) 
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Price must be between 50 and 5000"))
	})
}