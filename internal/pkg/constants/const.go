package constants

const (
	// DatasetJobSecretsMountPath is the path to the directory where the
	// dataset job secrets are mounted.
	DatasetJobSecretsMountPath                  string = "/run/dataset/secrets" // #nosec G101
	DatasetJobCondaConfigDir                    string = "/run/dataset/conda"
	DatasetJobCondaCondaEnvironmentYAMLFilename string = "environment.yaml"
	DatasetJobCondaPipRequirementsTxtFilename   string = "requirements.txt"
	DatasetJobCondaCondaEnvironmentYAMLPath     string = DatasetJobCondaConfigDir + "/" + DatasetJobCondaCondaEnvironmentYAMLFilename
	DatasetJobCondaPipRequirementsTxtPath       string = DatasetJobCondaConfigDir + "/" + DatasetJobCondaPipRequirementsTxtFilename

	DatasetJobCondaMountDir = "/opt/baize-runtime-env"

	HamiVGPUTypeAnnotationName = "nvidia.com/use-gputype"
)

const (
	// The default baize-base env path for the conda env,
	// used for tensorboard, etc.
	//
	// Currently analysis component is using this path.
	CondaEnvBaizeBase    string = "/opt/conda/envs/baize-base"
	CondaEnvBaizeBaseBin string = CondaEnvBaizeBase + "/bin"

	DatasetNameLabel = "baize.io/dataset-name"
)
