# Software: nano v.60
# Developer: Marjon Cajocon
# -> to easily calculate age of an object or things from the past time

from datetime import datetime, timedelta
from re import match

date_config = {
	'time_zone': 8
}

def age(date_str: str, last_date = None) -> dict:
	res: dict = {
		'ok': True,
		'status': None,
		'date': [
			['year',   0, 31536000, 'Y'],
			['month',  0, 2592000,  'M'],
			['week',   0, 604800,   'W'],
			['day',    0, 86400,    'd'],
			['hour',   0, 3600,     'h'],
			['minute', 0, 60,       'm'],
			['second', 0, 1,        's']
		]
	}


	n: datetime = datetime.now() + timedelta(hours = date_config['time_zone'])

	if last_date:
		pass

	c: datetime = None

	# condition for dynamic date checking 
	if match(r'^[0-9]{4,5}[\-][0-9]{2}[\-][0-9]{2}[ ][0-9]{2}[\:][0-9]{2}[\:][0-9]{2}$', date_str):
		c = datetime.strptime(date_str, '%Y-%m-%d %H:%M:%S')
	elif match(r'^[0-9]{4,5}[\-][0-9]{2}[\-][0-9]{2}[ ][0-9]{2}[\:][0-9]{2}$', date_str):
		c = datetime.strptime(date_str, '%Y-%m-%d %H:%M')
	elif match(r'^[0-9]{4,5}[\-][0-9]{2}[\-][0-9]{2}$', date_str):
		c = datetime.strptime(date_str, '%Y-%m-%d')
	else:
		res['ok'] = False
		res['status'] = 'Invalid Date'
		return res

	delta_t: timedelta = n - c

	d: int =  int(delta_t.total_seconds()) #convert float to int

	date_list     = res['date']
	date_list_len = len(date_list)

	i:int = 0

	while i < date_list_len:
		r: int = d // date_list[i][2] # double // mean result return to int
		d = d % date_list[i][2]
		date_list[i][1] = r
		i += 1

	# print(res['date']) for debugging

	return  res

def ago_simple(date_str) -> str:
	ret = []

	data:dict = age(date_str)

	date: list = data['date']
	date_len: int = len(date)

	i: int  = 0

	while i < date_len:

		if date[i][1] != 0:

			j: int = i + 1

			ret.append({
				'v': date[i][1],
				'n': date[i][3],
				'f': date[i][0]
			})

			if j < date_len:
				ret.append({
					'v': date[j][1],
					'n': date[j][3],
					'f': date[j][0]
				})

			break

		i += 1

	ret_str = ''

	for x in ret:
		ret_str += f'{x["v"]}{x["n"]} '

	return ret_str
