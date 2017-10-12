## GRPC Playground

This is my first attempt at writing a service in GRPC.

I have a [gist here](https://gist.github.com/cecyc/6e1215eed1f3a59ab86d3ebf84ca6b46) with a ton of resources on getting started, because I found it a bit challenging to find the information I needed to learn this.

Hopefully this helps someone!

Structure:

```
proto-fun
├── client
|	└── main.go
├── doggos
|	└── proto
|		└── doggos.pb.go
├── server
|	└── main.go
└── vendor
	└── googleapis
```
The client folder contains the client that makes requests to the server.

The server folder contains the server that handles requests.

The proto folder contains the protocol buffer definition for the service.

The doggos folder contains the generated golang library that client/main.go and server/main.go use.

### Lessons learned 

1. Client and server can be different languages

One thing that stumped me is that the client and the server can actually be written in different languages.

You can generate libraries in languages like Ruby, Python, PHP (though you can't write a server in PHP yet), C#, and Javascript.

So technically, I could have made the client in a differen language than the server. They don't have to be both the same language, and I found that a little confusing when I started.

2. The server must implement all methods

This is kind of a dumb one, and a gimmie, but I didn't realize at first my server would need to implement all the server methods described in the server interface.

```
type DoggosServer interface {
	CreateNewDoggo(context.Context, *NewDoggoRequest) (*Doggo, error)
	GetDoggo(context.Context, *GetDoggoRequest) (*Doggo, error)
	DeleteDoggo(context.Context, *DeleteDoggoRequest) (*DeleteDoggoResponse, error)
	ListDoggos(context.Context, *ListDoggosRequest) (*ListDoggosResponse, error)
}
```

I typically start of my coding by doing one thing, making sure it works, then moving on to the next thing. So when I started coding and testing my server, I had only implemented CreateNewDoggo, so I was getting errors that the other methods weren't implemented.

What I did to get around that was simply implement the methods so I could satisfy the interface, but simply return `nil, nil`, and then build the functionality later as I tested.

3. Vendoring

I had to vendor some stuff from google in order to have access to the protocol buffers annotations (which is optional). There may be a better way to do this.