package system

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ChangeLogLevelRequest Temporarily changes level of logs. New level is not persistent and is lost when restarting server. Requires system administration permission.
type ChangeLogLevelRequest struct {
	Level string `form:"level,omitempty"` // The new level. Be cautious: DEBUG, and even more TRACE, may have performance impacts.
}

// DbMigrationStatusRequest Display the database migration status of SonarQube.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul>
// Deprecated: this action has been deprecated since version 10.6
type DbMigrationStatusRequest struct{}

// DbMigrationStatusResponse is the response for DbMigrationStatusRequest
type DbMigrationStatusResponse struct {
	Message   string `json:"message,omitempty"`
	StartedAt string `json:"startedAt,omitempty"`
	State     string `json:"state,omitempty"`
}

// HealthRequest Provide health status of SonarQube.<p>Although global health is calculated based on both application and search nodes, detailed information is returned only for application nodes.</p><p>  <ul> <li>GREEN: SonarQube is fully operational</li> <li>YELLOW: SonarQube is usable, but it needs attention in order to be fully operational</li> <li>RED: SonarQube is not operational</li> </ul></p><br>Requires the 'Administer System' permission or system passcode (see WEB_SYSTEM_PASS_CODE in sonar.properties).<br>When SonarQube is in safe mode (waiting or running a database upgrade), only the authentication with a system passcode is supported.
type HealthRequest struct{}

// HealthResponse is the response for HealthRequest
type HealthResponse struct {
	Causes []struct {
		Message string `json:"message,omitempty"`
	} `json:"causes,omitempty"`
	Health string `json:"health,omitempty"`
	Nodes  []struct {
		Causes []struct {
			Message string `json:"message,omitempty"`
		} `json:"causes,omitempty"`
		Health    string  `json:"health,omitempty"`
		Host      string  `json:"host,omitempty"`
		Name      string  `json:"name,omitempty"`
		Port      float64 `json:"port,omitempty"`
		StartedAt string  `json:"startedAt,omitempty"`
		Type      string  `json:"type,omitempty"`
	} `json:"nodes,omitempty"`
}

// InfoRequest Get detailed information about system configuration.<br/>Requires 'Administer' permissions.
type InfoRequest struct{}

