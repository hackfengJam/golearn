# -*- encoding: utf-8 -*-

import re
import sys
import os
import json

# Author : 青峰<me@heytaoge.com>


root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
default_port = 80

def main():
	args = sys.argv
	if len(args)<2:
		return default_port
	service = args[1]
	conf_file = os.path.join(os.path.join(os.path.join(os.path.join(root, "src"), "ptapp.cn"),service),"%s-prod.conf"%service)
	if not os.path.exists(conf_file):
		sys.stderr.write("filepath %s not exists"%conf_file)
		return default_port

	with open(conf_file, "r") as f:
		try:
			config = json.loads(f.read())
		except Exception as e:
			sys.stderr.write(str(e))
			config = {}
	return config.get("prometheus_exposer",{}).get("addr",":80")[1:]



if __name__ == '__main__':
	print(main())
