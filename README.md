![Alt Text](assets/Software-Engineering.png)

1.Install and configure Go 

2.Setup Git using these steps on your PC

$ ssh-keygen -t  ed25519 -C "cedric.kiama@gmail.com"
$ ssh-add -D
$ eval "$(ssh-agent -s)"
$ ssh-add
$ ssh -T git@github.com

3.Create the project directory

$ mkdir Golang && cd Golang
$ git clone git@github.com:cedric-kiama-wachira/Software-Engineering.git

fIXING mAIL SERVER, sOURCE BUILD DIDNT WORK I WILL DO A CLEANUP TODAY AND START FROM SCRATCH.
Decided to setup an Ubuntu 20.04 separately to install mail Server
Mail server is working as it should.
Now ensuring the mail server has enterprise grade security to it.
I think now the Mail and Web Server setup are 95% ready. Will fix it to attain 100% readiness in the comming weeks.
Reading about MCP IN ai.
Setting up an ETCD node for My K8s cluster on Hetzner cloud.
Setting up HAProxy and SSl certificate for it
Setting Up SSl via letsencrypt for ETCD nodes, hope it works.
COnfiguring SSL services further.
Letsencrypt for ETCD doesn't work, reverting to using Self Signed and see how to combine that with LetsEncrypt certificates.
Changing the configs for ssl again.
