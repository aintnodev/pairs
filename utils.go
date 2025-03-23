package main

import (
	"hash/fnv"
)

func genPercent(alname, blname string) float32 {
	h := fnv.New32a()
	h.Write([]byte(alname + blname))
	hashValue := h.Sum32()
	return float32(hashValue%100) + float32(float32(hashValue%100)/100.0)
}

func avgPercent(alname, blname string) float32 {
	return (genPercent(alname, blname) + genPercent(alname, blname)) / 2
}
