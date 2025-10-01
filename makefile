benchmark:
	mv ./builder.go ./concat.go ../ > /dev/null
	go test -bench=. -v -count=1  # -benchmem
	mv ../builder.go ../concat.go ./ > /dev/null

all: builder concat


builder: builder.go
	go build -o builder builder.go


concat: concat.go
	go build -o concat concat.go


test: run
run: all
	sleep 1
	/bin/time ./builder
	sleep 1
	/bin/time ./concat


clean:
	rm -rf builder concat