// InfoResponse is the response for InfoRequest
type InfoResponse struct {
	ALMs struct {
		GithubConfig string `json:"Github Config,omitempty"`
		GitlabConfig string `json:"Gitlab Config,omitempty"`
	} `json:"ALMs,omitempty"`
	Bundled struct {
		Config     string `json:"config,omitempty"`
		Csharp     string `json:"csharp,omitempty"`
		Flex       string `json:"flex,omitempty"`
		Go         string `json:"go,omitempty"`
		Iac        string `json:"iac,omitempty"`
		Jacoco     string `json:"jacoco,omitempty"`
		Java       string `json:"java,omitempty"`
		Javascript string `json:"javascript,omitempty"`
		Kotlin     string `json:"kotlin,omitempty"`
		Php        string `json:"php,omitempty"`
		Python     string `json:"python,omitempty"`
		Ruby       string `json:"ruby,omitempty"`
		Sonarscala string `json:"sonarscala,omitempty"`
		Text       string `json:"text,omitempty"`
		Vbnet      string `json:"vbnet,omitempty"`
		Web        string `json:"web,omitempty"`
		Xml        string `json:"xml,omitempty"`
	} `json:"Bundled,omitempty"`
	ComputeEngineDatabaseConnection struct {
		PoolActiveConnections  float64 `json:"Pool Active Connections,omitempty"`
		PoolIdleConnections    float64 `json:"Pool Idle Connections,omitempty"`
		PoolMaxConnections     float64 `json:"Pool Max Connections,omitempty"`
		PoolMaxLifetimems      float64 `json:"Pool Max Lifetime (ms),omitempty"`
		PoolMaxWaitms          float64 `json:"Pool Max Wait (ms),omitempty"`
		PoolMinIdleConnections float64 `json:"Pool Min Idle Connections,omitempty"`
		PoolTotalConnections   float64 `json:"Pool Total Connections,omitempty"`
	} `json:"Compute Engine Database Connection,omitempty"`
	ComputeEngineJVMProperties struct {
		AwtToolkit                 string `json:"awt.toolkit,omitempty"`
		ComRedhatFips              string `json:"com.redhat.fips,omitempty"`
		ComZaxxerHikariPoolNumber  string `json:"com.zaxxer.hikari.pool_number,omitempty"`
		FileEncoding               string `json:"file.encoding,omitempty"`
		FileSeparator              string `json:"file.separator,omitempty"`
		FtpNonProxyHosts           string `json:"ftp.nonProxyHosts,omitempty"`
		GopherProxySet             string `json:"gopherProxySet,omitempty"`
		HttpNonProxyHosts          string `json:"http.nonProxyHosts,omitempty"`
		JavaAwtGraphicsenv         string `json:"java.awt.graphicsenv,omitempty"`
		JavaAwtHeadless            string `json:"java.awt.headless,omitempty"`
		JavaAwtPrinterjob          string `json:"java.awt.printerjob,omitempty"`
		JavaClassPath              string `json:"java.class.path,omitempty"`
		JavaClassVersion           string `json:"java.class.version,omitempty"`
		JavaHome                   string `json:"java.home,omitempty"`
		JavaIoTmpdir               string `json:"java.io.tmpdir,omitempty"`
		JavaLibraryPath            string `json:"java.library.path,omitempty"`
		JavaRuntimeName            string `json:"java.runtime.name,omitempty"`
		JavaRuntimeVersion         string `json:"java.runtime.version,omitempty"`
		JavaSpecificationName      string `json:"java.specification.name,omitempty"`
		JavaSpecificationVendor    string `json:"java.specification.vendor,omitempty"`
		JavaSpecificationVersion   string `json:"java.specification.version,omitempty"`
		JavaVendor                 string `json:"java.vendor,omitempty"`
		JavaVendorUrl              string `json:"java.vendor.url,omitempty"`
		JavaVendorUrlBug           string `json:"java.vendor.url.bug,omitempty"`
		JavaVendorVersion          string `json:"java.vendor.version,omitempty"`
		JavaVersion                string `json:"java.version,omitempty"`
		JavaVersionDate            string `json:"java.version.date,omitempty"`
		JavaVmCompressedOopsMode   string `json:"java.vm.compressedOopsMode,omitempty"`
		JavaVmInfo                 string `json:"java.vm.info,omitempty"`
		JavaVmName                 string `json:"java.vm.name,omitempty"`
		JavaVmSpecificationName    string `json:"java.vm.specification.name,omitempty"`
		JavaVmSpecificationVendor  string `json:"java.vm.specification.vendor,omitempty"`
		JavaVmSpecificationVersion string `json:"java.vm.specification.version,omitempty"`
		JavaVmVendor               string `json:"java.vm.vendor,omitempty"`
		JavaVmVersion              string `json:"java.vm.version,omitempty"`
		JdkDebug                   string `json:"jdk.debug,omitempty"`
		JdkVendorVersion           string `json:"jdk.vendor.version,omitempty"`
		LineSeparator              string `json:"line.separator,omitempty"`
		OsArch                     string `json:"os.arch,omitempty"`
		OsName                     string `json:"os.name,omitempty"`
		OsVersion                  string `json:"os.version,omitempty"`
		PathSeparator              string `json:"path.separator,omitempty"`
		SocksNonProxyHosts         string `json:"socksNonProxyHosts,omitempty"`
		SunArchDataModel           string `json:"sun.arch.data.model,omitempty"`
		SunBootLibraryPath         string `json:"sun.boot.library.path,omitempty"`
		SunCpuEndian               string `json:"sun.cpu.endian,omitempty"`
		SunCpuIsalist              string `json:"sun.cpu.isalist,omitempty"`
		SunIoUnicodeEncoding       string `json:"sun.io.unicode.encoding,omitempty"`
		SunJavaCommand             string `json:"sun.java.command,omitempty"`
		SunJavaLauncher            string `json:"sun.java.launcher,omitempty"`
		SunJnuEncoding             string `json:"sun.jnu.encoding,omitempty"`
		SunManagementCompiler      string `json:"sun.management.compiler,omitempty"`
		SunOsPatchLevel            string `json:"sun.os.patch.level,omitempty"`
		UserCountry                string `json:"user.country,omitempty"`
		UserDir                    string `json:"user.dir,omitempty"`
		UserHome                   string `json:"user.home,omitempty"`
		UserLanguage               string `json:"user.language,omitempty"`
		UserName                   string `json:"user.name,omitempty"`
		UserTimezone               string `json:"user.timezone,omitempty"`
	} `json:"Compute Engine JVM Properties,omitempty"`
	ComputeEngineJVMState struct {
		FreeMemoryMB       float64 `json:"Free Memory (MB),omitempty"`
		HeapCommittedMB    float64 `json:"Heap Committed (MB),omitempty"`
		HeapInitMB         float64 `json:"Heap Init (MB),omitempty"`
		HeapMaxMB          float64 `json:"Heap Max (MB),omitempty"`
		HeapUsedMB         float64 `json:"Heap Used (MB),omitempty"`
		MaxMemoryMB        float64 `json:"Max Memory (MB),omitempty"`
		NonHeapCommittedMB float64 `json:"Non Heap Committed (MB),omitempty"`
		NonHeapInitMB      float64 `json:"Non Heap Init (MB),omitempty"`
		NonHeapUsedMB      float64 `json:"Non Heap Used (MB),omitempty"`
		SystemLoadAverage  string  `json:"System Load Average,omitempty"`
		Threads            float64 `json:"Threads,omitempty"`
	} `json:"Compute Engine JVM State,omitempty"`
	ComputeEngineLogging struct {
		LogsDir   string `json:"Logs Dir,omitempty"`
		LogsLevel string `json:"Logs Level,omitempty"`
	} `json:"Compute Engine Logging,omitempty"`
	ComputeEngineTasks struct {
		InProgress           float64 `json:"In Progress,omitempty"`
		LongestTimePendingms float64 `json:"Longest Time Pending (ms),omitempty"`
		MaxWorkerCount       float64 `json:"Max Worker Count,omitempty"`
		Pending              float64 `json:"Pending,omitempty"`
		ProcessedWithError   float64 `json:"Processed With Error,omitempty"`
		ProcessedWithSuccess float64 `json:"Processed With Success,omitempty"`
		ProcessingTimems     float64 `json:"Processing Time (ms),omitempty"`
		WorkerCount          float64 `json:"Worker Count,omitempty"`
		WorkersPaused        bool    `json:"Workers Paused,omitempty"`
	} `json:"Compute Engine Tasks,omitempty"`
	Database struct {
		Database                    string `json:"Database,omitempty"`
		DatabaseVersion             string `json:"Database Version,omitempty"`
		DefaultTransactionIsolation string `json:"Default transaction isolation,omitempty"`
		Driver                      string `json:"Driver,omitempty"`
		DriverVersion               string `json:"Driver Version,omitempty"`
		URL                         string `json:"URL,omitempty"`
		Username                    string `json:"Username,omitempty"`
	} `json:"Database,omitempty"`
	Health        string   `json:"Health,omitempty"`
	HealthCauses  []string `json:"Health Causes,omitempty"`
	Plugins       struct{} `json:"Plugins,omitempty"`
	SearchIndexes struct {
		IndexComponentsDocs           float64 `json:"Index components - Docs,omitempty"`
		IndexComponentsShards         float64 `json:"Index components - Shards,omitempty"`
		IndexComponentsStoreSize      string  `json:"Index components - Store Size,omitempty"`
		IndexIssuesDocs               float64 `json:"Index issues - Docs,omitempty"`
		IndexIssuesShards             float64 `json:"Index issues - Shards,omitempty"`
		IndexIssuesStoreSize          string  `json:"Index issues - Store Size,omitempty"`
		IndexMetadatasDocs            float64 `json:"Index metadatas - Docs,omitempty"`
		IndexMetadatasShards          float64 `json:"Index metadatas - Shards,omitempty"`
		IndexMetadatasStoreSize       string  `json:"Index metadatas - Store Size,omitempty"`
		IndexProjectmeasuresDocs      float64 `json:"Index projectmeasures - Docs,omitempty"`
		IndexProjectmeasuresShards    float64 `json:"Index projectmeasures - Shards,omitempty"`
		IndexProjectmeasuresStoreSize string  `json:"Index projectmeasures - Store Size,omitempty"`
		IndexRulesDocs                float64 `json:"Index rules - Docs,omitempty"`
		IndexRulesShards              float64 `json:"Index rules - Shards,omitempty"`
		IndexRulesStoreSize           string  `json:"Index rules - Store Size,omitempty"`
		IndexUsersDocs                float64 `json:"Index users - Docs,omitempty"`
		IndexUsersShards              float64 `json:"Index users - Shards,omitempty"`
		IndexUsersStoreSize           string  `json:"Index users - Store Size,omitempty"`
		IndexViewsDocs                float64 `json:"Index views - Docs,omitempty"`
		IndexViewsShards              float64 `json:"Index views - Shards,omitempty"`
		IndexViewsStoreSize           string  `json:"Index views - Store Size,omitempty"`
	} `json:"Search Indexes,omitempty"`
	SearchState struct {
		CPUUsage                          float64 `json:"CPU Usage (%),omitempty"`
		DiskAvailable                     string  `json:"Disk Available,omitempty"`
		FieldDataCircuitBreakerEstimation string  `json:"Field Data Circuit Breaker Estimation,omitempty"`
		FieldDataCircuitBreakerLimit      string  `json:"Field Data Circuit Breaker Limit,omitempty"`
		FieldDataMemory                   string  `json:"Field Data Memory,omitempty"`
		JVMHeapMax                        string  `json:"JVM Heap Max,omitempty"`
		JVMHeapUsage                      string  `json:"JVM Heap Usage,omitempty"`
		JVMHeapUsed                       string  `json:"JVM Heap Used,omitempty"`
		JVMNonHeapUsed                    string  `json:"JVM Non Heap Used,omitempty"`
		JVMThreads                        float64 `json:"JVM Threads,omitempty"`
		MaxFileDescriptors                float64 `json:"Max File Descriptors,omitempty"`
		OpenFileDescriptors               float64 `json:"Open File Descriptors,omitempty"`
		QueryCacheMemory                  string  `json:"Query Cache Memory,omitempty"`
		RequestCacheMemory                string  `json:"Request Cache Memory,omitempty"`
		RequestCircuitBreakerEstimation   string  `json:"Request Circuit Breaker Estimation,omitempty"`
		RequestCircuitBreakerLimit        string  `json:"Request Circuit Breaker Limit,omitempty"`
		State                             string  `json:"State,omitempty"`
		StoreSize                         string  `json:"Store Size,omitempty"`
		TranslogSize                      string  `json:"Translog Size,omitempty"`
	} `json:"Search State,omitempty"`
	ServerPushConnections struct {
		SonarLintConnectedClients float64 `json:"SonarLint Connected Clients,omitempty"`
	} `json:"Server Push Connections,omitempty"`
	Settings struct {
		DefaultNewCodeDefinition                               string `json:"Default New Code Definition,omitempty"`
		DevactivityStatus                                      string `json:"devactivity.status,omitempty"`
		HttpNonProxyHosts                                      string `json:"http.nonProxyHosts,omitempty"`
		ProcessGracefulStopTimeout                             string `json:"process.gracefulStopTimeout,omitempty"`
		ProcessIndex                                           string `json:"process.index,omitempty"`
		ProcessKey                                             string `json:"process.key,omitempty"`
		ProcessSharedDir                                       string `json:"process.sharedDir,omitempty"`
		ProjectsDefaultVisibility                              string `json:"projects.default.visibility,omitempty"`
		QualitygateDefault                                     string `json:"qualitygate.default,omitempty"`
		SonarAuthSamlCertificateSecured                        string `json:"sonar.auth.saml.certificate.secured,omitempty"`
		SonarAuthSamlEnabled                                   string `json:"sonar.auth.saml.enabled,omitempty"`
		SonarAuthSamlLoginUrl                                  string `json:"sonar.auth.saml.loginUrl,omitempty"`
		SonarAuthSamlProviderId                                string `json:"sonar.auth.saml.providerId,omitempty"`
		SonarAuthSamlProviderName                              string `json:"sonar.auth.saml.providerName,omitempty"`
		SonarAuthSamlSignatureEnabled                          string `json:"sonar.auth.saml.signature.enabled,omitempty"`
		SonarAuthSamlSpCertificateSecured                      string `json:"sonar.auth.saml.sp.certificate.secured,omitempty"`
		SonarAuthSamlSpPrivateKeySecured                       string `json:"sonar.auth.saml.sp.privateKey.secured,omitempty"`
		SonarAuthSamlUserLogin                                 string `json:"sonar.auth.saml.user.login,omitempty"`
		SonarAuthSamlUserName                                  string `json:"sonar.auth.saml.user.name,omitempty"`
		SonarAuthenticatorIgnoreStartupFailure                 string `json:"sonar.authenticator.ignoreStartupFailure,omitempty"`
		SonarAutoDatabaseUpgrade                               string `json:"sonar.autoDatabaseUpgrade,omitempty"`
		SonarBuildbreakerSkip                                  string `json:"sonar.buildbreaker.skip,omitempty"`
		SonarCPredefinedMacros                                 string `json:"sonar.c.predefinedMacros,omitempty"`
		SonarCeGracefulStopTimeOutInMs                         string `json:"sonar.ce.gracefulStopTimeOutInMs,omitempty"`
		SonarCeJavaAdditionalOpts                              string `json:"sonar.ce.javaAdditionalOpts,omitempty"`
		SonarCeJavaOpts                                        string `json:"sonar.ce.javaOpts,omitempty"`
		SonarClusterEnabled                                    string `json:"sonar.cluster.enabled,omitempty"`
		SonarClusterKubernetes                                 string `json:"sonar.cluster.kubernetes,omitempty"`
		SonarClusterName                                       string `json:"sonar.cluster.name,omitempty"`
		SonarClusterNodeName                                   string `json:"sonar.cluster.node.name,omitempty"`
		SonarClusterNodePort                                   string `json:"sonar.cluster.node.port,omitempty"`
		SonarClusterWebStartupLeader                           string `json:"sonar.cluster.web.startupLeader,omitempty"`
		SonarCoreId                                            string `json:"sonar.core.id,omitempty"`
		SonarCoreServerBaseURL                                 string `json:"sonar.core.serverBaseURL,omitempty"`
		SonarCoreStartTime                                     string `json:"sonar.core.startTime,omitempty"`
		SonarCoreTreemapColormetric                            string `json:"sonar.core.treemap.colormetric,omitempty"`
		SonarCoreTreemapSizemetric                             string `json:"sonar.core.treemap.sizemetric,omitempty"`
		SonarCpdCrossProject                                   string `json:"sonar.cpd.cross_project,omitempty"`
		SonarDbcleanerBranchesToKeepWhenInactive               string `json:"sonar.dbcleaner.branchesToKeepWhenInactive,omitempty"`
		SonarDbcleanerDaysBeforeDeletingInactiveBranchesAndPRs string `json:"sonar.dbcleaner.daysBeforeDeletingInactiveBranchesAndPRs,omitempty"`
		SonarDbcleanerMonthsBeforeDeletingAllSnapshots         string `json:"sonar.dbcleaner.monthsBeforeDeletingAllSnapshots,omitempty"`
		SonarDbcleanerWeeksBeforeDeletingAllSnapshots          string `json:"sonar.dbcleaner.weeksBeforeDeletingAllSnapshots,omitempty"`
		SonarDryRunCacheLastUpdate                             string `json:"sonar.dryRun.cache.lastUpdate,omitempty"`
		SonarEsPort                                            string `json:"sonar.es.port,omitempty"`
		SonarForceAuthentication                               string `json:"sonar.forceAuthentication,omitempty"`
		SonarGenericcoverageSuffixes                           string `json:"sonar.genericcoverage.suffixes,omitempty"`
		SonarGovernanceReportViewFrequency                     string `json:"sonar.governance.report.view.frequency,omitempty"`
		SonarJavaCoveragePlugin                                string `json:"sonar.java.coveragePlugin,omitempty"`
		SonarJavascriptJQueryObjectAliases                     string `json:"sonar.javascript.jQueryObjectAliases,omitempty"`
		SonarJdbcDriverPath                                    string `json:"sonar.jdbc.driverPath,omitempty"`
		SonarJdbcMaxActive                                     string `json:"sonar.jdbc.maxActive,omitempty"`
		SonarJdbcMaxWait                                       string `json:"sonar.jdbc.maxWait,omitempty"`
		SonarJdbcMinIdle                                       string `json:"sonar.jdbc.minIdle,omitempty"`
		SonarJdbcPassword                                      string `json:"sonar.jdbc.password,omitempty"`
		SonarJdbcUrl                                           string `json:"sonar.jdbc.url,omitempty"`
		SonarJdbcUsername                                      string `json:"sonar.jdbc.username,omitempty"`
		SonarLfEnableGravatar                                  string `json:"sonar.lf.enableGravatar,omitempty"`
		SonarLfLogoWidthPx                                     string `json:"sonar.lf.logoWidthPx,omitempty"`
		SonarLogJsonOutput                                     string `json:"sonar.log.jsonOutput,omitempty"`
		SonarOrganisation                                      string `json:"sonar.organisation,omitempty"`
		SonarPathData                                          string `json:"sonar.path.data,omitempty"`
		SonarPathHome                                          string `json:"sonar.path.home,omitempty"`
		SonarPathLogs                                          string `json:"sonar.path.logs,omitempty"`
		SonarPathTemp                                          string `json:"sonar.path.temp,omitempty"`
		SonarPathWeb                                           string `json:"sonar.path.web,omitempty"`
		SonarPlsqlFileSuffixes                                 string `json:"sonar.plsql.file.suffixes,omitempty"`
		SonarPluginsRiskConsent                                string `json:"sonar.plugins.risk.consent,omitempty"`
		SonarPreviewExcludePlugins                             string `json:"sonar.preview.excludePlugins,omitempty"`
		SonarReportDashboardName                               string `json:"sonar.report.dashboard.name,omitempty"`
		SonarReportFrequency                                   string `json:"sonar.report.frequency,omitempty"`
		SonarReportIgnoreSslErrors                             string `json:"sonar.report.ignoreSslErrors,omitempty"`
		SonarReportLastDate                                    string `json:"sonar.report.last_date,omitempty"`
		SonarReportLastDateDevReport                           string `json:"sonar.report.last_date.dev_report,omitempty"`
		SonarReportLastDateManagementReport                    string `json:"sonar.report.last_date.management_report,omitempty"`
		SonarReportLogin                                       string `json:"sonar.report.login,omitempty"`
		SonarReportSubject                                     string `json:"sonar.report.subject,omitempty"`
		SonarReports                                           string `json:"sonar.reports,omitempty"`
		SonarScmDisabled                                       string `json:"sonar.scm.disabled,omitempty"`
		SonarScmEnabled                                        string `json:"sonar.scm.enabled,omitempty"`
		SonarSearchHost                                        string `json:"sonar.search.host,omitempty"`
		SonarSearchJavaAdditionalOpts                          string `json:"sonar.search.javaAdditionalOpts,omitempty"`
		SonarSearchJavaOpts                                    string `json:"sonar.search.javaOpts,omitempty"`
		SonarSearchPort                                        string `json:"sonar.search.port,omitempty"`
		SonarTechnicalDebtRatingGrid                           string `json:"sonar.technicalDebt.ratingGrid,omitempty"`
		SonarTelemetryCompression                              string `json:"sonar.telemetry.compression,omitempty"`
		SonarTelemetryEnable                                   string `json:"sonar.telemetry.enable,omitempty"`
		SonarTelemetryFrequencyInSeconds                       string `json:"sonar.telemetry.frequencyInSeconds,omitempty"`
		SonarTelemetryUrl                                      string `json:"sonar.telemetry.url,omitempty"`
		SonarUpdatecenterActivate                              string `json:"sonar.updatecenter.activate,omitempty"`
		SonarWebGracefulStopTimeOutInMs                        string `json:"sonar.web.gracefulStopTimeOutInMs,omitempty"`
		SonarWebJavaAdditionalOpts                             string `json:"sonar.web.javaAdditionalOpts,omitempty"`
		SonarWebJavaOpts                                       string `json:"sonar.web.javaOpts,omitempty"`
		SonarWebSsoEmailHeader                                 string `json:"sonar.web.sso.emailHeader,omitempty"`
		SonarWebSsoEnable                                      string `json:"sonar.web.sso.enable,omitempty"`
		SonarWebSsoGroupsHeader                                string `json:"sonar.web.sso.groupsHeader,omitempty"`
		SonarWebSsoLoginHeader                                 string `json:"sonar.web.sso.loginHeader,omitempty"`
		SonarWebSsoNameHeader                                  string `json:"sonar.web.sso.nameHeader,omitempty"`
		SonarWebSsoRefreshIntervalInMinutes                    string `json:"sonar.web.sso.refreshIntervalInMinutes,omitempty"`
	} `json:"Settings,omitempty"`
	System struct {
		AcceptedExternalIdentityProviders                               string  `json:"Accepted external identity providers,omitempty"`
		DataDir                                                         string  `json:"Data Dir,omitempty"`
		Docker                                                          bool    `json:"Docker,omitempty"`
		Edition                                                         string  `json:"Edition,omitempty"`
		ExternalIdentityProvidersWhoseUsersAreAllowedToSignThemselvesUp string  `json:"External identity providers whose users are allowed to sign themselves up,omitempty"`
		ForceAuthentication                                             bool    `json:"Force authentication,omitempty"`
		HighAvailability                                                bool    `json:"High Availability,omitempty"`
		HomeDir                                                         string  `json:"Home Dir,omitempty"`
		OfficialDistribution                                            bool    `json:"Official Distribution,omitempty"`
		Processors                                                      float64 `json:"Processors,omitempty"`
		ServerID                                                        string  `json:"Server ID,omitempty"`
		TempDir                                                         string  `json:"Temp Dir,omitempty"`
		Version                                                         string  `json:"Version,omitempty"`
	} `json:"System,omitempty"`
	WebDatabaseConnection struct {
		PoolActiveConnections  float64 `json:"Pool Active Connections,omitempty"`
		PoolIdleConnections    float64 `json:"Pool Idle Connections,omitempty"`
		PoolMaxConnections     float64 `json:"Pool Max Connections,omitempty"`
		PoolMaxLifetimems      float64 `json:"Pool Max Lifetime (ms),omitempty"`
		PoolMaxWaitms          float64 `json:"Pool Max Wait (ms),omitempty"`
		PoolMinIdleConnections float64 `json:"Pool Min Idle Connections,omitempty"`
		PoolTotalConnections   float64 `json:"Pool Total Connections,omitempty"`
	} `json:"Web Database Connection,omitempty"`
	WebJVMProperties struct {
		AwtToolkit                                      string `json:"awt.toolkit,omitempty"`
		CatalinaBase                                    string `json:"catalina.base,omitempty"`
		CatalinaHome                                    string `json:"catalina.home,omitempty"`
		CatalinaUseNaming                               string `json:"catalina.useNaming,omitempty"`
		ComRedhatFips                                   string `json:"com.redhat.fips,omitempty"`
		ComZaxxerHikariPoolNumber                       string `json:"com.zaxxer.hikari.pool_number,omitempty"`
		FileEncoding                                    string `json:"file.encoding,omitempty"`
		FileSeparator                                   string `json:"file.separator,omitempty"`
		FtpNonProxyHosts                                string `json:"ftp.nonProxyHosts,omitempty"`
		GopherProxySet                                  string `json:"gopherProxySet,omitempty"`
		HttpAgent                                       string `json:"http.agent,omitempty"`
		HttpNonProxyHosts                               string `json:"http.nonProxyHosts,omitempty"`
		JavaAwtGraphicsenv                              string `json:"java.awt.graphicsenv,omitempty"`
		JavaAwtHeadless                                 string `json:"java.awt.headless,omitempty"`
		JavaAwtPrinterjob                               string `json:"java.awt.printerjob,omitempty"`
		JavaClassPath                                   string `json:"java.class.path,omitempty"`
		JavaClassVersion                                string `json:"java.class.version,omitempty"`
		JavaHome                                        string `json:"java.home,omitempty"`
		JavaIoTmpdir                                    string `json:"java.io.tmpdir,omitempty"`
		JavaLibraryPath                                 string `json:"java.library.path,omitempty"`
		JavaRuntimeName                                 string `json:"java.runtime.name,omitempty"`
		JavaRuntimeVersion                              string `json:"java.runtime.version,omitempty"`
		JavaSpecificationName                           string `json:"java.specification.name,omitempty"`
		JavaSpecificationVendor                         string `json:"java.specification.vendor,omitempty"`
		JavaSpecificationVersion                        string `json:"java.specification.version,omitempty"`
		JavaVendor                                      string `json:"java.vendor,omitempty"`
		JavaVendorUrl                                   string `json:"java.vendor.url,omitempty"`
		JavaVendorUrlBug                                string `json:"java.vendor.url.bug,omitempty"`
		JavaVendorVersion                               string `json:"java.vendor.version,omitempty"`
		JavaVersion                                     string `json:"java.version,omitempty"`
		JavaVersionDate                                 string `json:"java.version.date,omitempty"`
		JavaVmCompressedOopsMode                        string `json:"java.vm.compressedOopsMode,omitempty"`
		JavaVmInfo                                      string `json:"java.vm.info,omitempty"`
		JavaVmName                                      string `json:"java.vm.name,omitempty"`
		JavaVmSpecificationName                         string `json:"java.vm.specification.name,omitempty"`
		JavaVmSpecificationVendor                       string `json:"java.vm.specification.vendor,omitempty"`
		JavaVmSpecificationVersion                      string `json:"java.vm.specification.version,omitempty"`
		JavaVmVendor                                    string `json:"java.vm.vendor,omitempty"`
		JavaVmVersion                                   string `json:"java.vm.version,omitempty"`
		JdkDebug                                        string `json:"jdk.debug,omitempty"`
		JdkVendorVersion                                string `json:"jdk.vendor.version,omitempty"`
		LineSeparator                                   string `json:"line.separator,omitempty"`
		LogbackDisableServletContainerInitializer       string `json:"logbackDisableServletContainerInitializer,omitempty"`
		OrgApacheCatalinaStartupEXITONINITFAILURE       string `json:"org.apache.catalina.startup.EXIT_ON_INIT_FAILURE,omitempty"`
		OrgApacheTomcatUtilBufUDecoderALLOWENCODEDSLASH string `json:"org.apache.tomcat.util.buf.UDecoder.ALLOW_ENCODED_SLASH,omitempty"`
		OsArch                                          string `json:"os.arch,omitempty"`
		OsName                                          string `json:"os.name,omitempty"`
		OsVersion                                       string `json:"os.version,omitempty"`
		PathSeparator                                   string `json:"path.separator,omitempty"`
		SocksNonProxyHosts                              string `json:"socksNonProxyHosts,omitempty"`
		SunArchDataModel                                string `json:"sun.arch.data.model,omitempty"`
		SunBootLibraryPath                              string `json:"sun.boot.library.path,omitempty"`
		SunCpuEndian                                    string `json:"sun.cpu.endian,omitempty"`
		SunCpuIsalist                                   string `json:"sun.cpu.isalist,omitempty"`
		SunFontFontmanager                              string `json:"sun.font.fontmanager,omitempty"`
		SunIoUnicodeEncoding                            string `json:"sun.io.unicode.encoding,omitempty"`
		SunJavaCommand                                  string `json:"sun.java.command,omitempty"`
		SunJavaLauncher                                 string `json:"sun.java.launcher,omitempty"`
		SunJnuEncoding                                  string `json:"sun.jnu.encoding,omitempty"`
		SunManagementCompiler                           string `json:"sun.management.compiler,omitempty"`
		SunOsPatchLevel                                 string `json:"sun.os.patch.level,omitempty"`
		UserCountry                                     string `json:"user.country,omitempty"`
		UserDir                                         string `json:"user.dir,omitempty"`
		UserHome                                        string `json:"user.home,omitempty"`
		UserLanguage                                    string `json:"user.language,omitempty"`
		UserName                                        string `json:"user.name,omitempty"`
		UserTimezone                                    string `json:"user.timezone,omitempty"`
	} `json:"Web JVM Properties,omitempty"`
	WebJVMState struct {
		FreeMemoryMB       float64 `json:"Free Memory (MB),omitempty"`
		HeapCommittedMB    float64 `json:"Heap Committed (MB),omitempty"`
		HeapInitMB         float64 `json:"Heap Init (MB),omitempty"`
		HeapMaxMB          float64 `json:"Heap Max (MB),omitempty"`
		HeapUsedMB         float64 `json:"Heap Used (MB),omitempty"`
		MaxMemoryMB        float64 `json:"Max Memory (MB),omitempty"`
		NonHeapCommittedMB float64 `json:"Non Heap Committed (MB),omitempty"`
		NonHeapInitMB      float64 `json:"Non Heap Init (MB),omitempty"`
		NonHeapUsedMB      float64 `json:"Non Heap Used (MB),omitempty"`
		SystemLoadAverage  string  `json:"System Load Average,omitempty"`
		Threads            float64 `json:"Threads,omitempty"`
	} `json:"Web JVM State,omitempty"`
	WebLogging struct {
		LogsDir   string `json:"Logs Dir,omitempty"`
		LogsLevel string `json:"Logs Level,omitempty"`
	} `json:"Web Logging,omitempty"`
}

