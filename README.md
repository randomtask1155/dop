# Purpose

A basic doppler nozzle that can listen for events on the cf doppler endpoint 

# installation

make sure you have $GOPATH environment vairalbe set.  This will be the path used to install dop

```
export GOPATH=/my/path
export PATH=$GOPATH/bin/$PATH
```

This command will will insall dop under `$GOPATH/bin/dop`

```
go get github.com/randomtask1155/dop
```

# Usage 

```
Usage of dop:
  -token string
    	Provide an access token used to authenticate with doppler endpoint. Defaults to ~/.cf/config.json
  -url string
    	Web socket address example: wss://doppler.system.domain:443/apps/41abc841-cbc8-4cab-854d-640a7c8b6a5f/stream
```

**Note** you can copy the access token manually from your local `~/.cf/config.json` to the host you want to run dop on

```
~:> cat ~/.cf/config.json  | egrep AccessToken
  "AccessToken": "bearer eyJhbGciOiJSUzI1NiIsImtpZCI6Imtle....",
```

# Examples 

## Listing for events of a specific app 

```
$ dop -url wss://doppler.system.domain:443/apps/41abc841-cbc8-4cab-854d-640a7c8b6a5f/stream
starting output collector
starting read loop
Sending input
origin:"cloud_controller" eventType:LogMessage deployment:"cf" job:"cloud_controller" index:"2d760cbf-ea29-4e5f-8425-4412dc79cccd" ip:"10.193.70.13" logMessage:<message:"Updated app with guid 41abc841-cbc8-4cab-854d-640a7c8b6a5f ({\"state\"=>\"STARTED\"})" message_type:OUT timestamp:1487442274599132749 app_id:"41abc841-cbc8-4cab-854d-640a7c8b6a5f" source_type:"API" source_instance:"0" >
Sending input
origin:"rep" eventType:LogMessage timestamp:1487442491344232800 deployment:"cf" job:"ADFS" index:"0" ip:"10.152.9.20" logMessage:<message:"Downloading binary_buildpack..." message_type:OUT timestamp:1487442491344232800 app_id:"41abc841-cbc8-4cab-854d-640a7c8b6a5f" source_type:"STG" source_instance:"0" >
Sending input
origin:"rep" eventType:LogMessage timestamp:1487442491444240400 deployment:"cf" job:"ADFS" index:"0" ip:"10.152.9.20" logMessage:<message:"Downloading binary_buildpack failed" message_type:OUT timestamp:1487442491444240400 app_id:"41abc841-cbc8-4cab-854d-640a7c8b6a5f" source_type:"STG" source_instance:"0" >
Sending input
origin:"rep" eventType:LogMessage timestamp:1487442491602247800 deployment:"cf" job:"ADFS" index:"0" ip:"10.152.9.20" logMessage:<message:"Destroying container" message_type:OUT timestamp:1487442491602247800 app_id:"41abc841-cbc8-4cab-854d-640a7c8b6a5f" source_type:"STG" source_instance:"0" >
Sending input
origin:"rep" eventType:LogMessage timestamp:1487442491621249300 deployment:"cf" job:"ADFS" index:"0" ip:"10.152.9.20" logMessage:<message:"Failed to destroy container" message_type:ERR timestamp:1487442491621249300 app_id:"41abc841-cbc8-4cab-854d-640a7c8b6a5f" source_type:"STG" source_instance:"0" >
```

## Listing on firehose and greping for application guid

