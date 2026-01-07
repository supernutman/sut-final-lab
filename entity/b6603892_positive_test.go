package entity 

import (
	"testing"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
)

func TestFarmHasCow(t *testing.T) {
    g := NewGomegaWithT(t)
	t.Run("fetches the correct count", func(t *testing.T){
		book := Book{
			Title : "Name",
			Price : 51 ,
			Code : "BK123456",
		}
	ok,err := govalidator.ValidateStruct(book) 
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
	})
}