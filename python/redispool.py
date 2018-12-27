#!/usr/bin/env python
# -*- coding:utf-8 -*-
import logging
import redis
import redis.sentinel
import time
import threading
import random

LOG_DEBUG = False


class RedisPool(object):
    def __init__(self):
        self.redisClient = None
        self.use_redis = True  # true redis.false sentinal

    def start(self, redis_ip, redis_port, thread_num, id=0, redis_timeout=2):
        self.redisClient = redis.StrictRedis(
            host=redis_ip, port=redis_port, db=0, socket_timeout=redis_timeout)
        logging.info("RedisPool init ok,redis_ip : %s ,redis_port : %d,id : %d,redis_timeout:%s" % (
            redis_ip, redis_port, id, redis_timeout))
        self.use_redis = True

    def get_redis(self, index_num):
        if self.use_redis == False:
            if self.redisClient:
                if LOG_DEBUG:  # get_master_address比较耗性能 debug模式才打开
                    logging.debug("RedisPool--get redis-sentinel:%s" %
                                  (str(self.redisClient.connection_pool.get_master_address())))
        return self.redisClient

    def get_thread_index(self):
        return int(threading.currentThread().name.split(".")[1])

    def get_thread_redis(self):
        return self.redisClient

    def start_sentinel(self, sentinel_config):
        self.use_redis = False
        sentinels_list = sentinel_config['sentinels']
        random.shuffle(sentinels_list)
        sentinel_pool = redis.sentinel.Sentinel(sentinels_list,
                                                socket_timeout=sentinel_config['socket_timeout'])
        self.redisClient = sentinel_pool.master_for(
            sentinel_config['cluster_tag'], db=0, password="")
        logging.info("RedisPool init ok,use sentinel:%s" % (sentinel_config))


GREDISPOOL = RedisPool()

if __name__ == "__main__":

    pass
