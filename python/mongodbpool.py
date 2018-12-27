#!/usr/bin/env python
# -*- coding:utf-8 -*-

import logging
from pymongo import MongoClient, ASCENDING, DESCENDING, read_preferences


class MongoDbPool(object):
    def __init__(self):
        self.mongo_db = None

    def start(self, HOST, PORT, thread_num, DB, USER, PWD):
        try:
            conn = MongoClient(HOST, PORT, thread_num)
            self.mongo_db = conn[DB]
            self.mongo_db.authenticate(USER, PWD)
            logging.info(
                "MongoDbPool init ok, host:%s, port:%s, db:%s" % (HOST, PORT, DB))
        except Exception, ex:
            logging.info(
                "MongoDbPool-connect error to db : %s[%d],to db : %s,%s" % (HOST, int(PORT), DB, ex))

    def get_mongo_db(self, index_num):
        return self.mongo_db


class MongoDbPoolSecondary(object):
    def __init__(self):
        self.mongo_db = None

    def start(self, HOST, PORT, thread_num, DB, USER, PWD):
        try:
            conn = MongoClient(
                HOST, PORT, thread_num, read_preference=read_preferences.ReadPreference.SECONDARY_PREFERRED)
            self.mongo_db = conn[DB]
            self.mongo_db.authenticate(USER, PWD)
            logging.info("MongoDbPoolSecondary init ok")
        except Exception, ex:
            logging.info(
                "MongoDbPoolSecondary-connect error to db : %s[%d],to db : %s,%s" % (HOST, int(PORT), DB, ex))

    def get_mongo_db(self, index_num):
        return self.mongo_db


GMongoDBPOOL = MongoDbPool()
GMongoDBPOOLSecondary = MongoDbPoolSecondary()

if __name__ == "__main__":
    import json
    from bson import ObjectId
    GMongoDBPOOL.start("192.168.229.100", 27017, 1, "db", "user", "pwd")
    mongo_db = GMongoDBPOOL.get_mongo_db(1)

    query = {}
    res = mongo_db["table_name"].find(query)
    #res = mongo_db.group_users.find()
    for key in res:
        print key
