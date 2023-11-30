# BlogPost

Simple Server that has Rest API and GRPC server. PostgreSQL is used as a database.

In the web app posts can be:
- Created
- Deleted
- Seen

posts.

The proto file is located in `internal/grpc/protos`

### API End points
- URL:`/api/post/new`
- Method: `POST`
- Request Body: ```{"title": "", "content": "", "created_at": "", "username":""}```


- URL:`/api/post/all`
- Method: `GET`


- URL:`/api/post/{id}`
- Method: `DELETE`