# If you don‘t want to build it youself, you can try `docker pull killercai/postgres`.
FROM healthcheck/postgres:latest

# China debian mirror
RUN sed -i s@/deb.debian.org/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN apt-get clean && apt-get update
RUN apt-get install -y wget git build-essential libpq-dev python-dev postgresql-server-dev-all
# SCWS (Simple Chinese Word Segmentation library)
RUN cd /tmp && wget -q -O - http://www.xunsearch.com/scws/down/scws-1.2.1.tar.bz2 | tar xjf - && cd scws-1.2.1 && ./configure && make install
# zhpaser (postgres plugin)
RUN cd /tmp && git clone https://github.com/amutu/zhparser.git && cd zhparser && make && make install