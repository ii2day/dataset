package conda

// Output of `conda env list --json` command
type CondaEnvListOutputRaw struct {
	Envs []string `json:"envs"`
}

type CondaEnvListOutput struct {
	Envs []CondaEnvListOutputEnv `json:"envs"`
}

type CondaEnvListOutputEnv struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// Output of `conda info --json` command
type CondaInfoOutputRaw struct {
	GID               int            `json:"GID"`
	UID               int            `json:"UID"`
	ActivePrefix      string         `json:"active_prefix"`
	ActivePrefixName  string         `json:"active_prefix_name"`
	AvDataDir         string         `json:"av_data_dir"`
	AvMetadataURLBase interface{}    `json:"av_metadata_url_base"`
	Channels          []string       `json:"channels"`
	CondaBuildVersion string         `json:"conda_build_version"`
	CondaEnvVersion   string         `json:"conda_env_version"`
	CondaLocation     string         `json:"conda_location"`
	CondaPrefix       string         `json:"conda_prefix"`
	CondaShlvl        int            `json:"conda_shlvl"`
	CondaVersion      string         `json:"conda_version"`
	ConfigFiles       []string       `json:"config_files"`
	DefaultPrefix     string         `json:"default_prefix"`
	EnvVars           map[string]any `json:"env_vars"`
	Envs              []string       `json:"envs"`
	EnvsDirs          []string       `json:"envs_dirs"`
	NetrcFile         interface{}    `json:"netrc_file"`
	Offline           bool           `json:"offline"`
	PkgsDirs          []string       `json:"pkgs_dirs"`
	Platform          string         `json:"platform"`
	PythonVersion     string         `json:"python_version"`
	RcPath            string         `json:"rc_path"`
	RequestsVersion   string         `json:"requests_version"`
	RootPrefix        string         `json:"root_prefix"`
	RootWritable      bool           `json:"root_writable"`
	SiteDirs          []interface{}  `json:"site_dirs"`
	Solver            struct {
		Default   bool   `json:"default"`
		Name      string `json:"name"`
		UserAgent string `json:"user_agent"`
	} `json:"solver"`
	SysExecutable string     `json:"sys.executable"`
	SysPrefix     string     `json:"sys.prefix"`
	SysVersion    string     `json:"sys.version"`
	SysRcPath     string     `json:"sys_rc_path"`
	UserAgent     string     `json:"user_agent"`
	UserRcPath    string     `json:"user_rc_path"`
	VirtualPkgs   [][]string `json:"virtual_pkgs"`
}

