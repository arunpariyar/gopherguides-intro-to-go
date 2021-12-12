package cli

type IOCommander interface { 
	Commander
	SetIO(oi IO)
}