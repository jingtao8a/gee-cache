package main

import pb "org/jingtao8a/gee-cache/geecachepb"

type PeerGetter interface { // HTTPGetter implements this interface
	Get(in *pb.Request, out *pb.Response) error
}

type PeerPicker interface { // HTTPPool implements this interface
	PickPeer(key string) (PeerGetter, bool)
}
