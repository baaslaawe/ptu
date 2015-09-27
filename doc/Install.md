## Installation

*This process is really, really simple and clear. Because it's not Ruby.* :trollface:

1. Install [Go Version Manager](https://github.com/moovweb/gvm) (Go is hellish without it, really)

2. Clone PTU repo and cd there `git clone git@github.com:ivanilves/ptu.git && cd ptu`

3. Install **Go 1.5.1** and use it by default:
```
gvm install go1.5.1
gvm use go1.5.1 --default
```

4. Get all required Golang packages:<br/>
`go get -u errors flag fmt golang.org/x/crypto/ssh golang.org/x/crypto/ssh/agent gopkg.in/yaml.v2 io io/ioutil log math/rand net net/http os os/user regexp strconv strings testing time`

5. Run the continuous integration suite to check compliance: `./script/ci`

6. Try to run freshly generated binaries from `./bin` directory.

7. PROFIT!!! :dancer:
