import * as cdk from '@aws-cdk/core';
import * as lambda from '@aws-cdk/aws-lambda';
import * as apigateway from "@aws-cdk/aws-apigateway";
import assets = require("@aws-cdk/aws-s3-assets")
import path = require("path")

export class Page2PodStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // Create the Page2PodFunction lambda
    const funcAsset = new assets.Asset(this, 'Page2PodZip', {
      path: path.join(__dirname, '../../page2pod.zip'),
    });
    const handler = new lambda.Function(this, "Page2PodFunction", {
      runtime: lambda.Runtime.GO_1_X,
      handler: "page2pod.linux",
      code: lambda.Code.fromBucket(
        funcAsset.bucket,
        funcAsset.s3ObjectKey
      ),
      environment: {
      }
    });

    // Set up the api gateway for the Page2PodFunction
    const api = new apigateway.RestApi(this, "Page2PodAPI", {
      restApiName: "Page2Pod Service",
      description: "This service creates podcast episodes."
    });
    const postIntegration = new apigateway.LambdaIntegration(handler, {
      requestTemplates: { "application/json": '{ "statusCode": "200" }' }
    });
    api.root.addMethod("POST", postIntegration);

  }
}
