FROM centos:7
MAINTAINER  leozvc 
#phantomjs
COPY ./phantomjs /usr/local/bin/ 

#ProxyPool
COPY ./ProxyPool /usr/local/bin/

#configure
RUN mkdir /data; cd /data

COPY ./config.json /data/

WORKDIR /data

EXPOSE 8080

#run 
CMD ProxyPool 




