// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/orb/test/bdd

go 1.17

require (
	github.com/cenkalti/backoff/v4 v4.1.2
	github.com/cucumber/godog v0.9.0
	github.com/cucumber/messages-go/v10 v10.0.3
	github.com/fsouza/go-dockerclient v1.6.5
	github.com/google/uuid v1.3.0
	github.com/greenpau/go-calculator v1.0.1
	github.com/hyperledger/aries-framework-go v0.1.9-0.20220425190801-5d239d542d7d
	github.com/hyperledger/aries-framework-go-ext/component/storage/couchdb v0.0.0-20220310013829-55b4443130f8
	github.com/hyperledger/aries-framework-go-ext/component/storage/mongodb v0.0.0-20220425190739-5c1b81a74df6
	github.com/hyperledger/aries-framework-go-ext/component/vdr/orb v0.1.4-0.20220412155039-c618d62571b8
	github.com/hyperledger/aries-framework-go-ext/component/vdr/sidetree v0.0.0-20220330145438-233eb7999876
	github.com/hyperledger/aries-framework-go/component/storageutil v0.0.0-20220425190801-5d239d542d7d
	github.com/hyperledger/aries-framework-go/spi v0.0.0-20220330140627-07042d78580c
	github.com/igor-pavlenko/httpsignatures-go v0.0.23
	github.com/ipfs/go-ipfs-api v0.2.0
	github.com/mr-tron/base58 v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/tidwall/gjson v1.7.4
	github.com/trustbloc/orb v0.1.4-0.20220412070324-edc9187e0c3e
	github.com/trustbloc/sidetree-core-go v1.0.0-rc.1
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/Microsoft/hcsshim v0.8.18 // indirect
	github.com/ThreeDotsLabs/watermill v1.2.0-rc.7 // indirect
	github.com/VictoriaMetrics/fastcache v1.5.7 // indirect
	github.com/bluele/gcache v0.0.2 // indirect
	github.com/btcsuite/btcd v0.22.0-beta // indirect
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/containerd/cgroups v1.0.1 // indirect
	github.com/containerd/containerd v1.5.5 // indirect
	github.com/crackcomm/go-gitignore v0.0.0-20170627025303-887ab5e44cc3 // indirect
	github.com/cucumber/gherkin-go/v11 v11.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/docker v20.10.7+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/evanphx/json-patch v4.9.0+incompatible // indirect
	github.com/fxamacker/cbor/v2 v2.3.0 // indirect
	github.com/go-kivik/couchdb/v3 v3.2.6 // indirect
	github.com/go-kivik/kivik/v3 v3.2.3 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/certificate-transparency-go v1.1.2-0.20210512142713-bed466244fa6 // indirect
	github.com/google/tink/go v1.6.1 // indirect
	github.com/google/trillian v1.3.14-0.20210520152752-ceda464a95a3 // indirect
	github.com/ipfs/go-cid v0.0.7 // indirect
	github.com/ipfs/go-ipfs-files v0.0.8 // indirect
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a // indirect
	github.com/kilic/bls12-381 v0.1.1-0.20210503002446-7b7597926c69 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/libp2p/go-buffer-pool v0.0.2 // indirect
	github.com/libp2p/go-flow-metrics v0.0.3 // indirect
	github.com/libp2p/go-libp2p-core v0.6.1 // indirect
	github.com/libp2p/go-openssl v0.0.7 // indirect
	github.com/lithammer/shortuuid/v3 v3.0.7 // indirect
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1 // indirect
	github.com/minio/sha256-simd v0.1.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/moby/sys/mount v0.2.0 // indirect
	github.com/moby/sys/mountinfo v0.4.1 // indirect
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/multiformats/go-base32 v0.0.3 // indirect
	github.com/multiformats/go-base36 v0.1.0 // indirect
	github.com/multiformats/go-multiaddr v0.3.0 // indirect
	github.com/multiformats/go-multiaddr-net v0.2.0 // indirect
	github.com/multiformats/go-multibase v0.0.3 // indirect
	github.com/multiformats/go-multihash v0.0.14 // indirect
	github.com/multiformats/go-varint v0.0.6 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/opencontainers/runc v1.0.2 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/piprate/json-gold v0.4.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/pquerna/cachecontrol v0.1.0 // indirect
	github.com/spacemonkeygo/spacelog v0.0.0-20180420211403-2296661a0572 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/square/go-jose/v3 v3.0.0-20200630053402-0a67ce9b0693 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/teserakt-io/golang-ed25519 v0.0.0-20210104091850-3888c087a4c8 // indirect
	github.com/tidwall/match v1.0.3 // indirect
	github.com/tidwall/pretty v1.1.0 // indirect
	github.com/trustbloc/edge-core v0.1.8 // indirect
	github.com/trustbloc/vct v1.0.0-rc1.0.20220426133750-e02ec33d2826 // indirect
	github.com/whyrusleeping/tar-utils v0.0.0-20180509141711-8c6c8ba81d5c // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	go.mongodb.org/mongo-driver v1.8.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220408201424-a24fb2fb8a0f // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20220222213610-43724f9ea8cf // indirect
	google.golang.org/grpc v1.44.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/trustbloc/orb => ../../
