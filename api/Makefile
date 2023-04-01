.PHONY: genall
genall: genproto genswagger

.PHONY: genproto
genproto:
	hack/genproto.sh

.PHONY: genswagger
genswagger:
	hack/genswagger.sh

install_protoc_dependency:
	hack/install_protoc_dependency.sh

.PHONY: clean
clean:
	find ./ | grep .pb.go |xargs rm -f
	find ./ | grep .pb.gw.go |xargs rm -f
	find ./ | grep .pb.validate.go |xargs rm -f

.PHONY: verify-proto
verify-proto:
	hack/verify-proto.sh

.PHONY: verify-swagger
verify-swagger:
	hack/verify-swagger.sh
