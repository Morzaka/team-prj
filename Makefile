serve:
	swagger serve -F=swagger ./swagger/swagger.yaml


validate:
	swagger validate ./swagger/swagger.yml


.PHONY:swaggerApi validate
