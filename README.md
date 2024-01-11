# count-min-sketch-go

Implemenatation of count min sketch algorithm in go. I am using seed value in murmur hash, different seed value will help to create different hash function. we run all hash function for each entry and maintain the count in the 2d slice.