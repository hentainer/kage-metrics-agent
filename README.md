# kage-metrics-agent

[![Go Report Card](https://goreportcard.com/badge/github.com/hentainer/kage-metrics-agent)](https://goreportcard.com/report/github.com/hentainer/kage-metrics-agent)
[![Build Status](https://travis-ci.org/hentainer/kage-metrics-agent.svg?branch=master)](https://travis-ci.org/hentainer/kage-metrics-agent)
[![Docker build](https://img.shields.io/docker/automated/hentainer/kage-metrics-agent.svg)](https://hub.docker.com/r/hentainer/kage-metrics-agent/)
[![Coverage Status](https://coveralls.io/repos/github/hentainer/kage-metrics-agent/badge.svg?branch=master)](https://coveralls.io/github/hentainer/kage-metrics-agent?branch=master)
[![GitHub release](https://img.shields.io/github/release/hentainer/kage-metrics-agent.svg)](https://github.com/hentainer/kage-metrics-agent/releases)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hentainer/kage-metrics-agent/master/LICENSE)

## Synopsis

Kage (as in \"Kafka AGEnt\") reads Offset- and Lag metrics from Kafka and writes them to an InfluxDB.

## Motivation

Running Kafka often requires robust monitoring solutions. One way to tackle this is by querying the beans directly