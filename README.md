# count-min-sketch-go

Implemenatation of count min sketch algorithm in go. I am using seed value in murmur hash, different seed value will help to create different hash function. we run all hash function for each entry and maintain the count in the 2d slice.

If map size is 1000
   When we save 90000  repeated items in 100000  items we get the count as 90000  which is approx correct
   when we save 8000   repeated items in 100000  items we get the count as 8090  
   When we save 800    repeated items in 10000   items we get the count as 810 