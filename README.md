# expvar
Go expvar wrapper with goroutines and uptime

## How to use

Add the code to your project:

```go
import "github.com/night-codes/expvar"

expvar.Start(":3000")
```

Next step, use Go-apps console monitoring tool [expvarmon](https://github.com/divan/expvarmon) :

```bash
go get github.com/divan/expvarmon

expvarmon -ports="3000"
```

Read more about [expvars](http://golang.org/pkg/expvar/)

## License
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
Version 2, December 2004

Copyright (C) 2015 Oleksiy Chechel <alex.mirrr@gmail.com>

Everyone is permitted to copy and distribute verbatim or modified
copies of this license document, and changing it is allowed as long
as the name is changed.

DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

 0. You just DO WHAT THE FUCK YOU WANT TO.


