package web

// API is the primary point for REST API registration
type API struct {
	// Controllers contains the registered controllers
	Controllers []Controller
}
