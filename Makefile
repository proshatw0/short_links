FILES = main.go
TARGET = bin/short_links.exe

.PHONY: clean

TARGET: 
	go build -o $(TARGET) $(FILES)

clean:
	rm -r $(TARGET)