```
$ dop -url wss://doppler.system.domain:443/firehose/test123 | egrep 41abc841-cbc8-4cab-854d-640a7c8b6a5f
origin:"gorouter" eventType:HttpStartStop timestamp:1487442041999551545 deployment:"cf" job:"router" index:"6bddeda2-ac99-4c94-b3f4-60fe37a6ea2f" ip:"10.193.70.24" httpStartStop:<startTimestamp:1487442041993595204 stopTimestamp:1487442041999537830 requestId:<low:8738335304478508222 high:13624360002285430131 > peerType:Client method:GET uri:"http://api.run-05.haas-59.pez.pivotal.io/internal/log_access/41abc841-cbc8-4cab-854d-640a7c8b6a5f" remoteAddress:"10.193.70.250:47334" userAgent:"Go-http-client/1.1" statusCode:200 contentLength:0 instanceId:"a7755f15-2957-47e1-61d1-621422cfc400" forwarded:"10.193.70.27" forwarded:"10.193.70.250" >
origin:"gorouter" eventType:HttpStartStop timestamp:1487442042000076854 deployment:"cf" job:"router" index:"6bddeda2-ac99-4c94-b3f4-60fe37a6ea2f" ip:"10.193.70.24" httpStartStop:<startTimestamp:1487442041993302252 stopTimestamp:1487442042000067559 requestId:<low:8738335304478508222 high:13624360002285430131 > peerType:Server method:GET uri:"http://api.run-05.haas-59.pez.pivotal.io/internal/log_access/41abc841-cbc8-4cab-854d-640a7c8b6a5f" remoteAddress:"10.193.70.250:47334" userAgent:"Go-http-client/1.1" statusCode:200 contentLength:0 forwarded:"10.193.70.27" >
origin:"cloud_controller" eventType:LogMessage deployment:"cf" job:"cloud_controller" index:"2d760cbf-ea29-4e5f-8425-4412dc79cccd" ip:"10.193.70.13" logMessage:<message:"Updated app with guid 41abc841-cbc8-4cab-854d-640a7c8b6a5f ({\"state\"=>\"STARTED\"})" message_type:OUT timestamp:1487442042659776418 app_id:"41abc841-cbc8-4cab-854d-640a7c8b6a5f" source_type:"API" source_instance:"0" >
origin:"gorouter" eventType:HttpStartStop timestamp:1487442042680274096 deployment:"cf" job:"router" index:"6bddeda2-ac99-4c94-b3f4-60fe37a6ea2f" ip:"10.193.70.24" httpStartStop:<startTimestamp:1487442042545296946 stopTimestamp:1487442042680249094 requestId:<low:11115129583103816559 high:14773095896116596340 > peerType:Client method:PUT uri:"http://api.run-05.haas-59.pez.pivotal.io/v2/apps/41abc841-cbc8-4cab-854d-640a7c8b6a5f" remoteAddress:"10.193.70.250:47340" userAgent:"go-cli 6.23.1+a70deb38f.2017-01-13 / darwin" statusCode:201 contentLength:4595 instanceId:"a7755f15-2957-47e1-61d1-621422cfc400" forwarded:"10.64.248.138" forwarded:"10.193.70.250" >
origin:"gorouter" eventType:HttpStartStop timestamp:1487442042680762876 deployment:"cf" job:"router" index:"6bddeda2-ac99-4c94-b3f4-60fe37a6ea2f" ip:"10.193.70.24" httpStartStop:<startTimestamp:1487442042545027525 stopTimestamp:1487442042680752815 requestId:<low:11115129583103816559 high:14773095896116596340 > peerType:Server method:PUT uri:"http://api.run-05.haas-59.pez.pivotal.io/v2/apps/41abc841-cbc8-4cab-854d-640a7c8b6a5f" remoteAddress:"10.193.70.250:47340" userAgent:"go-cli 6.23.1+a70deb38f.2017-01-13 / darwin" statusCode:201 contentLength:4595 forwarded:"10.64.248.138" >
origin:"gorouter" eventType:HttpStartStop timestamp:1487442043005005173 deployment:"cf" job:"router" index:"6bddeda2-ac99-4c94-b3f4-60fe37a6ea2f" ip:"10.193.70.24" httpStartStop:<startTimestamp:1487442042985935516 stopTimestamp:1487442043004980906 requestId:<low:17098900877812520117 high:9157057007867574102 > peerType:Client method:GET uri:"http://api.run-05.haas-59.pez.pivotal.io/v2/apps/41abc841-cbc8-4cab-854d-640a7c8b6a5f" remoteAddress:"10.193.70.250:47344" userAgent:"go-cli 6.23.1+a70deb38f.2017-01-13 / darwin" statusCode:200 contentLength:1837 instanceId:"a7755f15-2957-47e1-61d1-621422cfc400" forwarded:"10.64.248.138" forwarded:"10.193.70.250" >
origin:"gorouter" eventType:HttpStartStop timestamp:1487442043005514723 deployment:"cf" job:"router" index:"6bddeda2-ac99-4c94-b3f4-60fe37a6ea2f" ip:"10.193.70.24" httpStartStop:<startTimestamp:1487442042985664710 stopTimestamp:1487442043005504411 requestId:<low:17098900877812520117 high:9157057007867574102 > peerType:Server method:GET uri:"http://api.run-05.haas-59.pez.pivotal.io/v2/apps/41abc841-cbc8-4cab-854d-640a7c8b6a5f" remoteAddress:"10.193.70.250:47344" userAgent:"go-cli 6.23.1+a70deb38f.2017-01-13 / darwin" statusCode:200 contentLength:1837 forwarded:"10.64.248.138" >
origin:"rep" eventType:LogMessage timestamp:1487442259439952000 deployment:"cf" job:"ADFS" index:"0" ip:"10.152.9.20" logMessage:<message:"Downloading binary_buildpack..." message_type:OUT timestamp:1487442259439952000 app_id:"41abc841-cbc8-4cab-854d-640a7c8b6a5f" source_type:"STG" source_instance:"0" >
origin:"rep" eventType:LogMessage timestamp:1487442259539959400 deployment:"cf" job:"ADFS" index:"0" ip:"10.152.9.20" logMessage:<message:"Downloading binary_buildpack failed" message_type:OUT timestamp:1487442259539959400 app_id:"41abc841-cbc8-4cab-854d-640a7c8b6a5f" source_type:"STG" source_instance:"0" >
origin:"rep" eventType:LogMessage timestamp:1487442259706968400 deployment:"cf" job:"ADFS" index:"0" ip:"10.152.9.20" logMessage:<message:"Destroying container" message_type:OUT timestamp:1487442259706968400 app_id:"41abc841-cbc8-4cab-854d-640a7c8b6a5f" source_type:"STG" source_instance:"0" >
origin:"rep" eventType:LogMessage timestamp:1487442259723970800 deployment:"cf" job:"ADFS" index:"0" ip:"10.152.9.20" logMessage:<message:"Failed to destroy container" message_type:ERR timestamp:1487442259723970800 app_id:"41abc841-cbc8-4cab-854d-640a7c8b6a5f" source_type:"STG" source_instance:"0" >
origin:"gorouter" eventType:HttpStartStop timestamp:1487442048331215120 deployment:"cf" job:"router" index:"6bddeda2-ac99-4c94-b3f4-60fe37a6ea2f" ip:"10.193.70.24" httpStartStop:<startTimestamp:1487442048313629397 stopTimestamp:1487442048331204213 requestId:<low:670985736938183539 high:8151325518793386364 > peerType:Client method:GET uri:"http://api.run-05.haas-59.pez.pivotal.io/v2/apps/41abc841-cbc8-4cab-854d-640a7c8b6a5f" remoteAddress:"10.193.70.250:47372" userAgent:"go-cli 6.23.1+a70deb38f.2017-01-13 / darwin" statusCode:200 contentLength:1873 instanceId:"a7755f15-2957-47e1-61d1-621422cfc400" forwarded:"10.64.248.138" forwarded:"10.193.70.250" >
origin:"gorouter" eventType:HttpStartStop timestamp:1487442048331680599 deployment:"cf" job:"router" index:"6bddeda2-ac99-4c94-b3f4-60fe37a6ea2f" ip:"10.193.70.24" httpStartStop:<startTimestamp:1487442048313392423 stopTimestamp:1487442048331671180 requestId:<low:670985736938183539 high:8151325518793386364 > peerType:Server method:GET uri:"http://api.run-05.haas-59.pez.pivotal.io/v2/apps/41abc841-cbc8-4cab-854d-640a7c8b6a5f" remoteAddress:"10.193.70.250:47372" userAgent:"go-cli 6.23.1+a70deb38f.2017-01-13 / darwin" statusCode:200 contentLength:1873 forwarded:"10.64.248.138" >
origin:"gorouter" eventType:HttpStartStop timestamp:1487442053415493677 deployment:"cf" job:"router" index:"6bddeda2-ac99-4c94-b3f4-60fe37a6ea2f" ip:"10.193.70.24" httpStartStop:<startTimestamp:1487442041986333923 stopTimestamp:1487442053415450888 requestId:<low:6216517624099443007 high:8021246184761303400 > peerType:Server method:GET uri:"http://doppler.run-05.haas-59.pez.pivotal.io:443/apps/41abc841-cbc8-4cab-854d-640a7c8b6a5f/stream" remoteAddress:"10.193.70.250:47332" userAgent:"Go-http-client/1.1" statusCode:200 contentLength:0 forwarded:"10.64.248.138" forwarded:"10.193.70.250" >
```

## Conntecting directly to the loggregator component

if you wnat to bypass a load blanacer and go directly to loggregrate via the unecrypted path then try this

**Note** `ws` for unecrypted and `wss` for TLS encryption

streaming logs

```
./dop -url ws://10.193.67.65:8081/apps/<APP GUID>/stream
```

connect to firehose

```
./dop -url ws://10.193.67.65:8081/firehose/test123
```
