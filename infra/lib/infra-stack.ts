import * as cdk from '@aws-cdk/core';
import * as s3 from '@aws-cdk/aws-s3';
import * as cloudfront from '@aws-cdk/aws-cloudfront';
import * as origins from '@aws-cdk/aws-cloudfront-origins';
import * as lambda from '@aws-cdk/aws-lambda';
import * as apigateway from "@aws-cdk/aws-apigateway";
import * as s3deployment from "@aws-cdk/aws-s3-deployment";
import assets = require("@aws-cdk/aws-s3-assets")
import path = require("path")

export class Page2PodStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    /*
     * Set up a site that will host page for adding new episodes to the podcast
     * It will also host the podcasts index and audio files
     */
    // Create the S3 bucket
    const podcastBucket = new s3.Bucket(this, 'Page2PodBucket', {
      versioned: true,
    });
    // Serve the contents over HTTPS
    const podcastDistribution = new cloudfront.Distribution(this, 'Page2PodDist', {
        defaultBehavior: { origin: new origins.S3Origin(podcastBucket) },
    });


    /*
     * Set up the lambda/API that will take the text/url and add an episode to a podcast.
     */
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
    podcastBucket.grantReadWrite(handler);

    // Set up the api gateway for the Page2PodFunction
    const api = new apigateway.RestApi(this, "Page2PodAPI", {
      restApiName: "Page2Pod Service",
      description: "This service creates podcast episodes."
    });
    const postIntegration = new apigateway.LambdaIntegration(handler, {
      requestTemplates: { "application/json": '{ "statusCode": "200" }' }
    });
    api.root.addMethod("POST", postIntegration);

    // Deploy site files to S3
    new s3deployment.BucketDeployment(this, 'InexHtmle', {
      sources: [s3deployment.Source.asset("./site/")],
      destinationBucket: podcastBucket,
    });
  }
}
