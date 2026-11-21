# String-ul se codifică în valori ușor diferite cu
# encoder-ele base64 standard și URL (`+` față de `-` la
# final), dar ambele se decodifică, cum ne dorim, în
# string-ul original.
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
