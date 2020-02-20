package main

//create routes
func (s *Server) routes() {
	s.router.HandleFunc("/forgotpassword", s.handleforgotpassword()).Methods("POST")

}
