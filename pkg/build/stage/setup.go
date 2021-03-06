package stage

import (
	"github.com/flant/dapp/pkg/build/builder"
	"github.com/flant/dapp/pkg/config"
	"github.com/flant/dapp/pkg/image"
	"github.com/flant/dapp/pkg/util"
)

func GenerateSetupStage(dimgBaseConfig *config.DimgBase, gaPatchStageOptions *NewGaPatchStageOptions, baseStageOptions *NewBaseStageOptions) *SetupStage {
	b := getBuilder(dimgBaseConfig, baseStageOptions)
	if b != nil && !b.IsSetupEmpty() {
		return newSetupStage(b, gaPatchStageOptions, baseStageOptions)
	}

	return nil
}

func newSetupStage(builder builder.Builder, gaPatchStageOptions *NewGaPatchStageOptions, baseStageOptions *NewBaseStageOptions) *SetupStage {
	s := &SetupStage{}
	s.UserWithGAPatchStage = newUserWithGAPatchStage(builder, Setup, gaPatchStageOptions, baseStageOptions)
	return s
}

type SetupStage struct {
	*UserWithGAPatchStage
}

func (s *SetupStage) GetDependencies(_ Conveyor, _ image.Image) (string, error) {
	stageDependenciesChecksum, err := s.getStageDependenciesChecksum(Setup)
	if err != nil {
		return "", err
	}

	return util.Sha256Hash(s.builder.SetupChecksum(), stageDependenciesChecksum), nil
}

func (s *SetupStage) PrepareImage(c Conveyor, prevBuiltImage, image image.Image) error {
	if err := s.UserWithGAPatchStage.PrepareImage(c, prevBuiltImage, image); err != nil {
		return nil
	}

	if err := s.builder.Setup(image.BuilderContainer()); err != nil {
		return err
	}

	return nil
}
