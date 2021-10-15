# page2pod
Turn a page into a hosted podcast.

# Setting up a development environment
  1. Install node
  2. Install cdk
    > npm install -g aws-cdk
  3. Install dependencies
    > (cd infra && npm install)
  4. Bootstrap cdk
    > (cd infra && cdk bootstrap aws://ACCOUNT-NUMBER/REGION)
  5. Install mockgen
    > go install github.com/golang/mock/mockgen@v1.6.0
  6. Do the initial deploy
    > ./deploy
  7. Find the value of the Page2PodAPIEndpoint URL in the Page2PodStack and set it as an environment variable
    > export PAGE2POD_API_ENDPOINT=https://<blaa blaa blaa>.execute-api.us-west-2.amazonaws.com/prod/
  8. Redeploy page2pod to update the admin page to point to the Page2PodAPIEndpoint URL
    > ./deploy



# Deploying page2pod
  1. Run the deploy script
    > ./deploy
