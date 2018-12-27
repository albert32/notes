#!/usr/bin/env python
# -*- coding:utf-8 -*-

'''
    redis分布式锁简易实现
'''

import time
import logging

TTL = 2 * 60  # 加锁后2超时2分钟自动删除锁


class RedisLock(object):
    def __init__(self):
        pass

    def get_lock(self, redisClient, lock_key):
        if not redisClient:
            logging.error('[REDIS ERROR], could not get redis!')
            return False

        result = False
        try:
            now_time = int(time.time())
            result = redisClient.setnx(lock_key, now_time)
            if result:
                redisClient.expire(lock_key, TTL)
            else:
                res = redisClient.get(lock_key)
                if res:
                    last_time = int(res)
                    if (last_time+TTL) < now_time:
                        logging.warn("RedisLock, lock release expire.")
                        redisClient.delete(lock_key)
                        result = redisClient.setnx(lock_key, now_time)
                        redisClient.expire(lock_key, TTL)
        except Exception, ex:
            logging.error("RedisLock, get_lock error, ex: %s" % ex)
        return result

    def release_lock(self, redisClient, lock_key):
        if not redisClient:
            logging.error('[REDIS ERROR], could not get redis!')
            return
        try:
            redisClient.delete(lock_key)
        except Exception, ex:
            logging.error("RedisLock,release_lock error,ex:%s" % ex)


REDISLOCK = RedisLock()

if __name__ == "__main__":
    pass
