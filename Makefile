build: clean
	@go build -ldflags "-w" -o workflow/alfred-lab .
	@alfred pack
	@mv alfred-lab* workflow/

clean:
	@rm -f workflow/alfred-lab*
