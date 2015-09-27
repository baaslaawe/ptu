## Portable Tunneling Utility
**ptu** is a simple ad-hoc SSH-based TCP port tunneling application. Written in **Go**. :smiling_imp: <br />
It can expose **any server** in your network, including your localhost, to the Internet.

##### HOWTO: [Installation](https://github.com/ivanilves/ptu/blob/master/doc/Install.md)
##### HOWTO: [Tailoring](https://github.com/ivanilves/ptu/blob/master/doc/Tailor.md)

#### What problems does it solve?
Well, with **ptu** you can ...
* ... provide remote professionals with temporary access to your infrastructure.
* ... demo your work to customers and your team without deploying it anywhere.

Apart from this, it's just a good tool to open temporary holes in your firewall<br />
without polluting firewall configuration with temporary rules **you may forget**<br />
to remove after you don't need them.

Also **ptu** require no installation and no superuser rights, so you can<br />
just grab and use it right away **without** involving your IT department.

#### Reinventing the wheel?
**ptu** concepts look similar to **[ngrok](https://ngrok.com/)**, though there are some major differences:
* **ptu** is "serverless", it could use virtually any SSH server as a "backend".
* it's not about just tunneling to locahost, you can tunnel to **any server** you have access to.
* **ptu** is Open Source. You can extend and modify it. And for sure, it's backdoor-free.

#### Usage
```
ptu -s <ssh_server>[:<ssh_port>] [OPTIONS]
```
###### Options
`-t <target_host>:<target_port>`<br />
A target host:port for connection forwarding. This is usually your company server or your localhost.<br />
By-default `<target_host>` is **localhost** and `<taget_port>` is **22**.

`-e <expose_port>`<br />
What port do we want to expose to the Internet? This should be Integer value inside 1..65535 range.<br />
**NB!** Using port numbers below 1024 would require root access to SSH server, which is generally bad.<br />
By-default `<expose_port>` takes random value in **10000..19999** range.

`-u <ssh_username>`<br />
The username to connect an SSH server. By-default it's your **system username**.

`-p <ssh_password>`<br />
Use provided password to login into SSH server.<br />
**NB!** Using SSH password is possible, but highly undesirable due to security reasons.<br />
By-default superior option, an SSH authentication agent, is used to authenticate SSH.

`-c <config_name>`<br />
Load settings from `~/.ptu/<config_name>.yaml` (see this [example](https://github.com/ivanilves/ptu/blob/master/data/fistro.yaml)).<br />
If you create file `~/.ptu/default.yaml`, settings stored there will override program built-in defaults.<br/>
Settings passed by command line arguments always take precedence over settings loaded from file.

## NB!
* Please see SSH server [GatewayPorts](http://www.snailbook.com/faq/gatewayports.auto.html) option.
* **ptu** runs on your machine and nowhere else. :trollface:

## Typical use case
![](https://raw.githubusercontent.com/ivanilves/ptu/master/doc/how_it_works.png)
