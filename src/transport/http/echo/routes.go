package echo

func (s *httpServer) setRoutes() {
	//auth routes
	{
		s.auth.POST("/exists-email", s.handler.ExistsEmail)
		s.auth.POST("/exists-phone", s.handler.ExistsPhone)
		s.auth.POST("/register-by-phone", s.handler.RegisterByPhone)
		s.auth.POST("/resend-code", s.handler.ReSendCode)
		s.auth.POST("/forgot-password", s.handler.ForgotPassword)
		s.auth.POST("/change-password", s.handler.ChangePassword)
		s.auth.POST("/login", s.handler.Login)
		s.auth.POST("/verify-login", s.handler.VerifyLogin)
		s.auth.POST("/verify-register", s.handler.VerifyRegister)
	}
	//account routes
	{
		s.account.GET("/me", s.handler.AccountProfile)
		s.account.PUT("/edit-profile", s.handler.EditAccount)
		s.account.PUT("/set-email", s.handler.SetEmail)
		s.account.PUT("/set-phone", s.handler.SetPhone)
		s.account.POST("/verify-email", s.handler.VerifyEmail)
		s.account.POST("/verify-phone", s.handler.VerifyPhone)
		s.account.POST("/resend-code", s.handler.AccountReSendCode)
		s.account.PUT("/change-password", s.handler.ChangeAccountPassword)
	}
	//session routes
	{
		s.session.POST("/list", s.handler.SessionList)
		s.session.DELETE("/delete", s.handler.DeleteSession)
	}
}
