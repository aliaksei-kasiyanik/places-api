import random

id_get_places = "/places/{} get_by_id\n"
geo_get_places = "/places?lat={}&lon={}&rad=500 get_by_geo\n"


def generate_requests_id():
    res = []
    with open("ids.txt", "r") as id_file:
        for line in id_file:
            res.append(id_get_places.format(line.strip()))
    return res


def generate_requests_geo():
    res = []
    for i in range(0, 100):
        req = geo_get_places.format(random.uniform(53.843070, 53.974931), random.uniform(27.430154, 27.671853))
        res.append(req)
    return res


def write_requests_to_file(requests):
    with open("requests", "w") as r_file:
        for r in requests:
            r_file.write(r)


reqs_id = generate_requests_id()
reqs_geo = generate_requests_geo()

total = reqs_id + reqs_geo

random.shuffle(total)

write_requests_to_file(total)
