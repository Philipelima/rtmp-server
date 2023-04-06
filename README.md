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

```mermaid 

    graph TD;
        R[Font]-->e[Encoder]-->S[RMTP Server];
        S-->C[Client - HLS];
      
```