# time-sorted-list

[![CircleCI](https://circleci.com/gh/go-zen-chu/time-sorted-list.svg?style=svg)](https://circleci.com/gh/go-zen-chu/time-sorted-list)

Time sorted list in Golang.

## What this data structure for?

- in memory time series data store
- any data sorted in time series
- query data using from or until in unixtime

<img src="./doc/data_structure.jpg" width=400/>

## Background

When we handle time series data, we need something like a sorted list (In Golang, by implementing sort.Interface we can do this).

Most of cases, we need not only storing data in time sequence but also querying data by time.

This data structure is intented to handle such cases.

## On going
- User can decide old or new data should be dropped if list is filled
