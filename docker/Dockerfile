FROM centos:7
MAINTAINER  leozvc 

RUN mkdir /data/

ENV PATH /data:$PATH

#phantomjs
COPY ./phantomjs /data/ 
RUN yum install -y fontconfig 
#ProxyPool
COPY ./ProxyPool /data/

#configure
COPY ./config.json /data/

COPY ./phantomjs_fetcher.js /data/


WORKDIR /data

EXPOSE 8080

#run
CMD ProxyPool




