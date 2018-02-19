package a1
import "testing"
import "fmt"

func stackerTest(s Stacker) bool{
    good := true
    if !s.isEmpty(){
      good = false
      fmt.Println("isempty tests not passed")
      }
    s.push (1)
    sp,_ := s.peek()
    if sp != 1{
      good = false
      fmt.Println("push or peek tests not passed")
      fmt.Println(sp)
      fmt.Println(s.size())
    }
    if s.size()!= 1{
      good = false
      fmt.Println("size tests not passed")
    }
    s.push (3)
    sp,_ = s.peek()
    if sp != 3{
      good = false
      fmt.Println("push or peek tests not passed")
    }
    if s.size()!= 2{
      good = false
      fmt.Println("size tests not passed")
    }
    S := s.copy()
    sp,_ = S.pop()
    sp1, _ := s.pop()
    sp2, _ := S.peek()
    sp3, _ := s.peek()
    if  sp!= sp1 || sp2!=sp3{
      good = false
      fmt.Println("copy tests not passed")
    }
    if s.size() != 1 || sp3 !=1 {
      good = false
      fmt.Println("pop tests not passed")
    }
    return good
}

func popAllTest(s Stacker) bool{
	s.push(33)
	s.push(2)
	s.push(45)
	popAll(s)
	if s.isEmpty(){
	    fmt.Println("popAllTest passed")
	    return true
	}
	return false
}

func stackEqualstest(s Stacker) bool{
	s.push(3)
	s.push(14)
	CopyS := s.copy()
	if stackEquals(CopyS, s){
	    fmt.Println("stackEqualstest passed")
	    return true
	}
	return false
}

func popAll(s Stacker) {
    for !s.isEmpty() {
      s.pop()
	}
}

func stackEquals(s, t Stacker) bool {
    S := s.copy()
    T := t.copy()
    if S.size() != T.size(){
      return false
    }
    var sp, tp int
    for S.isEmpty() {
      sp, _ = S.pop()
      tp, _ = T.pop()
      if sp != tp{
      return false
	}
    }
return true
}

func stackSliceTest() {
    s := makeStackSlice()
    if stackerTest(s) {
	if (popAllTest(s) && stackEqualstest(s)){
	fmt.Println("all StackSlice tests passed")
		}
	}
}

func stackLinkedTest() {
    s := makeStackLinked()
    if stackerTest(s) {
	if (popAllTest(s) && stackEqualstest(s)){
	fmt.Println("all StackLinked tests passed")
		}
	}
}

func Testa1(t *testing.T) {
  stackSliceTest()
  stackLinkedTest()
}

