package core


const (
	Static = "static"
	Dynamic = "dynamic"
)

/*`{
    "appName":"app_1", //应用名
    "group":"group_1", //应用分组
    "gpu":1, //gpu资源
    "cpu":8, //HT资源
    "ram":16, //内存资源
    "disk":60, //磁盘资源
    "replicas":432 //此集群扩容的实例数
	}`
*/
type App struct {
	AppName  string `json:"appName"`
	Group    string `json:"group"`
	Gpu      int64  `json:"gpu"`
	Cpu      int64  `json:"cpu"`
	Ram      int64  `json:"ram"`
	Disk     int64  `json:"disk"`
	Replicas int64  `json:"replicas"`
}

/*`{
    "sn":"machine_1", // 此宿主机唯一表述。
    "smName":"smName_1", //宿主机机型，不同机型资源规格上会有差异
    "gpu":8,
    "cpu":104,
    "ram":186,
    "disk":1000,
    "eni":10, //弹性网卡ENI约束，约束当前宿主机上容器实例数不允许超过eni数量
    "topologies":[
        {
            "socket":0,
            "core":0,
            "cpu":0
        }
    ]
	}`
*/
type Node struct {
	Sn         string `json:"sn"`
	SmName     string `json:"smName"`
	Gpu        int64  `json:"gpu"`
	Cpu        int64  `json:"cpu"`
	Ram        int64  `json:"ram"`
	Disk       int64  `json:"disk"`
	Eni        int64  `json:"eni"`
	Topologies []Topology
}

type Topology struct {
	Socket int   `json:"socket"`
	Core   int64 `json:"core"`
	Cpu    int64 `json:"cpu"`
}

/* `{
       "appName":"app_1",
       "group":"group_1",
       "gpu":1,
       "cpu":4,
       "ram":8,
       "disk":60,
       "cpuIDs":[
           0,
           1,
           2,
           3
       ]
   }\
*/
type Pod struct {
	AppName string `json:"appName"`
	Group   string `json:"group"`
	Gpu     int64  `json:"gpu"`
	Cpu     int64  `json:"cpu"`
	Ram     int64  `json:"ram"`
	Disk    int64  `json:"disk"`
	CpuIDs  []int  `json:"cpuIDs"`
}

/*`{
      "group":"group_1",
      "compactness":true,
      "maxInstancePerHost":2
  }`
*/
type GroupMaxInstancePerHost struct {
	Group              string `json:"group"`
	Compactness        bool   `json:"compactness"`
	MaxInstancePerHost int64  `json:"maxInstancePerHost"`
}

/*`{
     "resource":"GPU",
     "smName":"smName_1",
     "weight":10
  }`
*/
type NodeResourceWeight struct {
	Resource string `json:"resource"`
	SmName   string `json:"smName"`
	Weight   int64  `json:"weight"`
}

/*`{
      "replicas":100,
      "restrain":"le",
      "maxInstancePerHost":1
  }`
*/
type ReplicasMaxInstancePerHost struct {
	Replicas           int64  `json:"replicas"`
	Restrain           string `json:"restrain"`
	MaxInstancePerHost int64  `json:"maxInstancePerHost"`
}

var RuleRes Rule

type Rule struct {
	DefaultMaxInstancePerHost   int64                        `json:"defaultMaxInstancePerHost"`
	GroupMaxInstancePerHosts    []GroupMaxInstancePerHost    `json:"groupMaxInstancePerHosts"`
	NodeResourceWeights         []NodeResourceWeight         `json:"nodeResourceWeights"`
	ReplicasMaxInstancePerHosts []ReplicasMaxInstancePerHost `json:"replicasMaxInstancePerHosts"`
	TimeLimitInMins             int64                        `json:"timeLimitInMins"`
}
