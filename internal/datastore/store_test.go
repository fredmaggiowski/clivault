package datastore

// import (
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// )

// var _ = Describe("Store", func() {
// 	Context("Save", func() {
// 		It("returns error if recordId already exists", func() {
// 			store := NewEmptyStore()

// 			store.dataStore["someRecord"] = DataRecord("someValue")

// 			err := store.Save("someRecord", "newValue")
// 			Expect(err).NotTo(BeNil())
// 			Expect(err).To(MatchError(ErrAlreadyExists))
// 		})

// 		It("saves recordo in dataStore", func() {
// 			savedRecord := DataRecord("newValue")

// 			store := NewEmptyStore()

// 			err := store.Save("someRecord", savedRecord)
// 			Expect(err).To(BeNil())
// 			Expect(store.dataStore["someRecord"]).To(Equal(savedRecord))
// 		})
// 	})
// })
