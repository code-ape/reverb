package parser_test

type TestObj struct {
  X int
}

func main() {

  test_obj_1 := NewTestObj()
  test_obj_2 := NewTestObj()

  factory := MakeTestFactory(MakeTestTemplate())  


  e1 := factory.NewExpectations()
  e1.AddValForKey("one", 1)
  factory.RunSuite(test_obj_1, e1)

  e2 := factory.NewExpectations()
  e2.AddValForKey("one", 1)
  factory.RunSuite(test_obj_1, e1)


  // t := factory.MakeSuite(test_obj_1)
  // t.Expectations(e1)
  // t.RunSuite()
  // t.Run("one")



}

func NewTestObj() *TestObj {
  return &TestObj{}
}


func MakeTestTemplate() {
  var a int

  template_obj := NewTemplate()

  e := template_obj.Expectations()
  e.RegisterKeys([]string{"one"})

  template_obj.BeforeEach(func() {
    a = 1
  })

  template_obj.It("one", "one = one", func() {
      Expect(1).Should(Equal(e.Val("one")))
    })
  })

  return &template_obj
}