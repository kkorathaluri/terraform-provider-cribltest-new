terraform {
  required_providers {
    cribl-terraform = {
      source = "kkorathaluri/cribltest-new"
    }
  }
}

provider "cribl-terraform" {
  organization_id = "beautiful-nguyen-y8y4azd"
  workspace_id = "main"
}

resource "cribl-terraform_pipeline" "leaderlogs2metrics" {
  id = "leaderlogs2metrics"
  group_id = "default"
  conf = {
    async_func_timeout = 1000
    description = "Pipeline for processing leader logs and converting to metrics"
    output = "default"
    streamtags = []
    functions = [
      {
        id = "eval"
        filter = "true"
        disabled = false
        final = false
        group_id = "default"
        description = "Initial field evaluation"
        conf = {
          add = jsonencode([
            { name = "saas_domain", value = "saas_domain || __metadata.env.CRIBL_CLOUD_DOMAIN" },
            { disabled = false, name = "tenantId", value = "tenantId || __metadata.env.CRIBL_CLOUD_TENANT_ID" },
            { disabled = false, name = "accountId", value = "accountId || __metadata.env.CRIBL_CLOUD_ACCOUNT_ID" },
            { disabled = false, name = "fleet", value = "fleet || __metadata.cribl.group" },
            { disabled = false, name = "host", value = "instance" },
            { disabled = false, name = "event_message", value = "message" },
            { disabled = false, name = "workspace", value = "workspace || __metadata.env.CRIBL_WORKSPACE_NAME" }
          ])
          remove = jsonencode([])
        }
      },
      {
        id = "eval"
        filter = "channel == \"ProcessMetrics\" && cid.startsWith(\"service\")"
        conf = {
          add = jsonencode([
            {
              disabled = false
              name = "service_name"
              value = "cid.split(\":\")[1]"
            },
            {
              disabled = false
              name = "process"
              value = "source.match(/\\d+/)? source.match(/\\d+/)[0]: 0"
            }
          ])
        }
        group_id = "sbrucB"
      },
      {
        id = "publish_metrics"
        filter = "channel == \"ProcessMetrics\" && cid.startsWith(\"service\")"
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "accountId",
            "ignoreAlerts",
            "license",
            "ownerIsCribl",
            "ownerIsTestUser",
            "region",
            "saas_domain",
            "state",
            "tenantId",
            "fleet",
            "input",
            "output",
            "input_type",
            "output_type",
            "worker_group",
            "worker_node",
            "id",
            "channel",
            "host",
            "oemId",
            "url",
            "method",
            "workspace",
            "service_name",
            "process"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "gauge"
              outFieldExpr = "'cribl_process_metrics_cpu_perc'"
              inFieldName = "cpuPerc"
            },
            {
              metricType = "gauge"
              outFieldExpr = "'cribl_process_metrics_mem_usage'"
              inFieldName = "mem.rss"
            },
            {
              metricType = "gauge"
              inFieldName = "eluPerc"
              outFieldExpr = "'cribl_process_metrics_elu_perc'"
            }
          ])
        }
        group_id = "sbrucB"
        final = false
      },
      {
        id = "publish_metrics"
        filter = "source.endsWith('service/metrics/cribl.log')"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "accountId",
            "ignoreAlerts",
            "license",
            "ownerIsCribl",
            "ownerIsTestUser",
            "region",
            "saas_domain",
            "state",
            "tenantId",
            "fleet",
            "input",
            "output",
            "input_type",
            "output_type",
            "worker_group",
            "worker_node",
            "id",
            "channel",
            "host",
            "oemId",
            "url",
            "method",
            "workspace"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "gauge"
              outFieldExpr = "'cribl_metric_service_num_metrics'"
              inFieldName = "numMetrics"
            },
            {
              metricType = "gauge"
              inFieldName = "status"
              outFieldExpr = "'cribl_metric_service_response_status'"
            },
            {
              metricType = "distribution"
              inFieldName = "response_time"
              outFieldExpr = "'cribl_metric_service_response_time'"
            }
          ])
        }
        description = "Publish metrics service to grafana"
        group_id = "DlTnZf"
        final = true
      },
      {
        id = "eval"
        filter = "source.endsWith('service/lease/cribl.log')"
        disabled = false
        conf = {
          add = jsonencode([
            {
              disabled = false
              value = "channel == \"lease\" && level == \"error\" ? 1 : 0"
              name = "lease_error"
            }
          ])
        }
        group_id = "LhagcY"
      },
      {
        id = "publish_metrics"
        filter = "source.endsWith('service/lease/cribl.log')"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "accountId",
            "ignoreAlerts",
            "license",
            "ownerIsCribl",
            "ownerIsTestUser",
            "region",
            "saas_domain",
            "state",
            "tenantId",
            "fleet",
            "input",
            "output",
            "input_type",
            "output_type",
            "worker_group",
            "worker_node",
            "id",
            "channel",
            "host",
            "oemId",
            "url",
            "method",
            "workspace"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "counter"
              inFieldName = "lease_error"
              outFieldExpr = "'cribl_lease_service_error_count'"
            }
          ])
        }
        description = "publish lease service to grafana"
        group_id = "LhagcY"
        final = true
      },
      {
        id = "eval"
        filter = "source.endsWith('service/notifications/cribl.log') && level=='error'"
        disabled = false
        conf = {
          add = jsonencode([
            {
              disabled = false
              value = "1"
              name = "notifications_error"
            }
          ])
        }
        group_id = "AHoX5q"
      },
      {
        id = "publish_metrics"
        filter = "source.endsWith('service/notifications/cribl.log')"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "accountId",
            "ignoreAlerts",
            "license",
            "ownerIsCribl",
            "ownerIsTestUser",
            "region",
            "saas_domain",
            "state",
            "tenantId",
            "fleet",
            "input",
            "output",
            "input_type",
            "output_type",
            "worker_group",
            "worker_node",
            "id",
            "channel",
            "host",
            "oemId",
            "url",
            "method",
            "workspace"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "counter"
              inFieldName = "notifications_error"
              outFieldExpr = "'cribl_notifications_service_error_count'"
            }
          ])
        }
        description = "publish notifications service to grafana"
        group_id = "AHoX5q"
        final = false
      },
      {
        id = "eval"
        filter = "/service\\/\\w*connections\\/\\d+\\/cribl\\.log$/.test(source)"
        disabled = true
        conf = {
          add = jsonencode([
            {
              disabled = true
              name = "process"
              value = "/(\\d+)\\/cribl\\.log/.exec(source)[1]"
            },
            {
              disabled = true
              name = "__service"
              value = "/(\\w+)\\/\\d+\\/cribl\\.log$/.exec(source)[1]"
            },
            {
              disabled = true
              name = "error"
              value = "Number(level == 'error')"
            },
            {
              disabled = false
              name = "stream_leader_ok_configures"
              value = "Number(channel===\"CriblMaster\" && level===\"info\" && message == 'finished config update')"
            },
            {
              disabled = false
              name = "stream_leader_failed_configures"
              value = "Number(channel===\"CriblMaster\" && level===\"warn\" && message == 'failed config update')"
            },
            {
              disabled = false
              name = "__connMetric"
              value = "1"
            }
          ])
        }
        group_id = "HNdEbQ"
        description = "TODO use regex extract to not call exec twice"
      },
      {
        id = "publish_metrics"
        filter = "__connMetric == 1 && channel===\"CriblMaster\" && (stream_leader_ok_configures || stream_leader_failed_configures)"
        disabled = true
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "saas_domain"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "counter"
              inFieldName = "stream_leader_ok_configures"
              outFieldExpr = "stream_leader_ok_configures"
            },
            {
              metricType = "counter"
              inFieldName = "stream_leader_failed_configures"
              outFieldExpr = "stream_leader_failed_configures"
            }
          ])
        }
        group_id = "HNdEbQ"
        description = "Control Plane conn service metrics"
      },
      {
        id = "publish_metrics"
        filter = "__service != null"
        disabled = true
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "accountId",
            "ignoreAlerts",
            "license",
            "ownerIsCribl",
            "ownerIsTestUser",
            "region",
            "saas_domain",
            "state",
            "tenantId",
            "fleet",
            "input",
            "output",
            "input_type",
            "output_type",
            "worker_group",
            "worker_node",
            "id",
            "channel",
            "host",
            "oemId",
            "url",
            "method",
            "process",
            "workspace"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "gauge"
              inFieldName = "cpuPerc"
              outFieldExpr = "'cribl_$${__service}_process_metrics_cpu_perc'"
            },
            {
              metricType = "gauge"
              inFieldName = "mem.rss"
              outFieldExpr = "'cribl_$${__service}_process_metrics_mem_usage'"
            },
            {
              metricType = "counter"
              inFieldName = "error"
              outFieldExpr = "'cribl_$${__service}_process_metrics_error_count'"
            },
            {
              metricType = "gauge"
              inFieldName = "status"
              outFieldExpr = "'cribl_$${__service}_process_metrics_status'"
            },
            {
              metricType = "gauge"
              inFieldName = "response_time"
              outFieldExpr = "'cribl_$${__service}_process_metrics_response_time'"
            }
          ])
        }
        group_id = "HNdEbQ"
        description = "This can replace the 3 eval/publish groups we have for the metrics right now. cc: Kiran"
      },
      {
        id = "eval"
        filter = "source.includes('/log/service/connections/')"
        disabled = false
        conf = {
          add = jsonencode([
            {
              disabled = false
              value = "cid == \"service:connections\" && level == \"error\" ? 1 : 0"
              name = "error"
            },
            {
              disabled = false
              name = "process"
              value = "source.match(/\\d+/)? source.match(/\\d+/)[0]: 0"
            }
          ])
        }
        group_id = "IUhOiC"
      },
      {
        id = "publish_metrics"
        filter = "source.includes('/log/service/connections/')"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "accountId",
            "ignoreAlerts",
            "license",
            "ownerIsCribl",
            "ownerIsTestUser",
            "region",
            "saas_domain",
            "state",
            "tenantId",
            "fleet",
            "input",
            "output",
            "input_type",
            "output_type",
            "worker_group",
            "worker_node",
            "id",
            "channel",
            "host",
            "oemId",
            "url",
            "method",
            "process",
            "workspace"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "counter"
              inFieldName = "error"
              outFieldExpr = "'cribl_connections_service_error_count'"
            },
            {
              metricType = "gauge"
              outFieldExpr = "'cribl_connections_service_response_status'"
              inFieldName = "status"
            },
            {
              metricType = "gauge"
              outFieldExpr = "'cribl_connections_service_response_time'"
              inFieldName = "response_time"
            }
          ])
        }
        group_id = "IUhOiC"
        final = true
      },
      {
        id = "eval"
        filter = "source.includes('/log/service/stream_connections/')"
        disabled = false
        conf = {
          add = jsonencode([
            {
              disabled = false
              value = "cid == \"service:stream_connections\" && level == \"error\" ? 1 : 0"
              name = "error"
            },
            {
              disabled = false
              name = "process"
              value = "source.match(/\\d+/)? source.match(/\\d+/)[0]: 0"
            }
          ])
        }
        group_id = "Cgqrmv"
      },
      {
        id = "publish_metrics"
        filter = "source.includes('/log/service/stream_connections/')"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "accountId",
            "ignoreAlerts",
            "license",
            "ownerIsCribl",
            "ownerIsTestUser",
            "region",
            "saas_domain",
            "state",
            "tenantId",
            "fleet",
            "input",
            "output",
            "input_type",
            "output_type",
            "worker_group",
            "worker_node",
            "id",
            "channel",
            "host",
            "oemId",
            "url",
            "method",
            "process",
            "workspace"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "counter"
              inFieldName = "error"
              outFieldExpr = "'cribl_stream_connections_service_error_count'"
            },
            {
              metricType = "gauge"
              outFieldExpr = "'cribl_stream_connections_service_response_status'"
              inFieldName = "status"
            },
            {
              metricType = "gauge"
              outFieldExpr = "'cribl_stream_connections_service_response_time'"
              inFieldName = "response_time"
            }
          ])
        }
        group_id = "Cgqrmv"
        final = true
      },
      {
        id = "eval"
        filter = "source.includes('/log/service/edge_connections/')"
        disabled = false
        conf = {
          add = jsonencode([
            {
              disabled = false
              value = "cid == \"service:edge_connections\" && level == \"error\" ? 1 : 0"
              name = "error"
            },
            {
              disabled = false
              name = "process"
              value = "source.match(/\\d+/)? source.match(/\\d+/)[0]: 0"
            }
          ])
        }
        group_id = "uBHX8z"
      },
      {
        id = "publish_metrics"
        filter = "source.includes('/log/service/edge_connections/')"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "accountId",
            "ignoreAlerts",
            "license",
            "ownerIsCribl",
            "ownerIsTestUser",
            "region",
            "saas_domain",
            "state",
            "tenantId",
            "fleet",
            "input",
            "output",
            "input_type",
            "output_type",
            "worker_group",
            "worker_node",
            "id",
            "channel",
            "host",
            "oemId",
            "url",
            "method",
            "process",
            "workspace"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "counter"
              inFieldName = "error"
              outFieldExpr = "'cribl_edge_connections_service_error_count'"
            },
            {
              metricType = "gauge"
              outFieldExpr = "'cribl_edge_connections_service_response_status'"
              inFieldName = "status"
            },
            {
              metricType = "gauge"
              outFieldExpr = "'cribl_edge_connections_service_response_time'"
              inFieldName = "response_time"
            }
          ])
        }
        group_id = "uBHX8z"
        final = true
      },
      {
        id = "eval"
        filter = "source.includes('/log/service/connection_proxy/')"
        disabled = false
        conf = {
          add = jsonencode([
            {
              disabled = false
              value = "channel == \"failover\" && level == \"error\" ? 1 : 0"
              name = "connection_proxy_error"
            }
          ])
        }
        group_id = "ZO930W"
        description = "Evaluate errors from active active proxy"
      },
      {
        id = "publish_metrics"
        filter = "source.includes('/log/service/connection_proxy/')"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "accountId",
            "ignoreAlerts",
            "license",
            "ownerIsCribl",
            "ownerIsTestUser",
            "region",
            "saas_domain",
            "state",
            "tenantId",
            "fleet",
            "input",
            "output",
            "input_type",
            "output_type",
            "worker_group",
            "worker_node",
            "id",
            "channel",
            "host",
            "oemId",
            "url",
            "method",
            "workspace"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "counter"
              inFieldName = "connection_proxy_error"
              outFieldExpr = "'cribl_connection_proxy_service_error_count'"
            },
            {
              metricType = "gauge"
              inFieldName = "activeCxn"
              outFieldExpr = "'cribl_connection_proxy_service_active_cxn'"
            },
            {
              metricType = "gauge"
              inFieldName = "remoteActiveCxn"
              outFieldExpr = "'cribl_connection_proxy_service_active_output_cxn'"
            }
          ])
        }
        description = "publish connection proxy service to grafana"
        group_id = "ZO930W"
        final = false
      },
      {
        id = "eval"
        filter = "source.endsWith('/service/events/cribl.log')"
        disabled = false
        conf = {
          add = jsonencode([
            {
              disabled = false
              name = "events_error"
              value = "cid == \"service:events\" && level == \"error\" ? 1 : 0"
            }
          ])
        }
        group_id = "5hX27Q"
      },
      {
        id = "publish_metrics"
        filter = "source.endsWith('/service/events/cribl.log')"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "accountId",
            "ignoreAlerts",
            "license",
            "ownerIsCribl",
            "ownerIsTestUser",
            "region",
            "saas_domain",
            "state",
            "tenantId",
            "fleet",
            "input",
            "output",
            "input_type",
            "output_type",
            "worker_group",
            "worker_node",
            "id",
            "channel",
            "host",
            "oemId",
            "url",
            "method",
            "workspace"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "counter"
              inFieldName = "events_error"
              outFieldExpr = "'cribl_events_service_error_count'"
            }
          ])
        }
        group_id = "5hX27Q"
        final = true
      },
      {
        id = "eval"
        filter = "channel == \"failover\" && event_message.includes(\"standby\")"
        disabled = false
        conf = {
          add = jsonencode([
            {
              disabled = false
              name = "failover_count"
              value = "1"
            }
          ])
        }
        group_id = "Yv1d7R"
      },
      {
        id = "publish_metrics"
        filter = "channel == \"failover\" && event_message.includes(\"standby\")"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "accountId",
            "ignoreAlerts",
            "license",
            "ownerIsCribl",
            "ownerIsTestUser",
            "region",
            "saas_domain",
            "state",
            "tenantId",
            "fleet",
            "channel",
            "host",
            "oemId",
            "workspace"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "counter"
              inFieldName = "failover_count"
              outFieldExpr = "'cribl_leader_failover_count'"
            }
          ])
        }
        group_id = "Yv1d7R"
        final = true
      },
      {
        id = "publish_metrics"
        filter = "channel == \"ProcessMetrics\" && source == \"/opt/cribl/log/cribl.log\""
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "accountId",
            "ignoreAlerts",
            "license",
            "ownerIsCribl",
            "ownerIsTestUser",
            "region",
            "saas_domain",
            "state",
            "tenantId",
            "fleet",
            "channel",
            "host",
            "oemId",
            "workspace"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "gauge"
              inFieldName = "cpuPerc"
              outFieldExpr = "'cribl_api_service_process_metrics_cpu_perc'"
            },
            {
              metricType = "gauge"
              outFieldExpr = "'cribl_api_service_process_metrics_mem_rss'"
              inFieldName = "mem.rss"
            },
            {
              metricType = "gauge"
              outFieldExpr = "'cribl_api_service_process_metrics_mem_heap'"
              inFieldName = "mem.heap"
            },
            {
              metricType = "gauge"
              inFieldName = "eluPerc"
              outFieldExpr = "'cribl_api_service_process_metrics_elu_perc'"
            }
          ])
        }
        group_id = "UQpuD2"
        final = true
      },
      {
        id = "publish_metrics"
        filter = "segvkill_count > 0"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "accountId",
            "ignoreAlerts",
            "license",
            "ownerIsCribl",
            "ownerIsTestUser",
            "region",
            "saas_domain",
            "state",
            "tenantId",
            "fleet",
            "channel",
            "host",
            "oemId",
            "workspace"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "counter"
              inFieldName = "segvkill_count"
              outFieldExpr = "'cribl_system_process_crash_count'"
            }
          ])
        }
        group_id = "NAHxrw"
        final = true
      },
      {
        id = "aggregation"
        filter = "source.endsWith('/access.log')"
        disabled = false
        conf = {
          passthrough = true
          preserveGroupBys = false
          sufficientStatsOnly = false
          metricsMode = true
          timeWindow = jsonencode("60s")
          aggregations = jsonencode(["count().as(stream_leader_api_hit)"])
          cumulative = false
          flushOnInputClose = true
          groupbys = jsonencode(["status", "saas_domain"])
        }
        group_id = "Kn689s"
        description = "count http request by response"
        final = true
      },
      {
        id = "aggregation"
        filter = "source.endsWith('audit.log') && action==='deploy' && type==='groups'"
        disabled = false
        conf = {
          passthrough = false
          preserveGroupBys = false
          sufficientStatsOnly = false
          metricsMode = true
          timeWindow = jsonencode("60s")
          aggregations = jsonencode(["count().as(stream_leader_deploy_count)"])
          cumulative = false
          flushOnInputClose = true
          groupbys = jsonencode(["saas_domain"])
        }
        description = "count group deploys"
        group_id = "Kn689s"
        final = true
      },
      {
        id = "aggregation"
        filter = "level=='error' && source.endsWith(\"cribl.log\") && ([\"output:system_email\", \"CriblMaster\", \"Bundler\", \"bundleStore\", \"ConfigHelperKeyring\", \"PackExport\", \"CommittedCacheReaper\", \"PackMgr\", \"Configuration\", \"rest:OpsCloudUpgrade\", \"cloudAPI\"].includes(channel) || channel.startsWith(\"conf:\"))"
        disabled = false
        conf = {
          passthrough = false
          preserveGroupBys = false
          sufficientStatsOnly = false
          metricsMode = true
          timeWindow = jsonencode("60s")
          aggregations = jsonencode(["count().as(stream_leader_error)"])
          cumulative = false
          flushOnInputClose = true
          groupbys = jsonencode(["channel", "saas_domain"])
        }
        group_id = "Kn689s"
        description = "Control plane errors"
        final = true
      },
      {
        id = "publish_metrics"
        filter = "method=='POST' && source.endsWith('/access.log') && url.endsWith('/groups')"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "status",
            "saas_domain"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "gauge"
              inFieldName = "response_time"
              outFieldExpr = "'stream_leader_group_create_dur'"
            }
          ])
        }
        group_id = "Kn689s"
        description = "Group creation duration"
        final = false
      },
      {
        id = "publish_metrics"
        filter = "channel==='Bundler' && source.endsWith('cribl.log') && level=='info' && message=='finished creating group bundle'"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "saas_domain"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "gauge"
              inFieldName = "elapsed"
              outFieldExpr = "'stream_leader_bundling_dur'"
            }
          ])
        }
        group_id = "Kn689s"
        description = "deploy bundle creation time"
        final = true
      },
      {
        id = "publish_metrics"
        filter = "channel==='cribl' && cid=='api' && level=='info' && source.endsWith('cribl.log') && message=='config helper started'"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "saas_domain"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "gauge"
              inFieldName = "elapsed"
              outFieldExpr = "'stream_leader_ch_startup'"
            }
          ])
        }
        group_id = "Kn689s"
        description = "config helper startup time"
        final = true
      },
      {
        id = "publish_metrics"
        filter = "channel==='cribl' && cid=='api' && source.endsWith('cribl.log') && message=='API server started'"
        disabled = false
        conf = {
          overwrite = false
          dimensions = jsonencode([
            "saas_domain"
          ])
          removeMetrics = jsonencode([])
          removeDimensions = jsonencode([])
          fields = jsonencode([
            {
              metricType = "gauge"
              inFieldName = "uptime"
              outFieldExpr = "'stream_leader_api_startup'"
            }
          ])
        }
        group_id = "Kn689s"
        description = "leader api startup"
        final = true
      },
      {
        id = "eval"
        filter = "true"
        disabled = false
        conf = {
          remove = jsonencode([
            "appMetadata",
            "cid",
            "cribl_*",
            "error",
            "infrastructure",
            "infrastructureConfig",
            "level",
            "metronomeCustomerId",
            "oemCustomerId",
            "organization*",
            "source",
            "time",
            "message",
            "event_message"
          ])
        }
      },
      {
        id = "redis"
        filter = "true"
        disabled = false
        conf = {
          commands = jsonencode([
            {
              keyExpr = "`$${saas_domain}-organizations-$${tenantId}`"
              command = "get"
              outField = "organization"
            }
          ])
          deploymentType = jsonencode("standalone")
          authType = jsonencode("none")
          maxBlockSecs = jsonencode(5)
          tlsOptions = jsonencode({
            rejectUnauthorized = false
          })
          url = jsonencode("`redis://$${C.Lookup('config.csv', 'workspace').match(C.env.CRIBL_WORKSPACE_NAME, 'redis_host_read')}:6379`")
          enableClientSideCaching = jsonencode(true)
        }
        group_id = "56Hwu0"
      },
      {
        id = "serde"
        filter = "true"
        disabled = false
        conf = {
          mode = jsonencode("extract")
          type = jsonencode("json")
          srcField = jsonencode("organization")
          fields = jsonencode([])
          remove = jsonencode([
            "message",
            "channel"
          ])
        }
        description = "flatten organization"
        group_id = "56Hwu0"
      }
    ]
    output = "default"
    streamtags = []
  }
}

# Output the pipeline details
output "pipeline_details" {
  value = cribl-terraform_pipeline.leaderlogs2metrics
} 