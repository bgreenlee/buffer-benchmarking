package main

import (
	"testing"
)

var num = 1000000

func BenchmarkLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linkedListBuffer(num)
	}
}

func BenchmarkChannelBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		channelBufferSelect(num)
	}
}

func BenchmarkChannelBufferIfLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		channelBufferIfLen(num)
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceBuffer(num)
	}
}

func BenchmarkSliceMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceBufferMake(num)
	}
}

func BenchmarkDeque(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dequeBuffer(num)
	}
}

func BenchmarkDeque2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deque2Buffer(num)
	}
}
