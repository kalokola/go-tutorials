module example.com/hello

go 1.25

// for local modules and not just files, but modules :)

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
