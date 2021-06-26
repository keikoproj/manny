package mocks

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/mock"
)

type S3API struct {
	mock.Mock
	s3iface.S3API
}

func (m *S3API) AbortMultipartUpload(*s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	return nil, nil
}
func (m *S3API) AbortMultipartUploadWithContext(aws.Context, *s3.AbortMultipartUploadInput, ...request.Option) (*s3.AbortMultipartUploadOutput, error) {
	return nil, nil
}
func (m *S3API) AbortMultipartUploadRequest(*s3.AbortMultipartUploadInput) (*request.Request, *s3.AbortMultipartUploadOutput) {
	return nil, nil
}

func (m *S3API) CompleteMultipartUpload(*s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	return nil, nil
}
func (m *S3API) CompleteMultipartUploadWithContext(aws.Context, *s3.CompleteMultipartUploadInput, ...request.Option) (*s3.CompleteMultipartUploadOutput, error) {
	return nil, nil
}
func (m *S3API) CompleteMultipartUploadRequest(*s3.CompleteMultipartUploadInput) (*request.Request, *s3.CompleteMultipartUploadOutput) {
	return nil, nil
}

func (m *S3API) CopyObject(*s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	return nil, nil
}
func (m *S3API) CopyObjectWithContext(aws.Context, *s3.CopyObjectInput, ...request.Option) (*s3.CopyObjectOutput, error) {
	return nil, nil
}
func (m *S3API) CopyObjectRequest(*s3.CopyObjectInput) (*request.Request, *s3.CopyObjectOutput) {
	return nil, nil
}

func (m *S3API) CreateBucket(*s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	return nil, nil
}
func (m *S3API) CreateBucketWithContext(aws.Context, *s3.CreateBucketInput, ...request.Option) (*s3.CreateBucketOutput, error) {
	return nil, nil
}
func (m *S3API) CreateBucketRequest(*s3.CreateBucketInput) (*request.Request, *s3.CreateBucketOutput) {
	return nil, nil
}

func (m *S3API) CreateMultipartUpload(*s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	return nil, nil
}
func (m *S3API) CreateMultipartUploadWithContext(aws.Context, *s3.CreateMultipartUploadInput, ...request.Option) (*s3.CreateMultipartUploadOutput, error) {
	return nil, nil
}
func (m *S3API) CreateMultipartUploadRequest(*s3.CreateMultipartUploadInput) (*request.Request, *s3.CreateMultipartUploadOutput) {
	return nil, nil
}

