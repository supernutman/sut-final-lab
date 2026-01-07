package entity 

import (
	"testing"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
)

func Test3(t *testing.T) {
    g := NewGomegaWithT(t)
	t.Run("fetches the correct count", func(t *testing.T){
		book := Book{
			Title : "Name",
			Price : 51 ,
			Code : "AK123456",
		}
	ok,err := govalidator.ValidateStruct(book) 
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))
	})
}