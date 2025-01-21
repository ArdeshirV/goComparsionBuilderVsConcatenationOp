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