// LogsRequest Get system logs in plain-text format. Requires system administration permission.
type LogsRequest struct {
	Name string `form:"name,omitempty"` // Name of the logs to get
}

// LogsResponse is the response for LogsRequest
type LogsResponse string

// MigrateDbRequest Migrate the database to match the current version of SonarQube.<br/>Sending a POST request to this URL starts the DB migration. It is strongly advised to <strong>make a database backup</strong> before invoking this WS.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul>
type MigrateDbRequest struct{}

// MigrateDbResponse is the response for MigrateDbRequest
type MigrateDbResponse struct {
	Message   string `json:"message,omitempty"`
	StartedAt string `json:"startedAt,omitempty"`
	State     string `json:"state,omitempty"`
}

// PingRequest Answers "pong" as plain-text
type PingRequest struct{}

// PingResponse is the response for PingRequest
type PingResponse string

// RestartRequest Restarts server. Requires 'Administer System' permission. Performs a full restart of the Web, Search and Compute Engine Servers processes. Does not reload sonar.properties.
type RestartRequest struct{}

// StatusRequest Get state information about SonarQube.<p>status: the running status <ul> <li>STARTING: SonarQube Web Server is up and serving some Web Services (eg. api/system/status) but initialization is still ongoing</li> <li>UP: SonarQube instance is up and running</li> <li>DOWN: SonarQube instance is up but not running because migration has failed (refer to WS /api/system/migrate_db for details) or some other reason (check logs).</li> <li>RESTARTING: SonarQube instance is still up but a restart has been requested (refer to WS /api/system/restart for details).</li> <li>DB_MIGRATION_NEEDED: database migration is required. DB migration can be started using WS /api/system/migrate_db.</li> <li>DB_MIGRATION_RUNNING: DB migration is running (refer to WS /api/system/migrate_db for details)</li> </ul></p>
type StatusRequest struct{}

