# aws tools

##AWS tools install

**install awscli**

```
sudo easy_install awscli

```

**conf aescli**
1>先在aws上创建访问秘钥
2>配置aws命令行

```
localhost:aws_workspace raylea$ aws configure
AWS Access Key ID [None]: AKIAP6RWFOFHKKGDFEEA
AWS Secret Access Key [None]: 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
Default region name [None]: cn-north-1
Default output format [None]: text

```

**ls aws s3**

```
aws s3 ls s3://microembedded
```

**sync aws**

```
aws s3 sync s3://microembedded .
```










