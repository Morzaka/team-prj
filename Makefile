validate:
	swagger validate ./swagger/swagger.yaml


serve:
	swagger serve -F=swagger ./swagger/swagger.yaml


all:
	swagger validate ./swagger/swagger.yaml
	swagger serve -F=swagger ./swagger/swagger.yaml