// StatusResponse is the response for StatusRequest
type StatusResponse struct {
	Id      string `json:"id,omitempty"`
	Status  string `json:"status,omitempty"`
	Version string `json:"version,omitempty"`
}

// UpgradesRequest Lists available upgrades for the SonarQube instance (if any) and for each one, lists incompatible plugins and plugins requiring upgrade.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.
type UpgradesRequest struct{}

// UpgradesResponse is the response for UpgradesRequest
type UpgradesResponse struct {
	InstalledVersionActive bool   `json:"installedVersionActive,omitempty"`
	LatestLTA              string `json:"latestLTA,omitempty"`
	LatestLTS              string `json:"latestLTS,omitempty"`
	UpdateCenterRefresh    string `json:"updateCenterRefresh,omitempty"`
	Upgrades               []struct {
		ChangeLogUrl string `json:"changeLogUrl,omitempty"`
		Description  string `json:"description,omitempty"`
		DownloadUrl  string `json:"downloadUrl,omitempty"`
		Plugins      struct {
			Incompatible []struct {
				Category         string `json:"category,omitempty"`
				Description      string `json:"description,omitempty"`
				EditionBundled   bool   `json:"editionBundled,omitempty"`
				Key              string `json:"key,omitempty"`
				License          string `json:"license,omitempty"`
				Name             string `json:"name,omitempty"`
				OrganizationName string `json:"organizationName,omitempty"`
				OrganizationUrl  string `json:"organizationUrl,omitempty"`
			} `json:"incompatible,omitempty"`
			RequireUpdate []struct {
				Category              string `json:"category,omitempty"`
				Description           string `json:"description,omitempty"`
				EditionBundled        bool   `json:"editionBundled,omitempty"`
				Key                   string `json:"key,omitempty"`
				License               string `json:"license,omitempty"`
				Name                  string `json:"name,omitempty"`
				OrganizationName      string `json:"organizationName,omitempty"`
				OrganizationUrl       string `json:"organizationUrl,omitempty"`
				TermsAndConditionsUrl string `json:"termsAndConditionsUrl,omitempty"`
				Version               string `json:"version,omitempty"`
			} `json:"requireUpdate,omitempty"`
		} `json:"plugins,omitempty"`
		ReleaseDate string `json:"releaseDate,omitempty"`
		Version     string `json:"version,omitempty"`
	} `json:"upgrades,omitempty"`
}
