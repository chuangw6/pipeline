// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// backed by a key store that you own and manage. When you use a KMS key in a
// custom key store for a cryptographic operation, the cryptographic operation is
// actually performed in your key store using your keys. KMS supports CloudHSM key
// stores (https://docs.aws.amazon.com/kms/latest/developerguide/keystore-cloudhsm.html)
// backed by an CloudHSM cluster (https://docs.aws.amazon.com/cloudhsm/latest/userguide/clusters.html)
// and external key stores (https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html)
// backed by an external key store proxy and external key manager outside of Amazon
// Web Services. This operation is part of the custom key stores (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// feature in KMS, which combines the convenience and extensive integration of KMS
// with the isolation and control of a key store that you own and manage. Before
// you create the custom key store, the required elements must be in place and
// operational. We recommend that you use the test tools that KMS provides to
// verify the configuration your external key store proxy. For details about the
// required elements and verification tests, see Assemble the prerequisites (for
// CloudHSM key stores) (https://docs.aws.amazon.com/kms/latest/developerguide/create-keystore.html#before-keystore)
// or Assemble the prerequisites (for external key stores) (https://docs.aws.amazon.com/kms/latest/developerguide/create-xks-keystore.html#xks-requirements)
// in the Key Management Service Developer Guide. To create a custom key store, use
// the following parameters.
//   - To create an CloudHSM key store, specify the CustomKeyStoreName ,
//     CloudHsmClusterId , KeyStorePassword , and TrustAnchorCertificate . The
//     CustomKeyStoreType parameter is optional for CloudHSM key stores. If you
//     include it, set it to the default value, AWS_CLOUDHSM . For help with
//     failures, see Troubleshooting an CloudHSM key store (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html)
//     in the Key Management Service Developer Guide.
//   - To create an external key store, specify the CustomKeyStoreName and a
//     CustomKeyStoreType of EXTERNAL_KEY_STORE . Also, specify values for
//     XksProxyConnectivity , XksProxyAuthenticationCredential , XksProxyUriEndpoint
//     , and XksProxyUriPath . If your XksProxyConnectivity value is
//     VPC_ENDPOINT_SERVICE , specify the XksProxyVpcEndpointServiceName parameter.
//     For help with failures, see Troubleshooting an external key store (https://docs.aws.amazon.com/kms/latest/developerguide/xks-troubleshooting.html)
//     in the Key Management Service Developer Guide.
//
// For external key stores: Some external key managers provide a simpler method
// for creating an external key store. For details, see your external key manager
// documentation. When creating an external key store in the KMS console, you can
// upload a JSON-based proxy configuration file with the desired values. You cannot
// use a proxy configuration with the CreateCustomKeyStore operation. However, you
// can use the values in the file to help you determine the correct values for the
// CreateCustomKeyStore parameters. When the operation completes successfully, it
// returns the ID of the new custom key store. Before you can use your new custom
// key store, you need to use the ConnectCustomKeyStore operation to connect a new
// CloudHSM key store to its CloudHSM cluster, or to connect a new external key
// store to the external key store proxy for your external key manager. Even if you
// are not going to use your custom key store immediately, you might want to
// connect it to verify that all settings are correct and then disconnect it until
// you are ready to use it. For help with failures, see Troubleshooting a custom
// key store (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html)
// in the Key Management Service Developer Guide. Cross-account use: No. You cannot
// perform this operation on a custom key store in a different Amazon Web Services
// account. Required permissions: kms:CreateCustomKeyStore (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (IAM policy). Related operations:
//   - ConnectCustomKeyStore
//   - DeleteCustomKeyStore
//   - DescribeCustomKeyStores
//   - DisconnectCustomKeyStore
//   - UpdateCustomKeyStore
func (c *Client) CreateCustomKeyStore(ctx context.Context, params *CreateCustomKeyStoreInput, optFns ...func(*Options)) (*CreateCustomKeyStoreOutput, error) {
	if params == nil {
		params = &CreateCustomKeyStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomKeyStore", params, optFns, c.addOperationCreateCustomKeyStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomKeyStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCustomKeyStoreInput struct {

	// Specifies a friendly name for the custom key store. The name must be unique in
	// your Amazon Web Services account and Region. This parameter is required for all
	// custom key stores.
	//
	// This member is required.
	CustomKeyStoreName *string

	// Identifies the CloudHSM cluster for an CloudHSM key store. This parameter is
	// required for custom key stores with CustomKeyStoreType of AWS_CLOUDHSM . Enter
	// the cluster ID of any active CloudHSM cluster that is not already associated
	// with a custom key store. To find the cluster ID, use the DescribeClusters (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_DescribeClusters.html)
	// operation.
	CloudHsmClusterId *string

	// Specifies the type of custom key store. The default value is AWS_CLOUDHSM . For
	// a custom key store backed by an CloudHSM cluster, omit the parameter or enter
	// AWS_CLOUDHSM . For a custom key store backed by an external key manager outside
	// of Amazon Web Services, enter EXTERNAL_KEY_STORE . You cannot change this
	// property after the key store is created.
	CustomKeyStoreType types.CustomKeyStoreType

	// Specifies the kmsuser password for an CloudHSM key store. This parameter is
	// required for custom key stores with a CustomKeyStoreType of AWS_CLOUDHSM . Enter
	// the password of the kmsuser crypto user (CU) account (https://docs.aws.amazon.com/kms/latest/developerguide/key-store-concepts.html#concept-kmsuser)
	// in the specified CloudHSM cluster. KMS logs into the cluster as this user to
	// manage key material on your behalf. The password must be a string of 7 to 32
	// characters. Its value is case sensitive. This parameter tells KMS the kmsuser
	// account password; it does not change the password in the CloudHSM cluster.
	KeyStorePassword *string

	// Specifies the certificate for an CloudHSM key store. This parameter is required
	// for custom key stores with a CustomKeyStoreType of AWS_CLOUDHSM . Enter the
	// content of the trust anchor certificate for the CloudHSM cluster. This is the
	// content of the customerCA.crt file that you created when you initialized the
	// cluster (https://docs.aws.amazon.com/cloudhsm/latest/userguide/initialize-cluster.html)
	// .
	TrustAnchorCertificate *string

	// Specifies an authentication credential for the external key store proxy (XKS
	// proxy). This parameter is required for all custom key stores with a
	// CustomKeyStoreType of EXTERNAL_KEY_STORE . The XksProxyAuthenticationCredential
	// has two required elements: RawSecretAccessKey , a secret key, and AccessKeyId ,
	// a unique identifier for the RawSecretAccessKey . For character requirements, see
	// XksProxyAuthenticationCredentialType . KMS uses this authentication credential
	// to sign requests to the external key store proxy on your behalf. This credential
	// is unrelated to Identity and Access Management (IAM) and Amazon Web Services
	// credentials. This parameter doesn't set or change the authentication credentials
	// on the XKS proxy. It just tells KMS the credential that you established on your
	// external key store proxy. If you rotate your proxy authentication credential,
	// use the UpdateCustomKeyStore operation to provide the new credential to KMS.
	XksProxyAuthenticationCredential *types.XksProxyAuthenticationCredentialType

	// Indicates how KMS communicates with the external key store proxy. This
	// parameter is required for custom key stores with a CustomKeyStoreType of
	// EXTERNAL_KEY_STORE . If the external key store proxy uses a public endpoint,
	// specify PUBLIC_ENDPOINT . If the external key store proxy uses a Amazon VPC
	// endpoint service for communication with KMS, specify VPC_ENDPOINT_SERVICE . For
	// help making this choice, see Choosing a connectivity option (https://docs.aws.amazon.com/kms/latest/developerguide/plan-xks-keystore.html#choose-xks-connectivity)
	// in the Key Management Service Developer Guide. An Amazon VPC endpoint service
	// keeps your communication with KMS in a private address space entirely within
	// Amazon Web Services, but it requires more configuration, including establishing
	// a Amazon VPC with multiple subnets, a VPC endpoint service, a network load
	// balancer, and a verified private DNS name. A public endpoint is simpler to set
	// up, but it might be slower and might not fulfill your security requirements. You
	// might consider testing with a public endpoint, and then establishing a VPC
	// endpoint service for production tasks. Note that this choice does not determine
	// the location of the external key store proxy. Even if you choose a VPC endpoint
	// service, the proxy can be hosted within the VPC or outside of Amazon Web
	// Services such as in your corporate data center.
	XksProxyConnectivity types.XksProxyConnectivityType

	// Specifies the endpoint that KMS uses to send requests to the external key store
	// proxy (XKS proxy). This parameter is required for custom key stores with a
	// CustomKeyStoreType of EXTERNAL_KEY_STORE . The protocol must be HTTPS. KMS
	// communicates on port 443. Do not specify the port in the XksProxyUriEndpoint
	// value. For external key stores with XksProxyConnectivity value of
	// VPC_ENDPOINT_SERVICE , specify https:// followed by the private DNS name of the
	// VPC endpoint service. For external key stores with PUBLIC_ENDPOINT
	// connectivity, this endpoint must be reachable before you create the custom key
	// store. KMS connects to the external key store proxy while creating the custom
	// key store. For external key stores with VPC_ENDPOINT_SERVICE connectivity, KMS
	// connects when you call the ConnectCustomKeyStore operation. The value of this
	// parameter must begin with https:// . The remainder can contain upper and lower
	// case letters (A-Z and a-z), numbers (0-9), dots ( . ), and hyphens ( - ).
	// Additional slashes ( / and \ ) are not permitted. Uniqueness requirements:
	//   - The combined XksProxyUriEndpoint and XksProxyUriPath values must be unique
	//   in the Amazon Web Services account and Region.
	//   - An external key store with PUBLIC_ENDPOINT connectivity cannot use the same
	//   XksProxyUriEndpoint value as an external key store with VPC_ENDPOINT_SERVICE
	//   connectivity in the same Amazon Web Services Region.
	//   - Each external key store with VPC_ENDPOINT_SERVICE connectivity must have its
	//   own private DNS name. The XksProxyUriEndpoint value for external key stores
	//   with VPC_ENDPOINT_SERVICE connectivity (private DNS name) must be unique in
	//   the Amazon Web Services account and Region.
	XksProxyUriEndpoint *string

	// Specifies the base path to the proxy APIs for this external key store. To find
	// this value, see the documentation for your external key store proxy. This
	// parameter is required for all custom key stores with a CustomKeyStoreType of
	// EXTERNAL_KEY_STORE . The value must start with / and must end with /kms/xks/v1
	// where v1 represents the version of the KMS external key store proxy API. This
	// path can include an optional prefix between the required elements such as
	// /prefix/kms/xks/v1 . Uniqueness requirements:
	//   - The combined XksProxyUriEndpoint and XksProxyUriPath values must be unique
	//   in the Amazon Web Services account and Region.
	XksProxyUriPath *string

	// Specifies the name of the Amazon VPC endpoint service for interface endpoints
	// that is used to communicate with your external key store proxy (XKS proxy). This
	// parameter is required when the value of CustomKeyStoreType is EXTERNAL_KEY_STORE
	// and the value of XksProxyConnectivity is VPC_ENDPOINT_SERVICE . The Amazon VPC
	// endpoint service must fulfill all requirements (https://docs.aws.amazon.com/kms/latest/developerguide/create-xks-keystore.html#xks-requirements)
	// for use with an external key store. Uniqueness requirements:
	//   - External key stores with VPC_ENDPOINT_SERVICE connectivity can share an
	//   Amazon VPC, but each external key store must have its own VPC endpoint service
	//   and private DNS name.
	XksProxyVpcEndpointServiceName *string

	noSmithyDocumentSerde
}

type CreateCustomKeyStoreOutput struct {

	// A unique identifier for the new custom key store.
	CustomKeyStoreId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomKeyStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCustomKeyStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCustomKeyStore{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateCustomKeyStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomKeyStore(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateCustomKeyStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "CreateCustomKeyStore",
	}
}
