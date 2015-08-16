## Portable Tunneling Utility: tunnel everything, tunnel everywhere
**ptu** is a simple ad-hoc SSH-based TCP port tunneling application.  
It can expose **any server** in your network, including your localhost, to the Internet.

#### What problems does it solve?
Well, with **ptu** you can ...
* ... provide remote professionals with temporary access to your infrastructure.
* ... demo your work to customers and your team without deploying it anywhere.

Apart from this, it's just a good tool to open temporary holes in your firewall  
without polluting firewall configuration with temporary rules **you may forget**  
to remove after you don't need them.

Also **ptu** require no installation and no superuser rights, so you can just grab and use it right away without involving your IT department.  

#### Reinventing the wheel?
**ptu** concepts look similar to **[ngrok](https://ngrok.com/)**, though there are some major differences:
* **ptu** is "serverless", it could use virtually any SSH server as a "backend".
* it's not about just tunneling to locahost, you can tunnel to **any server** you have access to [from your laptop].
* **ptu** is Open Source. You can extend and modify it. And for sure, it's backdoor-free.

#### Usage
```
ptu -s <ssh_server>[:<ssh_port>] -f <remote_host>[:<remote_port>] [-e <expose_port>] [-u <ssh_username>] [-p <ssh_password>] [-n]
```
