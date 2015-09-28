## Installation

*This process is really, really simple and clear. Because it's not Ruby.* :trollface:

  1. Install [Go Version Manager](https://github.com/moovweb/gvm) (Go is hellish without it, really)

  2. Clone PTU repo and cd there `git clone git@github.com:ivanilves/ptu.git && cd ptu`

  3. Install **Go 1.5.1** and use it by default:
  ```
  gvm install go1.5.1
  gvm use go1.5.1 --default
  ```

  4. Get all required Golang packages: `./script/install`

  5. Run the continuous integration suite to check compliance: `./script/ci`<br/>
  **NB!** To run continuous integration suite successfully, you need to have **SSH agent** working<br/>
  and you also need to be able to login your localhost via SSH using your RSA/DSA public key.

  6. Try to run freshly generated binaries from `./bin` directory.

  7. PROFIT!!! :dancer:
