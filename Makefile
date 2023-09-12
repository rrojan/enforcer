run:
	./libtest/run.sh
ignore-go-mod:
	git update-index --assume-unchanged go.mod
track-go-mod:
	git update-index --no-assume-unchanged go.mod
