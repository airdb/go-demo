ver1:
	go test -bench=. -count=10 | tee ver1.txt

ver2:
	go test -bench=. -count=10 | tee ver2.txt

benchstat:
	benchstat ver1.txt ver2.txt
