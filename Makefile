build:
	@go build -o workflow/alfred-lab .
	@alfred pack
	@mv alfred-lab* workflow/
