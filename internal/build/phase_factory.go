package build

type DefaultPhaseFactory struct {
	lifecycle *Lifecycle
}

func NewDefaultPhaseFactory(lifecycle *Lifecycle) *DefaultPhaseFactory {
	return &DefaultPhaseFactory{lifecycle: lifecycle}
}

func (d *DefaultPhaseFactory) New(provider *PhaseConfigProvider) RunnerCleaner {
	return &Phase{
		ctrConf:    provider.ContainerConfig(),
		hostConf:   provider.HostConfig(),
		name:       provider.Name(),
		docker:     d.lifecycle.docker,
		logger:     d.lifecycle.logger,
		uid:        d.lifecycle.builder.UID(),
		gid:        d.lifecycle.builder.GID(),
		fileFilter: d.lifecycle.fileFilter,
	}
}
