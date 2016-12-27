import requests
import json

fs_url_ll = 'https://api.foursquare.com/v2/venues/search?ll={},{}&limit=50&client_id={}&client_secret={}&v=20161226'
fs_url_near = 'https://api.foursquare.com/v2/venues/search?near={}&intent=browse&limit=50&client_id={}&client_secret={}&v=20161226'
fs_url_region = 'https://api.foursquare.com/v2/venues/search?ne={},{}&sw={},{}&categoryId={}&intent=browse&limit=50&client_id={}&client_secret={}&v=20161226'
clientId = "BFZZSMOGHVWKEXNTMB5KR1MOARDSQFWC1QWTGQ4MGIA5YRTZ"
clientSecret = "2Q2V0BS5O2C51J1SUA2E5HOFXDN2TQ5T3OSQIXENLGQQVMLR"

categories = [
	# "4d4b7104d754a06370d81259" Arts
	"4d4b7105d754a06379d81259", # Travel and transport
	# "4d4b7105d754a06377d81259"  # Outdoor and recreation
]


def pretty_json(res):
	print(json.dumps(res, sort_keys=True, indent=4, ensure_ascii=False))


def get_fs_response_ll(lat, lon):
	url = fs_url_ll.format(lat, lon, clientId, clientSecret)
	print(url)
	res = requests.get(url)
	if res.status_code != 200:
		return None
	return requests.get(url).json()


def get_fs_response_near(near):
	url = fs_url_near.format(near, clientId, clientSecret)
	print(url)
	res = requests.get(url)
	if res.status_code != 200:
		return None
	return requests.get(url).json()


def get_fs_response_region(ne_lon, ne_lat, sw_lon, sw_lat):
	cat_ids = ",".join(categories)
	url = fs_url_region.format(ne_lon, ne_lat, sw_lon, sw_lat, cat_ids, clientId, clientSecret)
	print(url)
	res = requests.get(url)
	if res.status_code != 200:
		print(res.status_code, res.text)
		return None
	return requests.get(url).json()


def create_new_place(lon, lat, name, categories, fs_id):
	place = {}
	location = {}
	location["Type"] = "Point"
	location["coordinates"] = [lon, lat]
	place["location"] = location
	place["name"] = name
	place["categories"] = categories
	place["fsId"] = fs_id
	return place


def process_response(res):
	places = []
	venues = res["response"]["venues"]
	print(len(venues))
	for venue in venues:
		fs_id = venue["id"]
		location = venue["location"]
		name = venue["name"]
		if not name:
			continue
		lat = location["lat"]
		lng = location["lng"]
		categories = []
		fsCategories = venue["categories"]
		for fsCat in fsCategories:
			categories.append(fsCat["name"])
		places.append(create_new_place(lng, lat, name, categories, fs_id))

	return places


def get_map_grid(ne_lon, ne_lat, sw_lon, sw_lat, step=0.01):
	loc_lon = ne_lon
	loc_lat = ne_lat
	while loc_lat > sw_lat:
		while loc_lon > sw_lon:
			yield (loc_lon, loc_lat)
			loc_lon -= step
		loc_lat -= step
		loc_lon = ne_lon


def run():
	# res = get_fs_response_ll(53.8835622, 27.3147977)
	# res = get_fs_response_near("Minsk")
	# lon = 53.924015
	# lat = 27.549630
	step = 0.005
	for loc in get_map_grid(53.974931, 27.671853, 53.843070, 27.430154, step):
		res = get_fs_response_region(loc[0], loc[1], loc[0] - step, loc[1] - step)
		if res is not None:
			# pretty_json(res)
			places = process_response(res)
			# pretty_json(places)
			for place in places:
				r = requests.post("http://gotravel.today/places", json=place)
				if r.status_code != 201:
					print(r.status_code, r.text)


run()

