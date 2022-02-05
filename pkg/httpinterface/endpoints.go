package httpinterface

// infrastructure related
const (
	// endPointGroupInfrastructure concentrates routes related to infrastructure.
	endPointGroupInfrastructure = "/infra"
	// endpointIsReady is endpoint to get if app ready to operate.
	endpointIsReady = "/ready"
	endpointVersion = "/version"
)

// authorization related
const (
	// endPointGroupAuthorization concentrates routes for authorization scenarios.
	endPointGroupAuthorization = "/auth"
	// endPointLogin is endpoint for login.
	endPointLogin      = "/login"
	endPointCreateUser = "/createuser"
)
