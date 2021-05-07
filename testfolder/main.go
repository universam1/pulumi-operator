package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cluster, err := elasticache.NewCluster(ctx, "sam", &elasticache.ClusterArgs{
			ApplyImmediately: pulumi.BoolPtr(true),
			Engine:           pulumi.String("redis"),
			// EngineVersion:      pulumi.String("5.0.6"),
			// ParameterGroupName: pulumi.String("default.redis5.0"),
			EngineVersion:      pulumi.String("6.x"),
			ParameterGroupName: pulumi.String("default.redis6.x"),
			NodeType:           pulumi.String("cache.t3.micro"),
			NumCacheNodes:      pulumi.Int(1),
			Port:               pulumi.Int(6379),
			SecurityGroupIds:   pulumi.ToStringArray([]string{"sg-456db431"}),
		})
		if err != nil {
			return err
		}
		cn := cluster.CacheNodes
		ctx.Export("ClusterAddress", cn.Index(pulumi.Int(0)).Address())
		return nil
	})
}
