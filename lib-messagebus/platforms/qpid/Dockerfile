################################################################################
# (C) Copyright [2020] Anand S <Anand.S.Main@gmail.com>

# Licensed under the Apache License, Version 2.0 (the "License"); you may
# not use this file except in compliance with the License. You may obtain
# a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations
# under the License.
################################################################################
# BUILD : docker build --tag dockqpid:1.0 .
# RUN   : docker run -ti dockqpid
################################################################################

FROM ubuntu:latest
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
	apt-get install -y \
	gcc \
	g++ \
	automake \
	libwebsockets-dev \
	libtool \
	zlib1g-dev \
	cmake \
	libsasl2-dev \
	libssl-dev \
	python \
	python-dev \
	libuv1-dev \
	sasl2-bin \
	swig \
	maven \
	git && \
	apt-get -y clean

RUN git clone https://gitbox.apache.org/repos/asf/qpid-dispatch.git \
	&& cd /qpid-dispatch \
	&& git submodule add https://gitbox.apache.org/repos/asf/qpid-proton.git \
	&& git submodule update --init

WORKDIR /qpid-dispatch

RUN mkdir qpid-proton/build \
	&& cd qpid-proton/build \
	&& cmake .. -DSYSINSTALL_BINDINGS=ON -DCMAKE_INSTALL_PREFIX=/usr \
        -DSYSINSTALL_PYTHON=ON && make install

WORKDIR /qpid-dispatch

RUN mkdir build && cd build && \
	cmake .. -DCMAKE_INSTALL_PREFIX=/usr -DUSE_VALGRIND=NO && \
	cmake --build . --target install

ENV PYTHONPATH=/usr/lib/python2.7/site-packages
WORKDIR /qpid-dispatch
ENTRYPOINT ["qdrouterd"]
