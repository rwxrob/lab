# Learning Lab: Setup Icinga Monitoring

> ⚠️
> There are better modern alternatives than Icinga (such as Prometheus).
> But some of us have to understand and support Icinga anyway.

By the end of this lab you should have a working Icinga 2 monitoring
system setup with all its dependencies on a single server (virtual
machine) suitable for sending test alert messages.

> 💡
> Note that this lab is for learning about how to install and configure
> Icinga, not just to get it up so that you can test the REST API (for
> example). For such things consider using the Docker compose file
> instead (see Related).

> ⚠️
> Avoid the temptation to get Icinga working in Kubernetes. It's not
> supported and has no precedent (yet).

1. **Understand what Icinga is and how it came to be**

1. **Understand the system dependencies for an Icinga installation**

1. **Provision a virtual machine to host Icinga**

1. **Install Icinga dependencies**

1. **Install Icinga**

1. **Configure Icinga to receive alerts**

Related:

* <https://github.com/icinga/icinga-vagrant>
* Vagrant - Icinga Web 2  
  <https://icinga.com/docs/icinga-web-2/latest/doc/99-Vagrant/>
* Learning Icinga » Documentation, Live Demo, Trainings  
  <https://icinga.com/learn>
* Nagios vs Icinga Comparison & Differences for Network Monitoring in 2022  
  <https://www.ittsystems.com/nagios-vs-icinga-comparison/>
* Home to Use Icinga to Monitor Your Servers and Services  
  <https://www.digitalocean.com/community/tutorials/how-to-use-icinga-to-monitor-your-servers-and-services-on-ubuntu-14-04>
* How To Provision NGINX Using Vagrant - Vegibit  
  <https://vegibit.com/how-to-provision-nginx-using-vagrant/>
* Installation - Icinga 2  
  <https://icinga.com/docs/icinga-2/latest/doc/02-installation/>
* <https://github.com/lippserd/docker-compose-icinga>
* How to Install PostgreSQL on CentOS 7 -- Easy Step-by-Step Tutorial  
  <https://www.hostinger.com/tutorials/how-to-install-postgresql-on-centos-7/>
* Setting up the PostgreSQL database - Icinga Web 2 - Icinga Community  
  <https://community.icinga.com/t/setting-up-the-postgresql-database/1622>
* Advanced Topics - Icinga Web 2  
  <https://icinga.com/docs/icinga-web-2/latest/doc/20-Advanced-Topics/#automating-the-installation-of-icinga-web-2>
