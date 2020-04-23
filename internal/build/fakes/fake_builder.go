package fakes

import (
	"github.com/Masterminds/semver"
	"github.com/buildpacks/imgutil"
	ifakes "github.com/buildpacks/imgutil/fakes"

	"github.com/buildpacks/pack/internal/api"
	"github.com/buildpacks/pack/internal/builder"
)

type FakeBuilder struct {
	ReturnForImage               *ifakes.Image
	ReturnForUID                 int
	ReturnForGID                 int
	ReturnForLifecycleDescriptor builder.LifecycleDescriptor
}

func NewFakeBuilder(ops ...func(*FakeBuilder)) (*FakeBuilder, error) {
	infoVersion, err := semver.NewVersion("12.34")
	if err != nil {
		return nil, err
	}

	platformAPIVersion, err := api.NewVersion("23.45")
	if err != nil {
		return nil, err
	}

	buildpackVersion, err := api.NewVersion("34.56")
	if err != nil {
		return nil, err
	}

	fakeBuilder := &FakeBuilder{
		ReturnForImage: ifakes.NewImage("some-name", "", nil),
		ReturnForUID:  99,
		ReturnForGID:  99,
		ReturnForLifecycleDescriptor: builder.LifecycleDescriptor{
			API: builder.LifecycleAPI{
				BuildpackVersion: buildpackVersion,
				PlatformVersion:  platformAPIVersion,
			},
			Info: builder.LifecycleInfo{
				Version: &builder.Version{Version: *infoVersion},
			},
		},
	}

	for _, op := range ops {
		op(fakeBuilder)
	}

	return fakeBuilder, nil
}

func WithImage(image *ifakes.Image) func(*FakeBuilder) {
	return func(builder *FakeBuilder) {
		builder.ReturnForImage = image
	}
}

func WithPlatformVersion(version *api.Version) func(*FakeBuilder) {
	return func(builder *FakeBuilder) {
		builder.ReturnForLifecycleDescriptor.API.PlatformVersion = version
	}
}

func (b *FakeBuilder) Name() string {
	return b.ReturnForImage.Name()
}

func (b *FakeBuilder) Image() imgutil.Image {
	return b.ReturnForImage
}

func (b *FakeBuilder) UID() int {
	return b.ReturnForUID
}

func (b *FakeBuilder) GID() int {
	return b.ReturnForGID
}

func (b *FakeBuilder) LifecycleDescriptor() builder.LifecycleDescriptor {
	return b.ReturnForLifecycleDescriptor
}
