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


# Deploying page2pod
  1. Run the deploy script
    > ./deploy
