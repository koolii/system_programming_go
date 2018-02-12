listener, err := net.Listen("unix", "socketfile")
if err != nil {
	panic(err)
}
defer listener.Close()

conn, err := listener.Accept()
if err != nil {
	// error handling
}