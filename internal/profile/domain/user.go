package domain

type User struct {
	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username"`

	Following []User `protobuf:"bytes,4,rep,name=following,proto3" json:"following"`
	Followers []User `protobuf:"bytes,5,rep,name=followers,proto3" json:"followers"`
}
