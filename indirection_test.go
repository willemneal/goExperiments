package main

import "testing"

func TestStruct_GetXIndirect(t *testing.T) {
}

func TestStruct_GetX(t *testing.T) {
}

func BenchmarkGetXInderectStruct(b *testing.B) {
        var s = &Struct{}
        s.x = 10
        for n := 0; n < b.N; n++ {
                s.GetXIndirect()
        }
}

func BenchmarkGetXerectStruct(b *testing.B) {
        var s = Struct{}
        s.x = 10
        for n := 0; n < b.N; n++ {
                s.GetX()
        }
}


func BenchmarkGetXInderectInterface(b *testing.B) {
        var s Interface = &Struct{10}

        for n := 0; n < b.N; n++ {
                s.GetXIndirect()
        }
}

func BenchmarkGetXerectInterface(b *testing.B) {
        var s Interface = &Struct{10}
        for n := 0; n < b.N; n++ {
                s.GetX()
        }
}
