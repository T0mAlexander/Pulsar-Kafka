#!/bin/sh

cp -r apache-pulsar-3.1.1/* pulsar/
rm -rf apache-pulsar-3.1.1-bin.tar.gz apache-pulsar-3.1.1
pulsar standalone