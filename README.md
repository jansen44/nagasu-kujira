![Whale Render - Ruins of Atlas](https://i.imgur.com/g9eYbWN.jpg)
___Art by [NesoKaiyoH](https://www.deviantart.com/nesokaiyoh)___
# NagasuKujira - 長須鯨

### About
Kujira鯨 aims to be an OSS solution for Kanban/Scrum stuff. It aims to be easy to run, develop, and build extensions. This mainly aims to help individuals that has a lot of different projects and can't just manage them (myself for instance), but this doesn't mean that it won't work for teams also 😉 

### About - 長須鯨
NagasuKujira長須鯨 is part of the Kujira project as the main backend service. 長須鯨 (Fin Whales) are the fastest whales of the world, and this service aims to be the fastest service of the Kujira stack, so it makes sense (I guess).

### 長須鯨 Roadmap
At first things may seem kinda simple, in general that's what I aim to do, but there are a lot of things I need/want to change in a near future.

You can check all roadmap and current development status on our [Github Project](https://github.com/jansen44/nagasu-kujira/projects/1)

### Running 長須鯨
Before running, you need to satisfy the selected driver and selected storage dependencies. 
Currently, there is only one driver (RestAPI Server) and one storage (MySQL). 
- RestAPI Server works as is, you can set the default listening port with some flags.
- For MySQL Storage you need one MySQL server running with a database called ```nagasu```. 
  - Now it's only able to run the server locally, but feel free to edit ```infra/repositories.go (line 13)``` to connect to a remote server. I intend to add a flag to set this in a near future, so don't worry about it.
- You also need to run database migrations before running... Or you can just add a ```-migrate``` flag and it will work as it was supposed to.

```bash
go run main.go -migrate -restPort=3030 -storage=mysql -driver=rest
```
The above command runs Nagasu with automatic storage migrations, using MySQL as storage and running a RestAPI Server (listening on port ```3030```).
