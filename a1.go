package a1
import "fmt"
import "errors"

type Stacker interface {

    isEmpty() bool
    size() int
    push(x int)
    peek() (int, error)
    pop() (int, error)
    copy() Stacker
}


type StackSlice struct {
        items []int
}


func makeStackSlice() Stacker {
    s := &StackSlice{
        items: make([]int, 0, 1024),
    }
    return s
}

func (s StackSlice) isEmpty() bool{
  size := len(s.items)
  if size != 0{
    return false
	}
  return true

}

func (s StackSlice) size() int{
  size := len(s.items)
  return size
}

func (s *StackSlice) push(x int) {
  slice := append(s.items,x)
  s.items = slice
}

func (s StackSlice) peek () (int, error){
  n:= s.size()
  if (n == 0) {
    return 0, errors.New("the stack is empty.")
  } else {
    return s.items[n-1] , nil
  }
}

func (s *StackSlice) pop() (int, error){
  size := s.size()
  if size == 0 {
    return 0,errors.New("the stack is empty.")
  } else{
    pop_one := s.items[size-1]
    s.items = s.items[0:size-1]
    return pop_one,nil
  }
}


func (s StackSlice) copy() Stacker{
  new_stack := &StackSlice{
      items: s.items,}
  return new_stack
  }

type StackLinked struct {
  head *TheNode
}

type TheNode struct {
    item int
    next * TheNode
}


func makeStackLinked() Stacker {
    head := &TheNode{0, nil}
    return &StackLinked{head}
}

func (s StackLinked) isEmpty() bool{
    return s.head.next == nil
}

func (s StackLinked) size() int{
    size := 0
    curr := s.head
    for curr.next != nil{
	size ++
	curr = curr.next
    } 
    return size
 }


func (s *StackLinked) push(x int) {
    curr := s.head
    for curr.next != nil{
        curr = curr.next
    }
    curr.next = &TheNode{x, nil}
}



func (s StackLinked) peek () (int, error){
  if s.isEmpty(){
    return 0,errors.New("the stack is empty.")
}else{
    curr := s.head
    for curr.next != nil {
      curr = curr.next
    }
    return curr.item,nil
  }
}

func (s StackLinked) pop() (int, error){
  if s.isEmpty(){
    return 0,errors.New("the stack is empty.")
}else{
    curr := s.head
    this := curr
    for curr.next != nil {
      this = curr
      curr = curr.next
      }
      ret := curr.item
      this.next = nil
      curr = nil
    return ret,nil
  }
}




func (s StackLinked) copy() Stacker{
  new_stack := makeStackLinked()
  curr := s.head
  for curr.next != nil {
     curr = curr.next
     new_stack.push(curr.item)
   }
  return new_stack
}




func (this StackSlice) String() string {
	peekone, _ := this.peek()
	return fmt.Sprintf("the peek one is %d", peekone)
}

func (this StackLinked) String() string {
	peekone, _ := this.peek()
	return fmt.Sprintf("the peek one is %d", peekone)
}
