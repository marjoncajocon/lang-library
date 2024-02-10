# Software: nano v.60
# Developer: Marjon Cajocon
# -> to easily calculate age of an object or things from the past time

from datetime import datetime, timedelta
from re import match

def age(date_str: str) -> dict:
	res: dict = {
		'ok': True,
		'status': None,
		'date': [
			['year',   0, 31536000],
			['month',  0, 2592000],
			['week',   0, 604800],
			['day',    0, 86400],
			['hour',   0, 3600],
			['minute', 0, 60],
			['second', 0, 1]
		]
	}

	n: datetime = datetime.now()
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
	
	i = 0	

	while i < date_list_len:
		r: int = d // date_list[i][2] # double // mean result return to int
		d = d % date_list[i][2]
		date_list[i][1] = r
		i += 1
	
	# print(res['date']) for debugging
	
	return  res
