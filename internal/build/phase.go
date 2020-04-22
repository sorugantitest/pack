package build

import (
	"context"

	"github.com/docker/docker/api/types"
	dcontainer "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"

	"github.com/buildpacks/pack/internal/container"
	"github.com/buildpacks/pack/logging"
)

type Phase struct {
	name       string
	logger     logging.Logger
	docker     client.CommonAPIClient
	ctrConf    *dcontainer.Config
	hostConf   *dcontainer.HostConfig
	ctr        dcontainer.ContainerCreateCreatedBody
	uid, gid   int
	fileFilter func(string) bool
}

func (p *Phase) Run(ctx context.Context) error {
	var err error
	p.ctr, err = p.docker.ContainerCreate(ctx, p.ctrConf, p.hostConf, nil, "")
	if err != nil {
		return errors.Wrapf(err, "failed to create '%s' container", p.name)
	}

	return container.Run(
		ctx,
		p.docker,
		p.ctr.ID,
		logging.NewPrefixWriter(logging.GetWriterForLevel(p.logger, logging.InfoLevel), p.name),
		logging.NewPrefixWriter(logging.GetWriterForLevel(p.logger, logging.ErrorLevel), p.name),
	)
}

func (p *Phase) Cleanup() error {
	return p.docker.ContainerRemove(context.Background(), p.ctr.ID, types.ContainerRemoveOptions{Force: true})
}