func (m *S3API) DeleteBucket(*s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketWithContext(aws.Context, *s3.DeleteBucketInput, ...request.Option) (*s3.DeleteBucketOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketRequest(*s3.DeleteBucketInput) (*request.Request, *s3.DeleteBucketOutput) {
	return nil, nil
}

func (m *S3API) DeleteBucketAnalyticsConfiguration(*s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketAnalyticsConfigurationWithContext(aws.Context, *s3.DeleteBucketAnalyticsConfigurationInput, ...request.Option) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketAnalyticsConfigurationRequest(*s3.DeleteBucketAnalyticsConfigurationInput) (*request.Request, *s3.DeleteBucketAnalyticsConfigurationOutput) {
	return nil, nil
}

func (m *S3API) DeleteBucketCors(*s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketCorsWithContext(aws.Context, *s3.DeleteBucketCorsInput, ...request.Option) (*s3.DeleteBucketCorsOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketCorsRequest(*s3.DeleteBucketCorsInput) (*request.Request, *s3.DeleteBucketCorsOutput) {
	return nil, nil
}

func (m *S3API) DeleteBucketEncryption(*s3.DeleteBucketEncryptionInput) (*s3.DeleteBucketEncryptionOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketEncryptionWithContext(aws.Context, *s3.DeleteBucketEncryptionInput, ...request.Option) (*s3.DeleteBucketEncryptionOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketEncryptionRequest(*s3.DeleteBucketEncryptionInput) (*request.Request, *s3.DeleteBucketEncryptionOutput) {
	return nil, nil
}

func (m *S3API) DeleteBucketInventoryConfiguration(*s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketInventoryConfigurationWithContext(aws.Context, *s3.DeleteBucketInventoryConfigurationInput, ...request.Option) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketInventoryConfigurationRequest(*s3.DeleteBucketInventoryConfigurationInput) (*request.Request, *s3.DeleteBucketInventoryConfigurationOutput) {
	return nil, nil
}

func (m *S3API) DeleteBucketLifecycle(*s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketLifecycleWithContext(aws.Context, *s3.DeleteBucketLifecycleInput, ...request.Option) (*s3.DeleteBucketLifecycleOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketLifecycleRequest(*s3.DeleteBucketLifecycleInput) (*request.Request, *s3.DeleteBucketLifecycleOutput) {
	return nil, nil
}

func (m *S3API) DeleteBucketMetricsConfiguration(*s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketMetricsConfigurationWithContext(aws.Context, *s3.DeleteBucketMetricsConfigurationInput, ...request.Option) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketMetricsConfigurationRequest(*s3.DeleteBucketMetricsConfigurationInput) (*request.Request, *s3.DeleteBucketMetricsConfigurationOutput) {
	return nil, nil
}

func (m *S3API) DeleteBucketPolicy(*s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketPolicyWithContext(aws.Context, *s3.DeleteBucketPolicyInput, ...request.Option) (*s3.DeleteBucketPolicyOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketPolicyRequest(*s3.DeleteBucketPolicyInput) (*request.Request, *s3.DeleteBucketPolicyOutput) {
	return nil, nil
}

func (m *S3API) DeleteBucketReplication(*s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketReplicationWithContext(aws.Context, *s3.DeleteBucketReplicationInput, ...request.Option) (*s3.DeleteBucketReplicationOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketReplicationRequest(*s3.DeleteBucketReplicationInput) (*request.Request, *s3.DeleteBucketReplicationOutput) {
	return nil, nil
}

func (m *S3API) DeleteBucketTagging(*s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketTaggingWithContext(aws.Context, *s3.DeleteBucketTaggingInput, ...request.Option) (*s3.DeleteBucketTaggingOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketTaggingRequest(*s3.DeleteBucketTaggingInput) (*request.Request, *s3.DeleteBucketTaggingOutput) {
	return nil, nil
}

func (m *S3API) DeleteBucketWebsite(*s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketWebsiteWithContext(aws.Context, *s3.DeleteBucketWebsiteInput, ...request.Option) (*s3.DeleteBucketWebsiteOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteBucketWebsiteRequest(*s3.DeleteBucketWebsiteInput) (*request.Request, *s3.DeleteBucketWebsiteOutput) {
	return nil, nil
}

func (m *S3API) DeleteObject(*s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteObjectWithContext(aws.Context, *s3.DeleteObjectInput, ...request.Option) (*s3.DeleteObjectOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteObjectRequest(*s3.DeleteObjectInput) (*request.Request, *s3.DeleteObjectOutput) {
	return nil, nil
}

func (m *S3API) DeleteObjectTagging(*s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteObjectTaggingWithContext(aws.Context, *s3.DeleteObjectTaggingInput, ...request.Option) (*s3.DeleteObjectTaggingOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteObjectTaggingRequest(*s3.DeleteObjectTaggingInput) (*request.Request, *s3.DeleteObjectTaggingOutput) {
	return nil, nil
}

func (m *S3API) DeleteObjects(*s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteObjectsWithContext(aws.Context, *s3.DeleteObjectsInput, ...request.Option) (*s3.DeleteObjectsOutput, error) {
	return nil, nil
}
func (m *S3API) DeleteObjectsRequest(*s3.DeleteObjectsInput) (*request.Request, *s3.DeleteObjectsOutput) {
	return nil, nil
}

func (m *S3API) DeletePublicAccessBlock(*s3.DeletePublicAccessBlockInput) (*s3.DeletePublicAccessBlockOutput, error) {
	return nil, nil
}
func (m *S3API) DeletePublicAccessBlockWithContext(aws.Context, *s3.DeletePublicAccessBlockInput, ...request.Option) (*s3.DeletePublicAccessBlockOutput, error) {
	return nil, nil
}
func (m *S3API) DeletePublicAccessBlockRequest(*s3.DeletePublicAccessBlockInput) (*request.Request, *s3.DeletePublicAccessBlockOutput) {
	return nil, nil
}

func (m *S3API) GetBucketAccelerateConfiguration(*s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketAccelerateConfigurationWithContext(aws.Context, *s3.GetBucketAccelerateConfigurationInput, ...request.Option) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketAccelerateConfigurationRequest(*s3.GetBucketAccelerateConfigurationInput) (*request.Request, *s3.GetBucketAccelerateConfigurationOutput) {
	return nil, nil
}

func (m *S3API) GetBucketAcl(*s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketAclWithContext(aws.Context, *s3.GetBucketAclInput, ...request.Option) (*s3.GetBucketAclOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketAclRequest(*s3.GetBucketAclInput) (*request.Request, *s3.GetBucketAclOutput) {
	return nil, nil
}

func (m *S3API) GetBucketAnalyticsConfiguration(*s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketAnalyticsConfigurationWithContext(aws.Context, *s3.GetBucketAnalyticsConfigurationInput, ...request.Option) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketAnalyticsConfigurationRequest(*s3.GetBucketAnalyticsConfigurationInput) (*request.Request, *s3.GetBucketAnalyticsConfigurationOutput) {
	return nil, nil
}

func (m *S3API) GetBucketCors(*s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketCorsWithContext(aws.Context, *s3.GetBucketCorsInput, ...request.Option) (*s3.GetBucketCorsOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketCorsRequest(*s3.GetBucketCorsInput) (*request.Request, *s3.GetBucketCorsOutput) {
	return nil, nil
}

func (m *S3API) GetBucketEncryption(*s3.GetBucketEncryptionInput) (*s3.GetBucketEncryptionOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketEncryptionWithContext(aws.Context, *s3.GetBucketEncryptionInput, ...request.Option) (*s3.GetBucketEncryptionOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketEncryptionRequest(*s3.GetBucketEncryptionInput) (*request.Request, *s3.GetBucketEncryptionOutput) {
	return nil, nil
}

func (m *S3API) GetBucketInventoryConfiguration(*s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketInventoryConfigurationWithContext(aws.Context, *s3.GetBucketInventoryConfigurationInput, ...request.Option) (*s3.GetBucketInventoryConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketInventoryConfigurationRequest(*s3.GetBucketInventoryConfigurationInput) (*request.Request, *s3.GetBucketInventoryConfigurationOutput) {
	return nil, nil
}

func (m *S3API) GetBucketLifecycle(*s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketLifecycleWithContext(aws.Context, *s3.GetBucketLifecycleInput, ...request.Option) (*s3.GetBucketLifecycleOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketLifecycleRequest(*s3.GetBucketLifecycleInput) (*request.Request, *s3.GetBucketLifecycleOutput) {
	return nil, nil
}

func (m *S3API) GetBucketLifecycleConfiguration(*s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketLifecycleConfigurationWithContext(aws.Context, *s3.GetBucketLifecycleConfigurationInput, ...request.Option) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketLifecycleConfigurationRequest(*s3.GetBucketLifecycleConfigurationInput) (*request.Request, *s3.GetBucketLifecycleConfigurationOutput) {
	return nil, nil
}

func (m *S3API) GetBucketLocation(*s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketLocationWithContext(aws.Context, *s3.GetBucketLocationInput, ...request.Option) (*s3.GetBucketLocationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketLocationRequest(*s3.GetBucketLocationInput) (*request.Request, *s3.GetBucketLocationOutput) {
	return nil, nil
}

func (m *S3API) GetBucketLogging(*s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketLoggingWithContext(aws.Context, *s3.GetBucketLoggingInput, ...request.Option) (*s3.GetBucketLoggingOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketLoggingRequest(*s3.GetBucketLoggingInput) (*request.Request, *s3.GetBucketLoggingOutput) {
	return nil, nil
}

func (m *S3API) GetBucketMetricsConfiguration(*s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketMetricsConfigurationWithContext(aws.Context, *s3.GetBucketMetricsConfigurationInput, ...request.Option) (*s3.GetBucketMetricsConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketMetricsConfigurationRequest(*s3.GetBucketMetricsConfigurationInput) (*request.Request, *s3.GetBucketMetricsConfigurationOutput) {
	return nil, nil
}

func (m *S3API) GetBucketNotification(*s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error) {
	return nil, nil
}
func (m *S3API) GetBucketNotificationWithContext(aws.Context, *s3.GetBucketNotificationConfigurationRequest, ...request.Option) (*s3.NotificationConfigurationDeprecated, error) {
	return nil, nil
}
func (m *S3API) GetBucketNotificationRequest(*s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfigurationDeprecated) {
	return nil, nil
}

func (m *S3API) GetBucketNotificationConfiguration(*s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error) {
	return nil, nil
}
func (m *S3API) GetBucketNotificationConfigurationWithContext(aws.Context, *s3.GetBucketNotificationConfigurationRequest, ...request.Option) (*s3.NotificationConfiguration, error) {
	return nil, nil
}
func (m *S3API) GetBucketNotificationConfigurationRequest(*s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfiguration) {
	return nil, nil
}

func (m *S3API) GetBucketPolicy(*s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketPolicyWithContext(aws.Context, *s3.GetBucketPolicyInput, ...request.Option) (*s3.GetBucketPolicyOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketPolicyRequest(*s3.GetBucketPolicyInput) (*request.Request, *s3.GetBucketPolicyOutput) {
	return nil, nil
}

func (m *S3API) GetBucketPolicyStatus(*s3.GetBucketPolicyStatusInput) (*s3.GetBucketPolicyStatusOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketPolicyStatusWithContext(aws.Context, *s3.GetBucketPolicyStatusInput, ...request.Option) (*s3.GetBucketPolicyStatusOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketPolicyStatusRequest(*s3.GetBucketPolicyStatusInput) (*request.Request, *s3.GetBucketPolicyStatusOutput) {
	return nil, nil
}

func (m *S3API) GetBucketReplication(*s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketReplicationWithContext(aws.Context, *s3.GetBucketReplicationInput, ...request.Option) (*s3.GetBucketReplicationOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketReplicationRequest(*s3.GetBucketReplicationInput) (*request.Request, *s3.GetBucketReplicationOutput) {
	return nil, nil
}

func (m *S3API) GetBucketRequestPayment(*s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketRequestPaymentWithContext(aws.Context, *s3.GetBucketRequestPaymentInput, ...request.Option) (*s3.GetBucketRequestPaymentOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketRequestPaymentRequest(*s3.GetBucketRequestPaymentInput) (*request.Request, *s3.GetBucketRequestPaymentOutput) {
	return nil, nil
}

func (m *S3API) GetBucketTagging(*s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketTaggingWithContext(aws.Context, *s3.GetBucketTaggingInput, ...request.Option) (*s3.GetBucketTaggingOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketTaggingRequest(*s3.GetBucketTaggingInput) (*request.Request, *s3.GetBucketTaggingOutput) {
	return nil, nil
}

func (m *S3API) GetBucketVersioning(*s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketVersioningWithContext(aws.Context, *s3.GetBucketVersioningInput, ...request.Option) (*s3.GetBucketVersioningOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketVersioningRequest(*s3.GetBucketVersioningInput) (*request.Request, *s3.GetBucketVersioningOutput) {
	return nil, nil
}

func (m *S3API) GetBucketWebsite(*s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketWebsiteWithContext(aws.Context, *s3.GetBucketWebsiteInput, ...request.Option) (*s3.GetBucketWebsiteOutput, error) {
	return nil, nil
}
func (m *S3API) GetBucketWebsiteRequest(*s3.GetBucketWebsiteInput) (*request.Request, *s3.GetBucketWebsiteOutput) {
	return nil, nil
}

func (m *S3API) GetObject(*s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return nil, nil
}
func (m *S3API) GetObjectWithContext(ctx aws.Context, input *s3.GetObjectInput, options ...request.Option) (*s3.GetObjectOutput, error) {
	args := m.Called(ctx, input, options)
	return args.Get(0).(*s3.GetObjectOutput), args.Error(1)
}
func (m *S3API) GetObjectRequest(*s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput) {
	return nil, nil
}

func (m *S3API) GetObjectAcl(*s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error) {
	return nil, nil
}
func (m *S3API) GetObjectAclWithContext(aws.Context, *s3.GetObjectAclInput, ...request.Option) (*s3.GetObjectAclOutput, error) {
	return nil, nil
}
func (m *S3API) GetObjectAclRequest(*s3.GetObjectAclInput) (*request.Request, *s3.GetObjectAclOutput) {
	return nil, nil
}

func (m *S3API) GetObjectLegalHold(*s3.GetObjectLegalHoldInput) (*s3.GetObjectLegalHoldOutput, error) {
	return nil, nil
}
func (m *S3API) GetObjectLegalHoldWithContext(aws.Context, *s3.GetObjectLegalHoldInput, ...request.Option) (*s3.GetObjectLegalHoldOutput, error) {
	return nil, nil
}
func (m *S3API) GetObjectLegalHoldRequest(*s3.GetObjectLegalHoldInput) (*request.Request, *s3.GetObjectLegalHoldOutput) {
	return nil, nil
}

func (m *S3API) GetObjectLockConfiguration(*s3.GetObjectLockConfigurationInput) (*s3.GetObjectLockConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) GetObjectLockConfigurationWithContext(aws.Context, *s3.GetObjectLockConfigurationInput, ...request.Option) (*s3.GetObjectLockConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) GetObjectLockConfigurationRequest(*s3.GetObjectLockConfigurationInput) (*request.Request, *s3.GetObjectLockConfigurationOutput) {
	return nil, nil
}

func (m *S3API) GetObjectRetention(*s3.GetObjectRetentionInput) (*s3.GetObjectRetentionOutput, error) {
	return nil, nil
}
func (m *S3API) GetObjectRetentionWithContext(aws.Context, *s3.GetObjectRetentionInput, ...request.Option) (*s3.GetObjectRetentionOutput, error) {
	return nil, nil
}
func (m *S3API) GetObjectRetentionRequest(*s3.GetObjectRetentionInput) (*request.Request, *s3.GetObjectRetentionOutput) {
	return nil, nil
}

func (m *S3API) GetObjectTagging(*s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error) {
	return nil, nil
}
func (m *S3API) GetObjectTaggingWithContext(aws.Context, *s3.GetObjectTaggingInput, ...request.Option) (*s3.GetObjectTaggingOutput, error) {
	return nil, nil
}
func (m *S3API) GetObjectTaggingRequest(*s3.GetObjectTaggingInput) (*request.Request, *s3.GetObjectTaggingOutput) {
	return nil, nil
}

func (m *S3API) GetObjectTorrent(*s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error) {
	return nil, nil
}
func (m *S3API) GetObjectTorrentWithContext(aws.Context, *s3.GetObjectTorrentInput, ...request.Option) (*s3.GetObjectTorrentOutput, error) {
	return nil, nil
}
func (m *S3API) GetObjectTorrentRequest(*s3.GetObjectTorrentInput) (*request.Request, *s3.GetObjectTorrentOutput) {
	return nil, nil
}

func (m *S3API) GetPublicAccessBlock(*s3.GetPublicAccessBlockInput) (*s3.GetPublicAccessBlockOutput, error) {
	return nil, nil
}
func (m *S3API) GetPublicAccessBlockWithContext(aws.Context, *s3.GetPublicAccessBlockInput, ...request.Option) (*s3.GetPublicAccessBlockOutput, error) {
	return nil, nil
}
func (m *S3API) GetPublicAccessBlockRequest(*s3.GetPublicAccessBlockInput) (*request.Request, *s3.GetPublicAccessBlockOutput) {
	return nil, nil
}

func (m *S3API) HeadBucket(*s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
	return nil, nil
}
func (m *S3API) HeadBucketWithContext(aws.Context, *s3.HeadBucketInput, ...request.Option) (*s3.HeadBucketOutput, error) {
	return nil, nil
}
func (m *S3API) HeadBucketRequest(*s3.HeadBucketInput) (*request.Request, *s3.HeadBucketOutput) {
	return nil, nil
}

func (m *S3API) HeadObject(*s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	return nil, nil
}
func (m *S3API) HeadObjectWithContext(aws.Context, *s3.HeadObjectInput, ...request.Option) (*s3.HeadObjectOutput, error) {
	return nil, nil
}
func (m *S3API) HeadObjectRequest(*s3.HeadObjectInput) (*request.Request, *s3.HeadObjectOutput) {
	return nil, nil
}

func (m *S3API) ListBucketAnalyticsConfigurations(*s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	return nil, nil
}
func (m *S3API) ListBucketAnalyticsConfigurationsWithContext(aws.Context, *s3.ListBucketAnalyticsConfigurationsInput, ...request.Option) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	return nil, nil
}
func (m *S3API) ListBucketAnalyticsConfigurationsRequest(*s3.ListBucketAnalyticsConfigurationsInput) (*request.Request, *s3.ListBucketAnalyticsConfigurationsOutput) {
	return nil, nil
}

func (m *S3API) ListBucketInventoryConfigurations(*s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	return nil, nil
}
func (m *S3API) ListBucketInventoryConfigurationsWithContext(aws.Context, *s3.ListBucketInventoryConfigurationsInput, ...request.Option) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	return nil, nil
}
func (m *S3API) ListBucketInventoryConfigurationsRequest(*s3.ListBucketInventoryConfigurationsInput) (*request.Request, *s3.ListBucketInventoryConfigurationsOutput) {
	return nil, nil
}

func (m *S3API) ListBucketMetricsConfigurations(*s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	return nil, nil
}
func (m *S3API) ListBucketMetricsConfigurationsWithContext(aws.Context, *s3.ListBucketMetricsConfigurationsInput, ...request.Option) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	return nil, nil
}
func (m *S3API) ListBucketMetricsConfigurationsRequest(*s3.ListBucketMetricsConfigurationsInput) (*request.Request, *s3.ListBucketMetricsConfigurationsOutput) {
	return nil, nil
}

func (m *S3API) ListBuckets(*s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	return nil, nil
}
func (m *S3API) ListBucketsWithContext(aws.Context, *s3.ListBucketsInput, ...request.Option) (*s3.ListBucketsOutput, error) {
	return nil, nil
}
func (m *S3API) ListBucketsRequest(*s3.ListBucketsInput) (*request.Request, *s3.ListBucketsOutput) {
	return nil, nil
}

func (m *S3API) ListMultipartUploads(*s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
	return nil, nil
}
func (m *S3API) ListMultipartUploadsWithContext(aws.Context, *s3.ListMultipartUploadsInput, ...request.Option) (*s3.ListMultipartUploadsOutput, error) {
	return nil, nil
}
func (m *S3API) ListMultipartUploadsRequest(*s3.ListMultipartUploadsInput) (*request.Request, *s3.ListMultipartUploadsOutput) {
	return nil, nil
}

func (m *S3API) ListMultipartUploadsPages(*s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool) error {
	return nil
}
func (m *S3API) ListMultipartUploadsPagesWithContext(aws.Context, *s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m *S3API) ListObjectVersions(*s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
	return nil, nil
}
func (m *S3API) ListObjectVersionsWithContext(aws.Context, *s3.ListObjectVersionsInput, ...request.Option) (*s3.ListObjectVersionsOutput, error) {
	return nil, nil
}
func (m *S3API) ListObjectVersionsRequest(*s3.ListObjectVersionsInput) (*request.Request, *s3.ListObjectVersionsOutput) {
	return nil, nil
}

func (m *S3API) ListObjectVersionsPages(*s3.ListObjectVersionsInput, func(*s3.ListObjectVersionsOutput, bool) bool) error {
	return nil
}
func (m *S3API) ListObjectVersionsPagesWithContext(aws.Context, *s3.ListObjectVersionsInput, func(*s3.ListObjectVersionsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m *S3API) ListObjects(*s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	return nil, nil
}
func (m *S3API) ListObjectsWithContext(aws.Context, *s3.ListObjectsInput, ...request.Option) (*s3.ListObjectsOutput, error) {
	return nil, nil
}
func (m *S3API) ListObjectsRequest(*s3.ListObjectsInput) (*request.Request, *s3.ListObjectsOutput) {
	return nil, nil
}

func (m *S3API) ListObjectsPages(*s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool) error {
	return nil
}
func (m *S3API) ListObjectsPagesWithContext(aws.Context, *s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m *S3API) ListObjectsV2(*s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return nil, nil
}
func (m *S3API) ListObjectsV2WithContext(aws.Context, *s3.ListObjectsV2Input, ...request.Option) (*s3.ListObjectsV2Output, error) {
	return nil, nil
}
func (m *S3API) ListObjectsV2Request(*s3.ListObjectsV2Input) (*request.Request, *s3.ListObjectsV2Output) {
	return nil, nil
}

func (m *S3API) ListObjectsV2Pages(*s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool) error {
	return nil
}
func (m *S3API) ListObjectsV2PagesWithContext(aws.Context, *s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool, ...request.Option) error {
	return nil
}

func (m *S3API) ListParts(*s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	return nil, nil
}
func (m *S3API) ListPartsWithContext(aws.Context, *s3.ListPartsInput, ...request.Option) (*s3.ListPartsOutput, error) {
	return nil, nil
}
func (m *S3API) ListPartsRequest(*s3.ListPartsInput) (*request.Request, *s3.ListPartsOutput) {
	return nil, nil
}

func (m *S3API) ListPartsPages(*s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool) error {
	return nil
}
func (m *S3API) ListPartsPagesWithContext(aws.Context, *s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m *S3API) PutBucketAccelerateConfiguration(*s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketAccelerateConfigurationWithContext(aws.Context, *s3.PutBucketAccelerateConfigurationInput, ...request.Option) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketAccelerateConfigurationRequest(*s3.PutBucketAccelerateConfigurationInput) (*request.Request, *s3.PutBucketAccelerateConfigurationOutput) {
	return nil, nil
}

func (m *S3API) PutBucketAcl(*s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketAclWithContext(aws.Context, *s3.PutBucketAclInput, ...request.Option) (*s3.PutBucketAclOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketAclRequest(*s3.PutBucketAclInput) (*request.Request, *s3.PutBucketAclOutput) {
	return nil, nil
}

func (m *S3API) PutBucketAnalyticsConfiguration(*s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketAnalyticsConfigurationWithContext(aws.Context, *s3.PutBucketAnalyticsConfigurationInput, ...request.Option) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketAnalyticsConfigurationRequest(*s3.PutBucketAnalyticsConfigurationInput) (*request.Request, *s3.PutBucketAnalyticsConfigurationOutput) {
	return nil, nil
}

func (m *S3API) PutBucketCors(*s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketCorsWithContext(aws.Context, *s3.PutBucketCorsInput, ...request.Option) (*s3.PutBucketCorsOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketCorsRequest(*s3.PutBucketCorsInput) (*request.Request, *s3.PutBucketCorsOutput) {
	return nil, nil
}

func (m *S3API) PutBucketEncryption(*s3.PutBucketEncryptionInput) (*s3.PutBucketEncryptionOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketEncryptionWithContext(aws.Context, *s3.PutBucketEncryptionInput, ...request.Option) (*s3.PutBucketEncryptionOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketEncryptionRequest(*s3.PutBucketEncryptionInput) (*request.Request, *s3.PutBucketEncryptionOutput) {
	return nil, nil
}

func (m *S3API) PutBucketInventoryConfiguration(*s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketInventoryConfigurationWithContext(aws.Context, *s3.PutBucketInventoryConfigurationInput, ...request.Option) (*s3.PutBucketInventoryConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketInventoryConfigurationRequest(*s3.PutBucketInventoryConfigurationInput) (*request.Request, *s3.PutBucketInventoryConfigurationOutput) {
	return nil, nil
}

func (m *S3API) PutBucketLifecycle(*s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketLifecycleWithContext(aws.Context, *s3.PutBucketLifecycleInput, ...request.Option) (*s3.PutBucketLifecycleOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketLifecycleRequest(*s3.PutBucketLifecycleInput) (*request.Request, *s3.PutBucketLifecycleOutput) {
	return nil, nil
}

func (m *S3API) PutBucketLifecycleConfiguration(*s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketLifecycleConfigurationWithContext(aws.Context, *s3.PutBucketLifecycleConfigurationInput, ...request.Option) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketLifecycleConfigurationRequest(*s3.PutBucketLifecycleConfigurationInput) (*request.Request, *s3.PutBucketLifecycleConfigurationOutput) {
	return nil, nil
}

func (m *S3API) PutBucketLogging(*s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketLoggingWithContext(aws.Context, *s3.PutBucketLoggingInput, ...request.Option) (*s3.PutBucketLoggingOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketLoggingRequest(*s3.PutBucketLoggingInput) (*request.Request, *s3.PutBucketLoggingOutput) {
	return nil, nil
}

func (m *S3API) PutBucketMetricsConfiguration(*s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketMetricsConfigurationWithContext(aws.Context, *s3.PutBucketMetricsConfigurationInput, ...request.Option) (*s3.PutBucketMetricsConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketMetricsConfigurationRequest(*s3.PutBucketMetricsConfigurationInput) (*request.Request, *s3.PutBucketMetricsConfigurationOutput) {
	return nil, nil
}

func (m *S3API) PutBucketNotification(*s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketNotificationWithContext(aws.Context, *s3.PutBucketNotificationInput, ...request.Option) (*s3.PutBucketNotificationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketNotificationRequest(*s3.PutBucketNotificationInput) (*request.Request, *s3.PutBucketNotificationOutput) {
	return nil, nil
}

func (m *S3API) PutBucketNotificationConfiguration(*s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketNotificationConfigurationWithContext(aws.Context, *s3.PutBucketNotificationConfigurationInput, ...request.Option) (*s3.PutBucketNotificationConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketNotificationConfigurationRequest(*s3.PutBucketNotificationConfigurationInput) (*request.Request, *s3.PutBucketNotificationConfigurationOutput) {
	return nil, nil
}

func (m *S3API) PutBucketPolicy(*s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketPolicyWithContext(aws.Context, *s3.PutBucketPolicyInput, ...request.Option) (*s3.PutBucketPolicyOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketPolicyRequest(*s3.PutBucketPolicyInput) (*request.Request, *s3.PutBucketPolicyOutput) {
	return nil, nil
}

func (m *S3API) PutBucketReplication(*s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketReplicationWithContext(aws.Context, *s3.PutBucketReplicationInput, ...request.Option) (*s3.PutBucketReplicationOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketReplicationRequest(*s3.PutBucketReplicationInput) (*request.Request, *s3.PutBucketReplicationOutput) {
	return nil, nil
}

func (m *S3API) PutBucketRequestPayment(*s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketRequestPaymentWithContext(aws.Context, *s3.PutBucketRequestPaymentInput, ...request.Option) (*s3.PutBucketRequestPaymentOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketRequestPaymentRequest(*s3.PutBucketRequestPaymentInput) (*request.Request, *s3.PutBucketRequestPaymentOutput) {
	return nil, nil
}

func (m *S3API) PutBucketTagging(*s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketTaggingWithContext(aws.Context, *s3.PutBucketTaggingInput, ...request.Option) (*s3.PutBucketTaggingOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketTaggingRequest(*s3.PutBucketTaggingInput) (*request.Request, *s3.PutBucketTaggingOutput) {
	return nil, nil
}

func (m *S3API) PutBucketVersioning(*s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketVersioningWithContext(aws.Context, *s3.PutBucketVersioningInput, ...request.Option) (*s3.PutBucketVersioningOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketVersioningRequest(*s3.PutBucketVersioningInput) (*request.Request, *s3.PutBucketVersioningOutput) {
	return nil, nil
}

func (m *S3API) PutBucketWebsite(*s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketWebsiteWithContext(aws.Context, *s3.PutBucketWebsiteInput, ...request.Option) (*s3.PutBucketWebsiteOutput, error) {
	return nil, nil
}
func (m *S3API) PutBucketWebsiteRequest(*s3.PutBucketWebsiteInput) (*request.Request, *s3.PutBucketWebsiteOutput) {
	return nil, nil
}

func (m *S3API) PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return nil, nil
}
func (m *S3API) PutObjectWithContext(aws.Context, *s3.PutObjectInput, ...request.Option) (*s3.PutObjectOutput, error) {
	return nil, nil
}
func (m *S3API) PutObjectRequest(*s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput) {
	return nil, nil
}

func (m *S3API) PutObjectAcl(*s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error) {
	return nil, nil
}
func (m *S3API) PutObjectAclWithContext(aws.Context, *s3.PutObjectAclInput, ...request.Option) (*s3.PutObjectAclOutput, error) {
	return nil, nil
}
func (m *S3API) PutObjectAclRequest(*s3.PutObjectAclInput) (*request.Request, *s3.PutObjectAclOutput) {
	return nil, nil
}

func (m *S3API) PutObjectLegalHold(*s3.PutObjectLegalHoldInput) (*s3.PutObjectLegalHoldOutput, error) {
	return nil, nil
}
func (m *S3API) PutObjectLegalHoldWithContext(aws.Context, *s3.PutObjectLegalHoldInput, ...request.Option) (*s3.PutObjectLegalHoldOutput, error) {
	return nil, nil
}
func (m *S3API) PutObjectLegalHoldRequest(*s3.PutObjectLegalHoldInput) (*request.Request, *s3.PutObjectLegalHoldOutput) {
	return nil, nil
}

func (m *S3API) PutObjectLockConfiguration(*s3.PutObjectLockConfigurationInput) (*s3.PutObjectLockConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutObjectLockConfigurationWithContext(aws.Context, *s3.PutObjectLockConfigurationInput, ...request.Option) (*s3.PutObjectLockConfigurationOutput, error) {
	return nil, nil
}
func (m *S3API) PutObjectLockConfigurationRequest(*s3.PutObjectLockConfigurationInput) (*request.Request, *s3.PutObjectLockConfigurationOutput) {
	return nil, nil
}

func (m *S3API) PutObjectRetention(*s3.PutObjectRetentionInput) (*s3.PutObjectRetentionOutput, error) {
	return nil, nil
}
func (m *S3API) PutObjectRetentionWithContext(aws.Context, *s3.PutObjectRetentionInput, ...request.Option) (*s3.PutObjectRetentionOutput, error) {
	return nil, nil
}
func (m *S3API) PutObjectRetentionRequest(*s3.PutObjectRetentionInput) (*request.Request, *s3.PutObjectRetentionOutput) {
	return nil, nil
}

func (m *S3API) PutObjectTagging(*s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error) {
	return nil, nil
}
func (m *S3API) PutObjectTaggingWithContext(aws.Context, *s3.PutObjectTaggingInput, ...request.Option) (*s3.PutObjectTaggingOutput, error) {
	return nil, nil
}
func (m *S3API) PutObjectTaggingRequest(*s3.PutObjectTaggingInput) (*request.Request, *s3.PutObjectTaggingOutput) {
	return nil, nil
}

func (m *S3API) PutPublicAccessBlock(*s3.PutPublicAccessBlockInput) (*s3.PutPublicAccessBlockOutput, error) {
	return nil, nil
}
func (m *S3API) PutPublicAccessBlockWithContext(aws.Context, *s3.PutPublicAccessBlockInput, ...request.Option) (*s3.PutPublicAccessBlockOutput, error) {
	return nil, nil
}
func (m *S3API) PutPublicAccessBlockRequest(*s3.PutPublicAccessBlockInput) (*request.Request, *s3.PutPublicAccessBlockOutput) {
	return nil, nil
}

func (m *S3API) RestoreObject(*s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
	return nil, nil
}
func (m *S3API) RestoreObjectWithContext(aws.Context, *s3.RestoreObjectInput, ...request.Option) (*s3.RestoreObjectOutput, error) {
	return nil, nil
}
func (m *S3API) RestoreObjectRequest(*s3.RestoreObjectInput) (*request.Request, *s3.RestoreObjectOutput) {
	return nil, nil
}

func (m *S3API) SelectObjectContent(*s3.SelectObjectContentInput) (*s3.SelectObjectContentOutput, error) {
	return nil, nil
}
func (m *S3API) SelectObjectContentWithContext(aws.Context, *s3.SelectObjectContentInput, ...request.Option) (*s3.SelectObjectContentOutput, error) {
	return nil, nil
}
func (m *S3API) SelectObjectContentRequest(*s3.SelectObjectContentInput) (*request.Request, *s3.SelectObjectContentOutput) {
	return nil, nil
}

func (m *S3API) UploadPart(*s3.UploadPartInput) (*s3.UploadPartOutput, error) {
	return nil, nil
}
func (m *S3API) UploadPartWithContext(aws.Context, *s3.UploadPartInput, ...request.Option) (*s3.UploadPartOutput, error) {
	return nil, nil
}
func (m *S3API) UploadPartRequest(*s3.UploadPartInput) (*request.Request, *s3.UploadPartOutput) {
	return nil, nil
}

func (m *S3API) UploadPartCopy(*s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
	return nil, nil
}
func (m *S3API) UploadPartCopyWithContext(aws.Context, *s3.UploadPartCopyInput, ...request.Option) (*s3.UploadPartCopyOutput, error) {
	return nil, nil
}
func (m *S3API) UploadPartCopyRequest(*s3.UploadPartCopyInput) (*request.Request, *s3.UploadPartCopyOutput) {
	return nil, nil
}

func (m *S3API) WaitUntilBucketExists(*s3.HeadBucketInput) error {
	return nil
}
func (m *S3API) WaitUntilBucketExistsWithContext(aws.Context, *s3.HeadBucketInput, ...request.WaiterOption) error {
	return nil
}

func (m *S3API) WaitUntilBucketNotExists(*s3.HeadBucketInput) error {
	return nil
}
func (m *S3API) WaitUntilBucketNotExistsWithContext(aws.Context, *s3.HeadBucketInput, ...request.WaiterOption) error {
	return nil
}

func (m *S3API) WaitUntilObjectExists(*s3.HeadObjectInput) error {
	return nil
}
func (m *S3API) WaitUntilObjectExistsWithContext(aws.Context, *s3.HeadObjectInput, ...request.WaiterOption) error {
	return nil
}

func (m *S3API) WaitUntilObjectNotExists(*s3.HeadObjectInput) error {
	return nil
}
func (m *S3API) WaitUntilObjectNotExistsWithContext(aws.Context, *s3.HeadObjectInput, ...request.WaiterOption) error {
	return nil
}
