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
**NB!**<br/>
Windows users, who are not skilled enough to work with "normal" SSH, will benefit most from the tailoring feature.
<br/>

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
Due to obvious names of the variables I will not provide any descriptions. :wink:<br/>
<br/>

#### Post exec commands
If you want to launch some commands after tailoring process finish,<br/>
for instance upload newly built PTU binaries to S3, you need to set<br/>
environment variable `PTU_POST_EXEC` content to your desired command.<br/>
Actually contents of every environment variable with its name starting from<br/>
`PTU_POST_EXEC` will be executed, so you may launch multiple commands.<br/>
