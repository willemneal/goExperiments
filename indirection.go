package main

import "fmt"

func main() {
    
}

type Interface interface{
  GetX() uint
  GetXIndirect() uint
}

type Struct struct {
  x uint
}

func (s Struct) GetXIndirect() uint {
  return s.x
}

func (s *Struct) GetX() uint {
  return s.x
}
