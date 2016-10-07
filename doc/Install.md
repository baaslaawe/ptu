## Installation

**NB!** Despite the fact **ptu** binaries do work on Windows, only Linux and MacOSX are supported for development.

#### The process is really, really simple:

  1. Install [Go Version Manager](https://github.com/moovweb/gvm) (Go is hellish without it, really)

  2. Install "provisional" Go 1.4 to build working version of Go with it during next step:
  ```
  gvm install go1.4
  gvm use go1.4
  ```

  3. Install **Go 1.7.1** and use it by default:
  ```
  gvm install go1.7.1
  gvm use go1.7.1 --default
  ```

  4. Clone PTU repo and cd there `git clone git@github.com:ivanilves/ptu.git && cd ptu`

  5. Get all required Golang packages: `./script/install`

  6. Run the continuous integration suite to check compliance: `./script/ci`<br/>
  **NB!** To run continuous integration suite successfully, you need to have **SSH agent** working<br/>
  and you also need to be able to login your localhost via SSH using your RSA/DSA public key. :scream:

  7. Try to run freshly generated binaries from `./bin` directory.

  8. PROFIT!!! :dancer:
