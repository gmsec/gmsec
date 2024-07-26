"""
服务端,向etcd注册服务。
"""
import time
import sched
import threading

import etcd3
from flask import Flask, request, jsonify

app = Flask(__name__)

# etcd 连接对象
client = etcd3.client(host="192.155.1.121", port=30400)  # etcd默认的启动配置
service_name = "gmsec.srv.pytest"

# etcd key的过期时间,又叫租期续约时间
lease_timeout = 15
lease = client.lease(lease_timeout)
is_lease = False

# 定时任务对象
schedule = sched.scheduler(time.time, time.sleep)


def something_schedule():
    t = threading.Thread(target=schedul_job)
    t.setDaemon(True)
    t.start()


def schedul_job():
    inc = 3  # 定时任务,循环时间
    schedule.enter(0, 0, lease_server, (inc, schedule))  # 加入定时任务
    schedule.run()  # 执行


def lease_server(inc, schedule):
    global is_lease
    # 进行租期约需
    path = service_name
    value = "127.0.0.1:9001"
    key = path + "@" + value
    if is_lease:
        print("------------续约---------", time.time())
        lease.refresh()  # 刷新续约时间, 如果直接put,会触发监听变化,续约则不会。
    else:
        client.put(key, value, lease=lease)
        is_lease = True
        print("-----------服务注册---------", time.time())
    # 此处让定时任务循环执行,类似一个递归的样子
    schedule.enter(inc, 0, lease_server, (inc, schedule))


@app.route("/")
def server_views():
    print("------------headers---------------")
    print(str(request.headers))
    print("-----------request body-----------")
    print(request.data.decode())
    return jsonify({"code": 0, "msg": "from s1"})


if __name__ == '__main__':
    something_schedule()
    app.run(host="127.0.0.1", port=9001)
