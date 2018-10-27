VERSION := "latest"
IMAGE_APP_C  := "rsyslog-app-c"
IMAGE_APP_GO := "rsyslog-app-go"
IMAGE_APP_PY := "rsyslog-app-py"
IMAGE_RSYSLOG := "rsyslog"
REPO := "danielkraic"

build:
	docker build -t $(REPO)/$(IMAGE_RSYSLOG):$(VERSION) rsyslog
	docker build -t $(REPO)/$(IMAGE_APP_C):$(VERSION) app-c
	docker build -t $(REPO)/$(IMAGE_APP_GO):$(VERSION) app-go
	docker build -t $(REPO)/$(IMAGE_APP_PY):$(VERSION) app-py

push: build
	docker push $(REPO)/$(IMAGE_RSYSLOG):$(VERSION)
	docker push $(REPO)/$(IMAGE_APP_C):$(VERSION)
	docker push $(REPO)/$(IMAGE_APP_GO):$(VERSION)
	docker push $(REPO)/$(IMAGE_APP_PY):$(VERSION)

clean:
	docker rmi $(REPO)/$(IMAGE_RSYSLOG):$(VERSION)
	docker rmi $(REPO)/$(IMAGE_APP_C):$(VERSION)
	docker rmi $(REPO)/$(IMAGE_APP_GO):$(VERSION)
	docker rmi $(REPO)/$(IMAGE_APP_PY):$(VERSION)


.PHONY: all build push release clean
