package main

type PeerGetter interface { // HTTPGetter implements this interface
	Get(group string, key string) ([]byte, error)
}

type PeerPicker interface { // HTTPPool implements this interface
	PickPeer(key string) (PeerGetter, bool)
}
