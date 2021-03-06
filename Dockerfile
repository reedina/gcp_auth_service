FROM golang

# Copy the local package files to the container's workspace
ADD . /go/src/github.com/reedina/gcp_auth_service

# Add Environment variables  -- ADDED in Deployment YAML file
#ENV AWS_DB_USERNAME mikerapuano
#ENV AWS_DB_PASSWORD sd4msd5m!2005
#ENV AWS_DB_NAME rapuano
#ENV AWS_DB_URL mikerap01.cc92ps1k0iaz.us-east-1.rds.amazonaws.com

# Add Golang packages
RUN go get github.com/go-sql-driver/mysql && go get github.com/gorilla/mux && go get github.com/rs/cors

# Build the spm command inside the container
RUN go install github.com/reedina/gcp_auth_service

# Run the spm command by default when the container starts
ENTRYPOINT /go/bin/gcp_auth_service

# Service listens on port 4040
EXPOSE 5040