// Output of `conda config --show --json` command
type CondaConfigShowOutputRaw struct {
	AddAnacondaToken         bool          `json:"add_anaconda_token"`
	AddPipAsPythonDependency bool          `json:"add_pip_as_python_dependency"`
	AggressiveUpdatePackages []string      `json:"aggressive_update_packages"`
	AllowCondaDowngrades     bool          `json:"allow_conda_downgrades"`
	AllowCycles              bool          `json:"allow_cycles"`
	AllowNonChannelUrls      bool          `json:"allow_non_channel_urls"`
	AllowSoftlinks           bool          `json:"allow_softlinks"`
	AllowlistChannels        []interface{} `json:"allowlist_channels"`
	AlwaysCopy               bool          `json:"always_copy"`
	AlwaysSoftlink           bool          `json:"always_softlink"`
	AlwaysYes                interface{}   `json:"always_yes"`
	AnacondaUpload           interface{}   `json:"anaconda_upload"`
	AutoActivateBase         bool          `json:"auto_activate_base"`
	AutoStack                int           `json:"auto_stack"`
	AutoUpdateConda          bool          `json:"auto_update_conda"`
	BldPath                  string        `json:"bld_path"`
	Changeps1                bool          `json:"changeps1"`
	ChannelAlias             struct {
		Auth            interface{} `json:"auth"`
		Location        string      `json:"location"`
		Name            string      `json:"name"`
		PackageFilename interface{} `json:"package_filename"`
		Platform        interface{} `json:"platform"`
		Scheme          string      `json:"scheme"`
		Token           interface{} `json:"token"`
	} `json:"channel_alias"`
	ChannelPriority  string        `json:"channel_priority"`
	ChannelSettings  []interface{} `json:"channel_settings"`
	Channels         []string      `json:"channels"`
	ClientSslCert    interface{}   `json:"client_ssl_cert"`
	ClientSslCertKey interface{}   `json:"client_ssl_cert_key"`
	Clobber          bool          `json:"clobber"`
	CondaBuild       struct {
	} `json:"conda_build"`
	CreateDefaultPackages []interface{} `json:"create_default_packages"`
	Croot                 string        `json:"croot"`
	CustomChannels        struct {
		PkgsMain struct {
			Auth            interface{} `json:"auth"`
			Location        string      `json:"location"`
			Name            string      `json:"name"`
			PackageFilename interface{} `json:"package_filename"`
			Platform        interface{} `json:"platform"`
			Scheme          string      `json:"scheme"`
			Token           interface{} `json:"token"`
		} `json:"pkgs/main"`
		PkgsPro struct {
			Auth            interface{} `json:"auth"`
			Location        string      `json:"location"`
			Name            string      `json:"name"`
			PackageFilename interface{} `json:"package_filename"`
			Platform        interface{} `json:"platform"`
			Scheme          string      `json:"scheme"`
			Token           interface{} `json:"token"`
		} `json:"pkgs/pro"`
		PkgsR struct {
			Auth            interface{} `json:"auth"`
			Location        string      `json:"location"`
			Name            string      `json:"name"`
			PackageFilename interface{} `json:"package_filename"`
			Platform        interface{} `json:"platform"`
			Scheme          string      `json:"scheme"`
			Token           interface{} `json:"token"`
		} `json:"pkgs/r"`
	} `json:"custom_channels"`
	CustomMultichannels struct {
		Defaults []struct {
			Auth            interface{} `json:"auth"`
			Location        string      `json:"location"`
			Name            string      `json:"name"`
			PackageFilename interface{} `json:"package_filename"`
			Platform        interface{} `json:"platform"`
			Scheme          string      `json:"scheme"`
			Token           interface{} `json:"token"`
		} `json:"defaults"`
		Local []interface{} `json:"local"`
	} `json:"custom_multichannels"`
	Debug           bool `json:"debug"`
	DefaultChannels []struct {
		Auth            interface{} `json:"auth"`
		Location        string      `json:"location"`
		Name            string      `json:"name"`
		PackageFilename interface{} `json:"package_filename"`
		Platform        interface{} `json:"platform"`
		Scheme          string      `json:"scheme"`
		Token           interface{} `json:"token"`
	} `json:"default_channels"`
	DefaultPython          string        `json:"default_python"`
	DefaultThreads         interface{}   `json:"default_threads"`
	DepsModifier           string        `json:"deps_modifier"`
	Dev                    bool          `json:"dev"`
	DisallowedPackages     []interface{} `json:"disallowed_packages"`
	DownloadOnly           bool          `json:"download_only"`
	DryRun                 bool          `json:"dry_run"`
	EnablePrivateEnvs      bool          `json:"enable_private_envs"`
	EnvPrompt              string        `json:"env_prompt"`
	EnvsDirs               []string      `json:"envs_dirs"`
	ErrorUploadURL         string        `json:"error_upload_url"`
	ExecuteThreads         int           `json:"execute_threads"`
	Experimental           []interface{} `json:"experimental"`
	ExtraSafetyChecks      bool          `json:"extra_safety_checks"`
	FetchThreads           int           `json:"fetch_threads"`
	Force                  bool          `json:"force"`
	Force32Bit             bool          `json:"force_32bit"`
	ForceReinstall         bool          `json:"force_reinstall"`
	ForceRemove            bool          `json:"force_remove"`
	IgnorePinned           bool          `json:"ignore_pinned"`
	JSON                   bool          `json:"json"`
	LocalRepodataTTL       int           `json:"local_repodata_ttl"`
	MigratedChannelAliases []interface{} `json:"migrated_channel_aliases"`
	MigratedCustomChannels struct {
	} `json:"migrated_custom_channels"`
	NoLock                  bool          `json:"no_lock"`
	NoPlugins               bool          `json:"no_plugins"`
	NonAdminEnabled         bool          `json:"non_admin_enabled"`
	NotifyOutdatedConda     bool          `json:"notify_outdated_conda"`
	NumberChannelNotices    int           `json:"number_channel_notices"`
	Offline                 bool          `json:"offline"`
	OverrideChannelsEnabled bool          `json:"override_channels_enabled"`
	PathConflict            string        `json:"path_conflict"`
	PinnedPackages          []interface{} `json:"pinned_packages"`
	PipInteropEnabled       bool          `json:"pip_interop_enabled"`
	PkgsDirs                []string      `json:"pkgs_dirs"`
	ProxyServers            struct {
	} `json:"proxy_servers"`
	Quiet                        bool          `json:"quiet"`
	RegisterEnvs                 bool          `json:"register_envs"`
	RemoteBackoffFactor          int           `json:"remote_backoff_factor"`
	RemoteConnectTimeoutSecs     float64       `json:"remote_connect_timeout_secs"`
	RemoteMaxRetries             int           `json:"remote_max_retries"`
	RemoteReadTimeoutSecs        float64       `json:"remote_read_timeout_secs"`
	RepodataFns                  []string      `json:"repodata_fns"`
	RepodataThreads              interface{}   `json:"repodata_threads"`
	ReportErrors                 interface{}   `json:"report_errors"`
	RestoreFreeChannel           bool          `json:"restore_free_channel"`
	RollbackEnabled              bool          `json:"rollback_enabled"`
	RootPrefix                   string        `json:"root_prefix"`
	SafetyChecks                 string        `json:"safety_checks"`
	SatSolver                    string        `json:"sat_solver"`
	SeparateFormatCache          bool          `json:"separate_format_cache"`
	Shortcuts                    bool          `json:"shortcuts"`
	ShortcutsOnly                []interface{} `json:"shortcuts_only"`
	ShowChannelUrls              interface{}   `json:"show_channel_urls"`
	SigningMetadataURLBase       interface{}   `json:"signing_metadata_url_base"`
	Solver                       string        `json:"solver"`
	SolverIgnoreTimestamps       bool          `json:"solver_ignore_timestamps"`
	SslVerify                    bool          `json:"ssl_verify"`
	Subdir                       string        `json:"subdir"`
	Subdirs                      []string      `json:"subdirs"`
	TargetPrefixOverride         string        `json:"target_prefix_override"`
	Trace                        bool          `json:"trace"`
	TrackFeatures                []interface{} `json:"track_features"`
	UnsatisfiableHints           bool          `json:"unsatisfiable_hints"`
	UnsatisfiableHintsCheckDepth int           `json:"unsatisfiable_hints_check_depth"`
	UpdateModifier               string        `json:"update_modifier"`
	UseIndexCache                bool          `json:"use_index_cache"`
	UseLocal                     bool          `json:"use_local"`
	UseOnlyTarBz2                bool          `json:"use_only_tar_bz2"`
	Verbosity                    int           `json:"verbosity"`
	VerifyThreads                int           `json:"verify_threads"`
}
