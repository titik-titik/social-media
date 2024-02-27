package route

type RootRoute struct {
	UserRoute *UserRoute
}

func NewRootRoute(
	userRoute *UserRoute,
) *RootRoute {
	rootRoute := &RootRoute{
		UserRoute: userRoute,
	}
	return rootRoute
}

func (rootRoute *RootRoute) Register() {
	rootRoute.UserRoute.Register()
}
