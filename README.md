# RTMP SERVER 

![NGINX](https://img.shields.io/static/v1?label=NGINX&labelColor=07b60b&message=NGX&color=000000&logo=NGINX&logoColor=ffffff&style=flat-square)
![](https://img.shields.io/static/v1?label=Real-Time+Messaging+Protocol&labelColor=06a7ac&message=RTMP&color=000000&logo=&logoColor=ffffff&style=flat-square)
![](https://img.shields.io/static/v1?label=Http+Live+Streaming&labelColor=06a7ac&message=HLS&color=000000&logo=&logoColor=ffffff&style=flat-square)
![go](https://img.shields.io/static/v1?label=Golang+1.20&labelColor=08dae1&message=Go&color=000000&logo=go&logoColor=ffffff&style=flat-square)

A simple live video streaming service using nginx, rtmp, hls and golang (for authentication service).

<br>

### What's RTMP ?

<br>

RTMP (Real-time messaging protocol) is a data transmission technology developed by Macromidia, initially dedicated to transmitting data between streaming servers and the old Adobe Flash Player.

Nowadays the RTMP protocol is constantly used by services such as Facebook, Twitter, and Twitch to deliver live video to their users. 

In RTMP the content is delivered in chunks, usually called "chunks", so the user can consume the content without having to wait for the complete download.

<br>

### How a RTMP transmission works?

<br>

A RTMP transmission works in three-steps:

1. **Handshake** - The client, which can be an encoder software or hardware, that wants to transmit data initiates a connection by exchanging 3 packets. 
    * **The first packet**  - sent by the client is the RTMP version that is being streamed. 
    * **The second packet** - sent by the client shortly after, is a timestamp. The server replies with a copy of the two data it has just received.
    * **The final packet**  - When the duplex communication is complete, the client sends a copy of the timestamp again. The server returns the copy. When the last packet is returned from the server, the handshake is successfully completed. 