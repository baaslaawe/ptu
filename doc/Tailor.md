## Tailoring

You can create tailored versions of PTU using `./script/tailor` utility.<br/>
Tailoring is done by creating custom software builds with defaults you set.<br/>
Tailored binaries may be shipped to your users, customers etc.<br/>
<br/>
The good thing about tailored binaries, **they work out of the box**, without any<br/>
client-side configuration, they do not burden you, they do not waste your time.<br/>
<br/>
Run `./script/tailor -h` to see all available options.<br/>
<br/>
e.g. `./script/tailor -n evilcorp -s ssh.evilcorp.com -u tunnelier -p 5ecuReP@s5 -t evilserver.com:443`<br/>
will create a custom build `evilcorp` with SSH server set to ` ssh.evilcorp.com`, SSH user `tunnelier`,<br/>
SSH password `5ecuReP@s5` and target (redirect) host set to `evilserver.com:443`, Every user who will<br/>
launch this custom build will start tunneling to `evilserver.com:443` via SSH server `ssh.evilcorp.com`.<br/>
*As you can see, flags used to tailor PTU are the same flags that are also used for PTU runtime configuration.*<br/>
<br/>
**NB!** Windows users not skilled enough to work with "normal" SSH will benefit most from tailoring feature.<br/>
**NB!** Tailored builds will ignore user YAML config files, because they strictly follow "no attachment" policy. <br/>
<br/>

#### How does it work?
`./script/tailor` gets the job done in a few very easy steps:
1. Checks out a new local Git branch from `master` (may override with `BASE_BRANCH` environment variable).
2. Substitutes default config settings in `lib/util/config/default_config.go` with settings you provided.
3. Runs a CI to verify nothing is broken. Fails the process, if CI script (`./script/ci`) returns non-zero code.
4. Commits the changes (`default_config.go` and binaries created by successful CI run) into local branch.
5. If `PTU_POST_EXEC` environment variable is set, runs its value as a command (e.g. to upload binaries to S3).

#### Environment variables
Besides command line arguments you may also set tailoring defaults via environment variables:
```
PTU_SSH_SERVER
PTU_SSH_USERNAME
PTU_SSH_PASSWORD
PTU_TARGET_HOST
PTU_EXPOSED_BIND
PTU_EXPOSED_PORT
```
You may set them anywhere, but for your convenience tailor script will<br/>
try to read environment from file `~/.ptu/tailor_profile`, if it exists.<br/>
I advice you to keep all your tailoring-related environment variables there.<br/>
*Due to obvious names of the variables I will not provide any descriptions.* :wink:<br/>
<br/>

#### Post exec commands
If you want to launch some commands after tailoring process finish,<br/>
for instance upload newly built PTU binaries to S3, you need to set<br/>
environment variable `PTU_POST_EXEC` content to your desired command.<br/>
Actually contents of every environment variable with its name starting from<br/>
`PTU_POST_EXEC` will be executed, so you may launch multiple commands.<br/>
Placing the line `PTU_POST_EXEC="git push my_remote $(git rev-parse --abbrev-ref HEAD)"`<br/>
inside the `~/.ptu/tailor_profile` will trigger git push of your tailored branch to your Git remote<br/>
after a finish of the tailoring process.
